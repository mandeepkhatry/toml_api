package initializer

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

//InitializeRoutes : route initialization
func InitializeRoutes(resourceLocation string) {
	var r map[string]interface{}
	if _, err := toml.DecodeFile(resourceLocation, &r); err != nil {
		fmt.Println("Error")
	} else {
		// var apiEndPoint interface{}
		var endPoints map[string]interface{}
		//endPoints : root, detail, approve
		endPoints = r["api"].(map[string]interface{})

		for endPoint := range endPoints {
			methods := endPoints[endPoint].(map[string]interface{})
			if _, ok := methods["endpoint"]; !ok {
				//error
				fmt.Println("No endpoint defined")
			} else {
				//methods[1:] = [get, post,...]
				fmt.Println("Handle all the handlers here")
			}
		}
	}
}
