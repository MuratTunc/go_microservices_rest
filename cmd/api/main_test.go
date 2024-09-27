package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func TestrenderTemplate(t *testing.T) {

	var tmpl_page = "../../templates/home_page.html"

	isFileExist := checkFileExists(tmpl_page)

	if isFileExist {
		fmt.Println("file exist")
	} else {

		fmt.Println("file not exists")
	}

}
