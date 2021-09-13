package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

type Visitors map[string]*Visitor
type OccurrenceMap map[*ast.FuncDecl][]string

type Visitor struct {
	Occurrence OccurrenceMap
}

var Blacklist = []string{"MustQuery", "MustExec", "mustExecute"}

func (v *Visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch n := node.(type) {
	case *ast.FuncDecl:
		if !strings.HasPrefix(n.Name.String(), "Test") {
			return v
		}
		v.Occurrence[n] = make([]string, 0)
		v.checkBlockStmt(n.Body, n)
	}
	return v
}

func (v *Visitor) checkBlockStmt(block *ast.BlockStmt, n *ast.FuncDecl) {
	for _, stmt := range block.List {
		switch e := stmt.(type) {
		case *ast.ExprStmt:
			if call, ok := e.X.(*ast.CallExpr); ok {
				v.checkCallExpr(call, n)
			}
		case *ast.IfStmt:
			v.checkBlockStmt(e.Body, n)
		case *ast.ForStmt:
			v.checkBlockStmt(e.Body, n)
		case *ast.RangeStmt:
			v.checkBlockStmt(e.Body, n)
		case *ast.GoStmt:
			v.checkCallExpr(e.Call, n)
		}
	}
}

func (v *Visitor) checkCallExpr(call *ast.CallExpr, n *ast.FuncDecl) {
	switch f := call.Fun.(type) {
	case *ast.Ident: // check MustExec(...)
		v.checkEvil(nil, f, n)
	case *ast.SelectorExpr: // check (Ident or SelectorExpr).MustExec(...)
		v.checkSelectorExpr(f, n)
	}
}

func (v *Visitor) checkSelectorExpr(expr *ast.SelectorExpr, node *ast.FuncDecl) {
	if c, ok := expr.X.(*ast.CallExpr); ok {
		if s, ok := c.Fun.(*ast.SelectorExpr); ok { // check session.MustExec(...).Check(...)
			v.checkSelectorExpr(s, node)
		}
	}
	if c, ok := expr.X.(*ast.Ident); ok {
		v.checkEvil(c, expr.Sel, node)
	} else {
		v.checkEvil(nil, expr.Sel, node)
	}
}

func (v *Visitor) checkEvil(callee *ast.Ident, name *ast.Ident, node *ast.FuncDecl) {
	n := name.Name
	for _, f := range Blacklist {
		if f == n {
			v.Occurrence[node] = append(v.Occurrence[node], callee.String())
		}
	}
}

func Calculate(paths []string) (Visitors, error) {
	visitors := make(Visitors)
	for _, path := range paths {
		visitor := &Visitor{Occurrence: make(OccurrenceMap)}
		p, _ := filepath.Abs(path)
		f, err := parser.ParseFile(token.NewFileSet(), p, nil, parser.AllErrors)
		// ast.Print(fset, f)
		if err != nil {
			return nil, err
		}
		ast.Walk(visitor, f)
		visitors[p] = visitor
	}
	return visitors, nil
}

func Print(paths []string) error {
	fset := token.NewFileSet()
	for _, path := range paths {
		p, _ := filepath.Abs(path)
		f, err := parser.ParseFile(fset, p, nil, parser.AllErrors)
		if err != nil {
			return err
		}
		ast.Print(fset, f)
	}
	return nil
}
