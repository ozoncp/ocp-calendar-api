package utils

import (
	"log"
	"os"
	"time"
)

const maxAttempts = 4
const attemptTimeout = 1 * time.Second

func ReadConfig(path string) (err error) {
	var file *os.File
	attempts := 1

	defer func() {
		if file != nil {
			log.Printf("close a \"%v\"", path)
			err = file.Close()
		}
	}()

	for attempts < maxAttempts && file == nil {
		if file, err = os.Open(path); err != nil {
			attempts++
			log.Printf("error: %v, retry (attempt #%d out of %d)\n", err, attempts, maxAttempts)
			time.Sleep(attemptTimeout)
		} else {
			log.Printf("file \"%v\" has beed opened", path)
		}
	}

	return
}
