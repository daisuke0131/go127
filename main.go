package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main(){
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset,"testdata/a.go",nil,0)
	if err != nil{
		panic(err)
	}

	for _, d := range f.Decls {
		switch d := d.(type){
		case *ast.GenDecl:
			for _, e := range d.Specs{
				switch e := e.(type){
				case *ast.TypeSpec:
					switch c := e.Type.(type) {
					case *ast.StructType:
						for _,f := range c.Fields.List{
							if f.Tag != nil {
								for _, r := range f.Names{
									if isLowerChar(rune(r.Name[0])){
										fmt.Println("private field with tag!!")
										fmt.Println("struct:" + e.Name.Name)
										fmt.Println("field:" + r.Name)
										fmt.Println("tag:" + f.Tag.Value)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func isLowerChar(c rune) bool{
	lowerA := rune('a')
	lowerZ := rune('z')
	if c >= lowerA && c <= lowerZ{
		return true
	}
	return false
}
