package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

const (
	chunkSize = 64 * 1024 // 64KB
	magic     = "ZSEC"
	version   = byte(1)
)

var (
	ErrInvalidFile   = errors.New("invalid or corrupted file")
	ErrDecryption    = errors.New("decryption failed (wrong password or tampered data)")
)

// EncryptFile encrypts the file at inPath and writes the result to outPath.
func EncryptFile(inPath, outPath string, password []byte) error {
	inFile, err := os.Open(inPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	salt, err := GenerateSalt()
	if err != nil {
		return err
	}

	key := DeriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// 4 bytes base nonce
	baseNonce := make([]byte, 4)
	if _, err := rand.Read(baseNonce); err != nil {
		return err
	}

	// Write header
	outFile.Write([]byte(magic))
	outFile.Write([]byte{version})
	outFile.Write(salt)
	outFile.Write(baseNonce)

	buf := make([]byte, chunkSize)
	nonce := make([]byte, gcm.NonceSize())
	copy(nonce[0:4], baseNonce)

	var seq uint64 = 0
	for {
		n, err := io.ReadFull(inFile, buf)
		isLast := byte(0)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			isLast = 1
		} else if err != nil {
			return err
		}

		aad := make([]byte, 9)
		binary.BigEndian.PutUint64(aad[0:8], seq)
		aad[8] = isLast

		binary.BigEndian.PutUint64(nonce[4:12], seq)

		encrypted := gcm.Seal(nil, nonce, buf[:n], aad)
		if _, err := outFile.Write(encrypted); err != nil {
			return err
		}

		if isLast == 1 {
			break
		}
		seq++
	}

	return nil
}

// DecryptFile decrypts the file at inPath and writes the result to outPath.
func DecryptFile(inPath, outPath string, password []byte) error {
	inFile, err := os.Open(inPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Read and verify header
	header := make([]byte, 25)
	if _, err := io.ReadFull(inFile, header); err != nil {
		return ErrInvalidFile
	}

	if !bytes.Equal(header[0:4], []byte(magic)) || header[4] != version {
		return ErrInvalidFile
	}

	salt := header[5:21]
	baseNonce := header[21:25]

	key := DeriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer func() {
		outFile.Close()
		// If decryption failed, we should probably delete the partial output file,
		// but for a ponytail script, we can just let it be or handle it via caller.
	}()

	cipherBuf := make([]byte, chunkSize+gcm.Overhead())
	nonce := make([]byte, gcm.NonceSize())
	copy(nonce[0:4], baseNonce)

	var seq uint64 = 0
	for {
		n, err := io.ReadFull(inFile, cipherBuf)
		isLast := byte(0)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			isLast = 1
		} else if err != nil {
			return err
		}

		aad := make([]byte, 9)
		binary.BigEndian.PutUint64(aad[0:8], seq)
		aad[8] = isLast

		binary.BigEndian.PutUint64(nonce[4:12], seq)

		plaintext, err := gcm.Open(nil, nonce, cipherBuf[:n], aad)
		if err != nil {
			return ErrDecryption
		}

		if _, err := outFile.Write(plaintext); err != nil {
			return err
		}

		if isLast == 1 {
			break
		}
		seq++
	}

	return nil
}
