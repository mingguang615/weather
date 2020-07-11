package weather

import (
	"fmt"
	"testing"
)

func TestGetCitKey(t *testing.T) {
	fmt.Println(GetCitKey("上海市", "上海"))
	fmt.Println(GetCitKey("安徽省", "合肥市"))
	fmt.Println(GetCitKey("安徽省", "合肥"))
	fmt.Println(GetCitKey("安徽", "合肥市"))
	fmt.Println(GetCitKey("安徽", "合肥"))
}
