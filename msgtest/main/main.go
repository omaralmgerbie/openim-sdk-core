package main

import (
	"context"
	"flag"
	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/msgtest/module"
	"time"
)

func init() {
	InitWithFlag()
	if err := log.InitFromConfig("sdk.log", "sdk", 5,
		true, false, "./", 2, 24); err != nil {
		panic(err)
	}
}

var (
	totalOnlineUserNum    int     // 总在线用户数
	samplingRate          float64 // 抽样率
	NotFriendMsgSenderNum int     // 非好友消息发送者数
	groupMsgSenderNum     int     // 群消息发送者数
	msgSenderNumEvreyUser int     // 每个用户的消息数
	fastenedUserNum       int     // 固定用户数

	//recvMsgUserNum int // 消息接收者数, 抽样账号
	isRegisterUser bool // 是否注册用户
)

func InitWithFlag() {
	flag.IntVar(&totalOnlineUserNum, "o", 20000, "total online user num")
	flag.Float64Var(&samplingRate, "f", 0.1, "sampling rate")
	flag.IntVar(&NotFriendMsgSenderNum, "n", 100, "not friend msg sender num")
	flag.IntVar(&groupMsgSenderNum, "g", 100, "group msg sender num")
	flag.IntVar(&msgSenderNumEvreyUser, "m", 100, "msg sender num evrey user")

	flag.BoolVar(&isRegisterUser, "r", false, "register user to IM system")
	flag.IntVar(&fastenedUserNum, "u", 300, "fastened user num")
}

func PrintQPS() {
	for {
		time.Sleep(1 * time.Second)
		log.ZDebug(context.Background(), "QPS", "qps", module.GetQPS())
	}
}

func main() {
	flag.Parse()
	ctx := context.Background()
	p := module.NewPressureTester()
	f, r, err := p.SelectSample(totalOnlineUserNum, 0.01)
	if err != nil {
		log.ZError(ctx, "Sample UserID failed", err)
		return
	}
	log.ZDebug(ctx, "Sample UserID", "sampleUserLength", len(r), "sampleUserID", r, "length", len(f))
	time.Sleep(10 * time.Second)
	//
	if isRegisterUser {
		if err := p.RegisterUsers(f, nil, nil); err != nil {
			log.ZError(ctx, "Sample UserID failed", err)
			return
		}
	}
	go PrintQPS()
	// init users
	p.InitUserConns(f)
	log.ZDebug(ctx, "all user init connect to server success,start send message")
	time.Sleep(10 * time.Second)
	p.SendSingleMessages(f, 10, time.Millisecond*100)
	time.Sleep(1 * time.Minute)
	p.CheckMsg(ctx)

	time.Sleep(time.Hour * 60)

}
