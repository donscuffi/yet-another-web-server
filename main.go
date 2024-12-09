package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages []Message

func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	messages = append(messages, message)
	return c.JSON(http.StatusCreated, Response{
		Status:  "Success",
		Message: "Added the message",
	})
}

func main() {
	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)

	e.Start(":8080")
}
