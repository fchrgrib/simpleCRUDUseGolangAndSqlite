package models

type UpdateDataUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description_user"`
}

type UpdateLocation struct {
	Name     string `json:"location"`
	Location string `json:"description_location"`
}

type UpdateSchool struct {
	Rank int    `json:"school_rank"`
	Name string `json:"school_name"`
}
