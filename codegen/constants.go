package main

import (
	"slices"
	"sort"
	"strings"

	"github.com/tidwall/gjson"
)

type Endpoint struct {
	Method string
	Path   string
	ID     string
}

type RouteConstant struct {
	Value            string
	TournamentRegion string
	Description      string
	Deprecated       bool
}

type GenericConstant struct {
	Value       string
	Description string
	IsInteger   bool
	Deprecated  bool
}

func getRouteConstants(routesTable gjson.Result, routeType string) map[string]RouteConstant {
	routes := make(map[string]RouteConstant, len(routesTable.Map()))

	for name, details := range routesTable.Get(routeType).Map() {
		value := name
		name = strings.ToUpper(name)
		description := normalizeDescription(details.Get("description").String())
		tournamentRegion := details.Get("tournamentRegion").String()
		deprecated := details.Get("deprecated").Bool()

		routes[name] = RouteConstant{
			Value:            value,
			TournamentRegion: tournamentRegion,
			Description:      description,
			Deprecated:       deprecated,
		}
	}

	return routes
}

func getGenericConstants(table gjson.Result, constName string) map[string]GenericConstant {
	consts := make(map[string]GenericConstant, len(table.Map()))

	for _, item := range table.Array() {
		value := item.Get("x-name").String()
		name := strings.ToUpper(value)
		description := normalizeDescription(item.Get("x-desc").String())
		deprecated := item.Get("x-deprecated").Bool()
		isInteger := false

		if slices.Contains([]string{"GameMode", "QueueType", "GameType"}, constName) {
			isInteger = false
		} else {
			value = item.Get("x-value").String()
			isInteger = true
		}

		name += "_" + strings.ToUpper(constName)
		name = strings.Replace(name, "_DEPRECATED", "", 1)

		consts[name] = GenericConstant{
			Value:       value,
			IsInteger:   isInteger,
			Description: description,
			Deprecated:  deprecated,
		}
	}

	return consts
}

func getAllEndpoints(paths gjson.Result) []Endpoint {
	endpoints := make([]Endpoint, 0, len(paths.Map()))
	for path, groups := range paths.Map() {
		for verb, method := range groups.Map() {
			if strings.HasPrefix(verb, "x-") {
				continue
			}

			verb = strings.ToUpper(verb)
			methodID := method.Get("operationId").String()

			endpoints = append(endpoints, Endpoint{
				Method: verb,
				Path:   path,
				ID:     methodID,
			})
		}
	}

	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].ID < endpoints[j].ID
	})

	return endpoints
}
