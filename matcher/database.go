package matcher

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

const (
	timeLayout = "2015-01-18 00:01:54"
)

func ProcessBreeder(breeds []*Breed, state string) {
	database := os.Getenv("CodeDogDatabase")
	password := os.Getenv("CodeDogPass")
	user := os.Getenv("CodeDogUser")
	log.Println("User : " + user + ":" + password + "@/" + database)
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var breederCount int

	var breederTotalCount int // track inserted count that have already been inserted into the database.

	for _, breed := range breeds {

		var breedID int

		err := db.QueryRow("Select id from Breed where name = ?", breed.Name).Scan(&breedID)

		if err != nil {
			log.SetPrefix("WARNING (NO BREED) ")
			log.Printf("err : %s for breed : %s \n", err, breed.Name)
		}

		if breedID > 0 {
			for _, breeder := range breed.Breeder {

				err := db.QueryRow("Select count(*) from User where email = ?", breeder.Email).Scan(&breederCount)

				if err != nil {
					log.Println(err)
				}

				if breederCount == 0 {

					if breeder.Email == "" {
						log.SetPrefix("WARNING (no Email) ")
						log.Println(breeder.ID + " has no email	")
						continue
					}

					result, err := db.Exec("INSERT INTO User (email, membership_no, class_name, created_at) VALUES (?, ?, ?, ?)",
						breeder.Email, breeder.ID, "DogBreeder", time.Now().Format(time.RFC3339))

					if err != nil {
						log.SetPrefix("ERROR ")
						log.Println(err)
					}

					row, _ := result.LastInsertId()

					/* Insert Metadata  Only the State*/
					db.Exec("INSERT INTO Metadata (user_id, state) VALUES (?,?)", row, state)

					/*
						Insert many to many relationship for Breed, User.
						Table: Breed_User
					*/
					db.Exec("set foreign_key_checks=0")

					_, err2 := db.Exec("INSERT INTO Breed_User (breed_id, user_id) VALUES( ?, ?)", breedID, row)

					if err2 != nil {
						// log.Println(err2)
					}

					db.Exec("set foreign_key_checks=1")

					breederTotalCount = breederTotalCount + 1
				}
			}
		}
	}

	log.SetPrefix("INFO ")
	log.Printf("totla %d \n", breederTotalCount)

}
