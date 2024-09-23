package hangmanclassic

import (
	/*"io/ioutil"
	"math/rand"*/
	"bufio"
	"os"
)

var maxtentative int = 10

func readWordsFromFile(filename string) ([]string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}
