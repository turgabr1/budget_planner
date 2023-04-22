package repository

var tableName = "expenses"

type Expense struct {
	*Entry
}

func (e *Expense) Create() error {
	err := e.Entry.Create(tableName)
	return err
}

func (e *Expense) Update() error {
	err := e.Entry.Update(tableName)
	return err
}

func (e *Expense) Delete() error {
	err := e.Entry.Delete(tableName)
	return err
}

func GetAllExpenses(budgetId int64) ([]*Expense, error) {
	entries, _ := GetAllEntry(tableName, budgetId)
	expenses := make([]*Expense, 0)
	for _, entry := range entries {
		expense := &Expense{entry}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func GetExpenseByID(id int64) (*Expense, error) {
	entry, err := GetByID(tableName, id)
	if err != nil {
		return nil, err
	}
	expense := &Expense{entry}
	return expense, nil
}
