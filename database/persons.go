package database
import(
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
	"raven-go/model"
	"fmt"
	"log"
)

func InsertPerson(Person model.Profile,scanName string){

	tx, _ := db.Begin()
	stmt, err_stmt := tx.Prepare(model.InsertPersonQuery)

	if err_stmt != nil {
		log.Fatal(err_stmt)
	}


	_, err := stmt.Exec(GetScanId(scanName),Person.Name,Person.Position,Person.Company)
	tx.Commit()
	if err != nil{
		log.Println("ERROR: Error inserting Person.")
	}else{

		log.Println(fmt.Sprintf("Inserting %s .",Person.Name))
	}
}


func GetPersons(scanName string ) []string{

	var Persons []string
	rows, err := db.Query(model.GetPersonsFromScanNameQuery, scanName)
	var person string
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&person)
		Persons = append(Persons, person)
		if err != nil {
			log.Fatal(err)
		}
	}


	return Persons
}