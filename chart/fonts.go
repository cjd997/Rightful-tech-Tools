package chart

import (
	"strings"

	"github.com/flopp/go-findfont"
)

func font(name string, file ...string) (string, string) {
	filePath := name + ".ttf"
	if len(file) > 0 {
		if strings.HasSuffix(file[0], ".ttf") {
			filePath = file[0]
		} else {
			filePath = file[0] + ".ttf"
		}
	}
	fontPath, err := findfont.Find(filePath)
	if err != nil {
		panic(err)
	}
	return name, fontPath
}
