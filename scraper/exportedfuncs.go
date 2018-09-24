package scraper

import (
	"fmt"
	"raven/model"

	"raven/database"
)
var LUsername string
var LPassword string

func ExecuteScan(Scan model.Scan){
	StartDriver()
	Page := InitializePage()
	URLs := ExtractLinkedin(Scan.Company,Scan.Domain,Scan.Pages_Number, *Page)

	DoLogin(LUsername,LPassword,*Page)

	for _,url := range URLs {
		Profile := ParseProfile(url,*Page)
		if Profile != (model.Profile{}){
			fmt.Println(Profile)
			database.InsertPerson(Profile,Scan.Scan_name)
		}

	}
	Exit()
}