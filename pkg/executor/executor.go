package executor

import (
	"fmt"
	"good-db/pkg/parser"
	"good-db/pkg/storage"
)

type Executor struct {
	storage storage.Storage
}

func NewExecutor(s storage.Storage) *Executor {
	return &Executor{
		storage: s,
	}
}

func (e *Executor) Execute(query *parser.Query) (interface{}, error) {
	switch query.Type {
	case parser.SELECT:
		return e.executeSelect(query)
	case parser.INSERT:
		return e.executeInsert(query)
	case parser.UPDATE:
		return e.executeUpdate(query)
	case parser.DELETE:
		return e.executeDelete(query)
	default:
		return nil, fmt.Errorf("unsupported query type")
	}
}

func (e *Executor) executeSelect(query *parser.Query) (interface{}, error) {
	return nil, fmt.Errorf("SELECT execution not implemented yet")
}

func (e *Executor) executeInsert(query *parser.Query) (interface{}, error) {
	return nil, fmt.Errorf("INSERT execution not implemented yet")
}

func (e *Executor) executeUpdate(query *parser.Query) (interface{}, error) {
	return nil, fmt.Errorf("UPDATE execution not implemented yet")
}

func (e *Executor) executeDelete(query *parser.Query) (interface{}, error) {
	return nil, fmt.Errorf("DELETE execution not implemented yet")
}