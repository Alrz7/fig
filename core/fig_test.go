package fig

import (
	"testing"
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
	appConfig := CreateNewHandler("./test/", "appConfig")
	mainField := appConfig.NewField("./test/", "app_api")

	mainField.Set("google", &google)
	google.Port = 5050

	seocndField := appConfig.NewField("./test/", "Appconf_second")
	seocndField.Set("porche911", &porche911)

	appConfig.PanicRestore()

}
