package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/tidwall/gjson"
)

type Model struct {
	Description string
	DTO         string
	Props       []ModelProperty
}

type ModelProperty struct {
	Name        string
	Type        string
	Description string
	JSONField   string
}

func getAPIModels(filteredEndpointGroups map[string][]string, schema map[string]gjson.Result) map[string]Model {
	apiModels := make(map[string]Model, len(filteredEndpointGroups))
	dtoCount := 0

	for _, dtos := range filteredEndpointGroups {
		for _, rawDTO := range dtos {
			dtoCount++
			dtoSplit := strings.Split(rawDTO, ".")
			dto, version := getDTOAndVersion(rawDTO)

			if _, ok := apiModels[dto]; ok {
				panic(fmt.Errorf("duplicate data object, needs to be renamed in models & format files: %s", dto))
			}

			schema := schema[rawDTO]
			schemaDescription := normalizeDescription(schema.Get("description").String())

			properties := schema.Get("properties").Map()
			var sortedKeys []string
			for key := range properties {
				sortedKeys = append(sortedKeys, key)
			}
			sort.Strings(sortedKeys)

			props := make([]ModelProperty, 0, len(properties))
			namesUsed := make([]string, 0, len(properties))

			for _, propKey := range sortedKeys {
				prop := properties[propKey]
				name, fieldType := getModelField(prop, propKey, version, dtoSplit[0])
				description := normalizeDescription(prop.Get("description").String())

				if slices.Contains(namesUsed, name) {
					name += "_"
					fmt.Printf("Duplicate property name for property '%s' in '%s'. Renamed: '%s'\n", propKey, rawDTO, name)

					if slices.Contains(namesUsed, name) {
						panic(fmt.Errorf("duplicate property name '%s' for property '%s'", name, propKey))
					}
				}
				namesUsed = append(namesUsed, name)

				props = append(props, ModelProperty{
					Name:        name,
					Type:        fieldType,
					Description: description,
					JSONField:   fmt.Sprintf("`json:\"%s,omitempty\"`", propKey),
				})
			}

			if len(props) != len(properties) {
				panic(fmt.Errorf("unexpected props length for %s: %d != %d", rawDTO, len(properties), len(props)))
			}

			sort.Slice(props, func(i, j int) bool {
				return props[i].Name < props[j].Name
			})

			apiModels[dto] = Model{
				Description: schemaDescription,
				DTO:         rawDTO,
				Props:       props,
			}
		}
	}

	if len(apiModels) != dtoCount {
		panic(fmt.Errorf("unexpected amount of data objects: %d != %d", dtoCount, len(apiModels)))
	}

	return apiModels
}

func normalizeDescription(desc string) string {
	if desc == "" {
		return ""
	}
	lines := strings.Split(desc, "\n")

	trimmedLines := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmedLines = append(trimmedLines, strings.TrimSpace(line))
	}

	return strings.Join(trimmedLines, "\r\n    //\n    // ")
}

func getDTOAndVersion(rawDTO string) (string, string) {
	parts := strings.Split(rawDTO, ".")
	endpoint := parts[0]
	name := parts[1]

	version := strcase.ToCamel(versionRegex.FindString(endpoint))
	if len(version) > 2 {
		version = version[len(version)-2:]
	}

	name = strings.ReplaceAll(name, "Dto", "")
	name = strings.ReplaceAll(name, "DTO", "")
	name = strings.ReplaceAll(name, version, "")
	name = strings.Replace(name, "ChampionInfo", "ChampionRotation", 1)
	name += version + "DTO"
	for _, v := range goTypes {
		if strings.HasSuffix(name, v+version+"DTO") {
			name = strings.Replace(name, version+"DTO", "", 1)
			break
		}
	}

	endpoint = clientRegex.ReplaceAllString(endpoint, "")
	endpoint = endpoint[:len(endpoint)-3]
	endpoint = strcase.ToCamel(endpoint)

	if endpoint == "TournamentStub" && strings.HasPrefix(name, "Tournament") {
		name = strings.Replace(name, "Tournament", "", 1)
	}

	if !strings.HasPrefix(name, endpoint) {
		name = endpoint + name
	}
	return name, version
}

func getModelField(prop gjson.Result, propKey string, version string, endpoint string) (string, string) {
	propType := stringifyType(prop)
	propType = cleanDTOPropType(prop, version, endpoint, propType)

	name := propKey
	if digitRegex.MatchString(propKey) {
		name = "X" + propKey
	}

	if name == "x" {
		return "X", propType
	} else if name == "y" {
		return "Y", propType
	}

	name = strcase.ToCamel(strings.ReplaceAll(name, "-", ""))
	switch name {
	case "TimeCcingOthers":
		return "TimeCCingOthers", propType
	case "TakedownsFirstXminutes":
		return "TakedownsFirstXMinutes", propType
	case "WardTakedownsBefore20":
		return "WardTakedownsBefore20M", propType
	case "Puuid":
		return "PUUID", propType
	case "Xp":
		return "XP", propType
	case "Id":
		return "ID", propType
	case "Lp":
		return "LP", propType
	case "Url":
		return "URL", propType
	}

	if strings.HasPrefix(name, "RiotId") {
		name = strings.Replace(name, "RiotId", "RiotID", 1)
	}
	if strings.HasSuffix(name, "Id") {
		name = name[:len(name)-2] + "ID"
	}

	name = strings.Replace(name, "Ids", "IDs", 1)
	name = strings.Replace(name, "Tft", "TFT", 1)

	return name, propType
}
