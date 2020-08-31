package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjustarTablaEntradaElemento_20190827_111007 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjustarTablaEntradaElemento_20190827_111007{}
	m.Created = "20190827_111007"

	migration.Register("AjustarTablaEntradaElemento_20190827_111007", m)
}

// Run the migrations
func (m *AjustarTablaEntradaElemento_20190827_111007) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN importacion TYPE VARCHAR(15);")
	m.SQL("ALTER TABLE entradas.entrada_elemento ADD consecutivo varchar(15) NOT NULL;")
	m.SQL("ALTER TABLE entradas.entrada_elemento RENAME movimiento_id TO elemento_id; ")
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN elemento_id TYPE INTEGER USING (elemento_id::integer);")
	m.SQL("ALTER TABLE entradas.entrada_elemento ADD vigencia varchar(4) NOT NULL;")
}

// Reverse the migrations
func (m *AjustarTablaEntradaElemento_20190827_111007) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN importacion TYPE BOOLEAN USING (importacion::boolean);")
	m.SQL("ALTER TABLE entradas.entrada_elemento DROP COLUMN consecutivo;")
	m.SQL("ALTER TABLE entradas.entrada_elemento RENAME elemento_id TO movimiento_id;")
	m.SQL("ALTER TABLE entradas.entrada_elemento ALTER COLUMN movimiento_id TYPE VARCHAR(25);")
	m.SQL("ALTER TABLE entradas.entrada_elemento DROP COLUMN vigencia")
}
