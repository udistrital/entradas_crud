-- Inserts tipo_entrada
INSERT INTO entradas.tipo_entrada (nombre, descripcion, codigo_abreviacion, activo, numero_orden, clasificacion, fecha_creacion, fecha_modificacion)
	VALUES ('Sobrante', 'Entrada de elemento por sobrante', 'Sobrante', true, 1, 'Interna', now(), now());
INSERT INTO entradas.tipo_entrada (nombre, descripcion, codigo_abreviacion, activo, numero_orden, clasificacion, fecha_creacion, fecha_modificacion)
	VALUES ('Producción Propia', 'Entrada de elemento por producción propia', 'ProdPropia', true, 2, 'Interna', now(), now());
INSERT INTO entradas.tipo_entrada (nombre, descripcion, codigo_abreviacion, activo, numero_orden, clasificacion, fecha_creacion, fecha_modificacion)
	VALUES ('Reposición', 'Entrada de elemento por reposicíón', 'Reposición', true, 3, 'Interna', now(), now());

INSERT INTO entradas.tipo_entrada (nombre, descripcion, codigo_abreviacion, activo, numero_orden, clasificacion, fecha_creacion, fecha_modificacion)
	VALUES ('Donación', 'Entrada de elemento por donacion', 'Donación', true, 4, 'Externa', now(), now());
INSERT INTO entradas.tipo_entrada (nombre, descripcion, codigo_abreviacion, activo, numero_orden, clasificacion, fecha_creacion, fecha_modificacion)
	VALUES ('Adquisicón', 'Entrada de elemento por adquisición', 'Adquisición', true, 5, 'Externa', now(), now());
INSERT INTO entradas.tipo_entrada (nombre, descripcion, codigo_abreviacion, activo, numero_orden, clasificacion, fecha_creacion, fecha_modificacion)
	VALUES ('Terceros', 'Entrada de elemento por terceros', 'Terceros', true, 6, 'Externa', now(), now());
