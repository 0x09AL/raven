package main

import (

	"raven/model"
	"gopkg.in/gcfg.v1"
	"log"
	"raven/terminal"
	"raven/scraper"
	"fmt"
	b64 "encoding/base64"
)


func main(){


	b64_banner := "X19fX19fX19fXyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBfX19fX19fICAgICAgIF9fX19fX19fICAKXF9fX19fXyAgIFxfX19fXyBfX18gIF9fIF9fX18gICBfX19fICAgICAgICAgICBcICAgXyAgXCAgICAgIFxfX19fXyAgXCAKIHwgICAgICAgXy9cX18gIFxcICBcLyAvLyBfXyBcIC8gICAgXCAgIF9fX19fXyAvICAvX1wgIFwgICAgICAgXyhfXyAgPCAKIHwgICAgfCAgIFwgLyBfXyBcXCAgIC9cICBfX18vfCAgIHwgIFwgL19fX19fLyBcICBcXy8gICBcICAgICAvICAgICAgIFwKIHxfX19ffF8gIC8oX19fXyAgL1xfLyAgXF9fXyAgPl9fX3wgIC8gICAgICAgICAgXF9fX19fICAvIC9cIC9fX19fX18gIC8KICAgICAgICBcLyAgICAgIFwvICAgICAgICAgIFwvICAgICBcLyAgICAgICAgICAgICAgICAgXC8gIFwvICAgICAgICBcLyAK"
	banner, _ := b64.StdEncoding.DecodeString(b64_banner)
	fmt.Println(string(banner))


	Configuration := model.CFG{}
	err := gcfg.ReadFileInto(&Configuration,"config.conf")
	if err != nil{
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}
	//scraper.StartDriver()
	scraper.LUsername = Configuration.Creds.Username
	scraper.LPassword = Configuration.Creds.Password

	terminal.StartTerminal()


}