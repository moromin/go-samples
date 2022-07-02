package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// model
// type ToDo struct {
// 	Id   int
// 	Name string
// }

// func Create(db *sqlx.DB, t *ToDo) error {
// 	_, err := db.Exec("INSERT INTO todo(name) VALUES ( ? )", t.Name)
// 	return err
// }

// func Read(db *sqlx.DB, id int) (*ToDo, error) {
// 	t := &ToDo{}
// 	err := db.QueryRow("SELECT id, name FROM todo WHERE id = ?", id).Scan(&t.Id, &t.Name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return t, nil
// }

// func Update(db *sqlx.DB, t *ToDo) error {
// 	_, err := db.Exec("UPDATE todo SET name = ? WHERE id = ?", t.Name, t.Id)
// 	return err
// }

// func Delete(db *sqlx.DB, id int) error {
// 	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
// 	return err
// }

// func main() {
// 	d, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/mydb?charset=utf8mb4")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	todo := &ToDo{Name: "read books"}
// 	err = Create(d, todo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// ...
// }
