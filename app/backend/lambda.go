package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "encoding/json"
    "os"
    "strconv"
)

var host     = os.Getenv("DBHOST")
var port, _  = strconv.ParseInt(os.Getenv("DBPT"), 0, 64)
var user     = os.Getenv("DBU")
var password = os.Getenv("DBPW")
var dbname   = os.Getenv("DBNAME")

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
    Headers    map[string]string `json:"headers"`
}

type Status struct {
    Station string `json:"station"`
    Gpst string `json:"gpst"`
    X string `json:"x"`
    Y string `json:"y"`
    Z string `json:"z"`
}

func Handler() (Response, error) {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
      panic(err)
    }
    defer db.Close()

    rows, qerr := db.Query("SELECT * FROM rtkstatuspage")
    if qerr != nil {
        panic(qerr)
    }

    statuses := []Status{}
    var station, gpst, x, y, z string
    for rows.Next() {
        rows.Scan(&station, &gpst, &x, &y, &z)
        statuses = append(statuses, Status{station, gpst, x, y, z})
    }

    statusesJson, err := json.Marshal(statuses)

	return Response{
		StatusCode: 200,
		Body:       string(statusesJson),
        Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}

func main() {
    lambda.Start(Handler)
}
