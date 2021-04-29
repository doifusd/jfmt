package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

var JsonStr string
var data map[string]interface{}

func init() {
	fmt.Println("Usage:")
	fmt.Println(" jfmt -json='json字符串'")
	fmt.Println("Result:")
	flag.StringVar(&JsonStr, "json", "", "json字符串")
}

func main() {
	flag.Parse()
	if len(JsonStr) > 1 {
		if err := json.Unmarshal([]byte(JsonStr), &data); err != nil {
			fmt.Println("格式错误")
		}
		js, _ := json.MarshalIndent(data, " ", " ")
		fmt.Printf("%c[%dm%s%c[0m", 0x1B, 32, string(js), 0x1B)
	} else {
		fmt.Println("请按照Usage正确使用!!!")
	}
}
