package scraper

import (
	"github.com/sclevine/agouti"
	"log"
	"strings"
	"raven/model"
)

var driver = agouti.ChromeDriver()
	//agouti.ChromeOptions("args", []string{"--headless","--silent"}),

func StartDriver(){
	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}
}
func InitializePage() *agouti.Page{


	Page, _ := driver.NewPage()

	return Page
}

func Exit(){
	driver.Stop()

}
func DoLogin(username string, password string,Page agouti.Page){


	if err := Page.Navigate("https://www.linkedin.com/"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}


	email := Page.FindByClass(`login-email`)
	email.Fill(username)
	pass := Page.FindByClass(`login-password`)
	pass.Fill(password)
	Page.FindByID(`login-submit`).Click()

	title,_ := Page.Title()
	if strings.Contains(title,"Login"){
		log.Fatal("Linkedin login failed. Please check your credentials in the config file.")

	}else{
		log.Println("Logged in sucessfuly")
	}

}

func ParseProfile(url string,Page agouti.Page) model.Profile{

	if err := Page.Navigate(url); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	found , _ := Page.FindByClass(`pv-top-card-section__name`).Count()
	if found > 0{

		name,err := Page.FindByClass(`pv-top-card-section__name`).Text()
		if err != nil{
			log.Println("Failed extracting name")
		}
		position,err := Page.FindByClass(`pv-top-card-section__headline`).Text()
		if err != nil{
			log.Println("Failed extracting position")
		}
		company,err := Page.FindByClass(`pv-top-card-v2-section__company-name`).Text()

		if err != nil{
			log.Println(err.Error())
			log.Println("Failed extracting company")
		}

		Profile := model.Profile{}
		Profile.Name = name
		Profile.Position = position
		Profile.Company = company

		return Profile;
		}

	return model.Profile{}

}