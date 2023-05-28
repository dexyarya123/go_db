package belajar_golang_db

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsernameAndId(name string, id string) (string, string) {
	return name, id
}
func TestConnectionDatabase(t *testing.T) {
	connection := getConnectionDb()
	defer connection.Close()
	ctx := context.Background()
	id, name := GetUsernameAndId("Papa", "joko widodo")
	stmt, err := connection.PrepareContext(ctx, "INSERT INTO customer (id, username) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id, name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Berhasil Insert Ke db")
}

func TestSeletDatabase(t *testing.T) {
	connection := getConnectionDb()
	defer connection.Close()
	ctx := context.Background()

	query := "SELECT id,username FROM customer"
	Rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var id, usename string
		err := Rows.Scan(&id, &usename)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", usename)
	}
	fmt.Println("Berhasil Insert Ke db")
}

func TestQuerySqlComplect(t *testing.T) {
	connection := getConnectionDb()
	defer connection.Close()
	ctx := context.Background()

	query := "SELECT id,username,email,balance,rating,birthdate,is_married,created_at FROM customer"
	Rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var id, usename, email string
		var balance int32
		var rating float64
		var is_married bool
		var birthdate, created_at time.Time
		err := Rows.Scan(&id, &usename, &email, &balance, &rating, &birthdate, &is_married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("============")
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", usename)
		fmt.Println("email : ", email)
		fmt.Println("balance : ", balance)
		fmt.Println("Rating : ", rating)
		fmt.Println("birtday : ", birthdate)
		fmt.Println("MARRIED : ", is_married)
		fmt.Println("created : ", created_at)

	}
	fmt.Println("Berhasil Insert Ke db")

}
func TestSQLInjection(t *testing.T) {
	connection := getConnectionDb()
	defer connection.Close()
	username := "RIFKIGANTENG'; #"
	password := "salah"

	ctx := context.Background()

	sqlQuery := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1 "

	rows, err := connection.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("GAGAL LOGIN")
	}
}
func TestSQLInjectionSafe(t *testing.T) {
	connection := getConnectionDb()
	defer connection.Close()
	username := "RIFKIGANTENG"
	password := "1029394"

	ctx := context.Background()

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := connection.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("GAGAL LOGIN")
	}
}

func TestAutoIncrement(t *testing.T) {
	connection := getConnectionDb()
	defer connection.Close()
	ctx := context.Background()
	email := "koko@gmail.com"
	comment := "Rifki gaanteng Banget sih"

	sqlQuery := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := connection.ExecContext(ctx, sqlQuery, email, comment)

	if err != nil {
		panic(err.Error())
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	sqlQuerySelect := "SELECT email FROM comments WHERE id = ?"
	rows, err := connection.QueryContext(ctx, sqlQuerySelect, insertId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var email string
		rows.Scan(&email)
		fmt.Println(email)
	}
}
func TestPrepareStatemant(t *testing.T) {
	db := getConnectionDb()

	ctx := context.Background()
	querySql := "INSERT INTO comments(email,comment) VALUES(?,?)"
	stmt, err := db.PrepareContext(ctx, querySql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		email := "koko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Rifki gaanteng Banget sih" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err.Error())
		}
		lastChildId, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ini adalah Id Terakhur", lastChildId)

	}

}

func TestTransaction(t *testing.T) {
	db := getConnectionDb()
	ctx := context.Background()

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	//  do transaction
	for i := 0; i < 10; i++ {
		email := "koko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Rifki gaanteng Banget sih" + strconv.Itoa(i)
		querySql := "INSERT INTO comments(email,comment) VALUES(?,?)"
		result, err := tx.ExecContext(ctx, querySql, email, comment)
		if err != nil {
			panic(err.Error())
		}
		lastChildId, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ini adalah Id Terakhur", lastChildId)

	}

	err = tx.Rollback()
	if err != nil {
		panic(err.Error())
	}
}
