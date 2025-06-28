package services

// 2.
// fun обработки запроса от клиента
import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"tg-bot/clients"
	"tg-bot/models"
)

func GetAnswer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"owner_id": "5635393648"})
}

func PostAnswer(c *gin.Context) {
	msg := models.Msg{}
	if err := c.ShouldBindBodyWithJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	client := clients.NewHttpClient()
	resp, err := client.SendMsg(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	c.String(resp.StatusCode, string(respBody)) //
}
