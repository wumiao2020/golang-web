package dao

import (
	"../models"
	"github.com/go-xorm/xorm"
)

type Page struct {
	engine *xorm.Engine
}

func NewPage(engine *xorm.Engine) *Page {
	return &Page{
		engine:engine,
	}
}

func (d *Page) Get(string string) *models.Page {
	data := &models.Page{Identifier:string}
	ok, err := d.engine.Join("LEFT","page_data","page.id = page_data.page_id").Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (d *Page) GetByUuid(string string) *models.Page {
	data := &models.Page{Uuid:string}
	ok, err := d.engine.Join("LEFT","page_data","page.id = page_data.page_id").Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (d *Page) GetList(parentId int) []models.Page {
	datalist := make([]models.Page, 0)
	err := d.engine.Where("parent_id=?",parentId).Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *Page) GetAll() []models.Page {
	datalist := make([]models.Page, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (d *Page) Delete(id int) error {
	data := models.Page{Id:id,Status:0}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}
func (d *Page) Update(data *models.Page, column []string) error {
	_, err := d.engine.Id(data.Id).MustCols(column...).Update(data)
	return err
}
func (d *Page) Create(data *models.Page) error {
	_, err := d.engine.Insert(data)
	return err
}