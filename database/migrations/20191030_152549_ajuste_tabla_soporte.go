package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteTablaSoporte_20191030_152549 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteTablaSoporte_20191030_152549{}
	m.Created = "20191030_152549"

	migration.Register("AjusteTablaSoporte_20191030_152549", m)
}

// Run the migrations
func (m *AjusteTablaSoporte_20191030_152549) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento ADD ordenador_id integer;")
	m.SQL("CREATE TABLE entradas.soporte_entrada (id serial NOT NULL,documento_id integer NOT NULL,activo boolean NOT NULL,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,entrada_elemento_id integer,CONSTRAINT pk_soporte_entrada PRIMARY KEY (id));")
	m.SQL("ALTER TABLE entradas.soporte_entrada ADD CONSTRAINT fk_soporte_entrada_entrada_elemento FOREIGN KEY (entrada_elemento_id)REFERENCES entradas.entrada_elemento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *AjusteTablaSoporte_20191030_152549) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE entradas.entrada_elemento DROP COLUMN ordenador_id;")
	m.SQL("ALTER TABLE entradas.soporte_entrada DROP CONSTRAINT IF EXISTS fk_soporte_entrada_entrada_elemento CASCADE;")
	m.SQL("DROP TABLE IF EXISTS entradas.soporte_entrada CASCADE;")
}
