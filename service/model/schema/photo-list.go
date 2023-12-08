package schema

type PhotoList struct {
	LastDate string  `json:"lastDate"`
	LastId   string  `json:"lastId"`
	Records  []Photo `json:"records"`
}
