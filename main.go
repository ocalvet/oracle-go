package main
 
import (
    "fmt"
    "database/sql"
    _ "gopkg.in/goracle.v2"
)
 
func main(){
 
    db, err := sql.Open("goracle", "localhost:1521/oracle")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
     
     
    rows,err := db.Query("select sysdate from dual")
    if err != nil {
        fmt.Println("Error running query")
        fmt.Println(err)
        return
    }
    defer rows.Close()
 
    var thedate string
    for rows.Next() {
        rows.Scan(&thedate)
    }
    fmt.Printf("The date is: %s\n", thedate)
}