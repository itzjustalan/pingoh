package controllers

import (
	"fmt"
	"pingoh/db"
	"pingoh/services"
	"strconv"
	"strings"
)

// r: resource,
// i: id,
// l: page_size = 10,
// c: page_count = 1,
// s: sorts,
// f: filters,
// ij: innerJoins,
type FetchParams struct {
	Resource  string `query:"r" validate:"required,oneof=tasks"`
	Id        int    `query:"i" vlaidate:"omitempty,gte=1"`
	PageSize  int    `query:"l" validate:"omitempty,gte=1"`
	PageCount int    `query:"c" validate:"omitempty,gte=1"`
	M         map[string]string
}

// Sorts      map[string]string `query:"s" validate:"omitempty"`
// Filters    map[string]string `query:"f" validate:"omitempty"`
// InnerJoins map[string]string `query:"ij" validate:"omitempty"`

type resourceConfig struct {
	Table  string
	Fields []string
}

var FetchConfig = map[string]resourceConfig{
	"tasks": {
		Table:  "tasks",
		Fields: []string{"id", "name", "active", "repeat", "interval", "description", "type", "tags", "created_at"},
	},
}

func Fetch(p *FetchParams) ([]map[string]interface{}, error) {
	q := "SELECT "
	config, ok := FetchConfig[p.Resource]
	if !ok {
		// TODO: create custom httperror
		return nil, fmt.Errorf("invalid resource")
	}
	selectFields := strings.Join(config.Fields, ", ")
	q += selectFields
	q += " FROM " + p.Resource
	wheres := " WHERE"
	wheres_added := false
	if p.Id > 0 {
		wheres += " id = " + strconv.Itoa(p.Id)
		wheres_added = true
	}
	for _, v := range config.Fields {
		f, ok := p.M["f["+v+"]"]
		if !ok {
			continue
		}
		if wheres_added {
			wheres += " AND"
		}
		if !services.IsAlphanumeric(f) {
			// TODO: create custom httperror
			continue
		}
		wheres += " " + v + " = '" + f + "'"
		wheres_added = true
	}
	if wheres_added {
		q += wheres
	}
	sorts := " ORDER BY"
	sorts_added := false
	for _, v := range config.Fields {
		s, ok := p.M["s["+v+"]"]
		if !ok {
			continue
		}
		if sorts_added {
			sorts += ","
		}
		sorts += " " + v
		if s == "d" {
			sorts += " DESC"
		} else {
			sorts += " ASC"
		}
		sorts_added = true
	}
	if sorts_added {
		q += sorts
	}
	if p.PageSize > 0 {
		q += " LIMIT " + strconv.Itoa(p.PageSize)
	} else {
		q += " LIMIT 10"
	}
	if p.PageCount > 0 {
		q += " OFFSET " + strconv.Itoa(p.PageSize*(p.PageCount-1))
	} else {
		q += " OFFSET 0"
	}
	q += ";"
	fmt.Println(q, p)
	return db.SelectGenericRows(q)
}
