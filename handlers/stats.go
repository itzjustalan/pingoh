package handlers

import "pingoh/db"

func HttpResultsByTaskID(tid int) ([]db.HttpResult, error) {
	res, err := db.SelectAllHttpResultsByTaskID(tid)
	return res, err
}
