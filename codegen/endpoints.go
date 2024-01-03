package main

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/tidwall/gjson"
)

var methodNamesMapping = map[string]string{
	"CurrentGameInfoBySummoner":   "CurrentGameBySummonerID",
	"ChampionMasteryScoreByPUUID": "MasteryScoreByPUUID",
	"AllChampionMasteriesByPUUID": "AllMasteriesByPUUID",
	"ChampionMasteryByPUUID":      "MasteryByPUUID",
	"TopChampionMasteriesByPUUID": "TopMasteriesByPUUID",
	"AllChampionMasteries":        "AllMasteriesBySummonerID",
	"ChampionMastery":             "MasteryBySummonerID",
	"TopChampionMasteries":        "TopMasteriesBySummonerID",
	"ChampionMasteryScore":        "ScoreBySummonerID",
	"PlayersByPUUID":              "SummonerEntriesByPUUID",
	"PlayersBySummoner":           "SummonerEntriesBySummonerID",
	"FeaturedGames":               "Featured",
	"ShardData":                   "Shard",
	"TeamByID":                    "TeamByTeamID",
	"ChampionInfo":                "Rotation",
	"BySummonerName":              "ByName",
	"MatchIdsByPUUID":             "ListByPUUID",
	"Matchlist":                   "ListByPUUID",
	"Configs":                     "ConfigByID",
	"PlayerData":                  "ByPUUID",
	"PlatformData":                "Platform",
	"EntriesForSummoner":          "SummonerEntries",
	"Challenger":                  "ChallengerByQueue",
	"Grandmaster":                 "GrandmasterByQueue",
	"Master":                      "MasterByQueue",
	"TournamentByID":              "ByID",
	"TournamentByTeam":            "ByTeamID",
	"Match":                       "ByID",
}

type EndpointGroup struct {
	Path   string
	Groups gjson.Result
}

type Methods struct {
	Name              string
	Arguments         string
	MethodReturnTuple string
	OperationID       string
	Route             string
	HTTPMethod        string
	NilValue          string
	ReturnType        string
	ValueReturn       string
	ErrorReturn       string
	URLPath           string
	Body              string
	Description       []string
	Headers           []string
	HeaderChecks      []string
	Queries           []string
	HasReturn         bool
}

func getAPIEndpoints(endpointGroup map[string][]EndpointGroup) map[string][]Methods {
	apiEndpoints := make(map[string][]Methods)

	for endpointKey, groups := range endpointGroup {
		structName := strcase.ToCamel(endpointKey)
		version := structName[len(structName)-2:]
		key := removeGameName(structName) + "|" + endpointKey

		apiEndpoints[key] = make([]Methods, 0)

		for _, group := range groups {
			route := group.Path
			for verb, operation := range group.Groups.Map() {
				if strings.HasPrefix(verb, "x-") {
					continue
				}

				operationID := operation.Get("operationId").String()

				methodName := getMethodName(operationID)
				resp200 := operation.Get("responses.200")

				var returnType string
				if resp200.Exists() && resp200.Get("content").Exists() {
					returnType = getReturnType(resp200, version, endpointKey)
				}
				hasReturn := returnType != ""

				var bodyType string
				if operation.Get("requestBody").Exists() {
					bodyType = getBodyType(operation, version, endpointKey)
				}

				normalizedRoute := getNormalizedRoute(operation)
				argBuilder := []string{
					"ctx context.Context",
					"route " + normalizedRoute + "Route",
				}

				body := "nil"
				if bodyType != "" {
					argBuilder = append(argBuilder, "body *"+bodyType)
					body = "body"
				}

				allParams := operation.Get("parameters")
				var queryParams []gjson.Result
				var headerParams []gjson.Result
				routeArgument := ""
				if allParams.Exists() {
					pathParams := getSortedParams(allParams, "path", route)
					queryParams = getParams(allParams, "query")
					headerParams = getParams(allParams, "header")

					for _, paramList := range [][]gjson.Result{pathParams, queryParams, headerParams} {
						for _, param := range paramList {
							argBuilder = append(argBuilder,
								normalizePropName(param.Get("name").String())+
									" "+
									stringifyType(param.Get("schema")))
						}
					}

					routeArgument = formatRouteArgument(pathParams, route)
				} else {
					routeArgument = formatRouteArgument([]gjson.Result{}, route)
				}

				isPrimitiveType := returnType != "" && slices.Contains(goTypes, returnType)
				nilValue := getNilValue(returnType)

				descArr := strings.Split(
					strings.ReplaceAll(
						operation.Get("description").String(),
						"## Implementation Notes",
						"\n# Implementation Notes\n",
					),
					"\n",
				)

				descArr = append(descArr,
					"",
					"# Parameters",
					"   - `route` - Route to query.",
				)

				if len(allParams.Array()) > 0 {
					for _, param := range allParams.Array() {
						requiredStr := ""
						required := param.Get("required").Bool()
						if !required {
							requiredStr = " (optional)"
						}

						desc := param.Get("description").String()
						if desc != "" {
							requiredStr += " -"
						}

						descArr = append(descArr,
							fmt.Sprintf("   - `%s`%s %s", param.Get("name").String(), requiredStr, desc),
						)
					}
				}

				descArr = append(descArr,
					"",
					"# Riot API Reference",
					"",
					fmt.Sprintf("[%s]", operationID),
					"",
					fmt.Sprintf("[%s]: %s", operationID, operation.Get("externalDocs.url")),
				)

				methodReturnTuple := "error"
				if hasReturn {
					star := ""
					if !isPrimitiveType {
						star = "*"
					}
					methodReturnTuple = fmt.Sprintf("(%s%s, error)", star, returnType)
				}

				if strings.HasPrefix(methodReturnTuple, "(*[]") || strings.HasPrefix(methodReturnTuple, "(*map") {
					methodReturnTuple = strings.Replace(methodReturnTuple, "*", "", 1)
					isPrimitiveType = true
				}

				errorReturn := getMethodErrReturn(hasReturn, isPrimitiveType, nilValue)
				valueReturn := getMethodValueReturn(hasReturn, isPrimitiveType)

				apiEndpoints[key] = append(apiEndpoints[key], Methods{
					Name:              methodName,
					Arguments:         strings.Join(argBuilder, ", "),
					MethodReturnTuple: methodReturnTuple,
					Description:       descArr,
					OperationID:       operationID,
					HTTPMethod:        strcase.ToCamel(verb),
					Route:             route,
					NilValue:          nilValue,
					URLPath:           routeArgument,
					Body:              body,
					HasReturn:         hasReturn,
					ReturnType:        returnType,
					ValueReturn:       valueReturn,
					ErrorReturn:       errorReturn,
					Queries:           formatAddQueryParam(queryParams),
					Headers:           formatAddHeaderParam(headerParams),
					HeaderChecks:      formatAddHeadersChecks(headerParams, nilValue),
				})
			}
		}

		sort.Slice(apiEndpoints[key], func(i, j int) bool {
			return apiEndpoints[key][i].Name < apiEndpoints[key][j].Name
		})
	}

	return apiEndpoints
}
func getSortedParams(allParams gjson.Result, paramType string, route string) []gjson.Result {
	params := make([]gjson.Result, 0)

	for _, param := range allParams.Array() {
		if param.Get("in").String() == paramType {
			params = append(params, param)
		}
	}

	sort.Slice(params, func(i, j int) bool {
		return strings.Index(route, params[i].Get("name").String()) < strings.Index(route, params[j].Get("name").String())
	})

	return params
}

func getParams(allParams gjson.Result, paramType string) []gjson.Result {
	params := make([]gjson.Result, 0)

	for _, param := range allParams.Array() {
		if param.Get("in").String() == paramType {
			params = append(params, param)
		}
	}

	return params
}
func formatRouteArgument(pathParams []gjson.Result, route string) string {
	if len(pathParams) == 0 {
		return fmt.Sprintf("\"%s\"", route)
	}

	args := make([]string, len(pathParams))
	for i, param := range pathParams {
		name := param.Get("name").String()
		if name == "Authorization" {
			name = "authorization"
		}
		if name == "type" {
			name = "type_"
		}
		args[i] = name
	}

	formattedRoute := regexp.MustCompile(`\{(\S+?)\}`).ReplaceAllString(route, "ARG$")

	counter := 1
	result := regexp.MustCompile(`ARG\$`).ReplaceAllStringFunc(formattedRoute, func(match string) string {
		if counter <= len(args) {
			newValue := "ARG" + strconv.Itoa(counter) + "$"
			counter++
			return newValue
		}
		return match
	})

	re := regexp.MustCompile(`ARG(\d+)\$`)

	counter = 0
	result = re.ReplaceAllStringFunc(result, func(match string) string {
		indexStr := re.FindStringSubmatch(match)[1]

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return match
		}

		if index >= 1 && index <= len(args) {
			newValue := args[index-1]
			schema := pathParams[index-1].Get("schema")
			paramType := schema.Get("type").String()
			if newValue == "championId" {
				counter++
				newValue = "strconv.FormatInt(int64(championId), 10)"
				return fmt.Sprintf("\", %s, \"", newValue)
			}

			if schema.Get("enum").Exists() && schema.Get("x-enum").Exists() {
				counter++
				return fmt.Sprintf("\", %s.String(), \"", newValue)
			}

			if paramType == "integer" {
				format := schema.Get("format").String()
				if format == "int64" {
					newValue = fmt.Sprintf("strconv.FormatInt(%v, 10)", newValue)
				} else {
					newValue = fmt.Sprintf("strconv.FormatInt(int64(%v), 10)", newValue)
				}
			}

			counter++
			return fmt.Sprintf("\", %s, \"", newValue)
		}
		return match
	})

	result = fmt.Sprintf("\"%v\"", result)
	if strings.HasSuffix(result, ", \"\"") {
		result = result[:len(result)-3]
	}
	return result
}

func formatAddHeadersChecks(params []gjson.Result, nilValue string) []string {
	checks := make([]string, 0, len(params))
	for _, param := range params {
		name := normalizePropName(param.Get("name").String())
		checks = append(checks, fmt.Sprintf(
			`if %s == "" {
    return %s, fmt.Errorf("'%s' header is required")
}`, name, nilValue, name))
	}
	return checks
}

func formatAddHeaderParam(params []gjson.Result) []string {
	headers := make([]string, 0, len(params))
	for _, param := range params {
		name := normalizePropName(param.Get("name").String())
		headerName := name
		if headerName == "authorization" {
			headerName = "Authorization"
		}
		headers = append(headers, fmt.Sprintf(`equinoxReq.Request.Header.Set("%s", %s)`, headerName, name))
	}
	return headers
}

func formatAddQueryParam(params []gjson.Result) []string {
	queries := make([]string, 0, len(params))
	for _, param := range params {
		name := normalizePropName(param.Get("name").String())
		prop := param.Get("schema")
		letHeaderName := name
		if letHeaderName == "type_" {
			letHeaderName = "type"
		}

		condition := ""
		propType := prop.Get("type").String()
		if propType == "string" {
			condition = fmt.Sprintf(`%s != ""`, name)
		} else if propType == "integer" {
			condition = fmt.Sprintf(`%s != -1`, name)
		} else {
			panic(fmt.Errorf("unknown prop type: %s", propType))
		}

		conversion := ""
		end := ""
		format := prop.Get("format").String()
		if format == "int32" {
			conversion = "strconv.FormatInt(int64("
			end = "), 10)"
		} else if format == "int64" {
			conversion = "strconv.FormatInt("
			end = ", 10)"
		} else {
			conversion = "fmt.Sprint("
			end = ")"
		}

		value := name
		if prop.Get("type").String() != "string" {
			value = fmt.Sprintf(`%s%s%s`, conversion, name, end)
		}

		queries = append(queries, fmt.Sprintf(`if %s {
    values.Set("%s", %s)
}`, condition, letHeaderName, value))
	}
	return queries
}

func normalizePropName(propName string) string {
	out := propName
	if digitRegex.MatchString(out) {
		out = "X" + out
	}
	if out == "Authorization" {
		return "authorization"
	}
	if out == "type" {
		return out + "_"
	}
	return out
}

func getReturnType(resp200 gjson.Result, version string, endpointID string) string {
	jsonInfo := resp200.Get("content.application/json")
	returnType, _ := normalizeDTOName(
		stringifyType(jsonInfo.Get("schema")),
		version,
		endpointID,
	)

	if strings.HasPrefix(endpointID, "league-exp") && strings.Contains(returnType, "LeagueEntry") {
		returnType = strings.Replace(returnType, "LeagueEntry", "ExpLeagueEntry", -1)
	}

	return cleanIfPrimitiveType(jsonInfo, version, endpointID, returnType)
}

func getBodyType(operation gjson.Result, version string, endpointID string) string {
	jsonInfo := operation.Get("requestBody.content.application/json")
	body, _ := normalizeDTOName(
		stringifyType(jsonInfo.Get("schema")),
		version,
		endpointID,
	)
	return body
}

func getNormalizedRoute(operation gjson.Result) string {
	route := strcase.ToCamel(operation.Get("x-route-enum").String())
	route = strings.Replace(route, "Regional", "api.Regional", -1)
	return strings.Replace(route, "ValPlatform", "Platform", -1)
}

func getNilValue(returnType string) string {
	nilValues := map[string]string{
		"int32":   "0",
		"int64":   "0",
		"float32": "0",
		"float64": "0",
	}
	if returnType == "string" {
		return `""`
	}
	if returnType == "bool" {
		return `false`
	}
	val := nilValues[returnType]
	if val == "" {
		return `nil`
	}
	return val
}

func getMethodErrReturn(hasReturn bool, isPrimitiveType bool, nilValue string) string {
	if hasReturn {
		if !isPrimitiveType {
			return "nil, err"
		} else {
			return fmt.Sprintf("%s, err", nilValue)
		}
	} else {
		return "err"
	}
}

func getMethodValueReturn(hasReturn bool, isPrimitiveType bool) string {
	if hasReturn {
		if !isPrimitiveType {
			return "&data, nil"
		} else {
			return "data, nil"
		}
	} else {
		return "nil"
	}
}

func getMethodName(operationID string) string {
	dotIndex := strings.Index(operationID, ".")
	method := strcase.ToCamel(operationID[dotIndex+1:])
	temp := regexp.MustCompile("^Get").ReplaceAllString(method, "")
	temp = regexp.MustCompile("League").ReplaceAllString(temp, "")
	temp = regexp.MustCompile("Id$").ReplaceAllString(temp, "ID")
	temp = regexp.MustCompile("Puuid$").ReplaceAllString(temp, "PUUID")
	temp = regexp.MustCompile("Rsopuuid$").ReplaceAllString(temp, "RSOPUUID")
	if name, ok := methodNamesMapping[temp]; ok {
		return name
	}
	return temp
}
