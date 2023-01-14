package models

import "gorm.io/gorm"

type DataUser struct {
	gorm.Model
	Uid         int    `json:"uid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"descrciption_user"`
}

type Location struct {
	Coordinate int    `json:"coordinate"`
	Name       string `json:"location"`
	Location   string `json:"description_location"`
}

type School struct {
	Rank       int    `json:"school_rank"`
	Name       string `json:"school_name"`
	Coordinate int    `json:"school_location"`
}

var Users = []DataUser{
	{Uid: 135001, Name: "Fahrian Afdholi", Email: "fahrian.alifudin@gmail.com", Description: "I am the best player in this conservation"},
	{Uid: 135002, Name: "Fahrian Alifudin", Email: "fahrian.afdholi@gmail.com", Description: "I am the best player in this conservation"},
}
