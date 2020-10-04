package main

import (
	"github.com/getevo/composer"
	"testing"
)

func TestAddInt(t *testing.T) {

	image := composer.Image{
		Width:  1050,
		Height: 700,
		Objects: []composer.Object{
			{
				Name: "Background",
				Type: composer.IMAGE,
				Left: 0,Top:0, Width:1050,Height: 700,
				Value: "https://images.unsplash.com/photo-1493494817959-6981ce4b2603?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1050&q=80",
				Cache: true,
			},
			{
				Name: "Logo With Background",
				Type: composer.IMAGE,
				Left: 0,Top:0, Width:1050,Height: 700,
				Value: "logo.jpg",
				Effect: "removebackground(200) floodfill(512,350,100,#FFFFFFFF)",
				Cache: true,
			},
			{
				Name: "Logo",
				Type: composer.IMAGE,
				Left: 525,Top:350, Width:300,Height: 500,HAlign: composer.CENTER, VAlign: composer.MIDDLE, WordWrap: true,
				Value: "./logo.png",
				Effect: "invert() brightness(2)",
			},
			{
				Name: "Text",
				Type: composer.TEXT,
				Left: 525,Top:350, Width:300,HAlign: composer.CENTER, VAlign: composer.MIDDLE,
				Color: "rgba(255,0,0,128)",
				Font: "./fonts/Roboto-Bold.ttf",
				FontSize: 36,
				Value: "Hello World",
			},
		},
	}

	err := image.Create()
	if err != nil{
		panic(err)
	}
	image.SavePNG("./output.png")
}
