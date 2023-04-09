package csv

import (
	"encoding/csv"
	"os"
)

func SaveResult() {
	file, err := os.Create("example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"Date", "Site"})
	writer.Write([]string{"09/04/2023", "site.com"})
	writer.Flush()
}
