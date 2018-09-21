Offer Crawler
===
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

This project (my first project using golang) intend to help anyone who wants to easily check offers of their interest in online stores, saving what found and send an email in case of find something, it support multiple stores and links, but is necessary to make some configuration to work.

Any pull-request is welcome.

## Installation

```
go get gopkg.in/Iwark/spreadsheet.v2
```

## Preparation

#### Step 1

This project read YAML files located in the folder process > config-files with this structure

```yaml
name: "Audiomusica" [Name of the store]
domain: "https://www.audiomusica.com/catalogo/" [Domain of the store]
links: [A list of all the links that you want to lookig for, remember to remove the domain]
   -"komplete-kontrol-s49-mk2-teclado-controlador-native-instruments.html"
tags: [HTML attributes to filter in what the crawler gets]
  title: ".product-name h1"
  code: ".detalle-producto-left.last"
  store: ".precio-tienda"
  internet: ".precio-internet"
  sales: ".precio-oferta"
  price: ".price-box"
```

You can have all the files and links do you want

#### Step 2

We use the package from [Iwark](https://github.com/Iwark/spreadsheet) to write the information into a Google Spreadsheet so you will need to follow their instructions about requesting API account and configure the oauth2 to do so 

### Step 3

To send the email you must have a Google Account and create an app password to change the value in the file emailSender.go 

## Usage

First after doing the 3 steps mentioned above you must create a Google Spreadsheet and share it with the service account that you have created in GCP (no matter if this email didn't receive emails), after that in the first row you must add these values

[Date, Product, Code, Store Price, Internet Price, Sales Price, URL]

Next you must use the same name of attribute **name** in YAML file for each config files that you have

@TODO: Fix that when the author of the library [answer my request](https://github.com/Iwark/spreadsheet/issues/30).

Just compile and run it, the email will be sent if for the same code (SKU) exist a different value in Sales Price in the newest value before the actual request.

## TODO
- Add unitary test when I understand better how to write them.
- Understand if is a bug or I'm using incorrectly the Spreadsheet library because I can't create new Spreadsheets.
- Improve documentation.

## License

Search Offers is released under the MIT License.
