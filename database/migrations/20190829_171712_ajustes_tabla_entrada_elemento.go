package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjustesTablaEntradaElemento_20190829_171712 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjustesTablaEntradaElemento_20190829_171712{}
	m.Created = "20190829_171712"

	migration.Register("AjustesTablaEntradaElemento_20190829_171712", m)
}

// Run the migrations
func (m *AjustesTablaEntradaElemento_20190829_171712) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN vigencia TYPE VARCHAR(4);")
}

// Reverse the migrations
func (m *AjustesTablaEntradaElemento_20190829_171712) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN vigencia TYPE NUMERIC(4) USING (vigencia::numeric(4,0));")
}
