package db

import "time"

type HttpResult struct {
	ID        int           `json:"-"`
	TaskID    int           `json:"task_id" db:"task_id"`
	Code      int           `json:"code"`
	Ok        bool          `json:"ok"`
	Duration  time.Duration `json:"duration" db:"duration_ns"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
}

func createHttpResultTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS http_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		code INTEGER NOT NULL,
		ok BOOLEAN NOT NULL,
		duration_ns INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := DB.Exec(q)
	return err
}

func AddHttpResult(r *HttpResult) (int64, error) {
	q := `
	INSERT INTO http_results (
		task_id,
		code,
		ok,
		duration_ns
	) VALUES (?, ?, ?, ?)
	`
	res, err := DB.Exec(q, r.TaskID, r.Code, r.Ok, r.Duration)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func SelectAllHttpResultsByTaskID(tid int) ([]HttpResult, error) {
	q := `select * from http_results where task_id = ? ORDER BY created_at DESC LIMIT 10`
	var res []HttpResult
	err := DB.Select(&res, q, tid)
	return res, err
}
