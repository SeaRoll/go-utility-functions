package migrator

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"os"
	"sort"
)

func getFilesFromDir() []string {
	files, err := os.ReadDir("./sql/")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var fileNames []string

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	sort.Strings(fileNames)
	return fileNames
}

func getFileContent(fileName string) string {
	file, err := os.ReadFile("./sql/" + fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return string(file)
}

func createMigrationTable(conn *sql.DB) {
	_, err := conn.Exec("CREATE TABLE IF NOT EXISTS migrations (id text PRIMARY KEY, hash text NOT NULL);")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Println("MIGRATOR: Created migrations table")
}

func generateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func executeMigrations(conn *sql.DB, files []string) {
	tx, err := conn.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for _, file := range files {
		fileContent := getFileContent(file)
		hash := generateMD5Hash(fileContent)

		// Check if migration has already been executed
		var id string
		var hashFromDB string
		err = tx.QueryRow("SELECT id, hash FROM migrations WHERE id = $1", file).Scan(&id, &hashFromDB)
		if err != nil && err != sql.ErrNoRows {
			defer tx.Rollback()
			log.Fatal(err)
			os.Exit(1)
		}

		if err == nil && hash == hashFromDB {
			log.Println("MIGRATOR: Migration already executed: " + file)
			continue
		}
		if hashFromDB != "" && hash != hashFromDB {
			log.Println("MIGRATOR: Migration hash is wrong for file: " + file)
			os.Exit(1)
		}

		_, err := tx.Exec(fileContent)
		if err != nil {
			defer tx.Rollback()
			log.Fatal("MIGRATOR: Failed to execute sql file: ", err)
			os.Exit(1)
		}
		_, err = tx.Exec("INSERT INTO migrations (id, hash) VALUES ($1, $2)", file, hash)
		log.Println("MIGRATOR: Executed migration: " + file)

		if err != nil {
			defer tx.Rollback()
			log.Fatal(err)
			os.Exit(1)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Runs all migrations in the /sql folder that the output file exists in
// db - database connection
func RunMigration(db *sql.DB) {
	files := getFilesFromDir()
	createMigrationTable(db)
	executeMigrations(db, files)
}
