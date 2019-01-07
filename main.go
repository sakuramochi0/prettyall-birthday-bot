package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type T struct {
	Title      string
	Characters []struct {
		Name        string
		Birthday    string
		Description string
	}
}

func main() {
	t := ReadConfig("king_of_prism")

	println("title:", t.Title)
	for i, chara := range t.Characters {
		println(i, chara.Name, chara.Birthday, chara.Description)
	}
}

func ReadConfig(title string) (t T) {
	filename := "pretty-all-friends-birthday-calendar/config/" + title + ".yml"
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("ReadFile error:", err)
	}

	t = T{}
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	return
}
