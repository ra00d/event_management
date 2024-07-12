package configs

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gofiber/template/django/v3"
)

func AttachTemplateEngine() *django.Engine {
	engine := django.New("./views", ".html")
	engine.AddFunc("getCssAsset", func(name string) (res template.HTML) {
		filepath.Walk("public/assets", func(path string, info os.FileInfo, err error) error {
			// fmt.Print(info.Name())
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			if info.Name() == name {
				fmt.Println(path)
				fmt.Println(name)
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
			}
			return nil
		})
		return
	})
	return engine
}
