package database

import (
	"raven-go/model"
	"log"
	"fmt"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"

	"github.com/olekukonko/tablewriter"
	"os"

)

func InsertScan(scan model.Scan){

	tx, _ := db.Begin()
	stmt, err_stmt := tx.Prepare(model.CreateScanQuery)

	if err_stmt != nil {
		log.Fatal(err_stmt)
	}


	_, err := stmt.Exec(scan.Scan_name,scan.Company,scan.Domain,scan.Pages_Number)
	tx.Commit()
	if err != nil{
		log.Println("ERROR: Error inserting scan.")
	}else{

		log.Println(fmt.Sprintf("Starting new scan %s .",scan.Scan_name))
	}
}

func GetScanId(scan string) string{
	var temp string
	err := db.QueryRow(model.GetIdByScanNameQuery, scan).Scan(&temp)
	if err != nil {
		log.Println(err)
		return ""
	}

	return temp
}

func ShowScans(){

		rows, err := db.Query(model.GetALLScans)

		scan := model.Scan{}
		if err != nil {
			panic(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Scan Name", "Company"})

		for rows.Next() {
			err := rows.Scan(&scan.Scan_id, &scan.Scan_name, &scan.Company, &scan.Domain, &scan.Pages_Number)
			table.Append([]string{scan.Scan_name,scan.Company})
			if err != nil {
				log.Fatal(err)
			}
		}
		table.Render()

}

func GetScanCompleters()  func(string) []string{

	return func(line string) []string {
		var Scans []string
		var temp string
		rows, err := db.Query(model.SelectScanNames)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			err := rows.Scan(&temp)
			Scans = append(Scans,temp)
			if err != nil {
				log.Fatal(err)
			}
		}
		return Scans
	}
}