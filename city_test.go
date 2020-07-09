package weather

import (
	"fmt"
	"testing"
)

func TestCity(t *testing.T) {
	fmt.Println(City["上海"]["上海"])
	fmt.Println(City["北京"]["北京"])
	fmt.Println(City["安徽省"]["合肥"])
}
