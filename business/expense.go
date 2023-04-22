package business

import "github.com/google/uuid"

type Expense struct {
	id      uuid.UUID
	name    string
	monthly float64
	yearly  float64
}

func NewExpense(name string, monthly float64, yearly float64) *Expense {
	newExpense := new(Expense)
	newExpense.id, _ = uuid.NewUUID()
	newExpense.name = name
	if monthly == 0 {
		newExpense.monthly = yearly / 12
		newExpense.yearly = yearly
	} else {
		newExpense.monthly = monthly
		newExpense.yearly = monthly * 12
	}
	return newExpense
}
