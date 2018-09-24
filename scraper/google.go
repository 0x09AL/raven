package scraper

import (
	"github.com/sclevine/agouti"
	"fmt"
	"regexp"
	"net/url"
)

func ExtractLinkedin(company string,state string,pagesNumber int,page agouti.Page) []string {
	var URLs []string



	for i := 0 ; i < pagesNumber; i++{

		UrlDork := fmt.Sprintf("https://www.google.com/search?q=site:%s.linkedin.com/in/ Current. %s&t=h_&ia=web&start=%d",state,url.QueryEscape(company),i*10)


		page.Navigate(UrlDork)
		UrlCount , _ := page.AllByClass(`f`).Count()

		for i:=0; i < UrlCount; i++ {
			temp,_ := page.AllByClass(`f`).At(i).Text()

			r := regexp.MustCompile(`https:\/\/([a-zA-Z]*)\.linkedin\.com([a-zA-Z0-9\/-]*)`)
			URLs = append(URLs,r.FindString(temp))

		}

	}



	return URLs
}
