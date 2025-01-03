package main

import (
	"Password/account"
	"Password/files"
	"fmt"
	"github.com/fatih/color"
)

func main() {
Menu:
	for {
		switch getMenu() {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var userChoice int

	fmt.Println(color.RedString("_Менеджер паролей_"))
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	fmt.Print("Выберите действие: ")
	fmt.Scan(&userChoice)

	return userChoice
}

// Создание аккаунта и запись в JSON файл
func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		switch err.Error() {
		case "LOGIN_EMPTY":
			fmt.Println("Пустой логин")
		case "INVALID_URL":
			fmt.Println("Неверный формат URL")
		}
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(data, "data.json")
}

// Поиск аккаунта
func findAccount() {}

// Удаление аккаунта
func deleteAccount() {}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
