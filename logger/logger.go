package main
import (
    "github.com/hpcloud/tail"
    "log"
    "strings"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"

    "os"
    "strconv"
)

var station  = os.Getenv("STATION")
var host     = os.Getenv("DBHOST")
var port, _  = strconv.ParseInt(os.Getenv("DBPT"), 0, 64)
var user     = os.Getenv("DBU")
var password = os.Getenv("DBPW")
var dbname   = os.Getenv("DBNAME")

func updateTable(db *sql.DB, gpst1 string, gpst2 string,
    ebaseline string, nbaseline string, ubaseline string, q string,
    ns string, sde string, sdn string, sdu string, sden string, sdnu string,
    sdue string, age string, ratio string) {

    sqlStatement := `
    INSERT INTO rtkstatuspage (station, gpst, ebaseline, nbaseline, ubaseline, q, ns, sde, sdn, sdu, sden, sdnu, sdue, age, ratio)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
    ON CONFLICT ON CONSTRAINT rtkstatuspage_pkey
    DO UPDATE SET station = $1, gpst = $2, ebaseline = $3, nbaseline = $4, ubaseline = $5, q = $6, ns = $7, sde = $8, sdn = $9, sdu = $10, sden = $11, sdnu = $12, sdue = $13, age = $14, ratio=$15`
    _, err := db.Exec(sqlStatement, station, gpst1 + gpst2, ebaseline, nbaseline, ubaseline, q, ns, sde, sdn, sdu, sden, sdnu, sdue, age, ratio)
    if err != nil {
      panic(err)
    }

}

func main() {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
      panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
      panic(err)
    }

    fmt.Println("Successfully connected!")

    // skip header
    t, err := tail.TailFile("logs/" + station + ".pos", tail.Config{Follow: true, ReOpen: true})
    for line := range t.Lines {
        fields := strings.Fields(line.Text)
        log.Println(fields)

        // skip comments
        if fields[0] == "%" {
            continue
        }

        gpst1, gpst2, ebaseline, nbaseline, ubaseline, q, ns, sde, sdn, sdu,
            sden, sdnu, sdue, age, ratio := fields[0], fields[1], fields[2],
                fields[3], fields[4], fields[5], fields[6], fields[7], fields[8],
                fields[9], fields[10], fields[11], fields[12], fields[13], fields[14]

        go updateTable(db, gpst1, gpst2, ebaseline, nbaseline, ubaseline, q, ns, sde, sdn, sdu, sden,
            sdnu, sdue, age, ratio)
    }
    log.Println(err)
}
