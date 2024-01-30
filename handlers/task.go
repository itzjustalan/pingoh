package handlers

import (
	"pingoh/db"
	"pingoh/services"
	"slices"

	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type TaskChannel struct {
	mu     sync.RWMutex
	Active bool
	Stop   chan bool
	Subs   []chan *db.HttpResult
}

var TaskChannels = make(map[int]*TaskChannel)

func (tch *TaskChannel) Deactivate() {
	tch.mu.Lock()
	defer tch.mu.Unlock()
	tch.Active = false
	tch.Stop <- true
}

func (tch *TaskChannel) Publish(r *db.HttpResult) {
	for i := 0; i < len(tch.Subs); i++ {
		go func(i int) { tch.Subs[i] <- r }(i)
	}
}

func (tch *TaskChannel) Subscribe() int {
	tch.mu.Lock()
	defer tch.mu.Unlock()
	ch := make(chan *db.HttpResult, 1)
	tch.Subs = append(tch.Subs, ch)
	return len(tch.Subs) - 1
}

func (tch *TaskChannel) Unsubscribe(subID int) {
	tch.mu.Lock()
	defer tch.mu.Unlock()
	tch.Subs = append(tch.Subs[:subID], tch.Subs[subID+1:]...)
}

type NewTask struct {
	Name        string   `json:"name" validate:"required"`
	Type        string   `json:"type" validate:"required,eq=http"`
	Repeat      bool     `json:"repeat" validate:"required,boolean"`
	Active      bool     `json:"active" validate:"required,boolean"`
	Interval    int      `json:"interval" validate:"required,gte=1"`
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
		Active:      t.Active,
		Interval:    t.Interval,
		Description: t.Description,
		Tags:        t.Tags,
		Type:        db.TaskType(t.Type),
	}
	tid, err := db.CreateTask(&tm)
	if err != nil {
		return err
	}
	tm.ID = int(tid)
	switch t.Type {
	case "http":
		slices.Sort(t.Http.AcceptedStatusCodes)
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
	go startTask(tm)
	return nil
}

func StartTasks() {
	log.Info().Msg("starting tasks")
	tasks, err := db.GetAllActiveTasks()
	if err != nil {
		log.Error().Err(err).Msg("failed fetching tasks from db")
		return
	}
	for i := 0; i < len(tasks); i++ {
		go startTask(tasks[i])
	}
}

func ActivateTaskByID(tid int) error {
	if v, ok := TaskChannels[tid]; ok && v.Active {
		return nil
	}
	err := db.ActivateTaskByID(tid)
	if err != nil {
		return err
	}
	task, err := db.GetTaskByID(tid)
	if err != nil {
		return err
	}
	go startTask(task)
	return nil
}

func DeactivateTaskByID(tid int) error {
	if v, ok := TaskChannels[tid]; ok {
		v.Deactivate()
	}
	return db.DeactivateTaskByID(tid)
}

func startTask(t db.Task) {
	log.Info().Msgf("starting task: %v - %v", t.Name, t.ID)
	ticker := time.NewTicker(time.Second * time.Duration(t.Interval))
	defer ticker.Stop()
	if v, ok := TaskChannels[t.ID]; ok {
		v.Active = t.Active
	} else {
		TaskChannels[t.ID] = &TaskChannel{
			Active: t.Active,
			Stop:   make(chan bool),
			Subs:   make([]chan *db.HttpResult, 0),
		}
	}

	for {
		select {
		case <-TaskChannels[t.ID].Stop:
			log.Info().Msgf("stopping task: %v - %v", t.Name, t.ID)
			if v, ok := TaskChannels[t.ID]; ok {
				v.Active = false
			}
			return
		case <-ticker.C:
			if v, ok := TaskChannels[t.ID]; ok && v.Active {
				log.Info().Msgf("running task: %v - %v", t.Name, t.ID)
				switch t.Type {
				case "http":
					runHttpTask(&t)
				}
				if !t.Repeat || t.Interval == 0 {
					return
				}
			} else {
				log.Info().Msgf("tried calling inactive task")
				v.Deactivate()
				return
			}
		}
	}
}

func runHttpTask(task *db.Task) {
	t, err := db.GetHttpTaskByTaskID(task.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed fetching task details from db")
		return
	}

	client := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(t.Url)
	req.SetBodyRaw([]byte(t.Body))
	req.Header.SetMethod(string(t.Method))
	req.Header.Set("Content-Type", services.HeaderForEncoding(string(t.Encoding)))
	for k, v := range t.Headers {
		req.Header.Set(k, v)
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	for i := 0; i <= t.Retries; i++ {
		startTime := time.Now()
		err = client.DoTimeout(req, res, time.Duration(t.Timeout*int(time.Second)))
		if err != nil {
			log.Error().Err(err).Msg("failed sending request")
			return
		}
		success := slices.Contains(t.AcceptedStatusCodes, res.StatusCode())
		if success || i == t.Retries {
			result := db.HttpResult{
				TaskID:   t.TaskID,
				Code:     res.StatusCode(),
				Ok:       success,
				Duration: time.Since(startTime),
			}
			TaskChannels[task.ID].Publish(&result)
			db.AddHttpResult(&result)
			break
		}
	}
}
