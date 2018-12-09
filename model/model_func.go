package model


const (
	db         = "doc-manager"
)


type ModelOperation interface {
	FindAll()([]interface{},error)
	Insert()error
	Update(id string)error
	Remove(id string)error
	FindById(id string)(interface{}, error)
}
