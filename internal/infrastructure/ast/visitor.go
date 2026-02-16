package ast

import (
	"go/ast"
	"go/token"
	"strconv"

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

	msg, pos, ok := v.extractMessageArg(callExpr.Args)
	if !ok {
		return true
	}

	entry := domain.LogEntry{
		Message:  msg,
		Function: funcName,
		Pos:      pos,
	}

	v.applyRules(entry)

	return true
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

func (v *LogVisitor) extractMessageArg(args []ast.Expr) (string, token.Pos, bool) {
	if len(args) == 0 {
		return "", 0, false
	}

	arg := args[0]

	lit, ok := arg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", 0, false
	}

	val, err := strconv.Unquote(lit.Value)
	if err != nil {
		return "", 0, false
	}

	return val, lit.Pos(), true
}

func (v *LogVisitor) isLogFunction(name string) bool {
	switch name {
	case "Debug", "Debugf", "Debugw",
		"Info", "Infof", "Infow",
		"Warn", "Warnf", "Warnw",
		"Error", "Errorf", "Errorw",
		"Fatal", "Fatalf", "Fatalw",
		"Panic", "Panicf", "Panicw":
		return true
	}
	return false
}

func (v *LogVisitor) applyRules(entry domain.LogEntry) {
	for _, rule := range v.rules {
		issues := rule.Check(entry)
		for _, issue := range issues {
			v.pass.Reportf(issue.Pos, "[%s] %s", rule.Name(), issue.Message)
		}
	}
}
