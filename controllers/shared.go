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
	Count     bool   `query:"count" validate:"omitempty"`
	M         map[string]string
}

// Sorts      map[string]string `query:"s" validate:"omitempty"`
// Filters    map[string]string `query:"f" validate:"omitempty"`
// InnerJoins map[string]string `query:"ij" validate:"omitempty"`

type resourceConfig struct {
	Table  string
	Fields map[string]string
}

var FetchConfig = map[string]resourceConfig{
	"users": {
		Table: "users",
		Fields: map[string]string{
			"id":         "int",
			"name":       "string",
			"email":      "string",
			"role":       "string",
			"access":     "json",
			"created_at": "time",
		},
	},
	"tasks": {
		Table: "tasks",
		Fields: map[string]string{
			"id":          "int",
			"name":        "string",
			"active":      "bool",
			"repeat":      "bool",
			"interval":    "int",
			"description": "string",
			"type":        "string",
			"tags":        "json",
			"created_at":  "time",
		},
	},
	"http_tasks": {
		Table: "http_tasks",
		Fields: map[string]string{
			"id":                    "int",
			"task_id":               "int",
			"method":                "string",
			"url":                   "string",
			"body":                  "string",
			"headers":               "json",
			"encoding":              "string",
			"retries":               "int",
			"timeout":               "int",
			"accepted_status_codes": "json",
			"auth_method":           "string",
		},
	},
	"http_auths": {
		Table: "http_auths",
		Fields: map[string]string{
			"id":                  "int",
			"task_id":             "int",
			"username":            "string",
			"password":            "string",
			"oauth_method":        "string",
			"oauth_url":           "string",
			"oauth_client_id":     "string",
			"oauth_client_secret": "string",
			"oauth_scope":         "string",
		},
	},
	"http_results": {
		Table: "http_results",
		Fields: map[string]string{
			"id":          "int",
			"task_id":     "int",
			"code":        "int",
			"ok":          "bool",
			"duration_ns": "int",
			"created_at":  "time",
		},
	},
}

func Fetch(p *FetchParams) ([]map[string]interface{}, error) {
	q := "SELECT "
	config, ok := FetchConfig[p.Resource]
	if !ok {
		// TODO: create custom httperror
		return nil, fmt.Errorf("invalid resource")
	}
	if p.Count {
		q += "COUNT(*) AS count"
	} else {
		keys := make([]string, 0, len(config.Fields))
		for key := range config.Fields {
			keys = append(keys, config.Table+"."+key)
		}
		selectFields := strings.Join(keys, ", ")
		q += selectFields
	}
	q += " FROM " + p.Resource
	wheres := " WHERE"
	wheres_added := false
	if p.Id > 0 {
		wheres += " id = " + strconv.Itoa(p.Id)
		wheres_added = true
	}
	for k, v := range config.Fields {
		f, ok := p.M["f["+k+"]"]
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
		if v == "string" {
			wheres += " " + k + " LIKE '" + f + "%'"
		} else {
			wheres += " " + k + " = '" + f + "'"
		}
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
