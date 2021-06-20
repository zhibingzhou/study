package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var ratelimit = time.Tick(300 * time.Millisecond)

// func main() {

// 	common.Test()

// }

func main() {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "50,400"),
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	taskCtx, cancel = context.WithTimeout(taskCtx, 100000*time.Second)
	defer cancel()

	// ensure that the browser process is started
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}

	// listen network event
	listenForNetworkEvent(taskCtx)

	chromedp.Run(taskCtx,
		network.Enable(),
		chromedp.Navigate(`https://xinghe.tv/play/1292001`),
		chromedp.WaitVisible(`body`, chromedp.BySearch),
		chromedp.Click(`document.querySelector("#__next > div.jsx-2684177236 > div > div.jsx-2803742412 > div > div.jsx-965897251.player-div > div.jsx-4030465136 > video:nth-child(2)")`, chromedp.BySearch),
	)
	for {

	}
}

//监听
func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {

		case *network.EventResponseReceived:
			resp := ev.Response
			if len(resp.Headers) != 0 {
				// log.Printf("received headers: %s", resp.Headers)

				if strings.Index(resp.URL, ".m3u8") != -1 {
					log.Printf("received headers: %s", resp.URL)
				}
			}
			fmt.Println(resp.URL)
		}
		// other needed network Event
	})
}

// //过程
// body, _ := common.Fetch(common.GoUrl + "/tianjin_hexi/")
// bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
// common.WriteToFile(bodystr)
// //结果
// result := common.PageUserList([]byte(bodystr))
// fmt.Println(result)

//爬数据
//common.Run("/map.asp")
//数据展示

// conf_byte, err := common.ReadFile("./conf/conf.json")

// if err != nil {
// 	panic(err)
// }
// var json_conf map[string]string
// //解析json格式
// err = json.Unmarshal(conf_byte, &json_conf)
// if err != nil {
// 	panic(err)
// }

// router.Router.Run(json_conf["port"])

// //common.KJ()
// //common.Baidu()

// bot, err := tegelam.GetBot()
// if err != nil {
// 	fmt.Println(err.Error())
// }

// data, err := bot.TGetUpdates()
// if err != nil {
// 	fmt.Println(err.Error())
// }

// var sumid []int64
// for _, value := range data {

// 	<-ratelimit
// 	if value.Message != nil {
// 		lock_res := tegelam.RemoveSame(sumid, value.Message.Chat.ID)
// 		fmt.Println("接收消息：", value.Message.Text)
// 		fmt.Println("接收消息：from ", fmt.Sprintf("%+v", value.Message.From))
// 		if lock_res {
// 			continue
// 		}
// 	} else {
// 		continue
// 	}

// 	sumid = append(sumid, value.Message.Chat.ID)
// 	fmt.Println(sumid, value.Message.Chat.ID)

// 	// //发送消息
// 	// err := bot.SendWithMessage(value.Message.Chat.ID, "im TestSendWithMessage")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	//图片发送
// 	//bot.TestSendWithNewPhoto(-1001183275123, "./img/error.jpg", "haha")

// 	//发送Emoji
// 	//bot.TestSendWithDiceWithEmoji(-1001183275123, "🎲")

// 	//回复消息
// 	//bot.TestSendWithMessageReply(value.Message.Chat.ID, 16, "im TestSendWithMessageReply")

// 	//图片回复消息
// 	//bot.TestSendWithNewPhotoReply(-1001183275123, 16, "./img/error.jpg")

// 	//发送图片文件
// 	//bot.TestSendWithNewDocument(1111640192, "./img/error.jpg")

// 	//转发消息
// 	//bot.TestSendWithMessageForward(1111640192, -1001183275123, 21)

// 	//发送图片，加上备注
// 	//bot.TestSendWithNewPhotoWithFileReader(1111640192, "./img/error.jpg", "error.jpg", "Hello , this is test form send pic")

// 	//发送audio 文件
// 	//bot.TestSendWithNewAudio(1111640192, "./img/audio.mp3", "听妈妈的话", "周杰伦")

// 	//发送voice 文件
// 	//bot.TestSendWithNewVoice(1111640192, "./img/voice.ogg")

// 	//发送通讯号码
// 	//bot.TestSendWithContact(1111640192, "123456789", "axb")

// 	//发送地址
// 	//bot.TestSendWithVenue(1111640192, "趴赛", "Arista Place", 14.5055231, 120.9957029)

// 	//发送视频
// 	//bot.TestSendWithNewVideoNote(1111640192, "./img/videonote.mp4", 240)

// 	//显示一的状态
// 	//bot.TestSendChatConfig(1111640192)

// 	//获取用户图片
// 	//bot.TestGetUserProfilePhotos(1252740888, -1001183275123)

// 	//Webhook注册
// 	//bot.TestSetWebhookWithCert("./img/PUBLIC.pem")
// 	//bot.TestSetWebhookWithoutCert()

// }

// //bot.ExampleNewBotAPI()
// // bot.ExampleAnswerInlineQuery()

// // bot.SendWithExistingVideo(BAACAgUAAxkBAAN6X6E1ctHacUUP-WLpopNSotYXYF8AAp0BAAK8bwhVYGqsSrAZZvceBA", "xiha")

// bot.SendWithExistingVideo(1111640192, "BAACAgQAAx0ESLf3vgABBPMHX6KD3ZVMqit7jJJ24o6_R8KKhmwAAicCAALOtx1RZo0S7q69JGseBA", "xixi")

// bot.PlayExampleNewBotAPI()
