package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error // THis line is a shitty shit!!!!!!!!
func main() {
	db, err = sql.Open("mysql", "root:maxes2312733@tcp(database-1.chqdh8gfrj4n.us-east-2.rds.amazonaws.com:3306)/test01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "at index")
	check(err)
}

/*TO DISPLAY THE RECORDS FROM A SPECIFIED DB*/

func amigos(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT amigoName FROM amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(res, s)
}

/*TO CREATE A NEW TABLE TO THE DATABASE*/
func create(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "Table created", n)
}

/*TO INSERT A DATA TO THE TABLE */
func insert(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("james");`)
	check(err)
	defer stmt.Close()
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "INSERTED RECORD", n)
}

/*READ FROM THE DATABASE*/
func read(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(res, "The records are", name)
	}
}

/*Update a record */
func update(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="mahmoud" WHERE name="james";`)
	check(err)
	defer stmt.Close()
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(res, "the updated record", n)
}

/*TO DELETE A RECORD*/
func del(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="mahmoud";`)
	check(err)
	defer stmt.Close()
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(res, "RECORD IS DELETED", n)
}

/*TO DROP A TABLE FROM THE DATA BASE*/

func drop(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer ;`)
	check(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	fmt.Fprintln(res, "Table is deleted")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
