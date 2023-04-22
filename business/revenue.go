package business

import "github.com/google/uuid"

type Revenue struct {
	id      uuid.UUID
	name    string
	monthly float64
	yearly  float64
}

func NewRevenue(name string, monthly float64, yearly float64) *Revenue {
	newRevenue := new(Revenue)
	newRevenue.id, _ = uuid.NewUUID()
	newRevenue.name = name
	if monthly == 0 {
		newRevenue.monthly = yearly / 12
		newRevenue.yearly = yearly
	} else {
		newRevenue.monthly = monthly
		newRevenue.yearly = monthly * 12
	}
	return newRevenue
}
