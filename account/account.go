package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand/v2"
	"net/url"
	_ "reflect"
	"time"
)

type Account struct {
	login    string `json:"login" xml:"test"` // тэги
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func newAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

func NewAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	//var field, _ = reflect.TypeOf(newAcc).Elem().FieldByName("login")
	//fmt.Println(string(field.Tag))

	if password == "" {
		newAcc.generatePassword(12)
		//newAcc.Account.generatePassword(12)
	}

	return newAcc, nil
}

func (acc accountWithTimeStamp) OutputPassword() {
	color.Red(acc.login)
	fmt.Printf("Login: %s\nPassword: %s\nURL: %s\ncreatedAt: %s\nupdatedAt: %s", acc.login, acc.password, acc.url, acc.createdAt, acc.updatedAt)
}

func (acc Account) OutputPassword() {
	fmt.Printf("Login: %s\nPassword: %s\nURL: %s", acc.login, acc.password, acc.url)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func (acc *Account) generatePassword(n int) {
	password := ""
	for i := 0; i < n; i++ {
		randomIndex := rand.IntN(len(letterRunes))
		password += string(letterRunes[randomIndex])
	}
	acc.password = password
}
