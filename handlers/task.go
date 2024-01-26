package handlers

import (
	"pingoh/db"

	"github.com/gofiber/fiber/v2"
)

type NewTask struct {
	Name        string   `json:"name" validate:"required"`
	Type        string   `json:"type" validate:"required,eq=http"`
	Repeat      bool     `json:"repeat" validate:"required,boolean"`
	Interval    int      `json:"interval" validate:"required,gte=10"`
	Description string   `json:"description" validate:"required,omitempty"`
	Tags        []string `json:"tags" validate:"required,dive,unique"`
	Http        struct {
		Method              string            `json:"method" validate:"required,oneof=get post put patch delete head options"`
		URL                 string            `json:"url" validate:"required,url"`
		Body                string            `json:"body"`
		Headers             map[string]string `json:"headers" validate:"required"`
		Encoding            string            `json:"encoding" validate:"required,oneof=none text html form json xml"`
		Retries             int               `json:"retries" validate:"gte=0"`
		Timeout             int               `json:"timeout" validate:"gte=0"`
		AcceptedStatusCodes []int             `json:"accepted_status_codes" validate:"required,unique,dive,number"`
		AuthMethod          string            `json:"auth_method" validate:"required,oneof=none basic oauth2"`
		BasicAuth           struct {
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		} `json:"basic_auth" validate:"required_if=AuthMethod basic"`
		OAuth2 struct {
			OauthMethod       string `json:"oauth_method" validate:"required"`
			OauthUrl          string `json:"oauth_url" validate:"required,url"`
			OauthClientID     string `json:"oauth_client_id" validate:"required"`
			OauthClientSecret string `json:"oauth_client_secret" validate:"required"`
			OauthClientScope  string `json:"oauth_client_scope" validate:"required"`
		} `json:"oauth2" validate:"required_if=AuthMethod oauth2"`
	} `json:"http" validate:"required_if=Method http"`
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
			Method:              db.HttpTaskMethod(t.Http.Method),
			Url:                 t.Http.URL,
			Body:                t.Http.Body,
			Encoding:            db.HttpTaskBodyEncoding(t.Http.Encoding),
			Headers:             t.Http.Headers,
			Retries:             t.Http.Retries,
			Timeout:             t.Http.Timeout,
			AcceptedStatusCodes: t.Http.AcceptedStatusCodes,
			AuthMethod:          db.HttpTaskAuthMethod(t.Http.AuthMethod),
		}
		_, err := db.CreateHttpTask(&ht)
		if err != nil {
			return err
		}
		switch t.Http.AuthMethod {
		case "none":
			break
		case "basic":
			ba := db.HttpBasicAuth{
				Username: t.Http.BasicAuth.Username,
				Password: t.Http.BasicAuth.Password,
			}
			_, err := db.CreateHttpBasicAuth(int(tid), &ba)
			if err != nil {
				return err
			}
		case "oauth2":
			oa := db.HttpOAuth2Auth{
				Method:       db.HttpOAuth2AuthMethod(t.Http.OAuth2.OauthMethod),
				Url:          t.Http.OAuth2.OauthUrl,
				ClientID:     t.Http.OAuth2.OauthClientID,
				ClientSecret: t.Http.OAuth2.OauthClientSecret,
				ClientScope:  t.Http.OAuth2.OauthClientScope,
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
