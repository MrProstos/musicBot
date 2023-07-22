package Utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

var envVariable = make(map[string]string)

func GetEnv(key string) string {
	if len(envVariable) == 0 {
		readEnvFile()
	}

	value, ok := envVariable[key]
	if !ok {
		panic("env key: " + key + " not exists")
	}

	return value
}

func readEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalln("env file not exists")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if line == nil {
			continue
		}

		data := bytes.Split(line, []byte{61})
		envVariable[string(data[0])] = string(data[1])
	}
}
