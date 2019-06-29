package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spfrank01/finalexam-basic-golang/database"
	"net/http"
	"strconv"
)

//GetAll get all todos
func (cus Customer) GetAll(conn *sql.DB) ([]database.DataLayer, error) {

	tt := []database.DataLayer{}
	rows, err := conn.Query("SELECT id, name, status FROM customers")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cusRes Customer
		if err := rows.Scan(&cusRes.ID, &cusRes.Name, &cusRes.Status); err != nil {
			return nil, err
		}
		tt = append(tt, database.IConv(cusRes))
	}

	return tt, err
}

//GetHandler gin api
func GetHandler(c *gin.Context) {

	t := Customer{}
	conn, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	tt, err := t.GetAll(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tt)
}

//GetByKey get todo by key
func (cus Customer) GetByKey(conn *sql.DB) (database.DataLayer, error) {

	row := conn.QueryRow("SELECT id, name, status FROM customers where id = $1", cus.ID)
	err := row.Scan(&cus.ID, &cus.Name, &cus.Status)
	if err != nil {
		return cus, err
	}
	return database.IConv(cus), err
}

//GetByIDHandler for retrive Todo by ID
func GetByIDHandler(c *gin.Context) {

	t := Customer{}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = id

	t2, err := database.GetByKey(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t2)
}
