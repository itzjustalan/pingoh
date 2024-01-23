package db

import "time"

type Task struct {
	ID          int       `json:"-"`
	UID         string    `json:"uid"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Repeat      bool      `json:"repeat"`
	Interval    int       `json:"interval"`
	Description string    `json:"description"`
	Tags        TaskTags  `json:"tags"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type HttpTask struct {
	TaskID        int                  `json:"-"`
	Method        HttpTaskMethod       `json:"method"`
	Url           string               `json:"url"`
	Body          string               `json:"body"`
	Headers       HttpTaskHeaders      `json:"headers"`
	Encoding      HttpTaskBodyEncoding `json:"encoding"`
	Retries       int                  `json:"retries"`
	Timeout       int                  `json:"timeout"`
	AcceptedCodes HttpTaskStatusCodes  `json:"accepted_codes"`
	AuthMethod    HttpTaskAuthMethod   `json:"auth_method"`
}

type HttpBasicAuth struct {
	TaskID   int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpOAuth2Auth struct {
	TaskID       int                  `json:"-"`
	Method       HttpOAuth2AuthMethod `json:"method"`
	TokenUrl     string               `json:"token_url"`
	ClientID     string               `json:"client_id"`
	ClientSecret string               `json:"client_secret"`
	ClientScope  string               `json:"client_scope"`
}

type TaskType string
type TaskTags []string
type HttpTaskStatusCodes []int
type HttpTaskMethod string
type HttpTaskAuthMethod string
type HttpTaskHeaders map[string]string
type HttpTaskBodyEncoding string
type HttpOAuth2AuthMethod string

const (
	TaskTypeHttp TaskType = "http"
)

const (
	HttpTaskMethodGet    HttpTaskMethod = "get"
	HttpTaskMethodPost   HttpTaskMethod = "post"
	HttpTaskMethodPut    HttpTaskMethod = "put"
	HttpTaskMethodPatch  HttpTaskMethod = "patch"
	HttpTaskMethodDelete HttpTaskMethod = "delete"
)
const (
	HttpOAuth2AuthMethodHeader   HttpOAuth2AuthMethod = "header"
	HttpOAuth2AuthMethodFormData HttpOAuth2AuthMethod = "form_data"
)

const (
	HttpTaskAuthMethodBasicAuth HttpTaskAuthMethod = "basic_auth"
	HttpTaskAuthMethodOAuth2    HttpTaskAuthMethod = "oauth2"
)

const (
	HttpTaskBodyEncodingJson HttpTaskBodyEncoding = "json"
	HttpTaskBodyEncodingXml  HttpTaskBodyEncoding = "xml"
	HttpTaskBodyEncodingText HttpTaskBodyEncoding = "text"
	HttpTaskBodyEncodingHtml HttpTaskBodyEncoding = "html"
	HttpTaskBodyEncodingForm HttpTaskBodyEncoding = "form"
)

func createTasksTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		repeat BOOLEAN NOT NULL,
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
