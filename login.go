package login

import (
	"bufio"
	"io"
	"encoding/csv"
	"log"
	"os"
)

type LoginDetail struct {
	Password  string
	RefreshToken string
}

var Logins = make(map[string]LoginDetail)

func ReadCsv() {
	csvFile, err := os.Open("people.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		Logins[line[0]] = LoginDetail{ Password: line[1] }
	}
}
