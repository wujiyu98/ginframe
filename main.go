package main

import (
	_ "github.com/wujiyu98/ginframe/config"
	"github.com/wujiyu98/ginframe/model"
	"github.com/wujiyu98/ginframe/notice"
)

func main() {

	var enquiry = model.Enquiry{
		Name:        "wujiyu",
		Email:       "284703576@qq.com",
		Comment:     "hello man",
		Country:     "china",
		MobilePhone: "1312313131",
		Company:     "SFA SDFASF",
		Products: []model.EnquiryProduct{
			{Title: "atmega1", Manufacturer: "ti", Summary: "asdf"},
			{Title: "atmega2", Manufacturer: "ti", Summary: "asdf"},
		},
	}
	notice.Enquiry(enquiry)

}
