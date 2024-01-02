package main

import (
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

func getMapKeys(endpointGroups map[string][]EndpointGroup) []string {
	keys := make([]string, 0, len(endpointGroups))
	for key := range endpointGroups {
		keys = append(keys, key)
	}
	return keys
}

// Read all files and return the filename and content
func readTemplateFiles(path string) ([][]string, error) {
	pattern := filepath.Join(path, "*.tmpl")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	templates := make([][]string, 0, len(files))
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		filename := filepath.Base(f)
		templates = append(templates, []string{filename[:len(filename)-5], string(b)})
	}
	return templates, nil
}

// Read all json files and return their json objects keyed by the filename
func readJSONSpecsFiles() (map[string]gjson.Result, error) {
	pattern := filepath.Join("./specs/", "*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	jsonFiles := make(map[string]gjson.Result, len(files))
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		filename := filepath.Base(f)
		jsonFiles[filename[:len(filename)-5]] = gjson.ParseBytes(content)
	}

	return jsonFiles, nil
}
