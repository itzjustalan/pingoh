package routes

import (
	"pingoh/controllers"
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
		res, err := controllers.HttpResultsByTaskID(tid)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	r.Get("/ws/task/:task_id", websocket.New(func(c *websocket.Conn) {
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
			ch <- "stop"
			c.Close()
			return nil
		})
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				c.WriteMessage(websocket.TextMessage, []byte("error reading command"))
				c.Close()
				return
			}
			switch mt {
			case 1:
				switch cmd := string(msg); cmd {
				case "start":
					res, err := controllers.HttpResultsByTaskID(tid)
					if err != nil {
						c.WriteMessage(websocket.TextMessage, []byte("error fetching results"))
					}
					c.WriteJSON(res)
					go controllers.SendHttpResultsUpdates(tid, c, ch)
				case "stop":
					ch <- "stop"
					return
				default:
					c.WriteMessage(websocket.TextMessage, []byte("unknown command"))
				}
			default:
				c.WriteMessage(websocket.TextMessage, []byte("unknown message type"))
			}
		}
	}))
}
