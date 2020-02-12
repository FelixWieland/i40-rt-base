package main

import (
	"i40rtbase/config"
	. "i40rtbase/graphql"
	"i40rtbase/server"
	"log"
	"net/http"

	"github.com/FelixWieland/kosmo"
)

func main() {

	dashboardsSchema := kosmo.Type(Dashboards{}).Queries(GetDashboards)
	dashboardSchema := kosmo.Type(Dashboard{}).Queries(GetDashboard)

	httpServer := config.Service.Schemas(dashboardsSchema, dashboardSchema).Server()

	mux := httpServer.Handler.(*http.ServeMux)
	mux.HandleFunc("/", server.ServeFrontend)

	log.Printf("\n%v", httpServer.ListenAndServe())
}
