package file

import (
	"os"
)

type person struct {
	Num int `json:"num"`
}

func (p person) adder() *int {
	return &p.Num
}

// func main() {
// 	// p1 := person{Num: 15}
// 	// // pt := p1.adder()
// 	// // fmt.Println(pt)
// 	// // fmt.Println(*pt)
// 	// jsn, _ := json.Marshal(p1)

// 	// sl := person{}

// 	// json.Unmarshal(jsn, &sl)

// 	// fmt.Printf("%+v \n", sl)
// 	// saveFile(jsn)

// }

func DoesExist(dir, name string) (bool, error) {
	_, err := os.Stat(dir + name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func Save(b []byte, dir, name string) error {
	// exist, err := DoesExist(dir, name)
	// if !exist {
	// 	return err
	// }
	var err error
	myfile, _ := os.Create(dir + name)
	defer myfile.Close()

	_, err = myfile.Write(b)
	if err != nil {
		return err
	}
	return nil
}
