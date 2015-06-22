package matcher

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//check if email exist in the database..

func ProcessBreeder(breeders []*Breeder) {
	db, err := sql.Open("mysql", "root:@/CodeDog")

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	stmtOut, err2 := db.Prepare("Select count(*) from User where email = ?")
	if err2 != nil {
		log.Fatalln(err2)
	}
	defer stmtOut.Close()

	var count int

	for _, breeder := range breeders {
		stmtOut.QueryRow(breeder.Email).Scan(count)
		if count == 0 {
			//Insert user into here
			log.Printf("%s count is %d \n", breeder.Email, count)
		}
	}

}
