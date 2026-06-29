package crypto

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptDecryptCycle(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := filepath.Join(tempDir, "test.txt")
	encryptedFile := filepath.Join(tempDir, "test.txt.enc")
	decryptedFile := filepath.Join(tempDir, "test.txt.dec")

	password := []byte("super_secure_password")
	data := []byte("hello world, this is a test of the emergency broadcast system. ")
	// make it span multiple chunks
	largeData := bytes.Repeat(data, 5000)

	if err := os.WriteFile(originalFile, largeData, 0644); err != nil {
		t.Fatal(err)
	}

	if err := EncryptFile(originalFile, encryptedFile, password); err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	if err := DecryptFile(encryptedFile, decryptedFile, password); err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	decryptedData, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(largeData, decryptedData) {
		t.Fatal("decrypted data does not match original")
	}
}

func TestWrongPassword(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := filepath.Join(tempDir, "test.txt")
	encryptedFile := filepath.Join(tempDir, "test.txt.enc")
	decryptedFile := filepath.Join(tempDir, "test.txt.dec")

	password := []byte("correct_password")
	os.WriteFile(originalFile, []byte("secret data"), 0644)
	EncryptFile(originalFile, encryptedFile, password)

	err := DecryptFile(encryptedFile, decryptedFile, []byte("wrong_password"))
	if err != ErrDecryption {
		t.Fatalf("expected ErrDecryption, got: %v", err)
	}
}

func TestTruncationAttack(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := filepath.Join(tempDir, "test.txt")
	encryptedFile := filepath.Join(tempDir, "test.txt.enc")
	decryptedFile := filepath.Join(tempDir, "test.txt.dec")

	password := []byte("password")
	os.WriteFile(originalFile, []byte("secret data"), 0644)
	EncryptFile(originalFile, encryptedFile, password)

	// Truncate the encrypted file by removing the last 10 bytes
	fileInfo, _ := os.Stat(encryptedFile)
	os.Truncate(encryptedFile, fileInfo.Size()-10)

	err := DecryptFile(encryptedFile, decryptedFile, password)
	if err != ErrDecryption {
		t.Fatalf("expected ErrDecryption due to truncation, got: %v", err)
	}
}

func TestChunkReorderingAttack(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := filepath.Join(tempDir, "test.txt")
	encryptedFile := filepath.Join(tempDir, "test.txt.enc")
	decryptedFile := filepath.Join(tempDir, "test.txt.dec")

	password := []byte("password")
	// Make sure the file is larger than chunkSize (64KB) to have at least 2 chunks
	data1 := bytes.Repeat([]byte("A"), chunkSize)
	data2 := bytes.Repeat([]byte("B"), 100)
	os.WriteFile(originalFile, append(data1, data2...), 0644)
	
	EncryptFile(originalFile, encryptedFile, password)

	// Read the encrypted file
	encData, err := os.ReadFile(encryptedFile)
	if err != nil {
		t.Fatal(err)
	}
	
	// Header is 25 bytes.
	// Chunk 1 is 64KB + 16 bytes overhead = 65552 bytes.
	headerLen := 25
	chunk1Len := chunkSize + 16
	
	if len(encData) < headerLen+chunk1Len+16 {
		t.Fatal("Encrypted data too small")
	}
	
	// Swap chunk 1 and chunk 2
	chunk1Start := headerLen
	chunk1End := headerLen + chunk1Len
	chunk2Start := chunk1End
	chunk2End := len(encData)
	
	tamperedData := make([]byte, 0, len(encData))
	tamperedData = append(tamperedData, encData[:headerLen]...)
	tamperedData = append(tamperedData, encData[chunk2Start:chunk2End]...) // Put chunk 2 first
	tamperedData = append(tamperedData, encData[chunk1Start:chunk1End]...) // Put chunk 1 second
	
	os.WriteFile(encryptedFile, tamperedData, 0644)

	err = DecryptFile(encryptedFile, decryptedFile, password)
	if err != ErrDecryption {
		t.Fatalf("expected ErrDecryption due to chunk reordering, got: %v", err)
	}
}
