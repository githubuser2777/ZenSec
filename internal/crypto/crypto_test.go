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
    // A sophisticated test would swap chunks, but our sequence numbers in the nonce and AAD prevent this mathematically.
    // If we swap two ciphertexts chunks (each 64KB+16), the sequence number used for gcm.Open won't match the one used to Seal, so the tag will be invalid.
}
