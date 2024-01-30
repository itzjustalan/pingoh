package api

import (
	"fmt"
	"pingoh/handlers"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func addStatsRoutes(api *fiber.Router) {
	r := (*api).Group("/stats")

	r.Get("/task/:task_id", func(c *fiber.Ctx) error {
		tid, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		res, err := handlers.HttpResultsByTaskID(tid)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	r.Get("/ws/task/:task_id", websocket.New(func(c *websocket.Conn) {
		fmt.Println("req rcvd")
		tid, err := strconv.Atoi(c.Params("task_id", ""))
		if err != nil {
			c.Close()
			return
		}
		var (
			mt  int
			msg []byte
		)
		ch := make(chan string)
		c.SetCloseHandler(func(code int, text string) error {
			fmt.Println("cls hndlr", code, text)
			ch <- "stop"
			fmt.Println("sent stp")
			c.Close()
			return nil
		})
		for {
			fmt.Println("ni aano??")
			if mt, msg, err = c.ReadMessage(); err != nil {
				c.WriteMessage(websocket.TextMessage, []byte("error reading command"))
				c.Close()
				return
			}
			fmt.Println("mtmsgerr", mt, msg, err)
			switch mt {
			case 1:
				switch cmd := string(msg); cmd {
				case "start":
					fmt.Println("start 10")
					res, err := handlers.HttpResultsByTaskID(tid)
					if err != nil {
						c.WriteMessage(websocket.TextMessage, []byte("error fetching results"))
					}
					c.WriteJSON(res)
					go handlers.SendHttpResultsUpdates(tid, c, ch)
				case "stop":
					ch <- "stop"
					return
				default:
					fmt.Println("undefined cmd: " + cmd)
					c.WriteMessage(websocket.TextMessage, []byte("unknown command"))
				}
			default:
				c.WriteMessage(websocket.TextMessage, []byte("unknown message type"))
			}
		}
	}))
}
