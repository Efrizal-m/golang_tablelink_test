package config

import (
	"fmt"
	// "github.com/qoharu/movie-app/internal/movie"

	"github.com/jinzhu/gorm"
)

// DBCon ... Database Connection Instance
var dbCon *gorm.DB

// InitDB ... function to initialize database
func InitDB() {
	var err error

	var dbHost = "localhost"
	var dbPort = "5432"
	var dbUser = "postgres"
	var dbPass = "postgres"
	var dbName = "golang"

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	//connect to postgres database
	dbCon, err = gorm.Open("postgres", dbString)

	//migrate database
	// dbCon.AutoMigrate(&movie.Movie{})

	if err != nil {
		panic(err)
	}

}

// GetDBConnection ...
func GetDBConnection() *gorm.DB {
	InitDB()

	return dbCon
}
