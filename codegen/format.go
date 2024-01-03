package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/tidwall/gjson"
)

var (
	goTypes = []string{
		"bool",
		"string",
		"int32",
		"int64",
		"float32",
		"float64",
	}

	digitRegex   = regexp.MustCompile(`^\d`)
	versionRegex = regexp.MustCompile(`v.*\d`)
	clientRegex  = regexp.MustCompile("(?i)(lor|riot|val|lol|tft)")
)

func preamble(packageName string, version string) string {
	return fmt.Sprintf(`package %s

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = %v`, packageName, version)
}

func getEndpointGroup(clientName string, spec gjson.Result) map[string][]EndpointGroup {
	endpoints := make(map[string][]EndpointGroup)

	paths := spec.Get("paths").Map()
	for path, endpoint := range paths {
		parts := strings.Split(path, "/")
		if parts[1] == clientName || (parts[1] == "fulfillment" && clientName == "lol") {
			endpointName := endpoint.Get("x-endpoint").String()
			endpoints[endpointName] = append(endpoints[endpointName], EndpointGroup{path, endpoint})
		}
	}

	return endpoints
}

func filterEndpointGroup(endpointGroup map[string][]EndpointGroup, schemas gjson.Result) map[string][]string {
	schemaKeyByEndpoint := make(map[string][]string)

	apis := make([]string, 0, len(endpointGroup))
	for key := range endpointGroup {
		if key != "Error" {
			apis = append(apis, key)
		}
	}

	for schemaKey := range schemas.Map() {
		if schemaKey != "Error" {
			parts := strings.Split(schemaKey, ".")
			groupKey := parts[0]
			if slices.Contains(apis, groupKey) {
				schemaKeyByEndpoint[groupKey] = append(schemaKeyByEndpoint[groupKey], schemaKey)
			}
		}
	}

	return schemaKeyByEndpoint
}

func removeGameName(str string) string {
	return clientRegex.ReplaceAllString(str, "")
}

func getNormalizedClientName(clientName string) string {
	if clientName == "riot" {
		return "Riot"
	} else {
		return strings.ToUpper(clientName)
	}
}

func formatEndpointName(endpointName string) string {
	var name strings.Builder

	for _, item := range strings.Split(endpointName, "-") {
		if clientRegex.MatchString(item) {
			continue
		}
		name.WriteString(strcase.ToCamel(item))
	}

	return name.String()
}

func normalizeDTOName(dto string, version string, endpoint string) (string, string) {
	if len(version) > 2 {
		version = version[len(version)-2:]
	}

	temp := dto
	temp = strings.ReplaceAll(temp, "Dto", "")
	temp = strings.ReplaceAll(temp, "DTO", "")

	if strings.Contains(temp, version+"Wrapper") {
		temp = strings.Replace(temp, version+"Wrapper", "Wrapper", 1)
	}

	if strings.Contains(endpoint, "tournament") && strings.Contains(endpoint, "stub") {
		if strings.HasPrefix(temp, "Tournament") || strings.HasPrefix(temp, "LobbyEvent") || strings.HasPrefix(temp, "Provider") {
			temp = "Stub" + temp
		}
	}

	temp = strings.ReplaceAll(temp, version, "")

	if strings.HasPrefix(endpoint, "league-exp") && (strings.HasPrefix(temp, "League") || strings.HasPrefix(temp, "Mini")) {
		temp = "Exp" + temp
	}
	if strings.HasPrefix(endpoint, "val-ranked") && strings.HasPrefix(temp, "Player") {
		temp = strings.Replace(temp, "Player", "LeaderboardPlayer", 1)
	}
	if strings.HasPrefix(endpoint, "lor-ranked") && strings.HasPrefix(temp, "Player") {
		temp = "Leaderboard" + temp
	}
	if strings.HasPrefix(endpoint, "val-status") && strings.HasPrefix(temp, "Content") {
		temp = "Status" + temp
	}

	temp += version + "DTO"
	temp = strings.Replace(temp, "ChampionInfoV", "ChampionRotationV", 1)
	temp = removeGameName(temp)
	for _, v := range goTypes {
		if strings.HasSuffix(temp, v+version+"DTO") {
			temp = strings.Replace(temp, version+"DTO", "", 1)
			break
		}
	}

	return temp, version
}

func stringifyType(prop gjson.Result) string {
	if prop.Get("anyOf").Exists() {
		prop = prop.Get("anyOf").Array()[0]
	}

	enumType := prop.Get("x-enum").String()
	if enumType != "" && enumType != "locale" {
		propType := prop.Get("x-type").String()
		if enumType == "champion" {
			if prop.Get("format").String() == "" {
				return "int64"
			}
			return prop.Get("format").String()
		}
		if propType == "" {
			enumType = strcase.ToCamel(enumType)
		}
		if propType == "string" {
			return strcase.ToCamel(enumType)
		}
		return prop.Get("format").String()
	}

	refType := prop.Get("$ref").String()
	if refType != "" {
		return refType[strings.LastIndex(refType, ".")+1:]
	}

	propType := prop.Get("type").String()
	if propType != "" {
		switch propType {
		case "boolean":
			return "bool"
		case "integer":
			if prop.Get("format").String() == "int32" {
				return "int32"
			}
			return "int64"
		case "number":
			if prop.Get("format").String() == "float" {
				return "float32"
			}
			return "float64"
		case "array":
			return "[]" + stringifyType(prop.Get("items"))
		case "string":
			return "string"
		case "object":
			keyType := stringifyType(prop.Get("x-key"))
			valueType := stringifyType(prop.Get("additionalProperties"))
			return "map[" + keyType + "]" + valueType
		default:
			return propType
		}
	}

	return ""
}

func cleanIfPrimitiveType(prop gjson.Result, version, endpoint, t string) string {
	if !slices.Contains(goTypes, t) && !prop.Get("x-enum").Exists() {
		t, _ = normalizeDTOName(t, version, endpoint)

		for _, pType := range goTypes {
			if strings.HasSuffix(t, pType+version+"DTO") {
				t = strings.Replace(t, version+"DTO", "", 1)
				break
			}
		}
	}

	return t
}