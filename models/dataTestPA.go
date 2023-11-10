package models

type DataTestPA struct {
	DateTime    string `bson:"dateTime" json:"dateTime"`
	Temperature int    `bson:"temperature" json:"temperature"`
}
