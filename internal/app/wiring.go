package app

import (
	"vyshkat/internal/di"
	"vyshkat/internal/inspection"
	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

type Options struct {
	Inspector inspection.Inspector // если nil - используется реальный, иначе подменили
}

// BuildContainer - строим DI контейнер, указываем зависимости только тут
func BuildContainer(opts Options) *di.Container {
	c := di.New()

	di.Register(c, func(c *di.Container) *park.Park {
		return park.NewPark()
	})

	if opts.Inspector != nil {
		di.RegisterInstance[inspection.Inspector](c, opts.Inspector)
	} else {
		di.Register(c, func(c *di.Container) inspection.Inspector {
			return inspection.NewServiceCenter(2.0)
		})
	}

	di.Register(c, func(c *di.Container) *park.Reception {
		r, err := park.NewReception(
			di.MustResolve[inspection.Inspector](c),
			di.MustResolve[*park.Park](c),
		)
		if err != nil {
			panic(err)
		}
		return r
	})

	di.Register(c, func(c *di.Container) *reports.Service {
		svc, err := reports.NewService(di.MustResolve[*park.Park](c))
		if err != nil {
			panic(err)
		}
		return svc
	})

	return c
}
