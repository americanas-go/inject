package main

import (
	"fmt"

	"github.com/americanas-go/log"
)

func main() {

	zerolog.NewLogger()

	spec, err := ParseDir("/Users/joao.faria/Projetos/github.com/americanas-go/inject/examples/simple")
	if err != nil {
		log.Error(err.Error())
	}

	j, _ := yaml.Marshal(spec)
	fmt.Println(string(j))
}
