package sqlc_test

import (
	"context"
	"database/sql"
	"log"

	"github.com/alex-paperspace/sqlc"
	"github.com/alex-paperspace/sqlc/example"
)

func Example_build() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "dsn")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := example.New(sqlc.Wrap(db))
	authors, err := query.ListAuthors(sqlc.Build(ctx, func(b *sqlc.Builder) {
		b.Where("age > $1", 10)
		b.Where("name = $2", "foo")
		b.In("id", 1, 2, 3)
		b.Order("age,name DESC")
		b.Limit(10)
	}))

	if err != nil {
		log.Fatalln("ListAuthors", err)
	}

	log.Printf("list %d authors", len(authors))
}
