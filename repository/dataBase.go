package repository

import (
	"budget_planner/logger"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

var dbFile string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	dbFile = viper.GetString("database_file")
	if dbFile == "" {
		logger.Log.Error("database_file not set in config file")
		os.Exit(1)
	}

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		// Create the database file if it doesn't exist
		file, err := os.Create(dbFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = file.Close()
		if err != nil {
			return
		}
	}
	//CreateBudgetSchema()
	CreateBudgetTables()
	CreateExpenseTables()
	CreateRevenueTables()
}

func CreateBudgetSchema() {
	db := OpenDatabase()
	createTableSQL := `CREATE SCHEMA budget;`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Budget schema created")
}

func CreateBudgetTables() {
	db := OpenDatabase()
	createTableSQL := `CREATE  TABLE IF NOT EXISTS budget ( 
	id                   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name                 TEXT(100)  NOT NULL     
 )`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Budget table created")
}

func CreateExpenseTables() {
	db := OpenDatabase()
	createTableSQL := `CREATE  TABLE IF NOT EXISTS expenses ( 
	id                   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name                 TEXT(100)  NOT NULL     ,
	amount               FLOAT UNSIGNED NOT NULL     ,
	occurrence           TEXT  NOT NULL     ,
	budget_id            INTEGER  NOT NULL     ,
	CONSTRAINT fk_expenses_budget FOREIGN KEY ( budget_id ) REFERENCES budget( id ) ON DELETE CASCADE ON UPDATE CASCADE
 )`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Expenses table created")
}

func CreateRevenueTables() {
	db := OpenDatabase()
	createTableSQL := `CREATE  TABLE IF NOT EXISTS revenues ( 
	id                   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name                 TEXT(100)  NOT NULL     ,
	amount               FLOAT UNSIGNED NOT NULL     ,
	occurrence           TEXT  NOT NULL     ,
	budget_id            INTEGER  NOT NULL     ,
	CONSTRAINT fk_revenues_budget FOREIGN KEY ( budget_id ) REFERENCES budget( id ) ON DELETE CASCADE ON UPDATE CASCADE
 )`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Revenues table created")
}

func OpenDatabase() *sql.DB {

	DB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		logger.Log.Error("Failed to open database connection", zap.Error(err))
		os.Exit(1)
	}
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		logger.Log.Error("Failed to close database connection", zap.Error(err))
	//	}
	//}(DB)
	return DB
}
