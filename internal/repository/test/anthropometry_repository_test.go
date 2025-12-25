package test

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	// "github.com/google/uuid"
	// "github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/pkg/database"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file untuk koneksi database
	if err := godotenv.Load("../../.env"); err != nil {
		panic("Error loading .env file for tests")
	}
}

func TestGetFromChildID(t *testing.T) {
	// Setup: koneksi ke database
	db := database.InitDB()
	repo := repository.NewAnthropometryRepository(db)

	// Arrange: UUID child yang sudah ada di database
	// GANTI dengan UUID child yang kamu insert via Beekeeper/HeidiSQL
	// childID := uuid.MustParse("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")

	// Act: jalankan fungsi yang ingin ditest
	// var bind dto.AnthropometryResponse
	result, err := repo.Get(context.Background())

	// Assert: cek hasilnya
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Log hasil (seperti dd() di Laravel)
	t.Logf("Result count: %d", len(result))
	for i, record := range result {
		t.Logf("Record[%d]: Weight=%.2f, Height=%.2f, ChildID=%v",
			i, record.Weight, record.Height, record.ChildID)
	}
	t.Log(spew.Sdump(result)) // Sdump returns string, Dump prints directly to stdout

	// Contoh assertion tambahan
	if len(result) == 0 {
		t.Log("WARNING: No records found. Make sure you have inserted test data.")
	}
}
