package exporter

import (
	"raven-go/model"
	"fmt"
	"raven-go/database"
	"strings"
	"time"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func HandleExport(export model.Export,scanName string,checkPwned bool){

	var Persons []model.Person
	var person model.Person
	var breaches []model.Breach
	for _, name := range database.GetPersons(scanName){

		temp := strings.Split(name," ")
		if len(temp) > 1{
			firstname := strings.ToLower(temp[0])
			lastname := strings.ToLower(temp[1])

			if export.Domain != ""{

				switch export.Format {
				case "{firstname}.{lastname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s@%s",firstname,lastname,export.Domain)}
					Persons = append(Persons, person)
				case "{lastname}.{firstname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s@%s",lastname,firstname,export.Domain)}
					Persons = append(Persons, person)
				case "{firstname}-{lastname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s-%s@%s",firstname,lastname,export.Domain)}
					Persons = append(Persons, person)
				case "{firstname[0]}{lastname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",string(firstname[0]),lastname,export.Domain)}
					Persons = append(Persons, person)
				case "{lastname}{firstname[0]}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",lastname,string(firstname[0]),export.Domain)}
					Persons = append(Persons, person)
				case "{lastname[0]}{firstname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",string(lastname[0]),firstname,export.Domain)}
					fmt.Println(person.Email)
					Persons = append(Persons, person)
				case "{firstname}{lastname[0]}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",firstname,string(lastname[0]),export.Domain)}
					Persons = append(Persons, person)
				case "ALL":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s@%s",firstname,lastname,export.Domain)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s@%s",lastname,firstname,export.Domain)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s-%s@%s",firstname,lastname,export.Domain)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",string(firstname[0]),lastname,export.Domain)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",lastname,string(firstname[0]),export.Domain)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",string(lastname[0]),firstname,export.Domain)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s@%s",firstname,string(lastname[0]),export.Domain)}
					Persons = append(Persons, person)
				}

			} else{

				switch export.Format {
				case "{firstname}.{lastname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s",firstname,lastname)}
					Persons = append(Persons, person)
				case "{lastname}.{firstname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s",lastname,firstname)}
					Persons = append(Persons, person)
				case "{firstname}-{lastname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s-%s",firstname,lastname)}
					Persons = append(Persons, person)
				case "{firstname[0]}{lastname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",string(firstname[0]),lastname)}
					Persons = append(Persons, person)
				case "{lastname}{firstname[0]}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",lastname,string(firstname[0]))}
					Persons = append(Persons, person)
				case "{lastname[0]}{firstname}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",string(lastname[0]),firstname)}
					fmt.Println(person.Email)
					Persons = append(Persons, person)
				case "{firstname}{lastname[0]}@{domain}":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",firstname,string(lastname[0]))}
					Persons = append(Persons, person)
				case "ALL":
					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s",firstname,lastname)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s.%s",lastname,firstname)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s-%s",firstname,lastname)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",string(firstname[0]),lastname)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",lastname,string(firstname[0]))}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",string(lastname[0]),firstname)}
					Persons = append(Persons, person)

					person = model.Person{FirstName:firstname,LastName:lastname,Email:fmt.Sprintf("%s%s",firstname,string(lastname[0]))}
					Persons = append(Persons, person)
				}

			}

		}

	}


	if checkPwned == true{
		url := "https://haveibeenpwned.com/api/v2/breachedaccount/%s?truncateResponse=true"
		fmt.Println("[+] Scanning for breached e-mails [+]")
			for _,person := range Persons{
				fmt.Println(fmt.Sprintf("[+] Checking if %s is breached. [+] ",person.Email))
				resp, err := http.Get(fmt.Sprintf(url,person.Email))

				if err != nil {
					fmt.Println(err)
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				if resp.StatusCode == 200 {
					err = json.Unmarshal(body, &breaches)
					if err != nil {
						fmt.Println(err)
					}
					for _,breach := range breaches{
						fmt.Println(fmt.Sprintf("\033[31m[!]\033[0;0m E-mail %s breached in %s breach. \033[31m[!]\033[0;0m",person.Email,breach.Name))
					}

				}
			}
	}else{
		if export.Output == ""{
			export.Output = fmt.Sprintf("scan-%s-%s.csv",scanName,time.Now().Format("02-01-2006"))
		}
		output,err := os.OpenFile(export.Output,os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println(err)
		}
		output.WriteString("Firstname, Lastname, Email\n")
		for _,person := range Persons{

			output.WriteString(fmt.Sprintf("%s,%s,%s\n",person.FirstName,person.LastName,person.Email))

		}
		if err := output.Close(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("[+] Exported Data on file : %s [+]",export.Output))
	}

}