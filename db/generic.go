package db

import (
	"strings"
)

const DbTableColumnSeperator = "____"

func SelectGenericRowsx(q string) ([]map[string]map[string]interface{}, error) {
	var res []map[string]map[string]interface{}

	rows, err := DB.Queryx(q)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		data := make(map[string]map[string]interface{})
		results := make(map[string]interface{})
		err := rows.MapScan(results)
		if err != nil {
			return res, err
		}
		for k, v := range results {
			table := strings.Split(k, DbTableColumnSeperator)[0]
			column := strings.Split(k, DbTableColumnSeperator)[1]
			if _, ok := data[table]; !ok {
				data[table] = make(map[string]interface{})
			}
			data[table][column] = v
		}
		res = append(res, data)
	}
	return res, err
}

func SelectGenericRowsAsMap(q string) ([]map[string]interface{}, error) {
	var res []map[string]interface{}

	rows, err := DB.Queryx(q)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		results := make(map[string]interface{})
		err := rows.MapScan(results)
		if err != nil {
			return res, err
		}
		res = append(res, results)
	}
	return res, err
}

func SelectGenericRowsCustom(q string) ([]map[string]interface{}, error) {
	var res []map[string]interface{}

	rows, err := DB.Query(q)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return res, err
	}

	for rows.Next() {

		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = &values[i]
		}

		err := rows.Scan(values...)
		if err != nil {
			return res, err
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = values[i]
		}

		res = append(res, rowMap)
	}
	return res, err
}
