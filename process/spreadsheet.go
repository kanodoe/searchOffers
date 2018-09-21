package process

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"log"
)

func AddRecords(collection []OfferCollection) {

	for _, item := range collection {
		sheet := initSpreadsheet(item.StoreName)

		var nextRowPos = getNextEmptyRow(sheet)

		for i, offer := range item.OfferDataCollection {

			if searchLastCode(sheet, offer) {
				log.Print("New price was found for product: ", offer.Name)
				SendEmail(offer)
			}

			sheet.Update(nextRowPos+i, 0, item.Date)
			sheet.Update(nextRowPos+i, 1, offer.Name)
			sheet.Update(nextRowPos+i, 2, offer.Code)
			sheet.Update(nextRowPos+i, 3, offer.StorePrice)
			sheet.Update(nextRowPos+i, 4, offer.InetPrice)
			sheet.Update(nextRowPos+i, 5, offer.InetOfferPrice)
			sheet.Update(nextRowPos+i, 6, offer.Uri)
		}

		// Make sure call Synchronize to reflect the changes
		err := sheet.Synchronize()
		checkError(err)
	}
}

func initSpreadsheet(sheetTitle string) *spreadsheet.Sheet {

	log.Print("Init spreadsheet service to connect GCP API")

	var sheet *spreadsheet.Sheet

	data, err := ioutil.ReadFile("client_secret.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)

	defer func() {
		if err := recover(); err != nil {
			sheet = createNewSheet(sheetTitle, service)
		}
	}()

	sheet = findSheetByTitle(sheetTitle, service)

	return sheet
}

func fetchSpreadsheet(service *spreadsheet.Service) spreadsheet.Spreadsheet {

	ss, err := service.FetchSpreadsheet("11y4NmD39gPSPcKAzMU0lsAoDWVQsXOVxJNNVZXV7xek")
	checkError(err)

	return ss
}

func findSheetByTitle(sheetTitle string, service *spreadsheet.Service) *spreadsheet.Sheet {

	ss := fetchSpreadsheet(service)

	log.Println("Finding sheet with name ", sheetTitle)
	sheet, err := ss.SheetByTitle(sheetTitle)
	checkError(err)

	return sheet
}

func createNewSheet(sheetTitle string, service *spreadsheet.Service) *spreadsheet.Sheet {

	log.Print("Sheet not found, creating sheet named: ", sheetTitle)

	ss, err := service.CreateSpreadsheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: sheetTitle,
		},
	})
	checkError(err)

	sheet, err := ss.SheetByTitle(sheetTitle)
	checkError(err)

	return sheet

}

func getNextEmptyRow(sheet *spreadsheet.Sheet) int {

	log.Println("Finding last row position with values")

	var counter uint

	for _, row := range sheet.Rows {
		for _, cell := range row {
			counter = cell.Row
		}
	}

	return int(counter + 1)
}

func searchLastCode(sheet *spreadsheet.Sheet, data OfferData) bool {
	var match bool

	log.Println("Search for the same code (newly added only) and then compare if the price has changed")

	for i := range sheet.Rows {
		if sheet.Rows[i][2].Value == data.Code {

			match = false

			if sheet.Rows[i][5].Value != data.InetOfferPrice {
				match = true
			}
		}
	}

	return match
}

func checkError(err error) {
	if err != nil {
		log.Fatal("error: ", err)
		panic(err.Error())
	}
}
