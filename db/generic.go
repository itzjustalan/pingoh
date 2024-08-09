package db

func SelectGenericRows(q string) ([]map[string]interface{}, error) {
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
