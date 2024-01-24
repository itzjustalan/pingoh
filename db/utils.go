package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func jsonSqlValuer[T any](v T) (driver.Value, error) {
	jstr, err := json.Marshal(v)
	if err != nil {
		return driver.Value(""), err
	}
	return driver.Value(string(jstr)), nil
}

// use fn[T ~[]string | ~[]int](v *T) -> for pointers
func jsonSqlScanner[T any](v *T, src interface{}) error {
	switch src := src.(type) {
	case string:
		err := json.Unmarshal([]byte(src), v)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("incompatible type for UserAccess: %T", src)
	}
	return nil
}
