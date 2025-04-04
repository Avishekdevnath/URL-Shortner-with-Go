package postgres

import (
	"testing"
	"Backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupPostgres(t *testing.T) *PostgresStore {
	dsn := "host=localhost user=postgres password=1234 dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to Postgres: %v", err)
	}
	db.AutoMigrate(&models.URL{}, &models.User{})
	return &PostgresStore{db: db}
}

func TestPostgresStore(t *testing.T) {
	store := setupPostgres(t)

	// Test Save
	req := &models.ShortURLRequest{URL: "https://example.com"}
	resp, err := store.Save(req)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}
	if resp.ShortURL == "" || resp.OriginalURL != req.URL {
		t.Errorf("Save returned invalid response: %+v", resp)
	}

	// Test Get
	got, err := store.Get(resp.ShortURL)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got.OriginalURL != req.URL {
		t.Errorf("Get returned wrong URL: got %s, want %s", got.OriginalURL, req.URL)
	}

	// Test GetAll
	urls, err := store.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}
	if len(urls) != 1 || urls[0].OriginalURL != req.URL {
		t.Errorf("GetAll returned wrong data: %+v", urls)
	}
	if urls[0].UserID != nil {
		t.Errorf("UserID should be nil for anonymous: got %v", urls[0].UserID)
	}
}