package main

import (
	"slices"
	"strings"

	"github.com/tidwall/gjson"
)

type RouteConstant struct {
	Value            string
	TournamentRegion string
	Description      string
	Deprecated       bool
}

type GenericConstant struct {
	Value       string
	IsInteger   bool
	Description string
	Deprecated  bool
}

func getRouteConstants(routesTable gjson.Result, routeType string) map[string]RouteConstant {
	routes := make(map[string]RouteConstant, len(routesTable.Map()))

	for name, details := range routesTable.Get(routeType).Map() {
		description := normalizeDescription(details.Get("description").String())
		tournamentRegion := details.Get("tournamentRegion").String()
		value := name
		name = strings.ToUpper(name)
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
		description := normalizeDescription(item.Get("x-desc").String())
		name := strings.ToUpper(item.Get("x-name").String())
		deprecated := item.Get("x-deprecated").Bool()
		value := name
		isInteger := false
		if slices.Contains([]string{"GameMode", "QueueType", "GameType"}, constName) {
			if constName == "GameType" {
				value = strings.Replace(value, "_GAME", "", 1)
			}
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
