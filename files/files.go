package files

import (
	"fmt"
	"os"
)

func ReadFile(path string) {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
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
