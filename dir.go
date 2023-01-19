package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/americanas-go/log"
)

const (
	an = "@B"
	cb = "//"
)

func ParseDir(path string) (Spec, error) {
	d, err := parseDir(path)
	if err != nil {
		return Spec{}, err
	}

	annon := filterAnnotations(d)
	spec, err := annon.ToSpec()
	if err != nil {
		return Spec{}, err
	}

	return spec, nil
}

func parseDir(path string) (map[string]*ast.Package, error) {

	log.Infof("parsing dir %s", path)

	fset := token.NewFileSet() // positions are relative to fset
	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func filterAnnotations(d map[string]*ast.Package) (m Annotations) {
	m = make(map[string][]string)
	for _, p := range d {
		log.Infof("parsing package %s", p.Name)
		for k, f := range p.Files {
			log.Infof("parsing file %s", k)
			for _, g := range f.Comments {
				var contains bool
				var cmts []string
				for _, c := range g.List {
					if strings.Contains(c.Text, an) {
						contains = true
						cmts = append(cmts,
							strings.ReplaceAll(c.Text,
								strings.Join([]string{cb, an, ""}, " "), ""))
					}
				}
				if contains {
					w := strings.Split(strings.ReplaceAll(g.List[0].Text, cb, ""), " ")
					n := w[1]
					m[n] = cmts
				}
			}
		}
	}
	return m
}
