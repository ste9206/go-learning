package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sign("http.log.gz"))
}

// SHA1Sign returns SHA1 signature of uncompressed file
// Decompress only if file name ends with ".gz"
// cat http.log.gz | gunzip | sha1sum
func SHA1Sign(fileName string) (string, error) {

	file, err := os.Open(fileName)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// := can create shadowing - if you want to not create another var and reuse the same
	// use = only

	if !strings.HasSuffix(fileName, ".gz") {
		return "", fmt.Errorf("%q - string has not .gz suffix", fileName)
	}

	r, err := gzip.NewReader(file)

	if err != nil {
		return "", fmt.Errorf("%q - gzip %w", fileName, err)
	}
	
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy %w", fileName, err)
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}