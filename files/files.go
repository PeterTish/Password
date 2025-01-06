package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	len, err := file.Write(content) // число байт записанных
	defer file.Close()              // выполнить вконце stack frame
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна, число байт ", len)
}
