package inject

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/americanas-go/annotation"
	"go/format"
	"golang.org/x/tools/go/packages"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ModuleData struct {
	PackageName  string
	FunctionName string
	ImportPath   string
	Modules      []ImportData
	Imports      []ImportData
	Alias        string
	Entry        annotation.Entry
	Type         string
}

type ImportData struct {
	Alias string
	Path  string
	Entry annotation.Entry
}

func WithPath(ctx context.Context, path string) error {
	annotation.WithLogger(log)
	collector, err := annotation.Collect(
		annotation.WithPath(path),
		annotation.WithFilters("Inject", "Provide", "Invoke"),
	)
	if err != nil {
		return err
	}

	entries := collector.Entries()

	graph, err := NewGraphFromEntries(ctx, entries)
	if err != nil {
		return err
	}

	moduleName, err := getModuleName(path)
	if err != nil {
		return err
	}

	for _, vert := range graph.VerticesWithNoIncomingEdges() {
		err := generateModuleFile(graph, vert, moduleName)
		if err != nil {
			log.Errorf("Error generating module file: %v", err)
		}
	}

	// graph.Print()
	// graph.ExportToGraphviz("examples/simple/simple.gv")

	j1, _ := yaml.Marshal(entries)
	fmt.Println(string(j1))

	return err
}

func getType(annons []annotation.Annotation) string {
	for _, ann := range annons {
		if strings.ToUpper(ann.Name) == AnnotationTypeINVOKE.String() {
			return AnnotationTypeINVOKE.String()
		}
	}

	return AnnotationTypePROVIDE.String()
}

func generateModuleFile(graph *Graph[annotation.Entry], vertex *Vertex[annotation.Entry], moduleName string) error {
	entry := vertex.Value

	packageName := filepath.Base(entry.Path)
	funcName := entry.Func.Name

	data := ModuleData{
		PackageName:  packageName,
		FunctionName: funcName,
		ImportPath:   entry.Path,
		Alias:        generateAlias(entry.Path),
		Entry:        entry,
		Type:         getType(entry.Annotations),
	}

	// Rastrear as importações únicas
	uniqueImports := make(map[string]struct{})

	// Processar cada vértice adjacente
	for _, v := range vertex.Incoming() {

		var alias string
		if v.Value.Package != packageName {
			alias = generateAlias(v.Value.Path)
		}

		data.Modules = append(data.Modules, ImportData{Alias: alias, Entry: v.Value})

		if v.Value.Package == packageName {
			continue
		}

		importPath := strings.ReplaceAll(v.Value.Path, "github.com/", "")
		fullImportPath := moduleName + "/gen/inject/" + importPath

		if _, exists := uniqueImports[alias]; !exists {
			uniqueImports[alias] = struct{}{}
			data.Imports = append(data.Imports, ImportData{Alias: alias, Path: fullImportPath, Entry: v.Value})
		}
	}

	tmpl, err := template.New("module").Parse(moduleTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	repoPath := strings.ReplaceAll(vertex.Value.Path, "github.com/", "")
	fileName := fmt.Sprintf("%s_module.go", strings.ToLower(funcName))
	filePath := filepath.Join("gen", "inject", repoPath, fileName)

	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directories: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(formatted)
	if err != nil {
		return err
	}

	for _, v := range vertex.Adjacent() {
		err := generateModuleFile(graph, v, moduleName)
		if err != nil {
			return fmt.Errorf("error generating module file: %v", err)
		}
	}

	return nil
}

func getModuleName(basePath string) (string, error) {
	cfg := &packages.Config{Mode: packages.NeedName | packages.NeedModule, Dir: basePath}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		return "", err
	}
	if len(pkgs) == 0 {
		return "", fmt.Errorf("no packages found in %s", basePath)
	}
	if pkgs[0].Module == nil {
		return "", fmt.Errorf("no module information found in %s", basePath)
	}
	return pkgs[0].Module.Path, nil
}

func generateAlias(packagePath string) string {
	hash := md5.Sum([]byte(packagePath))
	return hex.EncodeToString(hash[:])
}
