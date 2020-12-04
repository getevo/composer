package composer

import "gopkg.in/fogleman/gg.v1"

type Image struct {
	Width   int
	Height  int
	Objects []Object    `json:"objects"`
	Context *gg.Context `json:"-"`
}
type ObjectType string
type HAlign string
type VAlign string

const (
	TEXT   ObjectType = "text"
	IMAGE  ObjectType = "image"
	RECT   ObjectType = "rect"
	OVAL   ObjectType = "oval"
	LEFT   HAlign     = "left"
	RIGHT  HAlign     = "right"
	CENTER HAlign     = "center"
	TOP    VAlign     = "top"
	BOTTOM VAlign     = "bottom"
	MIDDLE VAlign     = "middle"
)

type Object struct {
	Name        string     `json:"name"`
	Type        ObjectType `json:"type"`
	HAlign      HAlign     `json:"halign"`
	VAlign      VAlign     `json:"valign"`
	Value       string     `json:"string"`
	Top         int        `json:"top"`
	Left        int        `json:"left"`
	Color       string     `json:"color"`
	Font        string     `json:"font"`
	FontSize    int        `json:"font_size"`
	Width       int        `json:"width"`
	Height      int        `json:"height"`
	Effect      string     `json:"effect"`
	Cache       bool       `json:"cache"`
	LineSpacing float64    `json:"line_spacing,omitempty"`
	WordWrap    bool       `json:"wordwrap,omitempty"`
}
