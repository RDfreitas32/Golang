package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Cocectar retorna um ponteiro de sql que sera nosso banco e um erro
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	// verifica o status do db dado pela função Ping, a função ping retorna um nulo ou erro
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	// A função Conectar deve retornar um valor de bd e caso for nil, retorna valor zero de db, sempre um dos dois
	return db, nil
}
