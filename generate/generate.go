package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rasa/vfsgen"
)

func doFiles(dir string, files string) {
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
		Include:         files, // requires https://github.com/shurcooL/vfsgen/pull/60
	})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	doFiles("conf", "config.json$")
	doFiles("static", "\\.(css|js)$")
	doFiles("template", "\\.html?$")
	doFiles("www", "\\.(ico|txt)$")
}
