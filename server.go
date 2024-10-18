package main

import (
	"timeline/backend/app"
	"timeline/backend/lib/utils"
)

func main() {
	application, err := app.NewApplication()
	if err != nil {
		utils.F("Coulnd not create application: %v", err)
	}
	application.Start()
	defer application.Closer()()
}
