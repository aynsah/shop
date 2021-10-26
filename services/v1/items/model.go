package items

import (
	"fmt"
	"net/http"
	"shop/database"
	"shop/services/v1/taxes"
	"time"
)

type Item struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Taxes []taxes.Tax `json:"taxes"`
}

func ItemList(items *[]Item) (int, error) {
	query := `SELECT i.id, i.name, t.id, t.name, t.rate
			  FROM items i 
			  JOIN items_taxes it 
			  JOIN taxes t 
			  ON i.id = it.item_id 
			  AND t.id = it.tax_id
			  WHERE i.deleted_at is null
			  AND t.deleted_at is null`

	rows, err := database.DB.Query(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	var current_item Item
	var new_item Item

	for rows.Next() {
		var tax taxes.Tax
		err := rows.Scan(&new_item.ID, &new_item.Name, &tax.ID, &tax.Name, &tax.Rate)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// value awal
		if current_item.ID == 0 {
			current_item = new_item
		}

		if current_item.ID != new_item.ID {
			*items = append(*items, current_item)
			current_item = new_item
		}

		current_item.Taxes = append(current_item.Taxes, tax)

		current_item.ID = new_item.ID
	}

	*items = append(*items, current_item)

	return http.StatusOK, nil
}

func (i *Item) Save() (int, error) {
	query := `INSERT INTO items (name)
			  VALUES (?)`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(i.Name)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	last_id, _ := result.LastInsertId()
	i.ID = int(last_id)

	err = i.savePivotTable()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (i *Item) Update() (int, error) {
	query := `UPDATE items SET name = ?
			  where id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(i.Name, i.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = i.deletePivotTable()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = i.savePivotTable()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (i Item) Delete() (int, error) {
	query := `UPDATE items SET deleted_at = ?
			  where id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), i.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (i Item) savePivotTable() error {
	query := "INSERT INTO items_taxes (item_id, tax_id) VALUES "

	values := []interface{}{}

	for _, tax := range i.Taxes {
		values = append(values, i.ID)
		values = append(values, tax.ID)
		query += "(?, ?), "
	}

	query = query[:len(query)-2]
	fmt.Println(query)

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(values...)

	if err != nil {
		return err
	}

	return nil
}

func (i Item) deletePivotTable() error {
	query := `DELETE FROM items_taxes
			  where item_id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(i.ID)
	if err != nil {
		return err
	}

	return nil
}
