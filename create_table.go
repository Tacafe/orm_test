package main

import (
	"context"
	"log"
	"os"

	"orm_test/ent" // current dirの名前（module名）をrootに、ormのmoduleを指定
	"orm_test/ent/migrate"

	_ "github.com/lib/pq"
)

func main() {
	// db client
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=orm_test_dev password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.WriteTo(ctx, os.Stdout, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	log.Print("ent sample done.")
}
