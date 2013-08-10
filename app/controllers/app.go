package controllers

import (
	"github.com/robfig/revel"
	"os"
	"bufio"
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
			path := "files/" + fileHeader.Filename
			outFile, _ := os.OpenFile(path, os.O_CREATE, 0)
			writer := bufio.NewWriter(outFile)

			file, _ := fileHeader.Open()
			reader := bufio.NewReader(file)
			for {
				b, err := reader.ReadByte()
				if err != nil {
					break
				}
				writer.WriteByte(b)
			}
			writer.Flush()
		}
	}

	return c.Render()
}
