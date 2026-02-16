package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/filimonq/go-log-lint/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
