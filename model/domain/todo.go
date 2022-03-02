package domain

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

//type Todo struct {
//	ID          string `json:"id" gorm:"primaryKey"`
//	Text        string `json:"text"`
//	Done        bool   `json:"done"`
//	UserID      string `json:"user"`
//	TimeTableID string `json:"timeTable"`
//}
