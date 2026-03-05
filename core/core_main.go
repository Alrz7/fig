package core

import (
	"encoding/json"
	"fig/file"
	"log"
	"strings"
	"time"
)

const (
	BaseDir = "./"
	ConfDir = "./config/"
)

type Handeler struct {
	Dir            string    `json:"dir"`
	Name           string    `json:"name"`
	Format         string    `json:"format"`
	LatstTmodified time.Time `json:"last_time_modified"`
	restored       bool
	// IntField       cfInt     `json:"intField"`
	// StringField    cfString  `json:"stringField"`
	Data cField `json:"data"`

	// type, last_modified, etc.
}

// var Handelers = map[string]handeler{}

func CreateNewHandeler(dir, name string) *Handeler {
	format := strings.Split(name, ".")[1]
	if format != "json" {
		log.Panicf("given format `%v` is not Supported by FIG", format)
	}
	hndlr := Handeler{Dir: dir, Name: name, Format: format, restored: false, Data: cField{}}
	// hndlr.PanicRestore()
	return &hndlr
}

func (h *Handeler) Set(key string, newValue any) { // NOTE : need to make a WARNING about NewValue Being a Pointer (&any_value)
	h.Data.Set(key, newValue)
	if h.restored {
		h.Save()
	}
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
		log.Panicf("there was an error while saving `%v` at `%v` : %v", h.Name, h.Dir, err)
	}
}

func (h *Handeler) Restore() error {
	exists, err := file.DoesExist(h.Dir, h.Name)
	if !exists {
		return err
	}
	bytes, err := file.Read(h.Dir, h.Name)
	var tempHandeler Handeler
	err = json.Unmarshal(bytes, &tempHandeler)
	if err != nil {
		return err
	}
	err = marsh(h, &tempHandeler.Data)
	if err != nil {
		return err
	}
	h.restored = true
	return nil

}

func marsh(h *Handeler, data *cField) error {
	for key, val := range *data {
		if _, ok := h.Data[key]; ok {
			// fmt.Println(key, "found!: ")
			b, err := json.Marshal(val)
			if err != nil {
				return err
			}
			err = json.Unmarshal(b, h.Data[key])
			// fmt.Println((h.Data[key]))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *Handeler) PanicRestore(hr *Handeler) {
	if err := h.Restore(); err != nil {
		log.Panicf("there was an error while Restoring `%v` from `%v` : %v", h.Name, h.Dir, err)
	}
}
