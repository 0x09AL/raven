package terminal

import (
	"github.com/chzyer/readline"
	"fmt"
	"io"
	"strings"
	"raven-go/model"

	"encoding/json"
	"log"
	"strconv"
	"raven/scraper"
	"raven/database"
	"raven/exporter"

)

var context string = "main"
var prompt string = "raven (\033[0;32m%s\033[0;0m)\033[31m >> \033[0;0m"


func backMain(l *readline.Instance){
	context = "main"
	l.SetPrompt(fmt.Sprintf(prompt,"main"))
	l.Config.AutoComplete = MainCompleter
}

func handleScanCreation(l *readline.Instance){
	l.SetPrompt(fmt.Sprintf(prompt,"scan"))
	l.Config.AutoComplete = ScanCompleter
	Scan := model.Scan{}
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
		if len(line) == 0 {
		break
	} else {
		continue
	}
	} else if err == io.EOF {
		break
	}

		line = strings.TrimSpace(line)

		temp := strings.Split(line," ")
		command := temp[0]
		switch command {
		case "back":
			backMain(l)
			return
		case "options":
			data, err := json.MarshalIndent(Scan,"", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", data)
		case "set":
			if len(temp) > 2{
				option := temp[1]
				param := strings.Join(temp[2:]," ")
				switch option {
				case "domain":
					Scan.Domain = param
				case "company":
					Scan.Company = param
				case "scan_name":
					Scan.Scan_name = param
				case "pages_number":
					Scan.Pages_Number,_ = strconv.Atoi(param)
				}
			}else{
				fmt.Println("Invalid option")
			}
		case "unset":
			if len(temp) > 2{
				option := temp[1]
				switch option {
				case "domain":
					Scan.Domain = ""
				case "company":
					Scan.Company = ""
				case "scan_name":
					Scan.Scan_name = ""
				case "pages_number":
					Scan.Pages_Number = 0
				}
			}else{
				fmt.Println("Invalid option")
			}
		case "start":
			log.Println("Starting scan...")
			database.InsertScan(Scan)
			scraper.ExecuteScan(Scan)
		case "":

		default:
			fmt.Println("Invalid command")

		}
		}
}

func handleExport(l *readline.Instance,scanName string){

	l.SetPrompt(fmt.Sprintf(prompt,"exporter"))
	l.Config.AutoComplete = ExportCompleter

	export := model.Export{}

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		temp := strings.Split(line," ")
		command := temp[0]
		switch command {
		case "back":
			backMain(l)
			return
		case "options":
			data, err := json.MarshalIndent(export, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", data)
		case "export":
			exporter.HandleExport(export,scanName,false)
		case "checkpwned":
			exporter.HandleExport(export,scanName,true)
		case "set":
			if len(temp) > 2 {
				option := temp[1]
				switch option {
				case "domain":
					export.Domain = temp[2]
				case "output":
					export.Output = temp[2]
				case "format":
					export.Format = temp[2]
				}
			} else {
				fmt.Println("Invalid option")
			}
		case "unset":
			if len(temp) > 2 {
				option := temp[1]
				switch option {
				case "domain":
					export.Domain = ""
				case "output":
					export.Output = ""
				case "format":
					export.Format = ""
				}
			}
		}
		}

}

func handleInput(line string,l *readline.Instance)  {
	var command string
	line = strings.TrimSpace(line)
	temp := strings.Split(line," ")
	if len(temp) >1 {

		command = strings.Join(temp[1:]," ")
	}

	switch {
		case strings.HasPrefix(line,"new"):
			switch command {
				case "scan":
					handleScanCreation(l)
				}
		case strings.HasPrefix(line,"scans"):
			database.ShowScans()
		case strings.HasPrefix(line,"use"):
			handleExport(l,command)
		case strings.HasPrefix(line,"exit"):
			return
	}


}

func StartTerminal()  {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          fmt.Sprintf(prompt,"main"),
		HistoryFile:     "history.tmp",
		InterruptPrompt: "CTRL+C pressed exiting",
		EOFPrompt:       "quit",
		AutoComplete:	 MainCompleter,

	})


	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {

		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		handleInput(line,l)


	}
}
