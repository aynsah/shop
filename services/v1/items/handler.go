package items

import (
	"net/http"
	"shop/services/v1/taxes"
	"shop/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ItemsHandler(c *gin.Context) {
	var items []Item

	status, err := ItemList(&items)
	if err != nil {
		utils.ShowErr(c, status, "could not retrieve data", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    items,
	})
}

func StoreHandler(c *gin.Context) {
	name := c.PostForm("name")
	taxes_id := c.PostFormArray("taxes")

	if len(taxes_id) < 2 {
		utils.ShowErr(c, http.StatusBadRequest, "Items are required to have at least 2 taxes", nil)
		return
	}

	if name == "" {
		utils.ShowErr(c, http.StatusBadRequest, "Item name is required", nil)
		return
	}

	item := Item{
		Name: name,
	}

	for _, tax_id := range taxes_id {
		tax_id_str, _ := strconv.Atoi(tax_id)
		item.Taxes = append(item.Taxes, taxes.Tax{
			ID: tax_id_str,
		})
	}

	status, err := item.Save()
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
	taxes_id := c.PostFormArray("taxes")

	if name == "" {
		utils.ShowErr(c, http.StatusBadRequest, "Item name is required", nil)
		return
	}

	id, _ := strconv.Atoi(id_str)

	item := Item{
		ID:   id,
		Name: name,
	}

	for _, tax_id := range taxes_id {
		tax_id_str, _ := strconv.Atoi(tax_id)
		item.Taxes = append(item.Taxes, taxes.Tax{
			ID: tax_id_str,
		})
	}

	status, err := item.Update()
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

	item := Item{
		ID: id,
	}

	status, err := item.Delete()
	if err != nil {
		utils.ShowErr(c, status, "could not delete the data", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
