package main

import (
	"context"
	"os"
	"os/exec"

	"github.com/americanas-go/inject"
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

func main() {

	inject.WithLogger(zerolog.NewLogger(zerolog.WithLevel("INFO")))

	basePath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	log.Infof("current path is %s", basePath)

	err = inject.WithPath(context.Background(), basePath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("go mod tidy failed: %v", err)
	}

	cmd = exec.Command("go", "mod", "vendor")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("go mod vendor failed: %v", err)
	}

}
