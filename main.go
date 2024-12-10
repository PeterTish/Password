package main

import (
	"Password/account"
	"Password/files"
	"fmt"
)

func main() {
	//defer fmt.Println(1) // будет выполнен последний
	//defer fmt.Println(2)
	files.WriteFile("Привет, я файл!", "file.txt")
	files.ReadFile("file.txt")

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		switch err.Error() {
		case "LOGIN_EMPTY":
			fmt.Println("Пустой логин")
		case "INVALID_URL":
			fmt.Println("Неверный формат URL")
		}
		return
	}

	myAccount.OutputPassword()

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
