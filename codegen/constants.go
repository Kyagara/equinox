package main

import "github.com/tidwall/gjson"

type RouteConstant struct {
	Name             string
	Description      string
	TournamentRegion string
	QueueType        string
	Deprecated       bool
}

type GenericConstant struct {
	Name        string
	Description string
	Deprecated  bool
}

func getRouteConstants(routesTable gjson.Result, routeType string) map[string]RouteConstant {
	routes := make(map[string]RouteConstant, len(routesTable.Map()))

	for name, details := range routesTable.Get(routeType).Map() {
		description := details.Get("description").String()
		deprecated := details.Get("deprecated").Bool()
		tournamentRegion := details.Get("tournamentRegion").String()
		queueType := details.Get("x-name").String()

		routes[name] = RouteConstant{
			Name:             name,
			Description:      description,
			Deprecated:       deprecated,
			TournamentRegion: tournamentRegion,
			QueueType:        queueType,
		}
	}

	return routes
}

func getGenericConstants(table gjson.Result) map[string]GenericConstant {
	consts := make(map[string]GenericConstant, len(table.Map()))

	for _, item := range table.Array() {
		desc := item.Get("x-desc").String()
		name := item.Get("x-name").String()
		deprecated := item.Get("x-deprecated").Bool()

		consts[name] = GenericConstant{
			Name:        name,
			Description: desc,
			Deprecated:  deprecated,
		}
	}

	return consts
}
