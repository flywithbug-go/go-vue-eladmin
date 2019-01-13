package model_menu

import (
	"encoding/json"
	"fmt"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/model_menu_role"
	"vue-admin/web_server/model/model_role"
	"vue-admin/web_server/model/mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	menuCollection = mongo_index.CollectionMenu
	//MenuMenuAll    = "MENU_ALL"
	MenuPermissionSelect = "MENU_SELECT"
	MenuPermissionCreate = "MENU_CREATE"
	MenuPermissionEdit   = "MENU_EDIT"
	MenuPermissionDelete = "MENU_DELETE"
)

type meta struct {
	Title      string `json:"title,omitempty" bson:"title,omitempty"`
	Icon       string `json:"icon,omitempty" bson:"icon,omitempty"`
	NoCache    bool   `json:"noCache,omitempty" bson:"noCache,omitempty"`
	Breadcrumb bool   `json:"breadcrumb,omitempty" bson:"breadcrumb,omitempty"`
}

type Menu struct {
	Id         int64             `json:"id,omitempty" bson:"_id,omitempty"`
	PId        int64             `json:"pid,omitempty" bson:"pid,omitempty"` //父节点ID
	Sort       int               `json:"sort,omitempty" bson:"sort,omitempty"`
	Icon       string            `json:"icon,omitempty" bson:"icon,omitempty"`
	Name       string            `json:"name,omitempty" bson:"name,omitempty"`
	Label      string            `json:"label,omitempty" bson:"label,omitempty"`
	Path       string            `json:"path,omitempty" bson:"path,omitempty"`
	AlwaysShow bool              `json:"always_show" bson:"always_show"`
	Component  string            `json:"component,omitempty" bson:"component,omitempty"`
	IFrame     string            `json:"iFrame,omitempty" bson:"iFrame,omitempty"`
	CreateTime int64             `json:"createTime,omitempty" bson:"createTime,omitempty"`
	Children   []Menu            `json:"children,omitempty" bson:"children,omitempty"`
	Roles      []model_role.Role `json:"roles,omitempty" bson:"roles,omitempty"`
	Meta       meta              `json:"meta,omitempty"`
}

func (m Menu) ToJson() string {
	js, _ := json.Marshal(m)
	return string(js)
}

func (m Menu) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), menuCollection, query)
}

func (m Menu) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), menuCollection, docs...)
}

func (m Menu) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), menuCollection, selector, update, true)
}

func (m Menu) findOne(query, selector interface{}) (Menu, error) {
	ap := Menu{}
	err := mongo.FindOne(shareDB.DBName(), menuCollection, query, selector, &ap)
	return ap, err
}
func (m Menu) findAll(query, selector interface{}) (results []Menu, err error) {
	results = []Menu{}
	err = mongo.FindAll(shareDB.DBName(), menuCollection, query, selector, &results)
	return results, err
}

func (m Menu) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), menuCollection, selector)
}

func (m Menu) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), menuCollection, selector)
}

func (m Menu) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), menuCollection, query, selector)
}

func (m Menu) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Menu, err error) {
	results = []Menu{}
	err = mongo.FindPage(shareDB.DBName(), menuCollection, page, limit, query, selector, &results, fields...)
	return
}

func (m Menu) pipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeAll(shareDB.DBName(), menuCollection, pipeline, result, allowDiskUse)
}

func (m Menu) pipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeOne(shareDB.DBName(), menuCollection, pipeline, result, allowDiskUse)
}

func (m Menu) explain(pipeline, result interface{}) (results []Menu, err error) {
	err = mongo.Explain(shareDB.DBName(), menuCollection, pipeline, result)
	return
}

func (m Menu) Exist() bool {
	return m.isExist(bson.M{"_id": m.Id})
}

func (m Menu) Insert() (int64, error) {
	if m.PId != 0 && !m.isExist(bson.M{"_id": m.PId}) {
		return -1, fmt.Errorf("pid  not exist")
	}
	list := m.Children
	m.Id, _ = mongo.GetIncrementId(menuCollection)
	m.CreateTime = time.Now().Unix() * 1000
	m.Children = nil
	err := m.insert(m)
	if err != nil {
		return -1, err
	}
	m.Children = list
	m.updateMenuRole()
	return m.Id, nil
}

func (m Menu) updateMenuRole() {
	mr := model_menu_role.MenuRole{}
	mr.RemoveMenuId(m.Id)
	for _, role := range m.Roles {
		if role.Exist() {
			mr.RoleId = role.Id
			mr.MenuId = m.Id
			mr.Insert()
		}
	}
}

func (m Menu) Update() error {
	m.updateMenuRole()
	m.Roles = nil
	selector := bson.M{"_id": m.Id}
	return m.update(selector, m)
}

func (m Menu) Remove() error {
	mr := model_menu_role.MenuRole{}
	mr.RemoveMenuId(m.Id)
	return m.remove(bson.M{"_id": m.Id})
}

func (m Menu) FindAll(query, selector interface{}) (results []Menu, err error) {
	return m.findAll(query, selector)
}

func (m Menu) TotalCount(query, selector interface{}) (int, error) {
	return m.totalCount(query, selector)
}

func (m Menu) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]Menu, error) {
	results, err := m.findPage(page, limit, query, selector, fields...)
	if err != nil {
		return nil, err
	}
	makeTreeList(results, selector)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (m Menu) FindPageTreeFilter(page, limit int, query, selector interface{}, fields ...string) ([]Menu, error) {
	results, err := m.findPage(page, limit, query, selector, fields...)
	if err != nil {
		return nil, err
	}
	makeTreeList(results, selector)
	return results, err
}

func (m Menu) FetchTreeList(selector interface{}) (results []Menu, err error) {
	results, err = m.findAll(bson.M{"pid": 0}, selector)
	if err != nil {
		return
	}
	makeTreeList(results, selector)
	return
}

func (m Menu) FindOneTree() (menu Menu, err error) {
	menu, err = m.findOne(bson.M{"_id": m.Id}, nil)
	if err != nil {
		return
	}
	list := []Menu{menu}
	makeTreeList(list, nil)
	return list[0], nil
}

func (m *Menu) findChildren(selector interface{}) error {
	results, err := m.findAll(bson.M{"pid": m.Id}, selector)
	if err != nil {
		return err
	}
	m.Children = results
	return nil
}

func makeTreeList(list []Menu, selector interface{}) {
	for index := range list {
		err := list[index].findChildren(selector)
		if err != nil {
			return
		}
		if selector == nil {
			list[index].Meta = meta{
				Title: list[index].Name,
				Icon:  list[index].Icon,
			}
		} else {
			list[index].Label = list[index].Name
			list[index].Name = ""
		}
		makeTreeList(list[index].Children, selector)
	}
}
