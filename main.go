package main

import (
	"Telebot/model"
	"Telebot/utils"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const replylogDst = "./logs2"

func main() {
	utils.LogGenerator(replylogDst, "\n----------------------------\n")

	bot1, err := tgbotapi.NewBotAPI("填写你的TG Bot Apikey") //
	if err != nil {
		log.Panic(err)
	}

	bot1.Debug = true

	//log.Printf("Authorized on account %s", bot1.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot1.GetUpdatesChan(u)

	for update := range updates {
		processControl(update, bot1)
	}
}

func processControl(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.Text[0] == 47 && update.Message.From.IsBot == false {
		if update.Message.ReplyToMessage == nil {
			if update.Message.Text == "/help@Gwyndolyn_bot" {
				rText := "天气报告 输入/ 城市名称 可以查询近1小时的基本天气情况，目前支持 曼谷 伦敦 大阪 马尼拉 上海 湛江。 API : OpenWeatherMap"

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, rText)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)

			} else {
				rText := weatherProcess(update)

				if rText == "nil" {

				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, rText)
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)

				}

			}
		} else if update.Message.ReplyToMessage != nil {
			rText := replyProcess(update)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, rText)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)

		}
		//记录 Bot Api 的返回Json 数据
		str := utils.SprintJSON(update) + "\n\n"
		utils.LogGenerator(replylogDst, str)
	}

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
	const apiKey string = "appid=填写你的 openweathermap Apikey"
	const apiUrl string = "https://api.openweathermap.org/data/2.5/weather?q="
	var cityName string
	var cityText = strings.TrimPrefix(update.Message.Text, "/")
	var reqUrl string

	switch cityText {
	case "曼谷":
		cityName = "Bangkok"
		reqUrl = apiUrl + cityName + "&" + "lang=zh_cn" + "&" + apiKey

	case "伦敦":
		cityName = "London"
		reqUrl = apiUrl + cityName + "&" + "lang=zh_cn" + "&" + apiKey

	case "马尼拉":
		cityName = "Manila"
		reqUrl = apiUrl + cityName + "&" + "lang=zh_cn" + "&" + apiKey

	case "大阪":
		cityName = "Osaka"
		reqUrl = apiUrl + cityName + "&" + "lang=zh_cn" + "&" + apiKey

	case "上海":
		cityName = "Shanghai"
		reqUrl = apiUrl + cityName + "&" + "lang=zh_cn" + "&" + apiKey

	case "湛江":
		cityName = "Zhanjiang"
		reqUrl = apiUrl + cityName + "&" + "lang=zh_cn" + "&" + apiKey

	default:
		return "nil"
	}

	var wd model.WeaData
	data := wd.ApiGetData(reqUrl)
	//fmt.Printf("%s \n", data)
	wd.ParseData(data)

	returnText := fmt.Sprintf("%v近1小时天气：  %s , 温度: %.2f℃, 体感: %.2f℃, 湿度: %s％，压力：%s Hpa，风速: %s 米/秒 ", cityText, wd.Des, wd.Tempo, wd.Feels_Like, wd.Humidity, wd.Pressure, wd.WindSpeed)
	return returnText
}

/* for update := range updates {
	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "我爱你" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, `我也爱你哦 ヽ(✿ﾟ▽ﾟ)ノ `)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}
*/
