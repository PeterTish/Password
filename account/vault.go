package account

import (
	"Password/output"
	"encoding/json"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	// Встроенные интерфейсы
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

// Создание нового vault
func NewVault(db Db) *VaultWithDb {
	// Если файла не существует (не прочитали), то создаем файл
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	// Если файл существует (прочитали), то в файл добавляем
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")
		// Возвращаем пустой vault
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

// Добавление аккаунт в vault
func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

// Преобразование структуры в json
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}
	return file, nil
}

// Поиск аккаунта по URL
func (vault *VaultWithDb) FindAccountsByURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

// Удаление аккаунта по URL
func (vault *VaultWithDb) DeleteAccountByURL(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	vault.db.Write(data)
}
