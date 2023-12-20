package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "math/rand"
  "net/http"
  "time"
)

const (
  ServerToken    = "abcde"
  mainServiceUrl = "http://127.0.0.1:8000/links/change_imo/"
)

func main() {
  r := InitRoutes()
  r.Run(":9000")
}


type Order_For_Ship struct {
	Order_id          int    `json:"order_id"`
	Ship_id            int    `json:"ship_id"`
	IMO string `json:"imo"`

}
func RandStringBytesRmndr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

func InitRoutes() *gin.Engine {
  r := gin.Default()

  r.PUT("/", func(c *gin.Context) {
    // условная проверка авторизации
    token := c.GetHeader("Server-Token")
    if token != ServerToken {
      c.AbortWithStatusJSON(http.StatusForbidden, "неверный токен авторизации")
      return
    }

    var request Order_For_Ship

    if err := c.BindJSON(&request); err != nil {
      c.AbortWithStatusJSON(http.StatusBadRequest, "неверный формат данных")
      return
    }
	request.IMO = RandStringBytesRmndr(3) + "12345678"
    body, err := json.Marshal(request)
    if err != nil {
      fmt.Println(err)
    }

    // асинхронное вычисление
    go func() {
      res := performTask()
	  if res{}else{}


      client := &http.Client{}
      req, err := http.NewRequest("PUT", mainServiceUrl, bytes.NewBuffer(body))
      if err != nil {
        fmt.Println("Error creating request:", err)
        return
      }

      req.Header.Set("Content-Type", "application/json")
      req.Header.Set("Server-Token", ServerToken)

      _, err = client.Do(req)
      if err != nil {
        fmt.Println("Error sending request:", err)
        return
      }
    }()

    c.JSON(http.StatusOK, gin.H{"message": "заявка принята в работу"})
  })

  return r
}

func performTask() bool {
  // Задержка от 6 до 10 секунд
  delay := rand.Intn(1) + 4
  time.Sleep(time.Duration(delay) * time.Second)

  // Генерируем случайное число от [0;4]
  // Если число меньше 3, возвращаем true (успех), иначе - false (неудача)
  // 60% - успех

  return true
}