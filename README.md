# entradas_crud

API CRUD para el módulo de entradas del sistema ARKA II.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
ENTRADAS_CRUD_PGDB=[nombre de la base de datos]
ENTRADAS_CRUD_PGPASS=[password del usuario]
ENTRADAS_CRUD_PGURLS=[direccion de la base de datos]
ENTRADAS_CRUD_PGUSER=[usuario con acceso a la base de datos]
ENTRADAS_CRUD_PGSCHEMA=[esquema donde se ubican las tablas]
ENTRADAS_CRUD_HTTP_PORT=[puerto de ejecucion] bee run
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con ENTRADAS_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/entradas_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/entradas_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
ENTRADAS_CRUD_PORT=8080 ENTRADAS_CRUD_DB_HOST=127.0.0.1:27017 ENTRADAS_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=entradas_crud . --no-cache
# docker run -p 80:80 entradas_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/entradas_crud

#2. Moverse a la carpeta del repositorio
cd entradas_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/entradas_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/entradas_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/entradas_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/entradas_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/entradas_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/entradas_crud) |

## Modelo de Datos
[Modelo de Datos API CRUD Novedades](https://user-images.githubusercontent.com/23342808/61749331-1e6a5580-ad68-11e9-8698-8523ffe3f792.png)


## Licencia

This file is part of entradas_crud.

entradas_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

entradas_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with entradas_crud. If not, see https://www.gnu.org/licenses/.
