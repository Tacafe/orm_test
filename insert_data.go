package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"orm_test/ent"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func readCSV(filename string, idx int) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// [][]stringなのでループする
	documents := []string{}
	for _, v := range rows {
		documents = append(documents, v[idx])
	}
	return documents
}

func main() {
	// read sample csv data
	filename := flag.String("f", "sample_data.csv", "filename")
	idx := flag.Int("i", 0, "field index")
	flag.Parse()

	documents := readCSV(*filename, *idx)

	// establish db connection
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=orm_test_dev password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	td, _ := client.TestData.
		Query().
		All(ctx)
	fmt.Printf("TestData row count: %d\n", len(td))

	timeTime := time.Date(2022, 4, 1, 9, 0, 0, 0, time.Local)
	for _, document := range documents {
		client.TestData.
			Create().
			SetText(document).
			SetCreatedAt(timeTime).
			SetUpdatedAt(timeTime).
			SaveX(ctx)
	}

	td, _ = client.TestData.
		Query().
		All(ctx)
	fmt.Printf("TestData row count: %d\n", len(td))
}
