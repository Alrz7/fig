package save

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	Num int `json:"num"`
}

func (p person) adder() *int {
	return &p.Num
}
func main() {
	p1 := person{Num: 15}
	// pt := p1.adder()
	// fmt.Println(pt)
	// fmt.Println(*pt)
	jsn, _ := json.Marshal(p1)

	sl := person{}

	json.Unmarshal(jsn, &sl)

	fmt.Printf("%+v \n", sl)
	saveFile(jsn)

}

func saveFile(b []byte) {
	myfile, _ := os.Create("./tests/starting.json")
	defer myfile.Close()
	n, err := myfile.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
