package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

func doDir(dir string) {
	path := "../" + dir
	file := path + "/assets_vfsdata.go"
	// Delete the old file.
	fmt.Printf("Removing %s\n", file)
	os.Remove(file)
	fs := http.Dir(path)
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:        file,
		PackageName:     dir,
		BuildTags:       "!dev",
		VariableName:    "Assets",
		VariableComment: "",
		// Exclude:         "\\.go$", // requires https://github.com/shurcooL/vfsgen/pull/60
	})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	doDir("www")
	doDir("static")
	doDir("template")
}
