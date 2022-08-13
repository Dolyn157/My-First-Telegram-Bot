#朵琳第一个 Telegram Bot
  这是我编写的第一个  Telegram Bot，主要有以下两个功能

  回复他人时输入 "/" + "任意字符" 并发送， bot 会发出 “回复者 任意字符了 被回复者” 的信息。 例如A用户在频道里向B用户发送 /拍 ，则bot回复 "a 拍了 b"
  输入"/" + 城市名称 ， 则bot回复该城市的天气、温度、湿度等信息。目前支持 曼谷 伦敦 大阪 马尼拉 上海 湛江。

  ## 使用到的一些库和 API
  天气 API：https://openweathermap.org/api
  json 解析库： "github.com/buger/jsonparser"