package services

import (
	"wumiao/app/dao"
	"wumiao/app/models"
	"wumiao/datasource"
)

type Menu interface {
	GetMenu() []models.Menu
}

type menu struct {
	dao *dao.Menu
}

func MenuService() Menu {
	return &menu{
		dao: dao.NewMenu(datasource.InstanceMaster()),
	}
}

func (m menu) GetMenu() []models.Menu {
	return m.dao.GetMenu()
}
