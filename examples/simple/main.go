package main

import (
	"fmt"
	"os"

	"github.com/americanas-go/inject"
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
	"gopkg.in/yaml.v3"
)

func main() {

	inject.WithLogger(zerolog.NewLogger(zerolog.WithLevel("TRACE")))

	basePath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	log.Infof("current path is %s", basePath)

	graph, err := inject.BananaWithPath(basePath + "/examples/simple")
	if err != nil {
		log.Error(err.Error())
	}

	j, _ := yaml.Marshal(graph)
	fmt.Println(string(j))
}
