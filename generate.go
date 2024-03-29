package main

import (
	"github.com/solo-io/solo-kit/pkg/code-generator/cmd"
	"log"
)

//go:generate go run generate.go

func main() {

	log.Printf("starting generate")
	if err := cmd.Generate(cmd.GenerateOptions{
		RelativeRoot:       ".",
		CompileProtos:      true,
		SkipGenMocks:       true,
		SkipGeneratedTests: true,
	}); err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
}
