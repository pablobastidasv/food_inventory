package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Hasher struct {
	H hash.Hash
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

func main() {
	hash := ""

	ticker := time.NewTicker(100 * time.Millisecond)

	for {
		<-ticker.C
		var err error

		hasher := Hasher{
			H: sha256.New(),
		}
		err = hasher.HashProjectFiles()
		if err != nil {
			log.Fatal(err)
		}

		newHash := hasher.ToString()
		if hash != newHash {
			start := time.Now()
			hash = newHash
			fmt.Println()
			fmt.Println("Running tests...")
			cmd := exec.Command("go", "test", "./...")
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Printf("%s%v%s\n", Red, err, Reset)
				fmt.Printf("%s%v%s\n", Red, string(stdout), Reset)
			} else {
				fmt.Printf("%s%v%s\n", Green, string(stdout), Reset)
			}
			fmt.Printf("took %d sec.\n", time.Since(start).Milliseconds())
		}
	}
}

func (h *Hasher) HashProjectFiles() error {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		h.hashGoFile(path)

		return nil
	})
	return err

}

func (h *Hasher) hashGoFile(path string) error {
	// Only go files
	if !strings.HasSuffix(path, ".go") {
		// interested only in chages in go files
		// fmt.Printf("File %s ignored\n", path)
		return nil
	}

	// Ignore this go files
	if result, _ := regexp.MatchString(".*_templ.go", path); result {
		// interested only in chages in go files
		// fmt.Printf("File %s ignored\n", path)
		return nil
	}

	return h.hashFile(path)
}

func (h *Hasher) hashFile(path string) error {
	// hash the path name
	h.H.Write([]byte(path))

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	h.H.Write(fileContent)

	return nil

}

func (h *Hasher) ToString() string {
	return hex.EncodeToString(h.H.Sum(nil))
}

func (h *Hasher) Reset() {
	h.Reset()
}
