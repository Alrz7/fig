package fig

import (
	"encoding/json"
	"time"

	"github.com/Alrz7/fig/echo"
	"github.com/Alrz7/fig/file"
)

const (
	DefultDir string = "./config/"
	Json      string = ".json"
)

type Handeler struct {
	BaseDir          string    `json:"dir"` // handeler might have diferent Info Datas in future so i decided to not use the Fieldinfo here
	Name             string    `json:"name"`
	BaseFormat       string    `json:"format"`
	LastModification time.Time `json:"last_time_modified"`

	FieldsInfo map[string]*FieldInfo `json:"fileds_info"`
	Fields     map[string]*Field     `json:"-"`
}

var logger = echo.DefultLogger

// Handeler is gonna be the first building block of your config (your config File btw :) (dir-string: ./foo1/foo2/ , name-string: [HandelerName].HandelerType)
// FIG only supportes Json yet so the name is going to be [HandelerName].json
func CreateNewHandeler(dir, name string) *Handeler {
	isThere, err := file.CheckDir(dir)
	logger.Error(err, "")
	if !isThere {
		// err = file.MakeDir(dir)
		// logger.Error(err, "")
		logger.NewError("There was No such a Directory called %v, or maby the Path is Wrong!", dir)
	}
	newHandeler := Handeler{
		BaseDir:    dir,
		Name:       name,
		BaseFormat: Json,
		FieldsInfo: map[string]*FieldInfo{},
		Fields:     map[string]*Field{},
	}
	return &newHandeler
}

// .SaveInfo() saves infos of fields that has previusly changed.
func (h *Handeler) SaveInfo() error {
	h.LastModification = time.Now()
	bytes, err := json.Marshal(h)
	if err != nil {
		return err
	}
	err = file.Save(bytes, h.BaseDir, h.Name+h.BaseFormat)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handeler) Save() error {
	for _, field := range h.Fields {
		field.Info.LastModification = time.Now()
		err := field.Save()
		if err != nil {
			return err
		}
	}
	err := h.SaveInfo()
	return err
}

func (h *Handeler) PanicSave() {
	if err := h.Save(); err != nil {
		logger.Error(err, "there was an error while saving `%v` at `%v`", h.Name, h.BaseDir)
	}
}

/*
Field.Restore() matches the Last saved datas on the Config file to the raw variables at the startup,
there should be a call to the restore function at the end of every config Initiation, so you can either put a field/handeler.Restore()
at the end of your initiation part of code or create a confInit() function for your Config Initiation and use a `defer field/handeler.Restore()`
to do the same, the second approache is Recommended.
like the .Save() method, .Restore() is also avalable on both Handelers and Fields and it will effect on single field if is called on a field &
will effect on all linked-Fields is is Called on a Handeler.
if there was a variable in config that had no matching Item in your config youill get an error of type Prameter not declared.
*/
func (h *Handeler) Restore() error {
	exists, err := file.DoesExist(h.BaseDir, h.Name+h.BaseFormat)
	if exists {
		bytes, err := file.Read(h.BaseDir, h.Name+h.BaseFormat)
		err = json.Unmarshal(bytes, &h)
		if err != nil {
			return err
		}

		for _, field := range h.Fields {
			err := field.Restore()
			if err != nil {
				return err
			}
			h.Fields[field.Info.Name] = field
			h.FieldsInfo[field.Info.Name] = &field.Info
		}
	} else if err != nil {
		return err
	} else {
		err = h.Save()
		if err != nil {
			return err
		}
	}
	return nil
}
