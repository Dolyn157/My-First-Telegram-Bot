package main

import (
	"Telebot/model"
	"Telebot/utils"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const replylogDst = "./msgs"

func main() {
	utils.LogGenerator(replylogDst, "\n----------------------------\n")

	bot1, err := tgbotapi.NewBotAPI("") //  你的 BOT APIID
	if err != nil {
		log.Panic(err)
	}

	bot1.Debug = true

	//log.Printf("Authorized on account %s", bot1.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates := bot1.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		processControl(update, bot1)
	}
}

func processControl(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if !strings.HasPrefix(update.Message.Text, "/") {
		return
	}
	if update.Message.From.IsBot {
		return
	}
	if update.Message.ReplyToMessage != nil {
		return
	}
	if update.Message.Text == "/help@Gwyndolyn_bot" {
		rText := "天气报告 输入/ 城市名称 可以查询近1小时的基本天气情况，目前支持 曼谷 伦敦 大阪 京都 马尼拉 上海 湛江 广州。 API : OpenWeatherMap"

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rText)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	} else {
		rText := weatherProcess(update)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rText)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}
	if update.Message.ReplyToMessage != nil {
		rText := replyProcess(update)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rText)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}
	//记录 Bot Api 的返回Json 数据
	str := utils.SprintJSON(update) + "\n\n"
	utils.LogGenerator(replylogDst, str)
}

func replyProcess(update tgbotapi.Update) string {
	//log.Printf("[%s] %s ; replyed to [%s]", update.Message.From.UserName, update.Message.Text, update.Message.ReplyToMessage.From.UserName)

	senderNickName := fmt.Sprint(update.Message.From.FirstName + update.Message.From.LastName)
	receiverNickName := fmt.Sprint(update.Message.ReplyToMessage.From.FirstName + update.Message.ReplyToMessage.From.LastName)

	action := strings.TrimPrefix(update.Message.Text, "/")
	text := fmt.Sprintf(senderNickName + " " + action + "了 " + receiverNickName + "！")
	return text
}

func weatherProcess(update tgbotapi.Update) string {
	const apiKey string = "appid=" //你的气象APIID
	const apiUrl string = "https://api.openweathermap.org/data/2.5/weather?q="
	var cityText = strings.TrimPrefix(update.Message.Text, "/")
	var reqUrl string

	mcityNames := map[string]string{
		"曼谷":  "Bangkok",
		"马尼拉": "Manila",
		"大阪":  "Osaka",
		"京都":  "Kyoto",
		"伦敦":  "London",
		"上海":  "Shanghai",
		"湛江":  "Zhanjiang",
		"广州":  "Guangzhou",
	}

	var wd model.WeaData
	var returnText string
	reqUrl = apiUrl + mcityNames[cityText] + "&" + "lang=zh_cn" + "&" + apiKey
	data := wd.ApiGetData(reqUrl)
	wd.ParseData(data)

	if wd.ErrMsg == "400" { //利用天气 Api 给我们返回的404信息来通知用户城市是否找到。
		returnText = "未找到您输入的城市或程序还未做名称映射。"
	} else {
		returnText = fmt.Sprintf("%v近1小时天气：  %s , 温度: %.2f℃, 体感: %.2f℃, 湿度: %s％，压力：%s Hpa，风速: %s 米/秒 ", cityText, wd.Des, wd.Tempo, wd.Feels_Like, wd.Humidity, wd.Pressure, wd.WindSpeed)
	}
	fmt.Println("错误代码", wd.ErrMsg)
	return returnText
}
