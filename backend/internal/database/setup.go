package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

var DB *Queries
var Qctx context.Context = context.Background()

func SetupDatabase() {
	db, err := sql.Open("sqlite", "petsaway.db")
	if err != nil {
		panic(err)
	}
	DB = New(db)

	seedAdmin()
}

func seedAdmin() {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASS")

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err := DB.InsertUserByName(Qctx, InsertUserByNameParams{
		ID:           uuid.NewString(),
		Username:     username,
		PasswordHash: string(hash),
	})

	if err != nil {
		log.Fatalln(err)
	}
}
