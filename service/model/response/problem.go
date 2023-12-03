package response

type Problem struct {
	Type     string
	Title    string
	Status   int
	Detail   string
	Instance string
}
