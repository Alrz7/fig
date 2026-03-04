package core

import (
	"encoding/json"
	"fig/file"
	"log"
	"time"
)

const (
	BaseDir = "./"
	ConfDir = "./config/"
)

type Handeler struct {
	Dir        string   `json:"dir"`
	Name       string   `json:"name"`
	IntData    cfInt    `json:"intData"`
	StringData cfString `json:"stringData"`
	ListData   cflist   `json:"listData"`
	// type, last_modified, etc.
}

// var Handelers = map[string]handeler{}

func CreateNewHandeler(dir, name string) *Handeler {
	hndlr := Handeler{Dir: dir, Name: name, IntData: cfInt{}, StringData: cfString{}, ListData: cflist{}}
	hndlr.PanicRestore()
	return &hndlr
}

func (h *Handeler) Save() error {
	bytes, err := json.Marshal(h)
	if err != nil {
		return err
	}
	err = file.Save(bytes, h.Dir, h.Name)
	if err != nil {
		return err
	}
	return nil

}

func (h *Handeler) PanicSave() {
	if err := h.Save(); err != nil {
		log.Panicf("%v : there was an error while saving `%v` at `%v` : %v", time.Now(), h.Name, h.Dir, err)
	}
}

func (h *Handeler) Restore() error {
	// var res Handeler
	exists, err := file.DoesExist(h.Dir, h.Name)
	if !exists {
		return err
	}
	bytes, err := file.Read(h.Dir, h.Name)
	err = json.Unmarshal(bytes, h)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handeler) PanicRestore() {
	if err := h.Restore(); err != nil {
		log.Panicf("%v : there was an error while Restoring `%v` from `%v` : %v", time.Now(), h.Name, h.Dir, err)
	}
}

// we do the initializayions here
// we make a manager to manage every file and it makes handeler for each
// we make a manager to manage the encoding and decoding all the configs for each files
// there should be a Global Get function to search all and a Specified get function for each type
