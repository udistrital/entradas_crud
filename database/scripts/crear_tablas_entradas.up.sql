-- object: entradas | type: SCHEMA --
-- DROP SCHEMA IF EXISTS entradas CASCADE;
CREATE SCHEMA entradas;
-- ddl-end --

SET search_path TO pg_catalog,public,entradas;
-- ddl-end --

-- object: entradas.entrada_elemento | type: TABLE --
-- DROP TABLE IF EXISTS entradas.entrada_elemento CASCADE;
CREATE TABLE entradas.entrada_elemento(
	id serial NOT NULL,
	solicitante integer,
	observacion varchar(250),
	importacion bool,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo bool NOT NULL,
	tipo_entrada_id integer NOT NULL,
	acta_recibido_id integer NOT NULL,
	contrato_id integer,
	movimiento_id varchar(25),
	documento_contable_id integer NOT NULL,
	CONSTRAINT pk_entrada PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE entradas.entrada_elemento IS 'Tabla para almacenar la información requerida para el procedimiento de Entrada de un elemento. Los campos se permiten NULL puesto que hay varios tipos de entrada y no todos requieren la misma información.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.solicitante IS 'Campo para almacenar el solicitante de la entrada. Solo aploca para una entrada por SOBRANTES.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.observacion IS 'Campo para almacenar las observaciones a tener en cuenta al momento de realizar una entrada.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.importacion IS 'Campo que permite almacenar si una entrada por ADQUSICIÓN es de importación.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.fecha_creacion IS 'Fecha en la que se registra la entrada. Para llevar una trazabilidad.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.fecha_modificacion IS 'Campo para almacenar la fecha en que se realiza una modificación. Para llevar una Trazabilidad.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.acta_recibido_id IS 'Referencia al Acta de recibido, se utiliza para todos los tipos de entradas.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.contrato_id IS 'Campo para almacenar el contrato que referencia los elementos a los cuales se les hará una entrada. Se referencia de ARGO. Es utilizado para una entrada por ADQUISICIÓN y TERCEROS.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.movimiento_id IS 'Campo para almacenar la placa de los elementos a los cuales se les hará entrada. Solo aplica para REPOSICIÓN. Estos elementos son diferentes a los que vienen del acta de recibido. Se busca obtener la PLACA del elemento asociada al movimiento.';
-- ddl-end --
COMMENT ON COLUMN entradas.entrada_elemento.documento_contable_id IS 'Campo para relacionar el documento contable (P8) a la entrada. Se referencia de KRONOS';
-- ddl-end --

-- object: entradas.tipo_entrada | type: TABLE --
-- DROP TABLE IF EXISTS entradas.tipo_entrada CASCADE;
CREATE TABLE entradas.tipo_entrada(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar,
	codigo_abreviacion varchar(50),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	clasificacion varchar,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_entrada PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE entradas.tipo_entrada IS 'Tabla paramétrica para almacenar los diferentes tipos de entradas. Pueden ser: SOBRANTE, PRODUCCIÓN PROPIA, RECUPERACIÓN, DONACIÓN, ADQUISICIÓN.';
-- ddl-end --
COMMENT ON COLUMN entradas.tipo_entrada.clasificacion IS 'Para tipificar si es interno o externo';
-- ddl-end --
COMMENT ON COLUMN entradas.tipo_entrada.fecha_creacion IS 'Campo para almacenar la fecha en que se realiza el registro. Permite llevar una trazabilidad.';
-- ddl-end --
COMMENT ON COLUMN entradas.tipo_entrada.fecha_modificacion IS 'Campo para almacenar las fechas en que se realizan cambios. Permite llevar una trazabilidad.';
-- ddl-end --

-- object: fk_entrada_elemento_tipo_entrada | type: CONSTRAINT --
-- ALTER TABLE entradas.entrada_elemento DROP CONSTRAINT IF EXISTS fk_entrada_elemento_tipo_entrada CASCADE;
ALTER TABLE entradas.entrada_elemento ADD CONSTRAINT fk_entrada_elemento_tipo_entrada FOREIGN KEY (tipo_entrada_id)
REFERENCES entradas.tipo_entrada (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA entradas TO desarrollooas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA entradas TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA entradas TO desarrollooas;
