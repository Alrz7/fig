package main

import (
	"fig/core"
	"fmt"
)

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

type Company struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Founded      string         `json:"founded"`
	IsActive     bool           `json:"is_active"`
	Employees    []driver       `json:"employees"`
	Revenue      float64        `json:"revenue"`
	Departments  map[string]int `json:"departments"`
	Headquarters struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"headquarters"`
	// Partners     []map[string]any `json:"partners"`
	// ExtraDetails any              `json:"extra_details"`
}

type company struct {
}

func main() {
	// Uber := Company{}
	// miles := driver{}
	// porche911 := car{}

	miles := driver{
		Name:   "miles",
		Age:    23,
		Gender: "male",
		Job:    "driver",
	}

	Uber := Company{
		ID:          8403453853045,
		Name:        "Uber",
		Founded:     "2001/12/july",
		IsActive:    true,
		Employees:   []driver{miles},
		Revenue:     352094284,
		Departments: map[string]int{"Florida": 4, "Newyork": 3, "vegas": 8},
		Headquarters: struct {
			City  string "json:\"city\""
			State string "json:\"state\""
		}{City: "miami", State: "Florida"},
	}

	porche911 := car{
		Name:  "porche911",
		Model: "porche",
		Year:  2003,
	}
	name := "navid"

	appConfig := core.CreateNewHandeler("./", "testConfig.json")

	initConfig := func() {
		appConfig.Set("Uber", &Uber)
		// appConfig.Set("miles", &miles)
		appConfig.Set("porche911", &porche911)
		appConfig.Set("CEO", &name)
		defer appConfig.PanicRestore(appConfig)
	}
	initConfig()
	fmt.Println(Uber.Employees[0])
	appConfig.Pop("miles")
	// name = "koroush"
	// appConfig.Save()
	// fmt.Println(name)

}
