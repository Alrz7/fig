package fig

import (
	"encoding/json"
	"time"
	"github.com/Alrz7/fig/file"
)

type FieldInfo struct {
	Dir              string    `json:"dir"`
	Name             string    `json:"name"`
	Format           string    `json:"format"`
	LastModification time.Time `json:"last_time_modified"`
}

type Field struct {
	Info              FieldInfo `json:"info"`
	restored          bool      // blocking Set() functions from saving datas to the File before Restoring
	Data              Topic     `json:"data"`
	updateHandelrInfo func() error
	// IntField       cfInt     `json:"intField"`
	// StringField    cfString  `json:"stringField"`
	// type, last_modified, etc.
}

func (h *Handeler) NewField(dir, name string) *Field {
	// there should be the Dir-Path Unit checking this part
	isThere, err := file.CheckDir(dir)
	logger.Error(err, "")
	if !isThere {
		// err = file.MakeDir(dir)
		// logger.Error(err, "")
		logger.NewError("There was No such a Directory called %v, or maby the Path is Wrong!", dir)
	}
	newInfo := FieldInfo{
		Dir:    dir,
		Name:   name,
		Format: Json,
	}
	newfield := Field{
		Info:              newInfo,
		restored:          false,
		Data:              Topic{},
		updateHandelrInfo: h.SaveInfo,
	}
	// hndlr.PanicRestore()
	h.Fields[name] = &newfield
	h.FieldsInfo[name] = &newfield.Info
	return &newfield
}

func (f *Field) Set(key string, newValue any) { // NOTE : need to make a WARNING about NewValue Being a Pointer (&any_value)
	f.Data.Set(key, newValue)
	if f.restored {
		f.Save()
	}
}
func (f *Field) Pop(key string) any {
	return f.Data.Pop(key)
}

func (f *Field) Save() error {
	f.Info.LastModification = time.Now()
	bytes, err := json.Marshal(f)
	if err != nil {
		return err
	}
	err = file.Save(bytes, f.Info.Dir, f.Info.Name)
	if err != nil {
		return err
	}
	err = f.updateHandelrInfo()
	if err != nil {
		return err
	}
	return nil
}

func (f *Field) PanicSave() {
	if err := f.Save(); err != nil {
		logger.Error(err, "there was an error while saving `%v` at `%v`", f.Info.Name, f.Info.Dir)
	}
}

func (f *Field) Restore() error {
	exists, err := file.DoesExist(f.Info.Dir, f.Info.Name)
	if exists {
		bytes, err := file.Read(f.Info.Dir, f.Info.Name)
		var tempField Field
		err = json.Unmarshal(bytes, &tempField)
		if err != nil {
			return err
		}
		err = marsh(f, &tempField.Data)
		if err != nil {
			return err
		}
		f.restored = true
		newObject := needToSave(f, &tempField.Data)
		if newObject {
			err = f.Save()
			if err != nil {
				return err
			}
		}
	} else if err != nil {
		return err
	} else {
		err = f.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handeler) PanicRestore() {
	if err := h.Restore(); err != nil {
		logger.Error(err, "there was an error while Restoring `%v` from `%v`", h.Name, h.BaseDir)
	}
}

// if a new data was added , there should be a call to Save() after Restoring
func needToSave(f *Field, data *Topic) bool {
	for key := range f.Data {
		if _, ok := (*data)[key]; !ok {
			// fmt.Println("there is a new Object : ", key)
			return true
		}
	}
	return false
}

func marsh(f *Field, data *Topic) error {
	for key, val := range *data {
		if _, ok := f.Data[key]; ok { // NOTE: IF OK was False: we should decode to either Pop that Object or create a warning
			// fmt.Println(key, "found!: ")
			b, err := json.Marshal(val)
			if err != nil {
				return err
			}
			err = json.Unmarshal(b, f.Data[key])
			// fmt.Println((h.Data[key]))
			if err != nil {
				return err
			}
		} else {
			logger.NewError("Not All of %v's parameters were declared in your Application: lost `%v`", f.Info.Name, key)
		}
	}
	return nil
}

// it restores the Handeler to restore any saved config values
func (f *Field) PanicRestore() {
	if err := f.Restore(); err != nil {
		logger.Error(err, "there was an error while Restoring `%v` from `%v`", f.Info.Name, f.Info.Dir)
	}
}
