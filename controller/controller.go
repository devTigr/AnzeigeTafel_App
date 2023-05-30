package controller

import (
	"AnzeigeTafel_App/fyneWindow"
	"AnzeigeTafel_App/models"
	"log"
)

// Run does the running of the fyneWindow application
func Run(enablePersistence bool) {
	if enablePersistence {
		models.EnableFilePersistence()
	} else {
		models.DisableFilePersistence()
	}

	err := models.Initialize()
	checkAndHandleErrorWithTermination(err)

	fyneWindow.Clear()
	fyneWindow.PrintMenu()

	for true {
		executeCommand()
	}
}

func checkAndHandleErrorWithTermination(err error) {
	if err != nil {
		fyneWindow.PrintError(err)
		log.Fatal(err)
	}
}

func checkAndHandleErrorWithoutTermination(err error) {
	if err != nil {
		fyneWindow.PrintMessage("The following error occurred:")
		fyneWindow.PrintError(err)
		fyneWindow.PrintMessage("Press c to continue!")
	}
}

// check the file for new Data
func checkFile() {
	data.CheckFile()
}
