package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/buger/jsonparser"
)

type WeaData struct {
	Des        string
	Tempo      float64
	Feels_Like float64
	Pressure   string
	Humidity   string
	WindSpeed  string
	ErrMsg     string
}

func (Wd *WeaData) ApiGetData(apiUrl string) []byte {
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	return data
}

func (Wd *WeaData) ParseData(data []byte) {
	//weather

	Wd.Des, _ = jsonparser.GetString(data, "weather", "[0]", "description")
	Tempo, _, _, _ := jsonparser.Get(data, "main", "temp")            //气温
	Feels_Like, _, _, _ := jsonparser.Get(data, "main", "feels_like") //体感
	Pressure, _, _, _ := jsonparser.Get(data, "main", "pressure")     //气压 hpa
	Humidity, _, _, _ := jsonparser.Get(data, "main", "humidity")     //湿度 hpa
	WindSpeed, _, _, _ := jsonparser.Get(data, "wind", "speed")       // 风速
	errMsg, _, _, _ := jsonparser.Get(data, "cod")

	strTempo := string(Tempo)
	strFeels := string(Feels_Like)

	Wd.Tempo, _ = strconv.ParseFloat(strTempo, 32)
	Wd.Feels_Like, _ = strconv.ParseFloat(strFeels, 32)
	Wd.Pressure = string(Pressure)
	Wd.Humidity = string(Humidity)
	Wd.WindSpeed = string(WindSpeed)

	Wd.Tempo = Wd.Tempo - 273.15
	Wd.Feels_Like = Wd.Feels_Like - 273.15
	Wd.ErrMsg = string(errMsg)
}
