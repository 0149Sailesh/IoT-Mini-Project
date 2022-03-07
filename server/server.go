package main

import(
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/0149Sailesh/iot-server/services"
)

func main()  {

	r := gin.Default()
	err := godotenv.Load()
	
	var l = log.WithFields(log.Fields{
		"method":"Init method",
	})

	if err != nil {
		l.Fatal("Error loading .env file")
	  }
	
	services.ConnectDB()

	url := os.Getenv("URL")
    r.Run(url)
}