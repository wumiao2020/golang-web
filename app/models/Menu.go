package models

type Menu struct {
	Id          int     `json:"id" xorm:"pk autoincr 'id'"`
	Title       string    `json:"title" xorm:"comment('主题') index VARCHAR(255)"`
	Identifier  string    `json:"identifier" xorm:"unique VARCHAR(36)"`
}