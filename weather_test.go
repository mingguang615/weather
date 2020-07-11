package weather

import (
	"fmt"
	"testing"
)

func TestNew_GetWeatherCity(t *testing.T) {
	cli := new(New)
	resp, err := cli.GetWeatherCityName("上海市")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	citykey, _ := CityKey["安徽省"]["合肥"]
	resp, err = cli.GetWeatherCityKey(citykey)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
