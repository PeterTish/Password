package main

import (
	"Password/account"
	"Password/cloud"
	"Password/output"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.RedString("_Менеджер паролей_"))
	//vault := account.NewVault(files.NewJsonDb("data.json"))
	vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))

Menu:
	for {
		switch getMenu() {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var userChoice int

	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	fmt.Print("Выберите действие: ")
	fmt.Scan(&userChoice)

	return userChoice
}

// Создание аккаунта и запись в JSON файл
func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		switch err.Error() {
		case "LOGIN_EMPTY":
			output.PrintError("Пустой логин")
		case "INVALID_URL":
			output.PrintError("Неверный формат URL")
		}
		return
	}

	vault.AddAccount(*myAccount)
}

// Поиск аккаунта по URL
func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

// Удаление аккаунта
func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
