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
	appConfig := core.CreateNewHandeler("", "")
	myInts := appConfig.Int("number", 3)
	fmt.Println(*myInts)
	myInts.Set("name", 3)
	fmt.Println(*myInts)

}
