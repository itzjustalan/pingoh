package db

import (
	"database/sql/driver"
	"time"
)

type Task struct {
	ID int `json:"-"`
	// UID         string    `json:"uid"`
	Name        string    `json:"name"`
	Active      bool      `json:"active"`
	Repeat      bool      `json:"repeat"`
	Interval    int       `json:"interval"`
	Description string    `json:"description"`
	Type        TaskType  `json:"type"`
	Tags        TaskTags  `json:"tags"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type HttpTask struct {
	ID                  int                  `json:"-"`
	TaskID              int                  `json:"task_id" db:"task_id"`
	Method              HttpTaskMethod       `json:"method"`
	Url                 string               `json:"url"`
	Body                string               `json:"body"`
	Headers             HttpTaskHeaders      `json:"headers"`
	Encoding            HttpTaskBodyEncoding `json:"encoding"`
	Retries             int                  `json:"retries"`
	Timeout             int                  `json:"timeout"`
	AcceptedStatusCodes HttpTaskStatusCodes  `json:"accepted_status_codes" db:"accepted_status_codes"`
	AuthMethod          HttpTaskAuthMethod   `json:"auth_method" db:"auth_method"`
}

type HttpBasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpOAuth2Auth struct {
	Method       HttpOAuth2AuthMethod `json:"method" db:"oauth_method"`
	Url          string               `json:"url"`
	ClientID     string               `json:"client_id"`
	ClientSecret string               `json:"client_secret"`
	ClientScope  string               `json:"client_scope"`
}

type HttpAuth struct {
	HttpBasicAuth
	HttpOAuth2Auth
	ID     int `json:"-"`
	TaskID int `json:"task_id" db:"task_id"`
}

type TaskType string
type TaskTags []string
type HttpTaskStatusCodes []int
type HttpTaskMethod string
type HttpTaskAuthMethod string
type HttpTaskHeaders map[string]string
type HttpTaskBodyEncoding string
type HttpOAuth2AuthMethod string

func (v TaskTags) Value() (driver.Value, error) {
	return jsonSqlValuer(v)
}
func (v *TaskTags) Scan(src interface{}) error {
	return jsonSqlScanner(v, src)
}

func (v HttpTaskStatusCodes) Value() (driver.Value, error) {
	return jsonSqlValuer(v)
}
func (v *HttpTaskStatusCodes) Scan(src interface{}) error {
	return jsonSqlScanner(v, src)
}

func (v HttpTaskHeaders) Value() (driver.Value, error) {
	return jsonSqlValuer(v)
}
func (v *HttpTaskHeaders) Scan(src interface{}) error {
	return jsonSqlScanner(v, src)
}

const (
	TaskTypeHttp TaskType = "http"
)

const (
	HttpTaskMethodGet     HttpTaskMethod = "get"
	HttpTaskMethodPost    HttpTaskMethod = "post"
	HttpTaskMethodPut     HttpTaskMethod = "put"
	HttpTaskMethodPatch   HttpTaskMethod = "patch"
	HttpTaskMethodDelete  HttpTaskMethod = "delete"
	HttpTaskMethodHead    HttpTaskMethod = "head"
	HttpTaskMethodOptions HttpTaskMethod = "options"
)
const (
	HttpOAuth2AuthMethodHeader HttpOAuth2AuthMethod = "header"
	HttpOAuth2AuthMethodForm   HttpOAuth2AuthMethod = "form"
)

const (
	HttpTaskAuthMethodNone   HttpTaskAuthMethod = "none"
	HttpTaskAuthMethodBasic  HttpTaskAuthMethod = "basic"
	HttpTaskAuthMethodOAuth2 HttpTaskAuthMethod = "oauth2"
)

const (
	HttpTaskBodyEncodingNone HttpTaskBodyEncoding = "none"
	HttpTaskBodyEncodingText HttpTaskBodyEncoding = "text"
	HttpTaskBodyEncodingHtml HttpTaskBodyEncoding = "html"
	HttpTaskBodyEncodingForm HttpTaskBodyEncoding = "form"
	HttpTaskBodyEncodingJson HttpTaskBodyEncoding = "json"
	HttpTaskBodyEncodingXml  HttpTaskBodyEncoding = "xml"
)

func createTasksTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		repeat BOOLEAN NOT NULL,
		active BOOLEAN NOT NULL,
		interval INTEGER NOT NULL,
		tags TEXT NOT NULL DEFAULT "[]",
		description TEXT NOT NULL DEFAULT "",
		type TEXT CHECK( type IN ('http') ) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := DB.Exec(q)
	return err
}

func createHttpTasksTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS http_tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		method TEXT NOT NULL DEFAULT 'get' CHECK ( method IN (
			'get',
			'post',
			'put',
			'patch',
			'delete',
			'head',
			'options'
		) ),
		url TEXT NOT NULL,
		body  TEXT NOT NULL DEFAULT "",
		headers TEXT NOT NULL DEFAULT "{}",
		encoding TEXT NOT NULL DEFAULT 'none' CHECK ( encoding IN (
			'none',
			'text',
			'html',
			'form',
			'json',
			'xml'
		) ),
		retries INTEGER NOT NULL DEFAULT 0,
		timeout INTEGER NOT NULL DEFAULT 60,
		accepted_status_codes TEXT NOT NULL DEFAULT "[]",
		auth_method TEXT NOT NULL DEFAULT "none" CHECK ( auth_method IN (
			'none',
			'basic',
			'oauth2'
		) ),
		CONSTRAINT fk_tasks FOREIGN KEY (task_id) REFERENCES tasks(id)
	)
	`
	_, err := DB.Exec(q)
	return err
}

func createHttpAuthsTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS http_auths (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		username  TEXT NOT NULL DEFAULT "",
		password  TEXT NOT NULL DEFAULT "",
		oauth_method TEXT NOT NULL DEFAULT "header" CHECK ( oauth_method IN ('header', 'form') ),
		oauth_url  TEXT NOT NULL DEFAULT "",
		oauth_client_id  TEXT NOT NULL DEFAULT "",
		oauth_client_secret  TEXT NOT NULL DEFAULT "",
		oauth_client_scope  TEXT NOT NULL DEFAULT "",
		CONSTRAINT fk_tasks FOREIGN KEY (task_id) REFERENCES tasks(id)
	)
	`
	_, err := DB.Exec(q)
	return err
}

func CreateTask(t *Task) (int64, error) {
	q := `
	INSERT INTO tasks (
		name,
		repeat,
		active,
		interval,
		description,
		tags,
		type
	) VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	res, err := DB.Exec(q, t.Name, t.Repeat, t.Active, t.Interval, t.Description, t.Tags, t.Type)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func CreateHttpTask(t *HttpTask) (int64, error) {
	q := `
	INSERT INTO http_tasks (
		task_id,
		method,
		url,
		body,
		headers,
		encoding,
		retries,
		timeout,
		accepted_status_codes,
		auth_method
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	res, err := DB.Exec(q, t.TaskID, t.Method, t.Url, t.Body, t.Headers, t.Encoding, t.Retries, t.Timeout, t.AcceptedStatusCodes, t.AuthMethod)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func CreateHttpBasicAuth(tid int, m *HttpBasicAuth) (int64, error) {
	q := `
	INSERT INTO http_auths (
		task_id,
		username,
		password
	) VALUES (?, ?, ?)
	`
	res, err := DB.Exec(q, tid, m.Username, m.Password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func CreateHttpOAuth(tid int, m *HttpOAuth2Auth) (int64, error) {
	q := `
	INSERT INTO http_auths (
		task_id,
		oauth_method,
		oauth_url,
		oauth_client_id,
		oauth_client_secret,
		oauth_client_scope
	) VALUES (?, ?, ?, ?, ?, ?)
	`
	res, err := DB.Exec(q, tid, m.Method, m.Url, m.ClientID, m.ClientSecret, m.ClientScope)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllActiveTasks() ([]Task, error) {
	q := `SELECT * FROM tasks WHERE active = true`
	var tasks []Task
	err := DB.Select(&tasks, q)
	return tasks, err
}

func GetTaskByID(id int) (Task, error) {
	q := `SELECT * FROM tasks WHERE id = ?`
	var task Task
	err := DB.Get(&task, q, id)
	return task, err
}

func ActivateTaskByID(id int) error {
	q := `UPDATE tasks SET active = true WHERE id = ?`
	_, err := DB.Exec(q, id)
	return err
}

func DeactivateTaskByID(id int) error {
	q := `UPDATE tasks SET active = false WHERE id = ?`
	_, err := DB.Exec(q, id)
	return err
}

func GetHttpTaskByTaskID(tid int) (HttpTask, error) {
	q := `SELECT * FROM http_tasks WHERE task_id = ?`
	var doc HttpTask
	err := DB.Get(&doc, q, tid)
	return doc, err
}
