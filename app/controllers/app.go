package controllers

import (
	"github.com/robfig/revel"
	"os"
	"io/ioutil"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Upload() revel.Result {
	for _, fileHeaders := range c.Params.Files {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path := "files/" + fileHeader.Filename
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
		}
	}

	return c.Render()
}
