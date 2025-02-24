package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	// Configure ScyllaDB Connection
	cluster := gocql.NewCluster("192.168.117.3:9042") // Change to the correct IP
	cluster.Keyspace = "test"                         // Ensure this keyspace exists
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 10 * time.Second

	// Create session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}
	defer session.Close()

	fmt.Println("✅ Successfully connected to ScyllaDB!")

	// Create a Table
	err = session.Query(`
		CREATE TABLE IF NOT EXISTS test.users (
			id UUID PRIMARY KEY,
			name TEXT,
			email TEXT
		);
	`).Exec()
	if err != nil {
		log.Fatalf("❌ Failed to create table: %v", err)
	}
	fmt.Println("✅ Table created successfully!")

	// Insert Data
	id := gocql.TimeUUID()
	err = session.Query(`
		INSERT INTO test.users (id, name, email) VALUES (?, ?, ?);
	`, id, "John Doe", "johndoe@example.com").Exec()
	if err != nil {
		log.Fatalf("❌ Failed to insert data: %v", err)
	}
	fmt.Println("✅ Data inserted successfully!")

	// Retrieve Data
	var userID gocql.UUID
	var name, email, phone string

	iter := session.Query(`SELECT id, name, email, phone FROM test.users`).Iter()
	for iter.Scan(&userID, &name, &email, &phone) {
		fmt.Printf("📌 User: %s, Email: %s, Phone: %s\n", name, email, phone)
	}

	if err := iter.Close(); err != nil {
		log.Fatalf("❌ Failed to retrieve data: %v", err)
	}
}
