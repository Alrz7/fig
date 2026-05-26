package fig

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Alrz7/fig/file"
)

type FieldInfo struct {
	Dir              string    `json:"dir"`
	Name             string    `json:"name"`
	Format           string    `json:"format"`
	LinkedToHandelr  bool      `json:"linked_to_Handler"`
	LastModification time.Time `json:"last_time_modified"`
}

type Field struct {
	Info              FieldInfo `json:"info"`
	restored          bool      // blocking Set() functions from saving datas to the File before Restoring
	Data              Topic     `json:"data"`
	updateHandelrInfo func() error
	// type, last_modified, etc.
}

/*
Use CreatNewField lets you create a self handeled field that has all the abilities but its Not linked to a Handler or a group of fields
dir: the directory that you want to save your config in, the Base directory will be the Directory of your Config.go file and
it should have a format like `./yourConfigDir/` |
name: a name for your config file, its recomended to be the same as your variable's name.
*/
func CreateNewField(dir, name string) *Field {
	// there should be the Dir-Path Unit checking this part
	isThere, err := file.CheckDir(dir)
	logger.Error(err, "")
	if !isThere {
		logger.NewError("There was No such a Directory called %v, or maby the Path is Wrong!", dir)
	}
	newInfo := FieldInfo{
		Dir:             dir,
		Name:            name,
		Format:          Json,
		LinkedToHandelr: false,
	}
	newfield := Field{
		Info:     newInfo,
		restored: false,
		Data:     Topic{},
	}
	return &newfield
}

/*
NewField lets you add a new field to the Handler you are calling the method on, it uses CreteNewField to initiate the new field
and then links it to the Handler
*/
func (h *Handler) NewField(dir, name string) *Field {
	newfield := CreateNewField(dir, name)
	newfield.Info.LinkedToHandelr = true
	newfield.updateHandelrInfo = h.SaveInfo
	h.Fields[name] = newfield
	h.FieldsInfo[name] = &(newfield.Info)
	return newfield
}

/*
Use .Set() to add Items to your Field by Passing Their Pointers [NOT Values] to .set() along with a name for the item.
Fig will keep eye on the item and by using .save() method on either the Field or the Handelr you can submit the changes to the Config File.
Notice that the .save() mthod will save all linked-fields if is called on a Handler and will save a single field if is called on a field.
*/
func (f *Field) Set(key string, newValue any) { // NOTE : need to make a WARNING about NewValue Being a Pointer (&any_value)
	f.Data.Set(key, newValue)
	if f.restored {
		f.Save()
	}
}

func (f *Field) Pop(key string) any {
	return f.Data.Pop(key)
}

/*
Field.save() lets you submit or save the changes from your app to your config File,
if the field was linked to a Handler it will call a method on Handler to update the Info(s) on Both sides.
*/
func (f *Field) Save() error {
	f.Info.LastModification = time.Now()
	bytes, err := json.Marshal(f)
	if err != nil {
		return err
	}
	err = file.Save(bytes, f.Info.Dir, f.Info.Name+f.Info.Format)
	if err != nil {
		return err
	}
	if f.Info.LinkedToHandelr {
		err = f.updateHandelrInfo()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Field) PanicSave() {
	if err := f.Save(); err != nil {
		logger.Error(err, "there was an error while saving `%v` at `%v`", f.Info.Name, f.Info.Dir)
	}
}

/*
Field.Restore() matches the Last saved datas on the Config file to the raw variables at the startup,
there should be a call to the restore function at the end of every config Initiation, so you can either put a field/Handler.Restore()
at the end of your initiation part of code or create a confInit() function for your Config Initiation and use a `defer field/Handler.Restore()`
to do the same, the second approache is Recommended.
like the .Save() method, .Restore() is also avalable on both Handlers and Fields and it will effect on single field if is called on a field &
will effect on all linked-Fields is is Called on a Handler.
if there was a variable in config that had no matching Item in your config youill get an error of type Prameter not declared.
*/
func (f *Field) Restore() error {
	exists, err := file.DoesExist(f.Info.Dir, f.Info.Name+f.Info.Format)
	if exists {
		bytes, err := file.Read(f.Info.Dir, f.Info.Name+f.Info.Format)
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
			fmt.Println(1)
			err = f.Save()
			if err != nil {
				return err
			}
		}
	} else if err != nil {
		return err
	} else {
		fmt.Println(2)
		err = f.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) PanicRestore() {
	if err := h.Restore(); err != nil {
		logger.Error(err, "there was an error while Restoring `%v` from `%v`", h.Name, h.BaseDir)
	}
}

// IF There Was a New Data To Save, .Restore() calls .Save() method To Save new Datas after Restoring.
func needToSave(f *Field, data *Topic) bool {
	for key := range f.Data {
		if _, ok := (*data)[key]; !ok {
			return true
		}
	}
	return false
}

/*
Field.Data's values can have any Complex type created by the user to save app's Config so wee need a Second-Step Decoding
to match the datas to the type it was created on.
*/
func marsh(f *Field, data *Topic) error {
	for key, val := range *data {
		if _, ok := f.Data[key]; ok { // NOTE: IF OK was False: we should decide to either Pop that Object or create a warning
			b, err := json.Marshal(val)
			if err != nil {
				return err
			}
			err = json.Unmarshal(b, f.Data[key])
			if err != nil {
				return err
			}
		} else {
			logger.NewError("Not All of %v's parameters were declared in your Application: lost `%v`", f.Info.Name, key)
		}
	}
	return nil
}

func (f *Field) PanicRestore() {
	if err := f.Restore(); err != nil {
		logger.Error(err, "there was an error while Restoring `%v` from `%v`", f.Info.Name, f.Info.Dir)
	}
}
