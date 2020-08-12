package services

import (
	"../../datasource"
	"../dao"
	"../models"
)

type Page interface {
	GetAll() []models.Page
	GetList(parentId int) []models.Page
	Get(string string) *models.Page
	GetByUuid(string string) *models.Page
	DeleteByID(id int) error
	Update(data *models.Page, columns []string)  error
	Create(data *models.Page) error
}

type page struct {
	dao *dao.Page
}

func PageService() Page {
	return &page{
		dao: dao.NewPage(datasource.InstanceMaster()),
	}
}

func (p page) GetAll() []models.Page {
	return p.dao.GetAll()
}
func (p page) GetList(parentId int) []models.Page {
	return p.dao.GetList(parentId)
}
func (p page) Get(string string) *models.Page {
	return p.dao.Get(string)
}
func (p page) GetByUuid(string string) *models.Page {
	return p.dao.GetByUuid(string)
}
func (p page) DeleteByID(id int) error {
	return p.dao.Delete(id)
}
func (p page) Update(data *models.Page, columns []string) error {
	return p.dao.Update(data, columns)
}
func (p page) Create(data *models.Page) error {
	return p.dao.Create(data)
}
