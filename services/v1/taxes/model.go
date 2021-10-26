package taxes

import (
	"net/http"
	"shop/database"
	"time"
)

type Tax struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

func (t Tax) Save() (int, error) {
	query := `INSERT INTO taxes (name, rate)
			  VALUES (?, ?)`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(t.Name, t.Rate)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *Tax) Update() (int, error) {
	query := `UPDATE taxes SET name = ?, rate = ?
			  where id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(t.Name, t.Rate, t.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t Tax) Delete() (int, error) {
	query := `UPDATE taxes SET deleted_at = ?
			  where id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), t.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
