package park

import (
	"errors"
	"vyshkat/internal/domain"
	"vyshkat/internal/inspection"
)

type Reception struct {
	inspector inspection.Inspector
	park      *Park
}

func NewReception(inspector inspection.Inspector, park *Park) (*Reception, error) {
	if inspector == nil {
		return nil, errors.New("inspector must not be nil")
	}
	if park == nil {
		return nil, errors.New("park must not be nil")
	}
	return &Reception{inspector: inspector, park: park}, nil
}

// Register — приёмка транспорта через техосмотр. Для вещей Park.AddItem напрямую
func (r *Reception) Register(v domain.Vehicle) (inspection.Verdict, error) {
	verdict := r.inspector.Inspect(v)
	if verdict.Accepted {
		err := r.park.addVehicle(v)
		if err != nil {
			return verdict, err
		}
	}
	return verdict, nil
}
