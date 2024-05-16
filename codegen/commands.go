package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/tidwall/gjson"
)

func DownloadAndSaveSpecs(spec []string) error {
	url, filePath := spec[0], spec[1]
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download '%s'. Error: '%w'", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		err = os.WriteFile(filePath, body, 0644)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}

		fmt.Printf("Saving '%s'\n", filePath)
		return nil
	}

	return fmt.Errorf("failed to download '%s'. Status code: '%d'", url, response.StatusCode)
}

// Generates the code from the specs
func Compile() error {
	err := readJSONSpecsFiles(specs)
	if err != nil {
		return err
	}

	specVersion := specs["spec"].Get("info.version").String()
	if specVersion == "" {
		return fmt.Errorf("spec version not found, is the spec up to date?")
	}

	fmt.Printf("Current spec version: %s\n", specVersion)

	fmt.Printf("Compiling API package\n")
	err = compileApi(specs, specVersion)
	if err != nil {
		return err
	}

	fmt.Printf("Compiling client packages\n")
	return compileClients(specs, specVersion)
}

// Formats the generated code with 'goimports', 'gofmt' and 'betteralign'
func Format() error {
	fmt.Printf("Running 'goimports'\n")
	err := runCommand(exec.Command("goimports", "-w", "../"))
	if err != nil {
		return err
	}

	fmt.Printf("Running 'gofmt'\n")
	err = runCommand(exec.Command("gofmt", "-w", "../"))
	if err != nil {
		return err
	}

	fmt.Printf("Running 'betteralign'\n")
	err = runCommand(exec.Command("betteralign", "--apply", "../..."))
	if err != nil {
		return err
	}
	return runCommand(exec.Command("betteralign", "--apply", "../..."))
}

func runCommand(cmd *exec.Cmd) error {
	if output, err := cmd.CombinedOutput(); err != nil {
		out := string(output)
		if cmd.Args[0] == "betteralign" && strings.Contains(out, "struct with") {
			return nil
		}
		fmt.Printf("Error executing command '%s'\nOutput:\n%s", cmd.String(), out)
		return err
	}
	return nil
}

func compileApi(specs map[string]gjson.Result, specVersion string) error {
	preamble := preamble("api", specVersion)
	regionalRoutes := getRouteConstants(specs["routesTable"], "regional")
	ctx := pongo2.Context{
		"Preamble":       preamble,
		"RegionalRoutes": regionalRoutes,
	}

	// Reading api templates
	templates, err := readTemplateFiles("./templates/api")
	if err != nil {
		return err
	}

	results := make(map[string][]byte, len(templates))

	// Compiling templates
	for filename, template := range templates {
		fmt.Printf("Generating '%s' file\n", filename)

		tmpl, err := pongo2.FromBytes(template)
		if err != nil {
			return err
		}

		result, err := tmpl.ExecuteBytes(ctx)
		if err != nil {
			return err
		}

		results[filename] = result
	}

	// Writing results
	for filename, result := range results {
		err := os.WriteFile("../api/"+filename+".go", result, 0644)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	return nil
}

func compileClients(specs map[string]gjson.Result, specVersion string) error {
	for _, clientName := range clients {
		fmt.Printf("Generating '%s' client\n", clientName)

		schemas := specs["spec"].Get("components.schemas")

		preamble := preamble(clientName, specVersion)
		normalizedClientName := getNormalizedClientName(clientName)

		valRoutes := getRouteConstants(specs["routesTable"], "val-platform")
		LOL_TFT_Routes := getRouteConstants(specs["routesTable"], "platform")
		maps := getGenericConstants(specs["maps"], "Map")
		queues := getGenericConstants(specs["queues"], "Queue")
		gameTypes := getGenericConstants(specs["gameTypes"], "GameType")
		gameModes := getGenericConstants(specs["gameModes"], "GameMode")
		queueTypes := getGenericConstants(specs["queueTypes"], "QueueType")

		endpointGroups := getEndpointGroup(clientName, specs["spec"])
		endpointGroupsKeys := getMapKeys(endpointGroups)
		filteredEndpointGroups := filterEndpointGroup(endpointGroups, schemas)

		models := getAPIModels(filteredEndpointGroups, schemas.Map())
		endpoints := getAPIEndpoints(endpointGroups)

		ctx := pongo2.Context{
			"Split":                strings.Split,
			"FilterTFT":            filterTFT,
			"Preamble":             preamble,
			"ClientName":           clientName,
			"NormalizedClientName": normalizedClientName,
			"FormatEndpointName":   formatEndpointName,
			"RemoveGameName":       removeGameName,

			"VALRoutes":      valRoutes,
			"LOL_TFT_Routes": LOL_TFT_Routes,
			"Maps":           maps,
			"Queues":         queues,
			"GameTypes":      gameTypes,
			"GameModes":      gameModes,
			"QueueTypes":     queueTypes,

			"EndpointGroups":        endpointGroups,
			"EndpointGroupKeys":     endpointGroupsKeys,
			"FilteredEndpointGroup": filteredEndpointGroups,

			"Models":    models,
			"Endpoints": endpoints,
		}

		// Reading clients templates
		templates, err := readTemplateFiles("./templates/clients")
		if err != nil {
			return err
		}

		results := make(map[string][]byte, len(templates))

		// Compiling templates
		for filename, template := range templates {
			if filename == "constants" && (clientName == "riot" || clientName == "lor") {
				continue
			}

			tmpl, err := pongo2.FromBytes(template)
			if err != nil {
				return err
			}

			result, err := tmpl.ExecuteBytes(ctx)
			if err != nil {
				return err
			}
			results[filename] = result
		}

		// Writing results
		for filename, result := range results {
			err := os.WriteFile("../clients/"+clientName+"/"+filename+".go", result, 0644)
			if err != nil {
				return fmt.Errorf("error writing to file: %w", err)
			}
		}
	}

	return nil
}
