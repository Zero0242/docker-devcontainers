<p align="center">
  <a href="https://code.visualstudio.com/docs/remote/containers" target="blank">
  <img src="https://miro.medium.com/v2/resize:fit:719/1*vmW2Gu0ElsVxq79gYRLTlw.png" height="200" alt="Devcontainers Logo" /></a>
</p>

# Pruebas con devcontainers

Algunas pruebas para probar los devcontainers en diferentes tipos de proyectos, cada app de la carpeta `/apps` tienen sus propias configuraciones de entorno

- Referencia a extensiones [aquí](extensions.md)

## Requisitos

1. [devcontainers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Referencias

1. [Trabajando con Devcontainers](https://learn.microsoft.com/en-us/samples/microsoft/vscode-remote-try-dab/devcontainer/)

# Apps

### Gin Go App

Configuracion de un devcontainer sencillo para probar un backend el lenguaje `Go`

| Lenguaje | Imagen             | Observaciones                             |
| -------- | ------------------ | ----------------------------------------- |
| Go       | golang:1.23-alpine | Uso directo de la imagen sin alteraciones |

### Dotnet App

Entorno de `dotnet` para probar un api sencillo.

| Lenguaje | Imagen                                       | Observaciones                                      |
| -------- | -------------------------------------------- | -------------------------------------------------- |
| C#       | mcr.microsoft.com/devcontainers/dotnet:1-8.0 | Configuracion creada con la extension devcontainer |

### Kotlin Spring Boot

Configuración de un entorno de desarrollo para una aplicación de Spring Boot con `Kotlin`.

| Lenguaje | Imagen                             | Observaciones                          |
| -------- | ---------------------------------- | -------------------------------------- |
| Kotlin   | openjdk:24-ea-17-jdk-slim-bullseye | Añadidas herramientas como Maven y Git |

### Java Spring Boot

Configuración de un entorno de desarrollo para una aplicación de Spring Boot con `Java`.

| Lenguaje | Imagen                             | Observaciones                          |
| -------- | ---------------------------------- | -------------------------------------- |
| Java     | openjdk:24-ea-17-jdk-slim-bullseye | Añadidas herramientas como Maven y Git |

### Node.js App

Configuración de un entorno de desarrollo para una aplicación de `Node.js`. Aca se personaliza la imagen con `oh my zsh` y `powerlevel10k`

| Lenguaje | Imagen         | Observaciones                                                                      |
| -------- | -------------- | ---------------------------------------------------------------------------------- |
| Node.js  | node:23-alpine | Instaladas herramientas como Git, Curl, Zip, Unzip y configurado Zsh con Oh My Zsh |

### Go Gorilla App

Entorno de desarrollo de `go` + `postgres` para desarrollo de backend con **gorilla + gorm + air**

- Docker Compose (app): solo el servicio de postgresql
- Docker Compose (devcontainer): combinacion del entorno de desarrollo y la base de datos

> La extension de postgresql nos permite ver si podemos conectarnos a la base de datos desde el devcontainer

| Servicio | Imagen             | Observaciones                   |
| -------- | ------------------ | ------------------------------- |
| Go       | golang:1.23-alpine | Entorno con oh-my-zsh y plugins |
| Postgres | postgres:14-alpine | Postgresql                      |
