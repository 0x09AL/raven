package scraper

import (
	"github.com/sclevine/agouti"
	"fmt"
	"regexp"
	"net/url"
)

// Taken from stackoverflow :p

func RemoveDuplicatesFromSlice(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; ok {
		} else {
			m[item] = true
		}
	}
	var result []string
	for item, _ := range m {
		result = append(result, item)
	}
	return result
}

func ExtractLinkedin(company string,state string,pagesNumber int,page agouti.Page) []string {
	var URLs []string



	for i := 0 ; i < pagesNumber; i++{

		UrlDork := fmt.Sprintf("https://www.google.com/search?q=site:%s.linkedin.com/in/ Current. %s&t=h_&ia=web&start=%d",state,url.QueryEscape(company),i*10)


		page.Navigate(UrlDork)
		//UrlCount , _ := page.AllByClass(`f`).Count()
		UrlCount , _ := page.AllByXPath(`//a[@href]`).Count()
		for i:=0; i < UrlCount; i++ {
			temp,_ := page.AllByXPath(`//a[@href]`).At(i).Text()
			r := regexp.MustCompile(`https:\/\/([a-zA-Z]*)\.linkedin\.com([a-zA-Z0-9\/-]*)`)
			URLs = append(URLs,r.FindString(temp))

		}

	}
	URLs = RemoveDuplicatesFromSlice(URLs)
	return URLs
}
