package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	var fileArray = GetFileNames()
	var offerCollections []OfferCollection

	for _, element := range fileArray {

		var config = LoadFile(element)
		var offerCol OfferCollection

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
		AddRecords(offerCollections)
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
	dir := pwd + "/config-files"

	files, _ := ioutil.ReadDir(dir)

	for _, f := range files {
		fileNames = append(fileNames, dir+"/"+f.Name())
	}

	return fileNames

}

/**
Retrieve the offers one by one and return an object with all them
*/
func GetOffers(config Yaml) []OfferData {
	var offers []OfferData
	var tags = Tags(config.Tags)

	for i := range config.Links {
		offers = append(offers, RetrieveData(config.Domain+config.Links[i], tags))
	}

	return offers
}
