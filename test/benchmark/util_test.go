package benchmark_test

import (
	"os"

	jsonv2 "github.com/go-json-experiment/json"
)

func ReadFile(filename string, target any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = jsonv2.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
