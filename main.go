package main

import (
	"fig/core"
	"fmt"
)

type Config interface {
	driver | car
}

type driver struct {
	Name   string `Json:"name"`
	Age    int    `Json:"age"`
	Gender string `Json:"gender"`
	Job    string `Json:"job"`
}

type car struct {
	Name  string `Json:"name"`
	Model string `Json:"model"`
	Year  int    `Json:"year"`
}

func main() {
	// miles := driver{
	// 	Name:   "miles",
	// 	Age:    23,
	// 	Gender: "male",
	// 	Job:    "driver",
	// }

	// porche911 := car{
	// 	Name:  "porche911",
	// 	Model: "porche",
	// 	Year:  2003,
	// }

	miles := driver{}
	porche911 := car{}

	appConfig := core.CreateNewHandeler("./", "testConfig.json")
	appConfig.Set("miles", &miles)
	appConfig.Set("porche911", &porche911)
	appConfig.PanicRestore(appConfig)
	fmt.Println(miles.Age)

	// appConfig.Save()
}
