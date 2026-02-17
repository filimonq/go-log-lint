package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/filimonq/go-log-lint/internal/domain"
)

type LogVisitor struct {
	pass              *analysis.Pass
	rules             []domain.Rule
	supportedPackages map[string]bool
}

func NewLogVisitor(pass *analysis.Pass, rules []domain.Rule) *LogVisitor {
	return &LogVisitor{
		pass:  pass,
		rules: rules,
		supportedPackages: map[string]bool{
			"log/slog":        true,
			"go.uber.org/zap": true,
		},
	}
}

func (v *LogVisitor) Walk() {
	for _, file := range v.pass.Files {
		ast.Inspect(file, v.inspectNode)
	}
}

func (v *LogVisitor) inspectNode(node ast.Node) bool {
	callExpr, ok := node.(*ast.CallExpr)
	if !ok {
		return true
	}

	if !v.isLogPackage(callExpr) {
		return true
	}

	funcName, ok := v.extractFunctionName(callExpr.Fun)
	if !ok {
		return true
	}

	if !v.isLogFunction(funcName) {
		return true
	}

	msg, msgNode, ok := v.extractMessageArg(callExpr.Args)
	if !ok {
		return true
	}

	entry := domain.LogEntry{
		Message:  msg,
		Function: funcName,
		Pos:      msgNode.Pos(),
	}

	v.applyRules(entry, msgNode)

	return true
}

func (v *LogVisitor) applyRules(entry domain.LogEntry, msgNode ast.Expr) {
	for _, rule := range v.rules {
		issues := rule.Check(entry)
		for _, issue := range issues {

			diag := analysis.Diagnostic{
				Pos:     issue.Pos,
				Message: fmt.Sprintf("[%s] %s", rule.Name(), issue.Message),
			}

			if issue.Replacement != "" {
				newText := fmt.Sprintf("%q", issue.Replacement)

				diag.SuggestedFixes = []analysis.SuggestedFix{{
					Message: fmt.Sprintf("Fix with '%s'", issue.Replacement),
					TextEdits: []analysis.TextEdit{{
						Pos:     msgNode.Pos(),
						End:     msgNode.End(),
						NewText: []byte(newText),
					}},
				}}
			}

			v.pass.Report(diag)
		}
	}
}

func (v *LogVisitor) isLogPackage(node *ast.CallExpr) bool {
	var ident *ast.Ident

	switch fn := node.Fun.(type) {
	case *ast.SelectorExpr:
		ident = fn.Sel
	case *ast.Ident:
		ident = fn
	default:
		return false
	}

	obj := v.pass.TypesInfo.ObjectOf(ident)
	if obj == nil || obj.Pkg() == nil {
		return false
	}

	pkgPath := obj.Pkg().Path()
	return v.supportedPackages[pkgPath]
}

func (v *LogVisitor) extractFunctionName(expr ast.Expr) (string, bool) {
	switch fn := expr.(type) {
	case *ast.SelectorExpr:
		return fn.Sel.Name, true
	case *ast.Ident:
		return fn.Name, true
	}
	return "", false
}

func (v *LogVisitor) extractMessageArg(args []ast.Expr) (string, ast.Expr, bool) {
	if len(args) == 0 {
		return "", nil, false
	}

	arg := args[0]
	val, _, ok := v.resolveStringValue(arg)
	if !ok {
		return "", nil, false
	}

	return val, arg, true
}

func (v *LogVisitor) resolveStringValue(expr ast.Expr) (string, token.Pos, bool) {
	switch t := expr.(type) {
	case *ast.BasicLit:
		if t.Kind == token.STRING {
			val, err := strconv.Unquote(t.Value)
			if err != nil {
				return "", 0, false
			}
			return val, t.Pos(), true
		}
	case *ast.BinaryExpr:
		if t.Op == token.ADD {
			left, pos, okL := v.resolveStringValue(t.X)
			right, _, okR := v.resolveStringValue(t.Y)
			if okL && okR {
				return left + right, pos, true
			}
		}
	}
	return "", 0, false
}

func (v *LogVisitor) isLogFunction(name string) bool {
	prefixes := []string{"Debug", "Info", "Warn", "Error", "Fatal", "Panic"}
	for _, p := range prefixes {
		if strings.HasPrefix(name, p) {
			return true
		}
	}
	return false
}
