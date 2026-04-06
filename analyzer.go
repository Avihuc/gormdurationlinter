// Package gormdurationlinter provides a Go analysis pass that reports
// time.Duration fields with gorm type:bigint missing serializer:nanoduration.
package gormdurationlinter

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gormduration",
	Doc:      "reports time.Duration fields with gorm type:bigint missing serializer:nanoduration",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		st := n.(*ast.StructType)
		for _, field := range st.Fields.List {
			if !isTimeDuration(pass, field.Type) {
				continue
			}
			if field.Tag == nil {
				continue
			}
			tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
			gormTag := tag.Get("gorm")
			if gormTag == "" {
				continue
			}
			if !hasGormOption(gormTag, "type", "bigint") {
				continue
			}
			if hasGormOption(gormTag, "serializer", "nanoduration") {
				continue
			}
			pass.Reportf(field.Pos(), "time.Duration field with gorm type:bigint must include serializer:nanoduration")
		}
	})

	return nil, nil
}

func isTimeDuration(pass *analysis.Pass, expr ast.Expr) bool {
	typ := pass.TypesInfo.TypeOf(expr)
	if typ == nil {
		return false
	}
	if ptr, ok := typ.(*types.Pointer); ok {
		typ = ptr.Elem()
	}
	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	return obj.Name() == "Duration" && obj.Pkg() != nil && obj.Pkg().Path() == "time"
}

func hasGormOption(gormTag, key, value string) bool {
	for _, part := range strings.Split(gormTag, ";") {
		part = strings.TrimSpace(part)
		if k, v, ok := strings.Cut(part, ":"); ok {
			if strings.TrimSpace(k) == key && strings.TrimSpace(v) == value {
				return true
			}
		}
	}
	return false
}
