package models

type Video struct {
	Id          int    `uri:"id"`
	Title       string `json:"title"`
	Description string `json:"desciption"`
}
