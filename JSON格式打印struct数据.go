package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 将struct中的数据转化为 JSON格式  只会对struct中被导出的字段进行编码（大写）
func main() {
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}
	//初始化
	curiosity := []location{
		{Name: "a", Lat: 1.15454, Long: -1.54541},
		{Name: "b", Lat: 2.25545, Long: -2.45462},
		{Name: "c", Lat: 3.15615, Long: -3.49852},
	}
	bytes, err := json.MarshalIndent(curiosity, "", " ")
	//exitOnError(err)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))

}
