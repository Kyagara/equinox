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
		return fmt.Errorf("failed to download '%s'. Error: '%s'", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %s", err)
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %s", err)
		}

		err = os.WriteFile(filePath, body, 0644)
		if err != nil {
			return fmt.Errorf("error writing to file: %s", err)
		}

		fmt.Printf("Saving '%s'\n", filePath)
		return nil
	}

	return fmt.Errorf("failed to download '%s'. Status code: '%d'", url, response.StatusCode)
}

func Compile() error {
	specs, err := readJSONSpecsFiles()
	if err != nil {
		return err
	}

	specVersion := specs["spec"].Get("info.version").String()
	if specVersion == "" {
		return fmt.Errorf("spec version not found")
	}

	fmt.Printf("Spec version: %s\n", specVersion)

	fmt.Printf("Compiling API module\n")
	err = compileApi(specs, specVersion)
	if err != nil {
		return err
	}

	fmt.Printf("Compiling client modules\n")
	return compileClients(specs, specVersion)
}

func Format() error {
	fmt.Printf("Formatting imports\n")
	err := runCommand(exec.Command("goimports", "-w", "../"))
	if err != nil {
		return err
	}

	fmt.Printf("Formatting code\n")
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
		fmt.Printf("Failed to execute command: %s, Output: %s\n", err, out)
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
	for _, template := range templates {
		tmpl, err := pongo2.FromString(template[1])
		if err != nil {
			return err
		}

		result, err := tmpl.ExecuteBytes(ctx)
		if err != nil {
			return err
		}

		results[template[0]] = result
	}

	// Writing results
	for filename, result := range results {
		if err := os.WriteFile("../api/"+filename+".go", result, 0644); err != nil {
			return err
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
		gameTypes := getGenericConstants(specs["gameTypes"])
		gameModes := getGenericConstants(specs["gameModes"])
		queueTypes := getGenericConstants(specs["queueTypes"])

		endpointGroups := getEndpointGroup(clientName, specs["spec"])
		endpointGroupsKeys := getMapKeys(endpointGroups)
		filteredEndpointGroups := filterEndpointGroup(endpointGroups, schemas)

		models := getAPIModels(filteredEndpointGroups, schemas.Map())
		endpoints := getAPIEndpoints(endpointGroups)

		ctx := pongo2.Context{
			"Split":                strings.Split,
			"Preamble":             preamble,
			"ClientName":           clientName,
			"NormalizedClientName": normalizedClientName,
			"FormatEndpointName":   formatEndpointName,
			"RemoveGameName":       removeGameName,

			"VALRoutes":      valRoutes,
			"LOL_TFT_Routes": LOL_TFT_Routes,
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
		for _, template := range templates {
			if template[0] == "constants" && (clientName == "riot" || clientName == "lor") {
				continue
			}

			tmpl, err := pongo2.FromString(template[1])
			if err != nil {
				return err
			}

			result, err := tmpl.ExecuteBytes(ctx)
			if err != nil {
				return err
			}
			results[template[0]] = result
		}

		// Writing results
		for filename, result := range results {
			if err := os.WriteFile("../clients/"+clientName+"/"+filename+".go", result, 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
