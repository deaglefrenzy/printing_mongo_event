package models

import "time"

type Printer struct {
	// Printer details
	UUID        string `bson:"-" json:"uuid"`
	IP          string `bson:"ip" json:"ip"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`

	// Printer settings
	PaperWidth        PaperWidthEnum `bson:"paper_width" json:"paper_width"`
	CharactersPerLine int            `bson:"characters_per_line" json:"characters_per_line"`
	DotsPerInch       int            `bson:"dots_per_inch" json:"dots_per_inch"`
	LargeTicketFont   bool           `bson:"large_ticket_font" json:"large_ticket_font"`

	Categories []ItemCategoryUUIDAndName `bson:"categories" json:"categories"`
	Zones      []ZoneUUIDAndName         `bson:"zones" json:"zones"`
	Prints     []PrintableEnum           `bson:"prints" json:"prints"`

	CreatedAt int64 `bson:"created_at" json:"created_at"`
}

type ItemCategoryUUIDAndName struct {
	UUID string `bson:"uuid" json:"uuid"`
	Name string `bson:"name" json:"name"`
}

type ZoneUUIDAndName struct {
	UUID string `bson:"uuid" json:"uuid"`
	Name string `bson:"name" json:"name"`
}

type PrintableEnum string

const (
	Tickets      PrintableEnum = "tickets"
	Bills        PrintableEnum = "bills"
	Shifts       PrintableEnum = "shifts"
	Transactions PrintableEnum = "transactions"
)

type PaperWidthEnum int

const (
	PaperWidth32 PaperWidthEnum = 32
	PaperWidth48 PaperWidthEnum = 48
	PaperWidth60 PaperWidthEnum = 60
	PaperWidth72 PaperWidthEnum = 72
	PaperWidth80 PaperWidthEnum = 80
)

type PrintRequest struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Payload string    `json:"payload"`
	Created time.Time `json:"created"`
}
