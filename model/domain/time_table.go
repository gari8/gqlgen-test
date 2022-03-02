package domain

type TimeTable struct {
	ID    string  `json:"id"`
	Todos []*Todo `json:"todos"`
	Start string  `json:"start"`
	End   string  `json:"end"`
}
