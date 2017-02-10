package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var (
		class_id   int
		class_name string
	)

	db, err := sql.Open("postgres", "postgres://kloopzcm:<password>@localhost/kloopzdb?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT class_id, class_name FROM md_classes")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&class_id, &class_name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(class_id, class_name)
	}
}
