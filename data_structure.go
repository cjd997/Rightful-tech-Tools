package main

type LeftRightDirection string

const (
	Left  LeftRightDirection = "left"
	Right LeftRightDirection = "right"
)

type DatasItem struct {
	Text      string             `json:"text`
	Direction LeftRightDirection `json:"direction"`
	Cost      float32            `json:"cost"`
}

type Datas []DatasItem
