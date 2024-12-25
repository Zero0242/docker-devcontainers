<p align="center">
  <a href="https://go.dev/" target="blank">
  <img src="https://www.nixsolutions.com/uploads/2020/07/Golang.png" height="200" alt="App Logo" /></a>
</p>

# Proyecto Servidor Go

Aplicacion de go + gorilla/mux + base de datos

> creado en go-language

## DEV

1. Clonar repositorio con `git clone`
2. Instalar los paquetes de go-language con `go get`
3. Ejecutar el proyecto con `go run .`

- Ejecutar modo desarrollo con `air .` (requiere instalar air)

## Requisitos

1. Tener instalado go-language

## Scripts

Algunos scripts que pueden ser utilizados

| Comando                                        | Descripcion                                |
| ---------------------------------------------- | ------------------------------------------ |
| `go run main.go`                               | Ejecuta el archivo principal de la app     |
| `air .`                                        | (AIR) Corre la app en modo **desarrollo**  |
| `go mod tidy`                                  | Instala las dependencias                   |
| `go build`                                     | Compila el código fuente                   |
| `go test`                                      | Ejecuta las pruebas                        |
| `go fmt`                                       | Formatea el código fuente                  |
| `go vet`                                       | Analiza el código en busca de errores      |
| `go clean`                                     | Elimina archivos generados                 |
| `go get`                                       | Descarga e instala paquetes y dependencias |
| `go mod init github.com/<user>/<project_name>` | Inicia un nuevo proyecto                   |

#### Otros Scripts

Otros scripts que pueden usar para fines de desarrollo, (acciones de paquetes)

| Comando | Descripcion                                                                |
| ------- | -------------------------------------------------------------------------- |
| `....`  | Insertar scripts que usen los paquetes de terceros si es necesario hacerlo |

## Documentacion

Links de librerias utilizadas

- [Go Language](https://go.dev/)
- [Air/hot-reload](https://github.com/air-verse/air)
  - `go install github.com/air-verse/air@latest` instalacion
- [Posgres/alpine](https://luppeng.wordpress.com/2020/02/28/install-and-start-postgresql-on-alpine-linux/)
- [Curso](https://youtu.be/B6gQ1B0cn4s?si=UmFd_hg8PxnBDUEy)
