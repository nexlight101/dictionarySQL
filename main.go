package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//dict struct maps DData map to mongodb struct
type dict struct {
	id       int
	word     string
	explain1 string
	explain2 string
	explain3 string
	explain4 string
	explain5 string
	explain6 string
}

// Read the command line
func getInput() string {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a word to search for")
	}
	var words string
	for _, v := range os.Args[1:] {
		words += v + " "
	}
	words = strings.TrimSuffix(words, " ")
	return words

}

// prints the meanings
func outPut(r dict) {
	fmt.Println()
	fmt.Printf("You searched for the meaning of (%s):\n", r.word)
	fmt.Println("Possible meanings: ")
	fmt.Printf(" %s\n", r.explain1)
	if r.explain2 != " " {
		fmt.Printf(" %s\n", r.explain2)
	}
	if r.explain3 != " " {
		fmt.Printf(" %s\n", r.explain3)
	}
	if r.explain4 != " " {
		fmt.Printf(" %s\n", r.explain4)
	}
	if r.explain5 != " " {
		fmt.Printf(" %s\n", r.explain5)
	}
}

func main() {
	// Start sql client
	fmt.Println("Connecting to mySQL")
	// Open up our database connection.
	connString := "dictionaryuser:Password10@tcp(127.0.0.1:3306)/wordsdb?charset=utf8mb4"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Cannot open connection: %v\n", err)
	}
	fmt.Println("opened connection")
	// defer the close till after the main function has finished
	defer db.Close()
	// test the link
	err = db.Ping()
	if err != nil {
		log.Fatalf("Not connected to mySQL: %v\n", err)
	}
	fmt.Println("Connected to mySQL")
	// read the commandline
	word := getInput()
	// Execute the query
	query := fmt.Sprintf("SELECT * FROM dictionary WHERE word=\"%s\"", word)

	results, err := db.Query(query)
	if err != nil {
		log.Fatalf("Cannot execute query: %v\n", err)
	}
	for results.Next() {
		record := dict{}
		// for each row, scan the result into our tag composite object
		err = results.Scan(&record.id, &record.word, &record.explain1, &record.explain2, &record.explain3, &record.explain4, &record.explain5, &record.explain6)
		if err != nil {
			log.Fatalf("Cannot Marshal record: %v\n", err)
		}
		outPut(record)
	}

}
