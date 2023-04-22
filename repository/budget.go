package repository

type Budget struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (e *Budget) CreateBudget() (int64, error) {
	db := OpenDatabase()
	result, err := db.Exec("INSERT INTO budget(name) VALUES (?)", e.Name)
	budgetId, _ := result.LastInsertId()
	return budgetId, err
}

func GetAllBudget() ([]*Budget, error) {
	db := OpenDatabase()
	rows, err := db.Query("SELECT id, name FROM budget")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	budgets := make([]*Budget, 0)
	for rows.Next() {
		budget := &Budget{}
		err := rows.Scan(&budget.ID, &budget.Name)
		if err != nil {
			return nil, err
		}
		budgets = append(budgets, budget)
	}

	return budgets, nil
}
