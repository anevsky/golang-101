package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("gif-list-vk.json")
	if err != nil {
		log.Fatal(err)
	}

	var f interface{}
	err = json.Unmarshal(file, &f)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := f.(map[string]interface{})
	categories := jsonData["categories"].(map[string]interface{})

	fmt.Println("Gif last_mod_date : " + jsonData["last_mod_date"].(string))

	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	for _, category := range categories {
		items := category.([]interface{})
		for _, categoryItem := range items {
			buffer.WriteString(categoryItem.(map[string]interface{})["name"].(string))
			buffer.WriteString("<br>")
			buffer.WriteString("<img src=\"" + categoryItem.(map[string]interface{})["preview_link"].(string) + "\">")
			buffer.WriteString(" ")
			buffer.WriteString("<img src=\"" + categoryItem.(map[string]interface{})["original_link"].(string) + "\">")
			buffer.WriteString("<br>")
			buffer.WriteString("<p>")
			var width = categoryItem.(map[string]interface{})["dimension"].(map[string]interface{})["width"]
			buffer.WriteString(fmt.Sprintf("%.0f", width))
			buffer.WriteString("x")
			var height = categoryItem.(map[string]interface{})["dimension"].(map[string]interface{})["height"]
			buffer.WriteString(fmt.Sprintf("%.0f", height))
			buffer.WriteString(", ")
			var size = categoryItem.(map[string]interface{})["size"]
			buffer.WriteString(fmt.Sprintf("%.3f mb", size))
			buffer.WriteString("</p>")
		}
	}
	buffer.WriteString("</body></html>")

	var outputFileName = "gif-list-vk.html"

	err = ioutil.WriteFile(outputFileName, buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Please see " + outputFileName)
	}

}
