package inspection

import (
	"vyshkat/internal/domain"
)

type Inspector interface {
	Inspect(v domain.Vehicle) Verdict
}

type Verdict struct {
	Accepted bool
	Reason   string
}
