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
}

type Status struct {
    Station string `json:"station"`
    Gpst string `json:"gpst"`
    Ebaseline string `json:"ebaseline"`
    Nbaseline string `json:"nbaseline"`
    Ubaseline string `json:"ubaseline"`
    Q string `json:"q"`
    Ns string `json:"ns"`
    Sde string `json:"sde"`
    Sdn string `json:"sdn"`
    Sdu string `json:"sdu"`
    Sden string `json:"sden"`
    Sdnu string `json:"sdnu"`
    Sdue string `json:"sdue"`
    Age string `json:"age"`
    Ratio string `json:"ratio"`
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
    var station, gpst, ebaseline, nbaseline, ubaseline, q, ns, sde, sdn, sdu, sden, sdnu, sdue, age, ratio string
    for rows.Next() {
        rows.Scan(&station, &gpst, &ebaseline, &nbaseline, &ubaseline, &q, &ns, &sde, &sdn, &sdu, &sden, &sdnu, &sdue, &age, &ratio)
        statuses = append(statuses, Status{station, gpst, ebaseline, nbaseline, ubaseline, q, ns, sde, sdn, sdu, sden, sdnu, sdue, age, ratio})
    }

    statusesJson, err := json.Marshal(statuses)

	return Response{
		StatusCode: 200,
		Body:       string(statusesJson),
	}, nil
}

func main() {
    lambda.Start(Handler)
}
