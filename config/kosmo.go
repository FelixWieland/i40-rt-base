package config

import "github.com/FelixWieland/kosmo"

//Service config for a kosmo graphql service
var Service = kosmo.Service{
	GraphQLConfig: kosmo.GraphQLConfig{
		RemoveResolverPrefixes: true,
		ResolverPrefixes:       []string{"Get", "Delete", "Add"},
	},
	HTTPConfig: kosmo.HTTPConfig{
		Playground: true,
		Port:       ":8080",
	},
}
