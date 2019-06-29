package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spfrank01/finalexam-basic-golang/database"
	"net/http"
	"strconv"
)

//Delete todo
func (cus Customer) Delete(conn *sql.DB) error {

	stmt, err := conn.Prepare("DELETE FROM customers WHERE id=$1;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cus.ID)
	return err
}

//DeleteByIDHandler gin api
func DeleteByIDHandler(c *gin.Context) {

	t := Customer{}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = id

	err = database.Delete(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
