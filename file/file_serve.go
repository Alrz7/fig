package file

import (
	"os"
)

func Save(b []byte, dir, name string) error {
	var err error
	myfile, _ := os.Create(dir + name + ".json")
	defer myfile.Close()
	_, err = myfile.Write(b)
	if err != nil {
		return err
	}
	return nil
}
