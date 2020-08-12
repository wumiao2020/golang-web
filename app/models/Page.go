package models

import "html/template"

type Page struct {
	Id          int     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Uuid        string    `json:"uuid" xorm:"not null default '' comment('UUID') unique VARCHAR(36)"`
	ParentId    int       `json:"parent_id" xorm:"not null default 0 comment('父级ID') TINYINT(4)"`
	Type        string    `json:"type" xorm:"not null default 'note' comment('文档类型') ENUM('doing','error','noneed','nonsupport','note','success','unknow','waiting')"`
	Status      int       `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Title       string    `json:"title" xorm:"comment('主题') index VARCHAR(255)"`
	Keywords    string    `json:"keywords" xorm:"comment('关键字') index VARCHAR(255)"`
	Description string    `json:"description" xorm:"comment('描述') VARCHAR(255)"`
	CreatedAt   string `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt   string `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	Identifier  string    `json:"identifier" xorm:"unique VARCHAR(36)"`
	Content   template.HTML    `json:"content" xorm:"comment('描述') TEXT"`
}