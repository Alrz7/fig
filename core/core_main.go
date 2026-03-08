package core

import (
	"encoding/json"
	"fig/echo"
	"fig/file"
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
	restored       bool      // blocking Set() functions from saving datas to the File before Restoring
	Data           cField    `json:"data"`

	// IntField       cfInt     `json:"intField"`
	// StringField    cfString  `json:"stringField"`
	// type, last_modified, etc.
}

var logger = echo.DefultLogger

// var Handelers = map[string]handeler{}

func CreateNewHandeler(dir, name string) *Handeler {
	format := strings.Split(name, ".")[1]
	if format != "json" {
		logger.Errort("given format `%v` is not Supported by FIG", format)
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
func (h *Handeler) Pop(key string) any {
	return h.Data.Pop(key)
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
		logger.Error(err, "there was an error while saving `%v` at `%v`", h.Name, h.Dir)
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
	newObject := needToSave(h, &tempHandeler.Data)
	if newObject {
		h.Save()
	}
	return nil

}

// if a new data was added , there should be a call to Save() after Restoring
func needToSave(h *Handeler, data *cField) bool {
	for key := range h.Data {
		if _, ok := (*data)[key]; !ok {
			// fmt.Println("there is a new Object : ", key)
			return true
		}
	}
	return false
}

func marsh(h *Handeler, data *cField) error {
	for key, val := range *data {
		if _, ok := h.Data[key]; ok { // NOTE: IF OK was False: we should decode to either Pop that Object or create a warning
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
		} else {
			logger.Errort("Not All of %v's parameters were declared in your Application: lost `%v`", h.Name, key)
		}
	}
	return nil
}

func (h *Handeler) PanicRestore(hr *Handeler) {
	if err := h.Restore(); err != nil {
		logger.Error(err, "there was an error while Restoring `%v` from `%v`", h.Name, h.Dir)
	}
}
