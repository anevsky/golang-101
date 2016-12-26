package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	fmt.Println("hello")

	langTypes := []string{"arabic", "azerty", "cyrillic", "qwerty"}

	arabicLangCodes := []string{"am", "il"}
	azertyLangCodes := []string{"fr"}
	cyrillicLangCodes := []string{"az", "by", "kir", "kz", "md", "ru", "tg", "uk", "uz"}
	qwertyLangCodes := []string{"az", "de", "en", "es", "fi", "it", "no", "pl", "pt", "ro", "se", "tm", "tr", "uz"}

	for _, langType := range langTypes {
		var langCodes []string

		switch {
		case langType == "arabic":
			langCodes = arabicLangCodes
		case langType == "azerty":
			langCodes = azertyLangCodes
		case langType == "cyrillic":
			langCodes = cyrillicLangCodes
		case langType == "qwerty":
			langCodes = qwertyLangCodes
		default:
		}

		for _, langCode := range langCodes {
			// var iPhone = 320
			// var iPhone6Length = 375
			var iPhone6PlusLength = 414

			// var iPhoneNumber = 1.0
			// var iPhone6Number = 1.171875
			var iPhone6PlusNumber = 1.104

			var inputDir = "iphone6/languages/iPhone6Layout/" + langType + "/"
			var outputDir = "iphone6plus/" + "languages/iPhone6PlusLayout/" + langType + "/"
			var inputFileName = langType + "_" + langCode + "_layout_configuration_iphone6.json"
			var outputFileName = langType + "_" + langCode + "_layout_configuration_iphone6plus.json"

			file, err := ioutil.ReadFile(inputDir + inputFileName)
			if err != nil {
				log.Fatal(err)
			}

			var f interface{}
			err = json.Unmarshal(file, &f)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("json readed")

			jsonData := f.(map[string]interface{})
			keysLayout := jsonData["keys_layout"].(map[string]interface{})

			fmt.Println("identifier : " + jsonData["identifier"].(string))

			var buffer bytes.Buffer
			buffer.WriteString("{")
			buffer.WriteString("\n")
			buffer.WriteString("\"identifier\": \"")
			buffer.WriteString(jsonData["identifier"].(string))
			buffer.WriteString("\",\n")
			buffer.WriteString("\"size\": { \"width\": ")
			buffer.WriteString(fmt.Sprintf("%d", iPhone6PlusLength))
			buffer.WriteString(", \"height\": 208 },\n")
			buffer.WriteString("\"params\": { \"type\": \"language\", \"language\": \"")
			buffer.WriteString(langCode)
			buffer.WriteString("\" },\n")
			buffer.WriteString("\n")
			buffer.WriteString("\"keys_layout\": {\n")

			for _, keysMap := range keysLayout {
				keyInfo := keysMap.(map[string]interface{})

				var keyCode = keyInfo["code"]
				var keyX = keyInfo["x"]
				var keyY = keyInfo["y"]
				var keyWidth = keyInfo["width"]
				var keyHeight = keyInfo["height"]

				buffer.WriteString("\"")
				buffer.WriteString(keyCode.(string))
				buffer.WriteString("\"")
				buffer.WriteString(" : {\"code\": ")
				buffer.WriteString("\"")
				buffer.WriteString(keyCode.(string))
				buffer.WriteString("\"")
				buffer.WriteString(", ")
				buffer.WriteString("\"x\": ")
				buffer.WriteString(fmt.Sprintf("%.0f", keyX.(float64)*iPhone6PlusNumber))
				buffer.WriteString(", ")
				buffer.WriteString("\"y\": ")
				buffer.WriteString(fmt.Sprintf("%.0f", keyY.(float64)-3))
				buffer.WriteString(", ")
				buffer.WriteString("\"width\": ")
				buffer.WriteString(fmt.Sprintf("%.0f", keyWidth.(float64)*iPhone6PlusNumber))
				buffer.WriteString(", ")
				buffer.WriteString("\"height\": ")
				buffer.WriteString(fmt.Sprintf("%.0f", keyHeight.(float64)*iPhone6PlusNumber-2))
				buffer.WriteString("},\n")
			}

			buffer.WriteString("}\n}")

			fmt.Println("builded")

			os.MkdirAll(string(outputDir), 0777)

			err = ioutil.WriteFile(outputDir+outputFileName, buffer.Bytes(), 0644)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("OK!")
			}
		}

	}

}
