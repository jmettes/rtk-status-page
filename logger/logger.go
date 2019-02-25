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
    "sync"
)

var stationsEnv  = os.Getenv("STATIONS")
var host     = os.Getenv("DBHOST")
var port, _  = strconv.ParseInt(os.Getenv("DBPT"), 0, 64)
var user     = os.Getenv("DBU")
var password = os.Getenv("DBPW")
var dbname   = os.Getenv("DBNAME")

func updateTable(db *sql.DB, station string, gpst string, x string, y string, z string) {

    sqlStatement := `
    INSERT INTO rtkstatuspage (station, gpst, x, y, z)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT ON CONSTRAINT rtkstatuspage_pkey
    DO UPDATE SET station = $1, gpst = $2, x = $3, y = $4, z = $5`
    _, err := db.Exec(sqlStatement, station, gpst, x, y, z)
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

    stations := strings.Split(stationsEnv, ";")
    var wg sync.WaitGroup

    for _, station := range stations {
        log.Println("logging ", station)
        wg.Add(1)
        go (func(station string) {
            defer wg.Done()
            t, err := tail.TailFile("logs/" + station + ".pos", tail.Config{Follow: true, ReOpen: true})
            for line := range t.Lines {
                fields := strings.Fields(line.Text)
                log.Println(station, fields)
                // skip comments
                if fields[0] == "%" {
                    continue
                }
                gpst1, gpst2, xecef, yecef, zecef := fields[0], fields[1], fields[2],
                        fields[3], fields[4] 
                go updateTable(db, station, gpst1 + " " + gpst2, xecef, yecef, zecef)
            }
            log.Println(err)
        })(station)
    }
    wg.Wait()

}
