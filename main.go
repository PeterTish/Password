package main

import (
	"Password/account"
	"Password/files"
	"Password/output"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.RedString("_Менеджер паролей_"))
	vault := account.NewVault(files.NewJsonDb("data.json"))
	//vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))

Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})

		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

// Создание аккаунта и запись в JSON файл
func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

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
	url := promptData([]string{"Введите URL для поиска"})
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
	url := promptData([]string{"Введите URL для поиска"})
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

//func promptData(prompt string) string {
//	fmt.Print(prompt + ": ")
//	var res string
//	fmt.Scanln(&res)
//	return res
//}

func promptData[T any](prompt []T) string {
	for i := 0; i < len(prompt); i++ {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", prompt[i])
		} else {
			fmt.Println(prompt[i])
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
