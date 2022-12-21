package src

import (
	"log"
	"os"

	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

func loadConfig() map[interface{}]map[interface{}]string {
	f, _ := os.ReadFile("config/config.yaml")
	m := make(map[interface{}]map[interface{}]string)
	yaml.Unmarshal(f, &m)
	return m
}

func GetFileToReadPath() string {
	m := loadConfig()
	return m["fileConfig"]["path"] + m["fileConfig"]["name"]
}

func Read_HtmlFile() *html.Node {
	f, _ := os.Open(GetFileToReadPath())
	n_file, err := html.Parse(f)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer f.Close()
	return n_file
}
