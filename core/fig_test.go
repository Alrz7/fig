package fig

import (
	"testing"

	"github.com/Alrz7/fig/loggy"
)

func TestMain(t *testing.T) {

	type api struct {
		Url  string `json:"url"`
		Port int    `json:"port"`
	}
	type car struct {
		Name  string `Json:"name"`
		Model string `Json:"model"`
		Year  int    `Json:"year"`
	}
	porche911 := car{
		Name:  "porche911",
		Model: "porche",
		Year:  2003,
	}
	google := api{}
	appConfig, err := CreateNewHandler("./test/", "appConfig")
	if err != nil {
		loggy.DefaultLogger.Error(err)
	}
	mainField, err := appConfig.NewField("./test/", "app_api")
	if err != nil {
		loggy.DefaultLogger.Error(err)
	}

	mainField.Set("google", &google)
	google.Port = 5050

	seocndField, err := appConfig.NewField("./test/", "Appconf_second")
	if err != nil {
		loggy.DefaultLogger.Error(err)
	}
	seocndField.Set("porche911", &porche911)

	appConfig.PanicRestore()

}
