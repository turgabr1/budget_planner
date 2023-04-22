package repository

import "fmt"

type IEntry interface {
	Create() error
	Update() error
	Delete() error
}

type Entry struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Amount     float64    `json:"amount"`
	Occurrence Occurrence `json:"occurrence"`
	BudgetId   int64      `json:"budget_id"`
}

func (e *Entry) Create(tableName string) error {
	db := OpenDatabase()
	queryString := fmt.Sprintf("INSERT INTO %s(name, amount, occurrence, budget_id) VALUES (?, ?, ?, ?)", tableName)
	_, err := db.Exec(queryString, e.Name, e.Amount, e.Occurrence, e.BudgetId)
	return err
}

func (e *Entry) Update(tableName string) error {
	db := OpenDatabase()
	queryString := fmt.Sprintf("UPDATE %s SET name=?, amount=?, occurrence=?, budget_id=? WHERE id=?", tableName)
	_, err := db.Exec(queryString, e.Name, e.Amount, e.Occurrence, e.BudgetId, e.ID)
	return err
}

func (e *Entry) Delete(tableName string) error {
	db := OpenDatabase()
	queryString := fmt.Sprintf("DELETE FROM %s WHERE id=?", tableName)
	_, err := db.Exec(queryString, e.ID)
	return err
}

func GetAllEntry(tableName string, budgetId int64) ([]*Entry, error) {
	db := OpenDatabase()
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE budget_id=?", tableName)
	rows, err := db.Query(queryString, budgetId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entrys := make([]*Entry, 0)
	for rows.Next() {
		entry := &Entry{}
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Amount, &entry.Occurrence, &entry.BudgetId)
		if err != nil {
			return nil, err
		}
		entrys = append(entrys, entry)
	}

	return entrys, nil
}

func GetByID(tableName string, id int64) (*Entry, error) {
	db := OpenDatabase()
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	row := db.QueryRow(queryString, id)
	entry := &Entry{}
	err := row.Scan(&entry.ID, &entry.Name, &entry.Amount, &entry.Occurrence, &entry.BudgetId)
	if err != nil {
		return nil, err
	}
	return entry, nil
}
