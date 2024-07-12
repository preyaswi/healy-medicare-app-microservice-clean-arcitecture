package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)
func ChatRoute(route *fiber.App, chatHander *handler.ChatHandler) {
	chat := route.Group("/chat")
	chat.Use(middleware.UserAuthMiddleware())
	chat.Get("", websocket.New(chatHander.FriendMessage))
	chat.Get("/messages", chatHander.GetChat)
}

