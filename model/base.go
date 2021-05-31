package model

type Service interface {
	Create(request interface{}) (interface{}, error)
	Update(request interface{}) (interface{}, error)
	GetById(id float64) (interface{}, error)
	GetByKey(keys interface{}) (interface{}, error)
	Delete(request interface{}) error
	List(request interface{}) (interface{}, error)
}
