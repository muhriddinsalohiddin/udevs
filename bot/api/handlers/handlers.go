package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhriddinsalohiddin/udevs/bot/client"
)

type Message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

// @BasePath /v1

// @Summary MessageSender
// @Description send message
// @Tags Massage
// @Accept json
// @Produce json
// @Param Massage request body Message true "message_content"
// @Success 200
// @Failure 400
// @Router /send/ [post]
func SendMessageAPI(c *gin.Context) {
	var newMessage Message
	err := c.ShouldBindJSON(&newMessage)
	if err != nil {
		log.Printf("Problem with getting message from API Server: %v", err)
	}
	status := client.Stub(newMessage.Text, newMessage.Priority)

	c.AbortWithStatus(status)
}
