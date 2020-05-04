package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	var payload map[string]interface{}
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("%v %v\n", payload["msg"], int64(payload["time"].(float64)))

	//fmt.Println([]byte(stream))
	//Duplicate() bool
	//Qos() byte
	//Retained() bool
	//Topic() string
	//MessageID() uint16
	//Payload() []byte
	//Ack()
}

func TestMqtt() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	// opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")
	opts := mqtt.NewClientOptions().AddBroker("ws://localhost:8083/mqtt").SetClientID("emqx_go_client")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	fmt.Println("###############mqtt.NewClient(opts)###############")
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//订阅主题
	fmt.Println("###############Subscribe()###############")
	if token := c.Subscribe("#", 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		//os.Exit(1)
	}

	// 发布消息
	//timeNow := time.Now().Unix()
	//var msg = map[string]interface{}{
	//	"time": timeNow,
	//	"msg":  "hello",
	//}


	//res, err := json.Marshal(msg)
	//if err != nil {
	//	fmt.Println("json.Marshal failed:", err)
	//	return
	//}
	//
	//fmt.Println("before pub", string(res))
	//token := c.Publish("goclient/1", 0, true, res)
	//token.Wait()
	//fmt.Println("after pub", string(res))

	// time.Sleep(6 * time.Second)

	// 取消订阅
	// fmt.Println("###############Unsubscribe()###############")
	// if token := c.Unsubscribe("emqx/#"); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	// 断开连接
	fmt.Println("###############Disconnect()###############")
	// c.Disconnect(250)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 20000,
	//	"msg":  "success",
	//	"data": gin.H{
	//		"time":    msg["time"],
	//		"message": msg["message"],
	//		"topic":   "goclient/1",
	//	},
	//})

}
