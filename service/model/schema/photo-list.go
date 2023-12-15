package schema

type PhotoList struct {
	LastDate string  `json:"lastDate"`
	LastID   string  `json:"lastID"`
	Records  []Photo `json:"records"`
}
