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
	Resource  string `query:"r" validate:"required,oneof=users tasks http_tasks http_auths http_results"`
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

func selectAllowedFiledsOf(resource string) (string, error) {
	config, ok := FetchConfig[resource]
	if !ok {
		return "", fmt.Errorf("invalid resource")
	}
	keys := make([]string, 0, len(config.Fields))
	for key := range config.Fields {
		keys = append(keys, config.Table+"."+key+" AS "+config.Table+db.DbTableColumnSeperator+key)
	}
	return strings.Join(keys, ", "), nil
}

func filtersFromParams(p *FetchParams) (string, error) {
	config, ok := FetchConfig[p.Resource]
	if !ok {
		return "", fmt.Errorf("invalid resource")
	}
	wheres := " WHERE"
	wheres_added := false
	if p.Id > 0 {
		wheres += " " + config.Table + ".id = " + strconv.Itoa(p.Id)
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
			// TODO: check if "f" has any uppercase letters first lol be smort!
			wheres += " LOWER(" + config.Table + "." + k + ") LIKE '" + f + "%'"
		} else {
			wheres += " " + config.Table + "." + k + " = '" + f + "'"
		}
		wheres_added = true
	}

	if wheres_added {
		return wheres, nil
	} else {
		return "", nil
	}
}

func sortsFromParams(p *FetchParams) (string, error) {
	config, ok := FetchConfig[p.Resource]
	if !ok {
		return "", fmt.Errorf("invalid resource")
	}
	sorts := " ORDER BY"
	sorts_added := false
	for k := range config.Fields {
		s, ok := p.M["s["+k+"]"]
		if !ok {
			continue
		}
		if sorts_added {
			sorts += ","
		}
		sorts += " " + config.Table + "." + k
		if s == "d" {
			sorts += " DESC"
		} else {
			sorts += " ASC"
		}
		sorts_added = true
	}

	if sorts_added {
		return sorts, nil
	} else {
		return "", nil
	}
}

func innerJoinsFromParams(p *FetchParams) (string, error) {
	config, ok := FetchConfig[p.Resource]
	if !ok {
		return "", fmt.Errorf("invalid resource")
	}
	innerJoins := ""
	innerJoins_added := false

	for k := range config.Fields {
		f, ok := p.M["ij["+k+"]"]
		if !ok {
			continue
		}
		table := strings.Split(f, ".")[0]
		_, ok = FetchConfig[table]
		if !ok {
			continue
		}
		innerJoins += " INNER JOIN " + table + " ON " + config.Table + "." + k + " = " + f
		innerJoins_added = true
	}

	if innerJoins_added {
		return innerJoins, nil
	} else {
		return "", nil
	}
}

func limitsFromParams(p *FetchParams) string {
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.PageCount < 1 {
		p.PageCount = 1
	}
	return " LIMIT " + strconv.Itoa(p.PageSize) + " OFFSET " + strconv.Itoa(p.PageSize*(p.PageCount-1))
}

func FetchQuery(p *FetchParams) (string, error) {
	q := "SELECT "

	if p.Count {
		q += "COUNT(*) AS count"
	} else {
		selectFields, err := selectAllowedFiledsOf(p.Resource)
		if err != nil {
			return "", err
		}
		q += selectFields
		for k, v := range p.M {
			if strings.HasPrefix(k, "ij[") {
				table := strings.Split(v, ".")[0]
				fields, err := selectAllowedFiledsOf(table)
				if err != nil {
					return "", err
				}
				q += ", " + fields
			}
		}
	}
	q += " FROM " + p.Resource

	innerJoins, err := innerJoinsFromParams(p)
	if err != nil {
		return "", err
	}
	q += innerJoins

	wheres, err := filtersFromParams(p)
	if err != nil {
		return "", err
	}
	q += wheres

	sorts, err := sortsFromParams(p)
	if err != nil {
		return "", err
	}
	q += sorts

	q += limitsFromParams(p) + ";"
	// fmt.Println(q, p)
	return q, nil
}

// func Fetch(p *FetchParams) ([]db.TypedTables, error) {
func Fetch(p *FetchParams) ([]map[string]map[string]interface{}, error) {
	q, err := FetchQuery(p)
	if err != nil {
		return nil, err
	}
	// var data []db.TypedTables
	// err = db.DB.Select(&data, q)
	// return data, err

	// return db.SelectGenericRowsAsMap(q)
	return db.SelectGenericRowsx(q)
}
