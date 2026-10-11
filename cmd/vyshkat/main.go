package main

import (
	"fmt"
	"os"

	"vyshkat/internal/app"
	"vyshkat/internal/console"
	"vyshkat/internal/di"
	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

func main() {
	c := app.BuildContainer(app.Options{})
	reception := di.MustResolve[*park.Reception](c)
	p := di.MustResolve[*park.Park](c)
	svc := di.MustResolve[*reports.Service](c)

	if err := console.Run(reception, p, svc, os.Stdin, os.Stdout, stdinIsTerminal()); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка чтения ввода:", err)
		os.Exit(1)
	}
}

func stdinIsTerminal() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}
