package main

import (
	"io/ioutil"
	"log"
	"os"
	"searchOffers/process"
	"time"
)

func main() {

	var fileArray = GetFileNames()
	var offerCollections []process.OfferCollection

	for _, element := range fileArray {

		var config = process.LoadFile(element)
		var offerCol process.OfferCollection

		if config.Name != "" {
			log.Println("Processing file: ", element)
			var now = time.Now()

			offerCol.StoreName = config.Name
			offerCol.Date = now.Format(time.RFC822Z)
			offerCol.OfferDataCollection = GetOffers(config)

			offerCollections = append(offerCollections, offerCol)
		}
	}

	if len(offerCollections) > 0 {
		process.AddRecords(offerCollections)
	} else {
		log.Print("There's no YAML files in config-files folder")
		os.Exit(1)
	}
}

/**
Retrieve file names from crawler-config-files directory
*/
func GetFileNames() []string {
	var fileNames []string

	pwd, _ := os.Getwd()
	dir := pwd + "/process/config-files"

	files, _ := ioutil.ReadDir(dir)

	for _, f := range files {
		fileNames = append(fileNames, dir+"/"+f.Name())
	}

	return fileNames

}

/**
Retrieve the offers one by one and return an object with all them
*/
func GetOffers(config process.Yaml) []process.OfferData {
	var offers []process.OfferData
	var tags = process.Tags(config.Tags)

	for i := range config.Links {
		offers = append(offers, process.RetrieveData(config.Domain+config.Links[i], tags))
	}

	return offers
}
