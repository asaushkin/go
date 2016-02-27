jsondb
==========

JsonDB is wrapper for the data from sql databases. With JsonDB you can retrieve 
data with the JSON wrap around.

Example
=======

Using the JsonDB is very simple:

```go
package main

import (
    "github.com/asaushkin/go/database/jsondb"
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "host=10.84.0.6 dbname=timeacc sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    j, err := jsondb.NewJsonDB(db)
    fmt.Println(j.Json("select * from goings limit 1"))
}
```
