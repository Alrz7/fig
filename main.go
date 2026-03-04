package main

import (
	"fig/core"
	"fmt"
)

type myconf struct {
	name string
}

func (c myconf) save() {
	fmt.Println("nothing but save")
}

func main() {

	// savingNumber := core.Int("num", 3)
	// fmt.Println(*savingNumber)
	// savingNumber.Set()
	// fmt.Println(*savingNumber)

	// type person string

	// savingContent1 := map[string]any{"number": person("navid")}
	// saved1 := core.List(savingContent1)
	// fmt.Println(*saved1)
	// fmt.Println((*saved1)["number"])
	// saved1.Set("number", 8)
	// fmt.Println((*saved1)["number"])
	appConfig := core.CreateNewHandeler("./", "testConfig.json")
	secondConfig := core.CreateNewHandeler(core.BaseDir, "secondConfig.json")

	fmt.Println(appConfig)

	fmt.Println(secondConfig)

	// appConfig.String("apiKey", "sgklfsflkflhk6h655jk7gjk5gj6gl345gh6jkg3l4")
	// appConfig.String("name", "navid")

	// type driver struct {
	// 	Name   string `Json:"name"`
	// 	Age    int    `Json:"age"`
	// 	Gender string `Json:"gender"`
	// 	Job    string `Json:"job"`
	// }
	// maria := driver{
	// 	Name:   "maria",
	// 	Age:    25,
	// 	Gender: "female",
	// 	Job:    "uberDriver",
	// }

	// secondConfig.List("newDriver", maria)
	// secondConfig.String("name", "maria")
	// appConfig.PanicSave()
	// secondConfig.PanicSave()
}
