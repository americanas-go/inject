package simple

import (
	"context"
	"net/http"
)

// ExampleStruct Lorem ipsum dolor sit amet, consectetur adipiscing elit
// @B Packagegithub.com/americanas-go/annotation
// @B RelativePackage examples/simple
// @B App xpto
// @B HandlerType HTTP
// @B Type Interface
type ExampleStruct struct {
}

func (t *ExampleStruct) FooStructMethod(ctx context.Context, r *http.Request) (interface{}, error) {
	return Response{
		Message: "Hello world",
	}, nil
}

// FooMethod Lorem ipsum dolor sit amet, consectetur adipiscing elit
// @B Packagegithub.com/americanas-go/annotation
// @B RelativePackage examples/simple
// @B App xpto
// @B HandlerType HTTP
// @B Type Function
// @B Path /foo
// @B Path /
// @B Method POST
// @B Consume application/json
// @B Consume application/yaml
// @B Produce application/json
// @B Param query foo bool true tiam sed efficitur purus
// @B Param query bar string true tiam sed efficitur purus
// @B Param path foo string tiam sed efficitur purus
// @B Param path bar string tiam sed efficitur purus
// @B Param header foo string true tiam sed efficitur purus
// @B Param header bar string true tiam sed efficitur purus
// @B Bodygithub.com/americanas-go/annotation/examples/simple.Request
// @B Response 201github.com/americanas-go/annotation/examples/simple.Response tiam sed efficitur purus, at lacinia magna
func FooMethod(ctx context.Context, r *http.Request) (interface{}, error) {
	return Response{
		Message: "Hello world",
	}, nil
}

type Response struct {
	Message string
}
