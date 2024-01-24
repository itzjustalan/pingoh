package handlers

import (
	"pingoh/db"

	"github.com/gofiber/fiber/v2"
)

type NewTask struct {
	Name                string            `json:"name"`
	Type                string            `json:"type"`
	Repeat              bool              `json:"repeat"`
	Interval            int               `json:"interval"`
	Description         string            `json:"description"`
	Tags                []string          `json:"tags"`
	Method              string            `json:"method"`
	URL                 string            `json:"url"`
	Body                string            `json:"body"`
	Headers             map[string]string `json:"headers"`
	Encoding            string            `json:"encoding"`
	Retries             int               `json:"retries"`
	Timeout             int               `json:"timeout"`
	AcceptedStatusCodes []int             `json:"accepted_status_codes"`
	AuthMethod          string            `json:"auth_method"`
	Auth                struct {
		Username          string `json:"username"`
		Password          string `json:"password"`
		OauthMethod       string `json:"oauth_method"`
		OauthUrl          string `json:"oauth_url"`
		OauthClientID     string `json:"oauth_client_id"`
		OauthClientSecret string `json:"oauth_client_secret"`
		OauthClientScope  string `json:"oauth_client_scope"`
	} `json:"auth"`
}

func CreateNewTask(t *NewTask) error {
	tm := db.Task{
		Name:        t.Name,
		Repeat:      t.Repeat,
		Interval:    t.Interval,
		Description: t.Description,
		Tags:        t.Tags,
		Type:        db.TaskType(t.Type),
	}
	tid, err := db.CreateTask(&tm)
	if err != nil {
		return err
	}
	switch t.Type {
	case "http":
		ht := db.HttpTask{
			TaskID:              int(tid),
			Method:              db.HttpTaskMethod(t.Method),
			Url:                 t.URL,
			Body:                t.Body,
			Encoding:            db.HttpTaskBodyEncoding(t.Encoding),
			Headers:             t.Headers,
			Retries:             t.Retries,
			Timeout:             t.Timeout,
			AcceptedStatusCodes: t.AcceptedStatusCodes,
			AuthMethod:          db.HttpTaskAuthMethod(t.AuthMethod),
		}
		_, err := db.CreateHttpTask(&ht)
		if err != nil {
			return err
		}
		switch t.AuthMethod {
		case "none":
			break
		case "basic":
			ba := db.HttpBasicAuth{
				Username: t.Auth.Username,
				Password: t.Auth.Password,
			}
			_, err := db.CreateHttpBasicAuth(int(tid), &ba)
			if err != nil {
				return err
			}
		case "oauth2":
			oa := db.HttpOAuth2Auth{
				Method:       db.HttpOAuth2AuthMethod(t.Auth.OauthMethod),
				Url:          t.Auth.OauthUrl,
				ClientID:     t.Auth.OauthClientID,
				ClientSecret: t.Auth.OauthClientSecret,
				ClientScope:  t.Auth.OauthClientScope,
			}
			_, err := db.CreateHttpOAuth(int(tid), &oa)
			if err != nil {
				return err
			}
		default:
			return fiber.NewError(fiber.ErrBadRequest.Code, "auth method not supported")
		}
	default:
		return fiber.NewError(fiber.ErrBadRequest.Code, "task type not supported")
	}
	return nil
}
