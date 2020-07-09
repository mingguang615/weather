# Weather

获取中国天气网天气

## Usage
```
func main()  {
    resp,err := new(NewClient).GetWeather("上海市")
    if err != nil{
        panic(err)
    }
    fmt.Println(resp)
}
```

## Installation
```
go get github.com/mingguang615/weather
```