package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjustesTablaEntradaElemento_20190828_170626 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjustesTablaEntradaElemento_20190828_170626{}
	m.Created = "20190828_170626"

	migration.Register("AjustesTablaEntradaElemento_20190828_170626", m)
}

// Run the migrations
func (m *AjustesTablaEntradaElemento_20190828_170626) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN importacion TYPE BOOLEAN USING (importacion::boolean);")
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN vigencia TYPE NUMERIC(4) USING (vigencia::numeric(4,0));")
}

// Reverse the migrations
func (m *AjustesTablaEntradaElemento_20190828_170626) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN importacion TYPE VARCHAR(15);")
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN vigencia TYPE VARCHAR(4);")
}
