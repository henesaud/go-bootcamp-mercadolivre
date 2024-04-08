package main

import (
	"os"
)

func OpenFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic("o arquivo indicado nao foi encontrado ou esta danificado.")
	}
	return file
}
