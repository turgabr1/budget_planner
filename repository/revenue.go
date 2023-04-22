package repository

var revenueTableName = "revenues"

type Revenue struct {
	*Entry
}

func (e *Revenue) Create() error {
	err := e.Entry.Create(revenueTableName)
	return err
}

func (e *Revenue) Update() error {
	err := e.Entry.Update(revenueTableName)
	return err
}

func (e *Revenue) Delete() error {
	err := e.Entry.Delete(revenueTableName)
	return err
}

func GetAllRevenues(budgetId int64) ([]*Revenue, error) {
	entries, _ := GetAllEntry(revenueTableName, budgetId)
	revenues := make([]*Revenue, 0)
	for _, entry := range entries {
		revenue := &Revenue{entry}
		revenues = append(revenues, revenue)
	}

	return revenues, nil
}

func GetRevenueByID(id int64) (*Revenue, error) {
	entry, err := GetByID(revenueTableName, id)
	if err != nil {
		return nil, err
	}
	revenue := &Revenue{entry}
	return revenue, nil
}
