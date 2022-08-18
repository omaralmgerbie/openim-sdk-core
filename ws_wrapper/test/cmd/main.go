package main

import (
	"flag"
	"fmt"

	"open_im_sdk/ws_wrapper/test"
	"open_im_sdk/ws_wrapper/test/client"
	"time"
)

var jssdkURL = flag.String("url", "ws://43.128.5.63:10003/", "jssdk URL")
var imAPI = flag.String("api", "http://43.128.5.63:10002", "openIM api")
var connNum = flag.Int("connNum", 1, "conn num")

func main() {

	fmt.Printf("simulation js client, user num: %d, jssdkURL:%s, apiURL:%s \n\n", *connNum, *jssdkURL, *imAPI)

	admin := client.NewIMClient("", "openIMAdmin", *imAPI, *jssdkURL, 1)
	var err error
	admin.Token, err = admin.GetToken()
	if err != nil {
		panic(err)
	}
	uidList, err := admin.GetUserIDList()
	if err != nil {
		panic(err)
	}
	l := uidList[0:*connNum]
	l = []string{"MTc3MjYzNzg0Mjg="}
	for num, userID := range l {
		time.Sleep(time.Second * 1)
		go test.StartSimulationJSClient(*imAPI, *jssdkURL, userID, num, l)
	}

	for {
		time.Sleep(time.Second * 100)
		fmt.Println("jssdk simulation is running")
	}

}
