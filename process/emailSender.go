package process

import (
	"log"
	"net/smtp"
)

func SendEmail(offer OfferData) {
	from := ""
	pass := ""
	subject := "" + offer.Name
	to := ""

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		"Producto: " + offer.Name + "\n" +
		"Código: " + offer.Code + "\n" +
		"Precio Tienda: " + offer.StorePrice + "\n" +
		"Precio Internet: " + offer.InetPrice + "\n" +
		"Precio Oferta: " + offer.InetOfferPrice + "\n" +
		"Url: " + offer.Uri + "\n"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
