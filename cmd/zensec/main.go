package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/githubuser2777/ZenSec/internal/crypto"
	"golang.org/x/term"
)

func main() {
	encryptCmd := flag.Bool("encrypt", false, "Encrypt the specified file")
	decryptCmd := flag.Bool("decrypt", false, "Decrypt the specified file")
	filePath := flag.String("file", "", "Path to the target file")
	keyfilePath := flag.String("keyfile", "", "Path to a key file to use instead of a password (optional)")

	flag.Parse()

	if !*encryptCmd && !*decryptCmd {
		fmt.Println("Error: Must specify either -encrypt or -decrypt")
		flag.Usage()
		os.Exit(1)
	}

	if *encryptCmd && *decryptCmd {
		fmt.Println("Error: Cannot specify both -encrypt and -decrypt")
		flag.Usage()
		os.Exit(1)
	}

	if *filePath == "" {
		fmt.Println("Error: Must specify a file using -file")
		flag.Usage()
		os.Exit(1)
	}

	var password []byte
	if *keyfilePath != "" {
		var err error
		password, err = os.ReadFile(*keyfilePath)
		if err != nil {
			fmt.Printf("Error reading keyfile: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print("Enter password: ")
		var err error
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Printf("\nError reading password: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()

		if *encryptCmd {
			fmt.Print("Confirm password: ")
			confirm, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Printf("\nError reading password confirmation: %v\n", err)
				os.Exit(1)
			}
			fmt.Println()

			if !bytes.Equal(password, confirm) {
				fmt.Println("Error: Passwords do not match")
				os.Exit(1)
			}
		}
	}

	if len(password) == 0 {
		fmt.Println("Error: Password/Key cannot be empty")
		os.Exit(1)
	}

	if *encryptCmd {
		outPath := *filePath + ".enc"
		if !checkOverwrite(outPath) {
			os.Exit(0)
		}
		fmt.Printf("Encrypting %s to %s...\n", *filePath, outPath)
		if err := crypto.EncryptFile(*filePath, outPath, password); err != nil {
			fmt.Printf("Encryption failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Encryption successful!")
	} else if *decryptCmd {
		outPath := strings.TrimSuffix(*filePath, ".enc")
		if outPath == *filePath {
			outPath = *filePath + ".dec" // fallback if doesn't end with .enc
		}
		if !checkOverwrite(outPath) {
			os.Exit(0)
		}
		fmt.Printf("Decrypting %s to %s...\n", *filePath, outPath)
		if err := crypto.DecryptFile(*filePath, outPath, password); err != nil {
			fmt.Printf("Decryption failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Decryption successful!")
	}
}

func checkOverwrite(path string) bool {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File %s already exists. Overwrite? [y/N]: ", path)
		var response string
		fmt.Scanln(&response)
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			fmt.Println("Operation cancelled.")
			return false
		}
	}
	return true
}
