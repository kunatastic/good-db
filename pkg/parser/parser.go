package parser

import (
	"fmt"
	"strings"
)

type QueryType int

const (
	SELECT QueryType = iota
	INSERT
	UPDATE
	DELETE
	CREATE_TABLE
	DROP_TABLE
)

type Query struct {
	Type   QueryType
	Table  string
	Fields []string
	Values []interface{}
	Where  map[string]interface{}
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(sql string) (*Query, error) {
	sql = strings.TrimSpace(sql)
	if sql == "" {
		return nil, fmt.Errorf("empty query")
	}

	parts := strings.Fields(strings.ToUpper(sql))
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid query")
	}

	query := &Query{}

	switch parts[0] {
	case "SELECT":
		query.Type = SELECT
		return p.parseSelect(sql, query)
	case "INSERT":
		query.Type = INSERT
		return p.parseInsert(sql, query)
	case "UPDATE":
		query.Type = UPDATE
		return p.parseUpdate(sql, query)
	case "DELETE":
		query.Type = DELETE
		return p.parseDelete(sql, query)
	default:
		return nil, fmt.Errorf("unsupported query type: %s", parts[0])
	}
}

func (p *Parser) parseSelect(sql string, query *Query) (*Query, error) {
	return query, fmt.Errorf("SELECT parsing not implemented yet")
}

func (p *Parser) parseInsert(sql string, query *Query) (*Query, error) {
	return query, fmt.Errorf("INSERT parsing not implemented yet")
}

func (p *Parser) parseUpdate(sql string, query *Query) (*Query, error) {
	return query, fmt.Errorf("UPDATE parsing not implemented yet")
}

func (p *Parser) parseDelete(sql string, query *Query) (*Query, error) {
	return query, fmt.Errorf("DELETE parsing not implemented yet")
}