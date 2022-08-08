package entities

type Team struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	City           string `json:"city"`
	State          string `json:"state"`
	Stadium        string `json:"stadium"`
	LeagueDivision string `json:"division"`
}
