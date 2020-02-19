package main

import (
	"log"
	"net/http"

	"github.com/FelixWieland/i40-rt-base/config"
	"github.com/FelixWieland/i40-rt-base/graphql"
	"github.com/FelixWieland/i40-rt-base/server"

	"github.com/FelixWieland/kosmo"
)

func main() {

	dashboardsSchema := kosmo.Type(graphql.Dashboards{}).Queries(graphql.GetDashboards)
	dashboardSchema := kosmo.Type(graphql.Dashboard{}).Queries(graphql.GetDashboard)

	httpServer := config.Service.Schemas(dashboardsSchema, dashboardSchema).Server()

	mux := httpServer.Handler.(*http.ServeMux)
	mux.HandleFunc("/", server.ServeFrontend)

	log.Printf("\n%v", httpServer.ListenAndServe())
}
