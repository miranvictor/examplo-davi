package service

import (
	db "davifiber/internal/database"
	d "davifiber/internal/domain"
)

func QueryFirst() (d.Funcionario, error) {
	var f d.Funcionario

	db.Conn.First(&f)
	//Configurar erro
	return f, nil
}

func QueryAll() ([]d.Funcionario, error) {
	var fs []d.Funcionario

	db.Conn.Find(&fs)
	//Configurar erro

	return fs, nil
}

func QueryOne(id int) (d.Funcionario, error) {
	var f d.Funcionario

	db.Conn.First(&f, id)
	//Configurar erro

	return f, nil
}
