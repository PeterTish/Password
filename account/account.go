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
	Login     string    `json:"login" xml:"test"` // тэги для записи в JSON
	Password  string    `json:"password" xml:"test"`
	Url       string    `json:"url" xml:"test"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		Url:       urlString,
		Login:     login,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
			Url:      urlString,
			Login:    login,
			Password: password,
		},
	}
	//var field, _ = reflect.TypeOf(newAcc).Elem().FieldByName("Login")
	//fmt.Println(string(field.Tag))

	if password == "" {
		newAcc.generatePassword(12)
		//newAcc.Account.generatePassword(12)
	}

	return newAcc, nil
}

func (acc accountWithTimeStamp) OutputPassword() {
	color.Red(acc.Login)
	fmt.Printf("Login: %s\nPassword: %s\nURL: %s\ncreatedAt: %s\nupdatedAt: %s", acc.Login, acc.Password, acc.Url, acc.createdAt, acc.updatedAt)
}

func (acc Account) Output() {
	fmt.Printf("Login: %s\nPassword: %s\nURL: %s\ncreatedAt: %s\nupdatedAt: %s\n", acc.Login, acc.Password, acc.Url, acc.CreatedAt, acc.UpdatedAt)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func (acc *Account) generatePassword(n int) {
	password := ""
	for i := 0; i < n; i++ {
		randomIndex := rand.IntN(len(letterRunes))
		password += string(letterRunes[randomIndex])
	}
	acc.Password = password
}
