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
		}
		var (
			mt  int
			msg []byte
			ch  chan string
		)
		c.SetCloseHandler(func(code int, text string) error {
			fmt.Println("cls hndlr", code, text)
			ch <- "stop"
			c.Close()
			return nil
		})
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				ch <- "stop"
				c.Close()
				break
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
