package main

import (
	"fmt"
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

	for _, dtos := range filteredEndpointGroups {
		for _, rawDTO := range dtos {
			dtoSplit := strings.Split(rawDTO, ".")
			dto, version := getDTOAndVersion(rawDTO)

			schema := schema[rawDTO]
			schemaDescription := normalizeDescription(schema.Get("description").String())

			properties := schema.Get("properties")
			props := make([]ModelProperty, 0, len(properties.Map()))
			for propKey, prop := range properties.Map() {
				name, fieldType := getModelField(prop, propKey, version, dtoSplit[0])
				description := normalizeDescription(prop.Get("description").String())

				props = append(props, ModelProperty{
					Name:        name,
					Type:        fieldType,
					Description: description,
					JSONField:   fmt.Sprintf("`json:\"%s,omitempty\"`", propKey),
				})
			}

			apiModels[dto] = Model{
				Description: schemaDescription,
				DTO:         dtoSplit[1],
				Props:       props,
			}
		}
	}

	return apiModels
}

func normalizeDescription(desc string) string {
	if desc == "" {
		return ""
	}
	lines := strings.Split(desc, "\n")
	var trimmedLines []string
	for _, line := range lines {
		trimmedLines = append(trimmedLines, strings.TrimSpace(line))
	}
	return strings.Join(trimmedLines, "\r\n    //\n    // ")
}

func getDTOAndVersion(rawDTO string) (string, string) {
	parts := strings.Split(rawDTO, ".")
	endpoint := parts[0]
	dto := parts[1]
	v := strcase.ToCamel(versionRegex.FindString(endpoint))
	name, version := normalizeDTOName(dto, v, endpoint)
	return name, version
}

func getModelField(prop gjson.Result, propKey string, version string, endpoint string) (string, string) {
	propType := stringifyType(prop)
	propType = cleanIfPrimitiveType(prop, version, endpoint, propType)

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
	return name, propType
}
