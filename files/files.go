package files

import (
	"Password/output"
	"github.com/fatih/color"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	len, err := file.Write(content) // число байт записанных
	defer file.Close()              // выполнить в конце stack frame
	if err != nil {
		output.PrintError(err)
		return
	}
	color.Green("Запись успешна, число байт ", len)
}
