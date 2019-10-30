package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjustesDatosParametricas_20190910_154158 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjustesDatosParametricas_20190910_154158{}
	m.Created = "20190910_154158"

	migration.Register("AjustesDatosParametricas_20190910_154158", m)
}

// Run the migrations
func (m *AjustesDatosParametricas_20190910_154158) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("UPDATE entradas.tipo_entrada	SET nombre = 'Adquisición' WHERE nombre = 'Adquisicón';")
}

// Reverse the migrations
func (m *AjustesDatosParametricas_20190910_154158) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("UPDATE entradas.tipo_entrada	SET nombre = 'Adquisicón' WHERE nombre = 'Adquisición';")
}
