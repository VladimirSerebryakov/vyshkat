package main

import (
	"os"
	"vyshkat/internal/app"
	"vyshkat/internal/demo"
	"vyshkat/internal/di"
	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

func main() {
	c := app.BuildContainer(app.Options{}) // Option с полем nil - не подменный inspector
	reception := di.MustResolve[*park.Reception](c)
	p := di.MustResolve[*park.Park](c)
	svc := di.MustResolve[*reports.Service](c)
	demo.Run(reception, p, svc, os.Stdout)
}
