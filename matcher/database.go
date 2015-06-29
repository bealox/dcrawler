package matcher

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//check if email exist in the database..

const (
	timeLayout = "2015-01-18 00:01:54"
)

func ProcessBreeder(breeds []*Breed) {
	db, err := sql.Open("mysql", "root:@/CodeDog")

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
			log.Printf("err : %s for breed : %s \n", err, breed.Name)
		}

		log.Printf("%s ===============> %d \n", breed.Name, breedID)

		/*
			TODO : <LOG> Breeds that don't exist in the system
			If Breed don't exist in the system, then skip
		*/

		if breedID > 0 {
			for _, breeder := range breed.Breeder {

				err := db.QueryRow("Select count(*) from User where email = ?", breeder.Email).Scan(&breederCount)

				if err != nil {
					log.Println(err)
				}

				if breederCount == 0 {
					/*
						TODO : <LOG>Logging breeder that use the same email address on differnet membership.
						Known issue:
						Some Breeders Share Email address with each other.
						I have found QLD Breeders use the same address for Dachshund breed
						Maybe need to tack this in the log and email.
					*/

					if breeder.Email == "" {
						continue
					}

					result, err := db.Exec("INSERT INTO User (email, membership_no, class_name, created_at) VALUES (?, ?, ?, ?)",
						breeder.Email, breeder.ID, "DogBreeder", time.Now().Format(time.RFC3339))

					if err != nil {
						log.Println(err)
					}

					row, _ := result.LastInsertId()

					/*
						Insert many to many relationship for Breed, User.
						Table: Breed_User
					*/
					db.Exec("set foreign_key_checks=0")

					_, err2 := db.Exec("INSERT INTO Breed_User (breed_id, user_id) VALUES( ?, ?)", breedID, row)

					if err2 != nil {
						log.Println(err2)
					}

					db.Exec("set foreign_key_checks=1")

					breederTotalCount = breederTotalCount + 1
					log.Printf("%s -- inserted %d \n", breeder.Email, row)

				}
			}
		}
	}

	log.Printf("totla %d \n", breederTotalCount)

}
