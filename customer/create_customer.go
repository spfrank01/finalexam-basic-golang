package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/spfrank01/finalexam-basic-golang/database"
)

func (cus Customer) Insert(conn *sql.DB) (database.DataLayer, error) {

	row := conn.QueryRow("INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id",
		cus.Name, cus.Email, cus.Status)
	err := row.Scan(&cus.ID)
	return database.IConv(cus), err
}

func CreateHandler(c *gin.Context) {
	cusReq := Customer{}
	if err := c.ShouldBind(&cusReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cusRes := database.Insert(cusReq)
}
