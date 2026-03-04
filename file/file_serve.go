package file

import (
	"os"
)

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
