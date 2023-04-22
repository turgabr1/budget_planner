package business

import "github.com/google/uuid"

type Calculation interface {
	calculate() string
}

type Action interface {
	addRevenue()
	addExpense()
}

func (b *Budget) AddRevenue(r *Revenue) {
	b.revenues = append(b.revenues, r)
	b.totalYearlyRevenues += r.yearly
	b.TotalMonthlyRevenues += r.monthly
}

func (b *Budget) AddExpense(e *Expense) {
	b.Expenses = append(b.Expenses, e)
	b.totalYearlyExpenses += e.yearly
	b.TotalMonthlyExpenses += e.monthly
}

type Budget struct {
	id                   uuid.UUID
	revenues             []*Revenue
	Expenses             []*Expense
	TotalMonthlyRevenues float64
	totalYearlyRevenues  float64
	TotalMonthlyExpenses float64
	totalYearlyExpenses  float64
}

func NewBudget() *Budget {
	newBudget := new(Budget)
	newBudget.id, _ = uuid.NewUUID()
	newBudget.revenues = []*Revenue{}
	newBudget.Expenses = []*Expense{}
	newBudget.TotalMonthlyRevenues = 0
	newBudget.totalYearlyRevenues = 0
	newBudget.TotalMonthlyExpenses = 0
	newBudget.totalYearlyExpenses = 0
	return newBudget
}
