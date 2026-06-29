#!/usr/bin/env bash
# ===============================
# ZenSec Batch Processor (Bash)
# ===============================

set -e

echo "==============================="
echo " ZenSec Batch Processor"
echo "==============================="
echo ""

read -p "Enter the full path to the folder to process: " folder
if [ ! -d "$folder" ]; then
    echo "Folder does not exist."
    exit 1
fi

read -p "Enter the full path to your keyfile: " keyfile
if [ ! -f "$keyfile" ]; then
    echo "Keyfile does not exist."
    exit 1
fi

read -p "Do you want to Encrypt (E) or Decrypt (D)? [E/D]: " mode

if [[ "$mode" == "E" || "$mode" == "e" ]]; then
    echo ""
    echo "Starting ENCRYPTION process..."
    # Find all files not ending in .enc
    find "$folder" -type f ! -name "*.enc" | while read -r file; do
        echo "Encrypting: $file"
        zensec -encrypt -file "$file" -keyfile "$keyfile"
    done
elif [[ "$mode" == "D" || "$mode" == "d" ]]; then
    echo ""
    echo "Starting DECRYPTION process..."
    # Find all files ending in .enc
    find "$folder" -type f -name "*.enc" | while read -r file; do
        echo "Decrypting: $file"
        zensec -decrypt -file "$file" -keyfile "$keyfile"
    done
else
    echo "Invalid choice."
    exit 1
fi

echo ""
echo "Process complete."
