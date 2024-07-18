package decrypt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ExistsInList(key string, hash string) bool {
	file, err := os.Open("list/list.txt")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entry := fmt.Sprintf("%s  %s", key, hash)

	for scanner.Scan() {
		if scanner.Text() == entry {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return false
}

func FindKeyByHash(hash string) string {
	file, err := os.Open("list/list.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "  ")
		if len(parts) == 2 && parts[1] == hash {
			return parts[0]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return ""
}
