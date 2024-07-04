package main

import (
	"fmt"
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

// Read all template files and return its content keyed by the filename
func readTemplateFiles(path string) (map[string][]byte, error) {
	pattern := filepath.Join(path, "*.tmpl")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	templates := make(map[string][]byte, len(files))
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		filename := filepath.Base(f)
		templates[filename[:len(filename)-5]] = b
	}
	return templates, nil
}

// Read all json files and return their json objects keyed by the filename
func readJSONSpecsFiles(jsonFiles map[string]gjson.Result) error {
	pattern := filepath.Join("./specs/", "*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("no specs found, make sure to use UPDATE_SPECS=1")
	}

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		filename := filepath.Base(f)
		jsonFiles[filename[:len(filename)-5]] = gjson.ParseBytes(content)
	}

	return nil
}
