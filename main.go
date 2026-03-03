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
	myInts := appConfig.Int("number", 3)
	fmt.Println(*myInts)
	myInts.Set("name", 3)
	fmt.Println(*myInts)

	//--

	type myconf struct {
		name string
		num  int
	}
	myStruct := myconf{name: "navid", num: 3}
	strcuts := appConfig.List("conf1", myStruct)

	fmt.Println(*strcuts)
	fmt.Println((*strcuts)["conf1"])
	fmt.Println(appConfig)
	appConfig.PanicSave()

}
