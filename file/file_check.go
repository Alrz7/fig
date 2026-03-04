package file

import (
	"os"
)

func DoesExist(dir, name string) (bool, error) {
	_, err := os.Stat(dir + name)
	if os.IsNotExist(err) { // file doesn't exist
		return false, err
	} else if err != nil { // another path error
		return false, err
	}
	return true, nil
}

func Read(dir, name string) ([]byte, error) {
	var err error
	myfile, _ := os.Open(dir + name)
	defer myfile.Close()

	info, _ := myfile.Stat()
	b := make([]byte, info.Size())

	_, err = myfile.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}
