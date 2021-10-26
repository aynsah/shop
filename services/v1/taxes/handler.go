package taxes

import (
	"net/http"
	"shop/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StoreHandler(c *gin.Context) {
	name := c.PostForm("name")
	rate_str := c.PostForm("rate")

	if name == "" {
		utils.ShowErr(c, http.StatusBadRequest, "Tax name is required", nil)
		return
	}

	if rate_str == "" {
		utils.ShowErr(c, http.StatusBadRequest, "Tax rate is required", nil)
		return
	}

	rate, _ := strconv.ParseFloat(rate_str, 32)

	tax := Tax{
		Name: name,
		Rate: rate,
	}

	status, err := tax.Save()
	if err != nil {
		utils.ShowErr(c, status, "could not save the data", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UpdateHandler(c *gin.Context) {
	id_str := c.Param("id")
	name := c.PostForm("name")
	rate_str := c.PostForm("rate")

	if name == "" {
		utils.ShowErr(c, http.StatusBadRequest, "Tax name is required", nil)
		return
	}

	if rate_str == "" {
		utils.ShowErr(c, http.StatusBadRequest, "Tax rate is required", nil)
		return
	}

	id, _ := strconv.Atoi(id_str)
	rate, _ := strconv.ParseFloat(rate_str, 32)

	tax := Tax{
		ID:   id,
		Name: name,
		Rate: rate,
	}

	status, err := tax.Update()
	if err != nil {
		utils.ShowErr(c, status, "could not update the data", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func DeleteHandler(c *gin.Context) {
	id_str := c.Param("id")

	id, _ := strconv.Atoi(id_str)

	tax := Tax{
		ID: id,
	}

	status, err := tax.Delete()
	if err != nil {
		utils.ShowErr(c, status, "could not delete the data", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
