package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	count, err := CountGoroutines("testf.go", "mutexIncrement")
	fmt.Println(count, err)
}

func CountGoroutines(fileName, funcName string) (uint16, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return 0, err
	}

	var count uint16
	for _, decl := range astFile.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if fn.Name.Name != funcName {
			continue
		}

		ast.Inspect(fn, func(n ast.Node) bool {
			// Find Go Statements
			_, ok := n.(*ast.GoStmt)
			if !ok {
				return true
			}
			count++
			return true
		})

		fmt.Println(fn.Name.Name)

	}
	return count, nil
}
