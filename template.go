package inject

var (
	temp = `package {{.Package}}
import (
{{range .Imports}}
"{{.}}"
{{end}}
)

func Name() string {
return "{{.Name}}"
}`
)

type Template struct {
	Package string
	Imports []string
	Name    string
}
