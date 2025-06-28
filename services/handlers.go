package services

// 2.
// fun обработки запроса от клиента
import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetAnsw(c *gin.Context) {
	c.JSON(200, gin.H{"owner_id": "5635393648"})
}

func PostAnsw(c *gin.Context) {
	const BotToken = "7574486002:AAElO_kif9X9jfx5uLhjMda7EJfyK9c54O4" // токен в конфиги, зачем с большой буквы?

	var json struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}

	//type Response struct {  нужно делать через type потом уносить в пакет models
	//	ChatID int64  `json:"chat_id"`
	//	Text   string `json:"text"`
	//}

	// var resp ... создаём структуру пустую

	if err := c.ShouldBindJSON(&json); err != nil { // тут вместо json будет resp
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	sendMessageURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BotToken)

	resp, err := http.PostForm(sendMessageURL, url.Values{
		"chat_id": {fmt.Sprintf("%d", json.ChatID)},
		"text":    {json.Text},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	defer resp.Body.Close()

	body, _ := c.GetRawData()               // почитать про библиотеку io. использовать c.GetRawData() вместо io.ReadAll (суть та же, но экономит строчки, можно внутри посмотреть, что оно делдает)
	c.String(resp.StatusCode, string(body)) //
}
