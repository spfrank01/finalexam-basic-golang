package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spfrank01/finalexam-basic-golang/database"
	"net/http"
	"strconv"
)

//Update todo
func (cus Customer) Update(conn *sql.DB) error {

	stmt, err := conn.Prepare("UPDATE todos SET name=$2, status=$3 WHERE id=$1;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cus.ID, cus.Name, cus.Status)

	return err
}

//UpdateHandler gin api
func UpdateHandler(c *gin.Context) {

	t := Customer{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = id

	err = database.Update(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}
