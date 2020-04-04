package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func TestMqtt(ctx *gin.Context) {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	// opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")
	opts := mqtt.NewClientOptions().AddBroker("ws://localhost:8083/mqtt").SetClientID("emqx_test_client")

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

	if token := c.Subscribe("emqx/#", 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		//os.Exit(1)
	}

	// 发布消息
	timeNow := time.Now().Unix()
	var msg = map[string]interface{}{
		"time":    timeNow,
		"message": "hello",
	}
	res, err := json.Marshal(msg)
	// b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println("before pub", string(res))
	token := c.Publish("goclient/1", 0, true, res)
	token.Wait()
	fmt.Println("after pub", res)

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
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "success",
		"data": gin.H{
			"time":    msg["time"],
			"message": msg["message"],
			"topic":   "goclient/1",
		},
	})
	// time.Sleep(1 * time.Second)
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	//opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")
	//
	//opts.SetKeepAlive(60 * time.Second)
	//// 设置消息回调处理函数
	//opts.SetDefaultPublishHandler(f)
	//opts.SetPingTimeout(1 * time.Second)
	//
	//c := mqtt.NewClient(opts)
	//if token := c.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}
	//
	//// 订阅主题
	//if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}
	//
	//// 发布消息
	//token := c.Publish("testtopic/1", 0, false, "Hello World testtttttttttt")
	//token.Wait()
	//
	//time.Sleep(6 * time.Second)
	//
	//// 取消订阅
	//if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}
	//
	//// 断开连接
	//c.Disconnect(250)
	//time.Sleep(1 * time.Second)
}
