package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const api = "http://wthrcdn.etouch.cn/weather_mini?city=%s"

type WeatherDesc struct {
	Date      string //日期
	High      string //高温
	Fengli    string //风力
	Low       string //低温
	Fengxiang string //风向
	Type      string //天气类型
}

type WeatherResult struct {
	Data struct {
		Yesterday *WeatherDesc   `json:"yesterday"` //昨日天气
		City      string         `json:"city"`      //城市
		Forecast  []*WeatherDesc `json:"forecast"`  //
		Ganmao    string         `json:"ganmao"`
		Wendu     string         `json:"wendu"`
	} `json:"data"`
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

type NewClient int

//获取天气
func (c *NewClient) GetWeather(city string) (*WeatherResult, error) {
	resp, err := http.Get(fmt.Sprintf(api, city))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	weather := new(WeatherResult)
	if err := json.Unmarshal(data, weather); err != nil {
		return nil, err
	}
	return weather, nil
}
