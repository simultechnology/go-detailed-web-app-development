package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Option struct {
	Id       string `json:"option_id"`
	Name     string `json:"option_name"`
	Value    string `json:"option_value"`
	Autoload bool   `json:"autoload"`
}

func main() {
	fmt.Println("start!!")

	ctx := context.Background()

	db, err := sql.Open("mysql", "wordpress:wordpress@tcp(localhost:33306)/wordpress")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.QueryContext(ctx, "SELECT * FROM wp_options")
	if errors.Is(err, sql.ErrNoRows) { // this part will never been executed.
		fmt.Println("no record")
	} else if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//stmtOut, err := db.Prepare("SELECT * FROM wordpress.wp_options")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//defer stmtOut.Close()
	//
	//var squareNum int // we "scan" the result in here
	//
	//// Query the square-number of 13
	//rows, err := stmtOut.Query() // WHERE number = 13
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	var options []Option
	for rows.Next() {
		option := Option{}
		rows.Scan(&option.Id, &option.Name, &option.Value, &option.Autoload)
		//fmt.Println(option)
		options = append(options, option)
	}
	//fmt.Printf("%v", options)

	b, err := json.Marshal(options)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// result := string(b)
	// fmt.Println(result)

	file, err := os.Create("result.json")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = file.Write(b)
	if err != nil {
		return
	}
}
