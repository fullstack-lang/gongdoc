package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// Note is a UML note in a class diagram
type Note struct {
	Name          string
	Content       string
	X, Y          float64
	Width, Heigth float64
}

// Marshall provides the element of note as declaration
func (note *Note) Marshall(file *os.File, nbIndentation int) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "{\n")
	{
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tName:    \"%s\",\n", note.Name)
	}
	{
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tContent: \"%s\",\n", note.Content)
	}
	if note.X != 0.0 {
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tX:       %f,\n", note.X)
	}
	if note.Y != 0.0 {
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tY:       %f,\n", note.Y)
	}
	if note.Width != 0.0 {
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tWidth:   %f,\n", note.Width)
	}
	if note.Heigth != 0.0 {
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tHeigth:  %f,\n", note.Heigth)
	}

	indent(file, nbIndentation)
	fmt.Fprintf(file, "}")

	return nil
}

// Unmarshall updates note values from an ast.Epr
func (note *Note) Unmarshall(modelPkg *gong_models.ModelPkg, expr ast.Expr, fset *token.FileSet) {

	// expression should be a composite literal expression
	// models.Note{Position: uml.Position{X: 10, Y: 12}, Struct: &(models.Point{})}
	cl, ok := expr.(*ast.CompositeLit)
	if !ok {
		log.Fatal("Expecting a composite litteral " +
			fset.Position(expr.Pos()).String())
	}

	for _, elt := range cl.Elts {
		var kve *ast.KeyValueExpr
		if kve, ok = elt.(*ast.KeyValueExpr); !ok {
			log.Fatal("Element should be a KeyValueExpr" + fset.Position(elt.Pos()).String())
		}

		var ident *ast.Ident
		if ident, ok = kve.Key.(*ast.Ident); !ok {
			log.Fatal("Element Key should be an Ident" + fset.Position(kve.Pos()).String())
		}

		switch bl := kve.Value.(type) {
		case *ast.BasicLit:
			switch ident.Name {
			case "Name":
				note.Name = strings.TrimPrefix(bl.Value, "\"")
				note.Name = strings.TrimSuffix(note.Name, "\"")
			case "Content":
				note.Content = strings.TrimPrefix(bl.Value, "\"")
				note.Content = strings.TrimSuffix(note.Content, "\"")
			case "X":
				var err error
				note.X, err = strconv.ParseFloat(bl.Value, 64)
				if err != nil {
					log.Panic("Problem parsing float" + fset.Position(bl.Pos()).String())
				}
			case "Y":
				var err error
				note.Y, err = strconv.ParseFloat(bl.Value, 64)
				if err != nil {
					log.Panic("Problem parsing float" + fset.Position(bl.Pos()).String())
				}
			case "Width":
				var err error
				note.Width, err = strconv.ParseFloat(bl.Value, 64)
				if err != nil {
					log.Panic("Problem parsing float" + fset.Position(bl.Pos()).String())
				}
			case "Heigth":
				var err error
				note.Heigth, err = strconv.ParseFloat(bl.Value, 64)
				if err != nil {
					log.Panic("Problem parsing float" + fset.Position(bl.Pos()).String())
				}
			default:
				log.Fatal("Unkown key" +
					fset.Position(kve.Pos()).String())
			}
		default:
			log.Fatal("Element should be a basic lit" + fset.Position(kve.Pos()).String())
		}

	}
}
