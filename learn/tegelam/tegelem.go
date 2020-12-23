package tegelam

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//var Token = "1153713930:AAHh747QK511gVqSwtTbPIvPKm9ExdgX8pA"

var Token = "1434867188:AAG-I6ww1gywCmqiWiUygL53OdoLdsLiO7c"

//获取Bot 信息
func GetBot() (*BotAPI, error) {
	bot, err := NewBotAPI(Token)

	if err != nil {
		fmt.Println(err)
	}

	return bot, err
}

func RemoveSame(data []int64, id int64) bool {
	re := false
	for _, value := range data {
		if value == id {
			return true
		}
	}
	return re
}

//获取信息
func (bot *BotAPI) TGetUpdates() ([]Update, error) {

	u := NewUpdate(0)

	get_data, err := bot.GetUpdates(u)

	if err != nil {
		fmt.Println(err)
	}

	return get_data, err
}

//发送消息

func (bot *BotAPI) SendWithMessage(ChatID int64, message string) error {
	msg := NewMessage(ChatID, message)
	msg.ParseMode = "markdown" //格式
	// keyboard := NewKeyboardButtonRow(KeyboardButton{Text: "测试", RequestContact: true, RequestLocation: true})
	// msg.ReplyMarkup = NewReplyKeyboard(keyboard) //键盘
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送超链接消息
func (bot *BotAPI) SendMessageWithVideo(ChatID int64, message, video string) error {

	msg := NewVideoUpload(ChatID, video)
	msg.ParseMode = "HTML"
	msg.Caption = message
	//三种键盘
	keyboard := NewInlineKeyboardRow(NewInlineKeyboardButtonData("显示数据", "funny_time"))
	keyboard2 := NewInlineKeyboardRow(NewInlineKeyboardButtonURL("跳转链接按键", "http://www.baidu.com"))
	keyboard3 := NewInlineKeyboardRow(NewInlineKeyboardButtonSwitch("转发内容机器人", "img"))
	msg.ReplyMarkup = NewInlineKeyboardMarkup(keyboard, keyboard2, keyboard3)

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//回复消息
/*
ChatID 聊天室id
ReplyToMessageID 回复  对应消息的id
*/
func (bot *BotAPI) SendWithMessageReply(ChatID int64, ReplyToMessageID int, message string) error {

	msg := NewMessage(ChatID, message)
	msg.ReplyToMessageID = ReplyToMessageID
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//图片回复
/*
ChatID 聊天室id
ReplyToMessageID 回复  对应消息的id
*/
func (bot *BotAPI) SendWithNewPhotoReply(ChatID int64, ReplyToMessageID int, filename string) error {

	msg := NewPhotoUpload(ChatID, filename)
	msg.ReplyToMessageID = ReplyToMessageID

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//图片发送
/*
ChatID  发消息的聊天室
imgfile 图片地址
note 图片说明
*/
func (bot *BotAPI) SendWithNewPhoto(ChatID int64, imgfile, note string) error {

	msg := NewPhotoUpload(ChatID, imgfile)
	msg.Caption = note

	//带上自己的信息
	// msg.ReplyMarkup = ReplyKeyboardRemove{
	// 	RemoveKeyboard: true,
	// 	Selective:      false,
	// }
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//图片发送 字节发送
/*
ChatID  发消息的聊天室
imgfile 图片地址
file_name 文件名
note 图片说明
*/
func (bot *BotAPI) SendWithNewPhotoWithFileBytes(ChatID int64, imgfile, file_name, note string) error {

	data, _ := ioutil.ReadFile(imgfile)
	b := FileBytes{Name: "file_name", Bytes: data}

	msg := NewPhotoUpload(ChatID, b)
	msg.Caption = note
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//图片发送 流的方式发送
/*
ChatID  发消息的聊天室
imgfile 图片地址
file_name 文件名
note 图片说明
*/
func (bot *BotAPI) SendWithNewPhotoWithFileReader(ChatID int64, imgfile, file_name, note string) error {

	f, _ := os.Open(imgfile)
	reader := FileReader{Name: file_name, Reader: f, Size: -1}

	msg := NewPhotoUpload(ChatID, reader)
	msg.Caption = note
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//转发图片 （有图片的id）
/*
ChatID  发消息的聊天室
ExistingPhotoFileID 图片的id
note 图片备注
*/
func (bot *BotAPI) SendWithExistingPhoto(ChatID int64, ExistingPhotoFileID, note string) error {

	msg := NewPhotoShare(ChatID, ExistingPhotoFileID)
	msg.Caption = note
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//消息转发
/*
ChatID  被转发消息的聊天室
fromChatID 要转发的消息聊天室id
ReplyToMessageID 转发的消息id
*/
func (bot *BotAPI) SendWithMessageForward(chatID, fromChatID int64, ReplyToMessageID int) error {

	msg := NewForward(chatID, fromChatID, ReplyToMessageID)
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err

}

//发送图片文件
/*
ChatID  被转发消息的聊天室
imgfile 图片路径
*/
func (bot *BotAPI) SendWithNewDocument(ChatID int64, imgfile string) error {

	msg := NewDocumentUpload(ChatID, imgfile)
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//发送图片文件（有 id 的文件）
/*
ChatID  被转发消息的聊天室
ExistingDocumentFileID 发送的文件id
*/
func (bot *BotAPI) SendWithExistingDocument(ChatID int64, ExistingDocumentFileID string) error {

	msg := NewDocumentShare(ChatID, ExistingDocumentFileID)
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err

}

//发送声音文件
/*
ChatID  被转发消息的聊天室
audiofile 发送的声音文件
title 歌名
Performer 歌手
*/
func (bot *BotAPI) SendWithNewAudio(ChatID int64, audiofile, title, Performer string) error {

	msg := NewAudioUpload(ChatID, audiofile)
	msg.Title = title
	msg.Duration = 10
	msg.Performer = Performer
	msg.MimeType = "audio/mpeg"
	msg.FileSize = 688
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//发送声音文件（有 id 的文件）
/*
ChatID  被转发消息的聊天室
ExistingDocumentFileID 发送的文件id
*/
func (bot *BotAPI) SendWithExistingAudio(ChatID int64, ExistingAudioFileID, title, Performer string) error {

	msg := NewAudioShare(ChatID, ExistingAudioFileID)
	msg.Title = title
	msg.Duration = 10
	msg.Performer = Performer

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//发送声音文件
/*
ChatID  被转发消息的聊天室
voicefile 发送的声音文件
*/

func (bot *BotAPI) SendWithNewVoice(ChatID int64, voicefile string) error {

	msg := NewVoiceUpload(ChatID, voicefile)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//发送声音文件（有 id 的文件）
/*
ChatID  被转发消息的聊天室
ExistingVoiceFileID 发送的声音文件
*/
func (bot *BotAPI) SendWithExistingVoice(ChatID int64, ExistingVoiceFileID string) error {
	msg := NewVoiceShare(ChatID, ExistingVoiceFileID)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

//发送通讯号码
/*
ChatID  被转发消息的聊天室
number 号码
name 名字
*/
func (bot *BotAPI) SendWithContact(ChatID int64, number, name string) error {

	contact := NewContact(ChatID, number, name)
	_, err := bot.Send(contact)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送位置
/*
ChatID  被转发消息的聊天室
x 经度
y 纬度
*/
func (bot *BotAPI) SendWithLocation(ChatID int64, x, y float64) error {

	_, err := bot.Send(NewLocation(ChatID, x, y))

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送位置
/*
ChatID  被转发消息的聊天室
x 经度
y 纬度
location 地址
address 街道
*/
func (bot *BotAPI) SendWithVenue(ChatID int64, location, address string, x, y float64) error {

	venue := NewVenue(ChatID, location, address, x, y)
	_, err := bot.Send(venue)

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
func (bot *BotAPI) SendWithNewVideo(ChatID int64, video, note string) error {

	msg := NewVideoUpload(ChatID, video)
	msg.Duration = 10
	msg.Caption = note

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送视频（有 id 的文件）
/*
ChatID  被转发消息的聊天室
ExistingVideoFileID 视频id
note 标题
*/

func (bot *BotAPI) SendWithExistingVideo(ChatID int64, ExistingVideoFileID, note string) error {

	msg := NewVideoShare(ChatID, ExistingVideoFileID)
	msg.Duration = 10
	msg.Caption = note

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送视频留言
/*
ChatID  被转发消息的聊天室
video 视频
length 大小
*/
func (bot *BotAPI) SendWithNewVideoNote(ChatID int64, video string, length int) error {

	msg := NewVideoNoteUpload(ChatID, length, video)
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送视频留言（有 id 的文件）
/*
ChatID  被转发消息的聊天室
ExistingVideoNoteFileID 视频id
length 大小
*/

func (bot *BotAPI) SendWithExistingVideoNote(ChatID int64, ExistingVideoNoteFileID string, length int) error {

	msg := NewVideoNoteShare(ChatID, 240, ExistingVideoNoteFileID)
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//发送emjoi
// Dice can have values 1-6 for “🎲” and “🎯”, and values 1-5 for “🏀”.
func (bot *BotAPI) SendWithDiceWithEmoji(ChatID int64, Emoji string) error {

	msg := NewDiceWithEmoji(ChatID, Emoji)
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println(err)
	}

	return err

}

//发送状态
/*
 */
func (bot *BotAPI) SendChatConfig(ChatID int64) error {

	_, err := bot.Send(NewChatAction(ChatID, ChatFindLocation))

	if err != nil {
		fmt.Println(err)
	}

	return err

}

//编辑信息，没什么用
func (bot *BotAPI) SendEditMessage(ChatID int64, message string) error {

	msg, err := bot.Send(NewMessage(ChatID, "ing editing."))
	if err != nil {
		fmt.Println(err)
	}

	edit := EditMessageTextConfig{
		BaseEdit: BaseEdit{
			ChatID:    ChatID,
			MessageID: msg.MessageID,
		},
		Text: "Updated text.",
	}

	_, err = bot.Send(edit)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//获取用户的头像
//可以get到三个不同的文件,不同大小，可以自己选
func (bot *BotAPI) TGetUserProfilePhotos(userID int, chatID int64) error {

	user_photo, err := bot.GetUserProfilePhotos(NewUserProfilePhotos(userID))

	if err != nil {
		fmt.Println(err)
	}
	for _, value := range user_photo.Photos {
		for _, values := range value {
			bot.SendWithExistingPhoto(chatID, values.FileID, "111")

		}
	}
	return err
}

//WebHook 注册
//有秘钥注册
func (bot *BotAPI) SetWebhookWithCert(file_name string) {

	time.Sleep(time.Second * 2)

	bot.RemoveWebhook()

	wh := NewWebhookWithCert("https://example.com/tgbotapi-/"+bot.Token, file_name)
	_, err := bot.SetWebhook(wh)
	if err != nil {
		fmt.Println(err)
	}
	_, err = bot.GetWebhookInfo()
	if err != nil {
		fmt.Println(err)
	}
	bot.RemoveWebhook()
}

// WebHook
//无秘钥注册
func (bot *BotAPI) SetWebhookWithoutCert() {

	time.Sleep(time.Second * 2)

	bot.RemoveWebhook()

	wh := NewWebhook("https://example.com/tgbotapi-/" + bot.Token)
	_, err := bot.SetWebhook(wh)
	if err != nil {
		fmt.Println(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		fmt.Println(err)
	}
	if info.MaxConnections == 0 {
		fmt.Println("Expected maximum connections to be greater than 0")
	}
	if info.LastErrorDate != 0 {
		fmt.Println("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	bot.RemoveWebhook()
}

//挂载监听
func (bot *BotAPI) UpdatesChan() {

	var ucfg UpdateConfig = NewUpdate(0)
	ucfg.Timeout = 60
	abc, err := bot.GetUpdatesChan(ucfg)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(abc)
}

//组发送网络图片信息或者视频
func (bot *BotAPI) SendWithMediaGroup(ChatID int64) {

	cfg := NewMediaGroup(ChatID, []interface{}{
		NewInputMediaPhoto("https://github.com/go-telegram-bot-api/telegram-bot-api/raw/0a3a1c8716c4cd8d26a262af9f12dcbab7f3f28c/s/image.jpg"),
		NewInputMediaVideo("https://github.com/go-telegram-bot-api/telegram-bot-api/raw/0a3a1c8716c4cd8d26a262af9f12dcbab7f3f28c/s/video.mp4"),
		NewInputMediaPhoto("https://seopic.699pic.com/photo/50139/0661.jpg_wh1200.jpg"),
	})
	_, err := bot.Send(cfg)
	if err != nil {
		fmt.Println(err)
	}
}

//监听回复信息

func (bot *BotAPI) ExampleNewBotAPI() {

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
	}
	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		bot.SendWithMessageForward(update.Message.Chat.ID, 1111640192, 250)

		msg := NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

// webhook 监听回复信息
func (bot *BotAPI) ExampleNewWebhook() {

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err := bot.SetWebhook(NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "./img/PUBLIC.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "./img/PUBLIC.pem", "./img/PRIVATE.key", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func (bot *BotAPI) ExampleWebhookHandler() {

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err := bot.SetWebhook(NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "./img/PUBLIC.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	http.HandleFunc("/"+bot.Token, func(w http.ResponseWriter, r *http.Request) {
		update, err := bot.HandleUpdate(r)
		if err != nil {
			log.Printf("%+v\n", err.Error())
		} else {
			log.Printf("%+v\n", *update)
		}
	})

	http.ListenAndServeTLS("0.0.0.0:8443", "./img/PUBLIC.pem", "./img/PRIVATE.key", nil)
}

//启用内联
func (bot *BotAPI) ExampleAnswerInlineQuery() {

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	fmt.Println(err)

	for update := range updates {
		if update.InlineQuery == nil { // if no inline query, ignore it
			continue
		}

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
	}
}

//启用内联
func (bot *BotAPI) GetChatInformation(ChatID int64, SuperGroupUsername string) {

	Chat, err := bot.GetChat(ChatConfig{ChatID: ChatID, SuperGroupUsername: SuperGroupUsername})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(Chat)

}
