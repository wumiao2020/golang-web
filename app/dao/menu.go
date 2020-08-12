package dao

import (
	"../models"
	"github.com/go-xorm/xorm"
)

type Menu struct {
	engine *xorm.Engine
}

func NewMenu(engine *xorm.Engine) *Menu {
	return &Menu{
		engine:engine,
	}
}

func (d *Menu) GetMenu() []models.Menu {
	datalist := make([]models.Menu, 0)
	err := d.engine.Asc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}