package inject

import (
	"strings"

	"github.com/americanas-go/annotation"
	ustrings "github.com/americanas-go/utils/strings"
	"github.com/autom8ter/dagger"
)

func BananaWithPath(path string) (*dagger.Graph, error) {
	blocks, err := annotation.Collect(path)
	if err != nil {
		return nil, err
	}

	return Banana(blocks)
}

func Banana(blocks []annotation.Block) (*dagger.Graph, error) {

	graph := dagger.NewGraph()

	for _, block := range blocks {
		if !block.IsFunc() {
			continue
		}

		path := dagger.Path{
			XID: xid(block),
		}

		attr := make(map[string]interface{})
		attr[ModuleAttrMODULE.String()] = block.Module
		attr[ModuleAttrPATH.String()] = block.Path
		attr[ModuleAttrPACKAGE.String()] = block.Package
		attr[ModuleAttrFUNC.String()] = block.Func

		var isValid bool
		for _, ann := range block.Annotations {
			if isValidAnnotation(ann.Name) {
				isValid = true

				annType, _ := ParseAnnotationType(strings.ToUpper(ann.Name))

				switch annType {
				case AnnotationTypePROVIDE:
					path.XType = "provider"
					a := Annotation{}
					err := ann.AsStruct(&a)
					if err != nil {
						return nil, err
					}

					to := dagger.Path{
						XID:   "",
						XType: "reference",
					}
					graph.SetEdge(path, to, dagger.Node{
						Path: dagger.Path{
							XType: "provide",
						}})
				case AnnotationTypeINJECT:

					a := Annotation{}
					err := ann.AsStruct(&a)
					if err != nil {
						return nil, err
					}

					from := dagger.Path{
						XID:   "",
						XType: "reference",
					}
					graph.SetEdge(from, path, dagger.Node{
						Path: dagger.Path{
							XType: "inject",
						}})
				case AnnotationTypeINVOKE:
					path.XType = "invoker"
				case AnnotationTypeMODULE:

				}

			}
		}

		if isValid {
			graph.SetNode(path, attr)
		}
	}

	return graph, nil

}

func xid(block annotation.Block) string {
	return strings.ToLower(
		strings.Join([]string{
			block.Module + block.Path,
			block.Func,
		}, "_"))
}

func isValidAnnotation(value string) bool {
	if ustrings.SliceContains([]string{
		AnnotationTypeMODULE.String(),
		AnnotationTypePROVIDE.String(),
		AnnotationTypeINJECT.String(),
		AnnotationTypeINVOKE.String()},
		strings.ToUpper(value)) {
		return true
	}
	return false
}
