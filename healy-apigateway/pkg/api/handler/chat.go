package handler

import (
	"fmt"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/helper"
	"healy-apigateway/pkg/logging"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatHandler struct {
	Grpc_Client interfaces.ChatClient
}

func NewChatHandler(ChatingClient interfaces.ChatClient) *ChatHandler {
	return &ChatHandler{
		Grpc_Client: ChatingClient,
	}
}

var User = make(map[string]*websocket.Conn)

// FriendMessage handles WebSocket connections for chat messages
// @Summary WebSocket connection for chat messages
// @Description Establish WebSocket connection for chat
// @Tags Chat
// @Security Bearer
// @Produce application/json
// @Router /chat [get]
func (ch *ChatHandler) FriendMessage(c *websocket.Conn) {
	var mu sync.Mutex
	userID := c.Locals("user_id").(string)
	mu.Lock()
	defer mu.Unlock()
	defer delete(User, userID)
	defer c.Close()

	User[userID] = c

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			response := response.ClientResponse("Details not in correct format", nil, err.Error())
			c.WriteJSON(response)
			return
		}
		helper.SendMessageToUser(User, msg, userID)
	}
}
// GetChat retrieves chat messages
// @Summary Get Chat Messages
// @Description Retrieve chat messages
// @Tags Chat
// @Security Bearer
// @Produce application/json
// @Param FriendID query string true "FriendID"
// @Param Offset query string true "Offset"
// @Param Limit query string true "Limit"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /chat/messages [get]
func (ch *ChatHandler) GetChat(c *fiber.Ctx) error {
	logEntry := logging.Logger().WithField("context", "GetChatHandler")
	logEntry.Info("Processing GetChat request")

	// Retrieve query parameters from the request
	friendID := c.Query("FriendID")
	offset := c.Query("Offset")
	limit := c.Query("Limit")

	// Validate parameters
	if friendID == "" || offset == "" || limit == "" {
		errMsg := "FriendID, Offset, or Limit missing or invalid"
		logEntry.Error(errMsg)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": errMsg,
			"data":    nil,
			"error":   "Bad Request",
		})
	}

	// Convert offset and limit to appropriate types if needed
	// ...

	userID := c.Locals("user_id").(string)

	// Prepare the chat request object
	chatRequest := models.ChatRequest{
		FriendID: friendID,
		Offset:   offset,
		Limit:    limit,
	}

	// Call your gRPC client method with the retrieved parameters
	result, err := ch.Grpc_Client.GetChat(userID, chatRequest)
	fmt.Println(result)
	if err != nil {
		logEntry.WithError(err).Error("Error in getting chat details")
		errs := response.ClientResponse("Failed to get chat details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	logEntry.Info("get chat successful")
	successRes := response.ClientResponse("successfully got all chat details", result, nil)
	return c.Status(http.StatusOK).JSON(successRes)
}
