package controllers

import (
	"github.com/robfig/revel"
	"os"
	"bufio"
	"log"
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
			path := fileHeader.Filename
			outFile, err := os.Create(path)
			if err != nil {
				log.Println(err)
			}
			writer := bufio.NewWriter(outFile)

			file, _ := fileHeader.Open()
			reader := bufio.NewReader(file)
			buf := make([]byte, 4 * 1024 * 1024)
			for {
				n, err := reader.Read(buf)
				if err != nil {
					break
				}
				_, err = writer.Write(buf[:n])
				if err != nil {
					log.Println(err)
					break
				}
			}
			writer.Flush()
		}
	}

	return c.Render()
}
