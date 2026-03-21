package file

import (
	"os"
)

func Normalized(dir, name string) (string, int) {
	ac := 0
	if dir[0] != '.' {
		dir = "." + dir
		ac++
	}
	if dir[1] != '/' && dir[1:3] != "./" {
		dir = "./" + dir[1:]
		ac++
	}
	if dir[len(dir)-1] != '/' {
		dir = dir + "/"
		ac++
	}
	return dir, ac

}

func Save(b []byte, dir, name string) error {
	var err error
	myfile, _ := os.Create(dir + name)
	defer myfile.Close()
	_, err = myfile.Write(b)
	if err != nil {
		return err
	}
	return nil
}
