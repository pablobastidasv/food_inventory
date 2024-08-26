package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	total := 0

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fmt.Println("========")
		fmt.Printf("path %s\n", path)

		if !strings.HasSuffix(info.Name(), ".go") {
			// interested only in chages in go files
            fmt.Println("File ignored")
			return nil
		}

		pathHash := sha256.New()
		pathHash.Write([]byte(path))

		fileHash := sha256.New()
		fileContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		fileHash.Write(fileContent)

		pathHashedValue := hex.EncodeToString(pathHash.Sum(nil))
		fileContextHashedValue := hex.EncodeToString(fileHash.Sum(nil))


        resultHasher := sha256.New()
		resultHasher.Write([]byte(pathHashedValue + fileContextHashedValue))
        hashResult := hex.EncodeToString(resultHasher.Sum(nil))

		// fmt.Printf("go hashed path? %s\n", pathHashedValue)
		// fmt.Printf("go hashed file content? %s\n", fileContextHashedValue)
		fmt.Printf("go hashed total? %s\n", hashResult)
		fmt.Println("========")
        // 8e1fc8655849d3db982c590424e0aa28b7dc7e5d647c0c0b90b77dcd539b9a95
        // 
		total++
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ammount of files go files found %d\n", total)
}
