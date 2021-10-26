package taxes

import (
	"net/http"
)

type Tax struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

func (t Tax) Save() (int, error) {

	return http.StatusOK, nil
}

func (t *Tax) Update() (int, error) {

	return http.StatusOK, nil
}

func (t Tax) Delete() (int, error) {

	return http.StatusOK, nil
}
