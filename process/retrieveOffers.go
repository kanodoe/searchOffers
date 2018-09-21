package process

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type OfferCollection struct {
	StoreName           string
	Date                string
	OfferDataCollection []OfferData
}

type OfferData struct {
	Name           string
	Code           string
	StorePrice     string
	InetPrice      string
	InetOfferPrice string
	Uri            string
	ErrorMessage   error
}

func RetrieveData(uri string, tags Tags) OfferData {

	log.Print("Retrieving data from uri: ", uri)

	var offer OfferData

	resp, err := http.Get(uri)

	if err != nil {
		offer.ErrorMessage = err
		return offer
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))

	offer.Name = removeDuplicatedWhitespace(doc.Find(tags.Title).Text())
	offer.Code = removeCodePrefix(spaceFieldsJoin(doc.Find(tags.Code).Text()))
	offer.StorePrice = spaceFieldsJoin(doc.Find(tags.Store).ChildrenFiltered(tags.Price).Text())
	offer.InetPrice = spaceFieldsJoin(doc.Find(tags.Internet).ChildrenFiltered(tags.Price).Text())
	offer.InetOfferPrice = spaceFieldsJoin(doc.Find(tags.Sales).ChildrenFiltered(tags.Price).Text())
	offer.Uri = uri

	return offer
}

func removeCodePrefix(code string) string {
	log.Println("Removing Code, or SKU word config")
	reg := regexp.MustCompile(`^[^:-].*?:\s*`)
	return reg.ReplaceAllString(code, "${1}")
}

func spaceFieldsJoin(str string) string {
	log.Println("Removing space fields")
	return strings.Join(strings.Fields(str), "")
}

func removeDuplicatedWhitespace(str string) string {
	log.Println("Removing whitespace form crawled data")
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(str, " ")
}
