package models

type DataUser struct {
	ID          int    `json:"ID" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description_user"`
}

type Location struct {
	Coordinate int    `json:"coordinate" gorm:"primaryKey"`
	Name       string `json:"location"`
	Location   string `json:"description_location"`
}

type School struct {
	Rank       int    `json:"school_rank"`
	Name       string `json:"school_name"`
	Coordinate int    `json:"school_location" gorm:"primaryKey"`
}

var Users = []DataUser{
	{ID: 135001, Name: "Fahrian Afdholi", Email: "fahrian.alifudin@gmail.com", Description: "I am the best player in this conservation"},
	{ID: 135002, Name: "Fahrian Alifudin", Email: "fahrian.afdholi@gmail.com", Description: "I am the best player in this conservations"},
}
