package main

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		event := make(chan string)
		go func() {
			for x := range time.Tick(time.Millisecond * 200) {
				event <- x.String()
			}
		}()
		c.Stream(func(w io.Writer) bool {
			c.SSEvent("message", <-event)
			return true
		})
	})
	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM 
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port) 
}
