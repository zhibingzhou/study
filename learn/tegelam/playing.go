package tegelam

import (
	"encoding/json"
	"fmt"
	"log"
)

//监听回复信息
func (bot *BotAPI) PlayExampleNewBotAPI() {

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
	}

	// // Optional: wait for updates and clear them if you don't want to handle
	// // a large backlog of old messages
	// time.Sleep(time.Millisecond * 500)
	// updates.Clear()

	for update := range updates {

		if update.Message.Video != nil {
			re, _ := json.Marshal(update.Message.Video)
			fmt.Println(string(re))
		}

		if update.ChannelPost != nil { //频道消息

			count_number, _ := bot.GetChatMembersCount(ChatConfig{ChatID: update.ChannelPost.Chat.ID})
			msg := NewMessage(update.ChannelPost.Chat.ID, NewFontStyle("我是频道：", "b")+fmt.Sprintf("这个群一共有 %d 人", count_number))
			msg.ReplyToMessageID = update.ChannelPost.MessageID
			msg.ParseMode = "HTML"
			bot.Send(msg)

			continue
		}

		if update.InlineQuery == nil { // 内敛消息

		} else {
			inlineConf := InlineConfig{
				InlineQueryID: update.InlineQuery.ID,
				IsPersonal:    true,
				CacheTime:     0,
				Results:       []interface{}{},
			}

			if update.InlineQuery.Query == "love" {
				article := NewInlineQueryResultArticle(update.InlineQuery.ID, "Surprise", "I love you so much Baby ", "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1603891948961&di=c3c7fc07cc8cee03b0789718f92b409a&imgtype=0&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2F4%2F53ace79df303f.jpg")
				article.Description = update.InlineQuery.Query

				inlineConf.Results = append(inlineConf.Results, article)
			}

			// if update.InlineQuery.Query == "jewel" {

			// }

			if update.InlineQuery.Query == "img" {
				articlephoto := NewInlineQueryResultPhoto(update.InlineQuery.ID, "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1603891948961&di=c3c7fc07cc8cee03b0789718f92b409a&imgtype=0&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2F4%2F53ace79df303f.jpg", "https://seopic.699pic.com/photo/50139/0661.jpg_wh1200.jpg")
				articlephoto.Description = update.InlineQuery.Query
				inlineConf.Results = append(inlineConf.Results, articlephoto)
			}

			// articles := NewInlineQueryResultArticle(update.InlineQuery.ID, "love", "nimabi")
			// articles.Description = update.InlineQuery.Query

			if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
				log.Println(err)
			}
			// KeyboardButton()
			// NewReplyKeyboard()
			continue
		}

		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "funny_time":
			bot.PlaySendWithNewVideo(update.Message.Chat.ID, "./play/jewel.MP4", "My Baby")
			continue
		case "photo":
			bot.PlaySendWithMediaGroup(update.Message.Chat.ID)
			continue
		case "Calculator":
			bot.MakeCalculator(update.Message.Chat.ID)
			continue
		case "test":
			fontb := NewFontStyle("我是粗体", "b")
			fonti := NewFontStyle("我是斜体", "i")
			fontu := NewFontStyle("我是下划线", "u")
			fonts := NewFontStyle("我是删除线文字", "u")
			jump1 := NewJumpHTMLString("https://t.me/test_001kon", "跳转")
			showmessage := fontb + fonti + fontu + fonts + jump1 + "🤣"
			bot.SendMessageWithVideo(update.Message.Chat.ID, showmessage, "./play/jewel.MP4")
			continue
		case "/shaizi":
			bot.SendWithDiceWithEmoji(update.Message.Chat.ID, "🎲")
			continue
		case "/biao":
			bot.SendWithDiceWithEmoji(update.Message.Chat.ID, "🎯")
			continue
		case "/lanqiu":
			bot.SendWithDiceWithEmoji(update.Message.Chat.ID, "🏀")
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		result, err := bot.GetStickerSet(GetStickerSetConfig{Name: update.Message.Text})

		emoji := "🙈"
		if len(result.Stickers) > 0 {
			for _, value := range result.Stickers {
				emoji = value.Emoji + "\n"
			}

		}
		count_number, _ := bot.GetChatMembersCount(ChatConfig{ChatID: update.Message.Chat.ID})
		msg := NewMessage(update.Message.Chat.ID, NewFontStyle("你搜索的表情结果是：", "b")+"\n"+emoji+"\n"+NewFontStyle("你还可以发送下面的词汇 ： "+"\n"+"\n funny_time"+"\n photo"+"\n Calculator\n"+fmt.Sprintf("这个群一共有 %d 人", count_number), "i"))
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "HTML"

		//keyboard := NewKeyboardButtonRow(KeyboardButton{Text: "测试1", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "测试2", RequestContact: false, RequestLocation: false})
		//msg.ReplyMarkup = NewReplyKeyboard(keyboard) //自定义键盘 可以 发送联系方式 和位置
		//msg.ReplyMarkup = NewOneTimeReplyKeyboard(keyboard)

		//三种键盘
		// keyboard := NewInlineKeyboardRow(NewInlineKeyboardButtonData("显示数据", "funny_time"), NewInlineKeyboardButtonURL("跳转链接按键", "http://www.baidu.com"), NewInlineKeyboardButtonSwitch("转发内容机器人", "img"))
		// msg.ReplyMarkup = NewInlineKeyboardMarkup(keyboard)

		msg.ReplyMarkup = ReplyKeyboardRemove{
			RemoveKeyboard: true,
			Selective:      false,
		}
		//更新群组头像
		//bot.SetPhotoforchat(update.Message.Chat.ID)

		fmt.Println(result, err)
		bot.Send(msg)
	}
}

//组发送网络图片信息或者视频
func (bot *BotAPI) PlaySendWithMediaGroup(ChatID int64) error {

	cfg := NewMediaGroup(ChatID, []interface{}{
		// NewInputMediaVideo("https://photos.google.com/photo/AF1QipNPhjq5_2qZIyOoZgukayI7t2SUyr4ZCF14gJbf"),
		NewInputMediaPhoto("https://lh3.googleusercontent.com/o9GGh9600FW5dTtSjmIZVM2BYJ9SwJ0SMSKlzli7wHvLyo8TuhKdnkkKV0p_VnLia0ACvWYyChTa4UpLr0SjoCyBfuX1foXK3lezfPOwuel9Zrp7Z2EFn8A1Nbuc1UAG10r0lcuULFYFcreH7wgLAhft9Dkge95aonD8urmi9qeoXxXtc6Nf0YZ1q9gujXz3IMD_0_gImLRj9Ssao_f-B58eXYrLd_ttd7etijSNQhXNCDTiwyxbaXDh34SB9CC0fR9-nK_AmWGQZlntmclZPW2U7ZytSBJr6WXM7s-YA21H6pPFxKjrOMwoxYdttFzYgN4Dce9EOJtkIPwFtV7lJS_hcpocPDQCQyetlQseGWM_AyXWLqpDNkOwpWuhdKQYOqwAR5utSKbtClOUtXt1wSFgJ_QfV44qEc4AbZvetBzCsbuDFicEObbozOfbPH9ZDFt2-Sg6zTKg-8NdrG61vaNpBm8xnqYmrnNSX0MoIyX6Ro36PftZRHlW5jhF2M57LHrTT6VkDtcvu01_nXx8YIwhCA7V2PITlZcp62IXzXoEGEA7LZuFQSScT0q76gdi6qQcCkVV-X9F-nUDbEsI26QrMncZSB3wQNV0OolT3llsEwHlJnJUsgmMoWi9udDes8fgZtkP9E0mru8CnHt2bk6rAxm_vH0CoUVaDxe8eWtfzyZX0P2FUAPms45a=w211-h238-no?authuser=0"),
	})
	_, err := bot.Send(cfg)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//发送视频
/*
ChatID  被转发消息的聊天室
video 视频
note 标题
*/
func (bot *BotAPI) PlaySendWithNewVideo(ChatID int64, video, note string) error {

	msg := NewVideoUpload(ChatID, video)
	msg.Duration = 10
	msg.Caption = note

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (bot *BotAPI) MakeCalculator(ChatID int64) {

	msg := NewMessage(ChatID, "这个是自定义键盘")
	msg.ParseMode = "HTML"

	keyboard := NewKeyboardButtonRow(KeyboardButton{Text: "7", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "8", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "9", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "+", RequestContact: false, RequestLocation: false})
	keyboard1 := NewKeyboardButtonRow(KeyboardButton{Text: "4", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "5", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "6", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "-", RequestContact: false, RequestLocation: false})
	keyboard2 := NewKeyboardButtonRow(KeyboardButton{Text: "1", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "2", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "3", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "*", RequestContact: false, RequestLocation: false})
	keyboard3 := NewKeyboardButtonRow(KeyboardButton{Text: "0", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: ".", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "=", RequestContact: false, RequestLocation: false}, KeyboardButton{Text: "/", RequestContact: false, RequestLocation: false})
	msg.ReplyMarkup = NewReplyKeyboard(keyboard, keyboard1, keyboard2, keyboard3) //自定义键盘 可以 发送联系方式 和位置

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

}

func (bot *BotAPI) SetPhotoforchat(ChatID int64) {

	re, _ := bot.SetChatPhoto(SetChatPhotoConfig{BaseFile: BaseFile{File: "./img/error.jpg", BaseChat: BaseChat{ChatID: ChatID}}})
	fmt.Println(re)
}
