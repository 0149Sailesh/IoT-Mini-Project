package main

import(
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/0149Sailesh/iot-server/services"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
    "github.com/gin-contrib/cors"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}


func sub(client mqtt.Client) {
    topic1 := "vella"
    token := client.Subscribe(topic1, 1, nil)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s", topic1)

}
func publish(client mqtt.Client) {
        token := client.Publish("ulla", 0, false, "1")
        token.Wait()
    
}
func main()  {
	var broker = "broker.mqtt-dashboard.com"
    var port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID("go_mqtt_client5787653854")
    opts.SetUsername("iot")
    opts.SetPassword("things")
    opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
  }

    sub(client)
	var l = log.WithFields(log.Fields{
		"method":"Init method",
	})
	r := gin.Default()
    r.Use(cors.Default())
    r.GET("/trigger",func(c *gin.Context){
        publish(client)
    })

	err := godotenv.Load()
	
	if err != nil {
		l.Fatal("Error loading .env file")
	  }
	
	services.ConnectInit()
    
    r.GET("/ws",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
    portNo := os.Getenv("PORT")
    r.Run(":"+portNo)
}