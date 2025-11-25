package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type VanityKeypair struct {
	ID                 string
	PublicKey          string
	SecretKeyBase58    string
	SecretKey          []byte
	IsUsed             bool
	CreatedAt          time.Time
	UsedAt             *time.Time
	VanityPattern      string
	GenerationAttempts int
}

type Database struct {
	conn *sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{conn: db}, nil
}

func (db *Database) CreateKeypair(publicKey, secretKeyBase58 string, secretKey []byte, pattern string, attempts int) error {
	query := `
		INSERT INTO vanity_keypairs (id, public_key, secret_key_base58, secret_key, isUsed, created_at, vanity_pattern, generation_attempts)
		VALUES (gen_random_uuid(), $1, $2, $3, FALSE, NOW(), $4, $5)
	`
	_, err := db.conn.Exec(query, publicKey, secretKeyBase58, secretKey, pattern, attempts)
	if err != nil {
		return fmt.Errorf("failed to insert keypair: %w", err)
	}
	return nil
}

func (db *Database) GetAvailableCount() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM vanity_keypairs WHERE isUsed = FALSE`
	err := db.conn.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count available keypairs: %w", err)
	}
	return count, nil
}

func (db *Database) KeypairExists(publicKey string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM vanity_keypairs WHERE public_key = $1)`
	err := db.conn.QueryRow(query, publicKey).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check keypair existence: %w", err)
	}
	return exists, nil
}

func (db *Database) GetStats() (available, total int, avgAttempts float64, err error) {
	query := `
		SELECT 
			COUNT(CASE WHEN isUsed = FALSE THEN 1 END) as available,
			COUNT(*) as total,
			COALESCE(AVG(generation_attempts), 0) as avg_attempts
		FROM vanity_keypairs
	`
	err = db.conn.QueryRow(query).Scan(&available, &total, &avgAttempts)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to get stats: %w", err)
	}
	return available, total, avgAttempts, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}
