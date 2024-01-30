package handlers

import (
	"pingoh/db"

	"github.com/gofiber/contrib/websocket"
)

func HttpResultsByTaskID(tid int) ([]db.HttpResult, error) {
	res, err := db.SelectAllHttpResultsByTaskID(tid)
	return res, err
}

func SendHttpResultsUpdates(tid int, c *websocket.Conn, ch chan string) {
	tch, ok := TaskChannels[tid]
	if !ok {
		err := ActivateTaskByID(tid)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
	}
	subID := tch.Subscribe()
	for {
		select {
		case data := <-tch.Subs[subID]:
			c.WriteJSON(data)
		case msg := <-ch:
			switch msg {
			case "stop":
				tch.Unsubscribe(subID)
				c.Close()
				return
			}
		}
	}
}
