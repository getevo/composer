package composer

import "C"
import (
	"errors"
	"image/color"
	"regexp"
	"strconv"
)

var errInvalidFormat = errors.New("invalid format")
var rgbaRegex = regexp.MustCompile(`(rgba|RGBA)\((\d+),(\d+),(\d+),(\d+)\)`)
var rgbRegex = regexp.MustCompile(`(rgb|RGB)\((\d+),(\d+),(\d+)\)`)
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 9:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		c.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
func ParseRGBColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	matches := rgbaRegex.FindAllStringSubmatch(s,1)
	if len(matches) == 1 && len(matches[0]) == 6{
		var cp int
		cp,err = strconv.Atoi(matches[0][2])
		if err != nil{
			return color.RGBA{}, err
		}
		c.R = uint8(cp)
		cp,err = strconv.Atoi(matches[0][3])
		if err != nil{
			return color.RGBA{}, err
		}
		c.G = uint8(cp)
		cp,err = strconv.Atoi(matches[0][4])
		if err != nil{
			return color.RGBA{}, err
		}
		c.B = uint8(cp)
		cp,err = strconv.Atoi(matches[0][5])
		if err != nil{
			return color.RGBA{}, err
		}
		c.A = uint8(cp)
		return c,nil
	}

	matches = rgbRegex.FindAllStringSubmatch(s,1)
	if len(matches) == 1 && len(matches[0]) == 5{
		var cp int
		cp,err = strconv.Atoi(matches[0][2])
		if err != nil{
			return color.RGBA{}, err
		}
		c.R = uint8(cp)
		cp,err = strconv.Atoi(matches[0][3])
		if err != nil{
			return color.RGBA{}, err
		}
		c.G = uint8(cp)
		cp,err = strconv.Atoi(matches[0][4])
		if err != nil{
			return color.RGBA{}, err
		}
		c.B = uint8(cp)
		return c,nil
	}

	return color.RGBA{},errInvalidFormat
}
func ParseColor(s string)(c color.RGBA, err error){
	if s[0] == '#'{
		return ParseHexColor(s)
	}
	return ParseRGBColor(s)

}