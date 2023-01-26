//go:generate go-enum -f=$GOFILE --marshal
package inject

// ENUM(MODULE,PROVIDE,INJECT,INVOKE)
type AnnotationType int

// ENUM(MODULE,PATH,PACKAGE,FUNC)
type ModuleAttr int

type Annotation struct {
	Instance string
	Name     string
	Group    string
}
