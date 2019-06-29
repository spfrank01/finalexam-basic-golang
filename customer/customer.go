package customer

type Customer struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

/*
Update(conn *sql.DB) error
	Delete(conn *sql.DB) error
	GetByKey(conn *sql.DB) (DataLayer, error)
	GetAll(conn *sql.DB) ([]DataLayer, error)
*/
/*
func (cus Customer) Update(conn *sql.DB) error {
	return nil
}

func (cus Customer) Delete(conn *sql.DB) error {
	return nil
}

func (cus Customer) GetByKey(conn *sql.DB) (database.DataLayer, error) {
	return nil, nil
}
func (cus Customer) GetAll(conn *sql.DB) ([]database.DataLayer, error) {
	return nil, nil
}
*/
