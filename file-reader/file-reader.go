package file_reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]string, error) {
	pwd, _ := os.Getwd()
	path, _ := filepath.Abs(pwd + fmt.Sprintf("/input/%s", filename))
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var fileContent []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	return fileContent, scanner.Err()
}
