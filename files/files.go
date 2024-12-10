package files

import (
	"fmt"
	"os"
)

func ReadFile(path string) {
	//file, err := os.Open("file.txt") // чтение файла по байтам
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	len, err := file.WriteString(content) // число байт записанных
	defer file.Close()                    // выполнить вконце stack frame
	if err != nil {
		//file.Close() // закрытие файла, может вернуть ошибку
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна, число байт ", len)
}
