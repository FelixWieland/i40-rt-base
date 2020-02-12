package main

import (
	"i40rtbase/config"
	. "i40rtbase/graphql"
	"log"

	"github.com/FelixWieland/kosmo"
)

func main() {

	dashboardsSchema := kosmo.Type(Dashboards{}).Queries(GetDashboards)
	dashboardSchema := kosmo.Type(Dashboard{}).Queries(GetDashboard)

	server := config.Service.Schemas(dashboardsSchema, dashboardSchema).Server()

	log.Printf("\n%v", server.ListenAndServe())
}
