package schema

type PhotoList struct {
	LastDate string  `json:"lastDate"`
	Records  []Photo `json:"records"`
}
