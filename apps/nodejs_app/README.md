<p align="center">
  <a href="https://github.com/" target="blank">
  <img src="https://cdn-icons-png.flaticon.com/512/25/25231.png" width="200" alt="App Logo" /></a>
</p>

# Proyecto Cron

Cron básico, creado durante el curso de [docker](https://www.udemy.com/share/107Ple3@bdv7jBU2dSlE6vZw8aLAd8Eq0N7Xt_pSOL83lFBOB34jOuRb_K8f1XJgFnBywx8J8g==/).

Configurado con [babel](https://mugan86.medium.com/configurar-babel-en-nodejs-525fd101990b) para la sintaxis moderna de javascript

> creado en node js

## DEV

1. Clonar repositorio con `git clone`
2. Instalar los paquetes de node js con `npm install`
3. Ejecutar el proyecto con `npm run start` o `npm run dev`
   - Opcional: ejecutar los comandos del `package.json` o usando `F5`

## Producción

1. Clonar repositorio con `git clone`
2. Instalar los paquetes de node js con `npm install`
   - APROX son +30mb de peso en paquetes
3. Correr las pruebas con `npm run test`
4. Generar build `npm run build`
5. Instalar solo los paquetes de produccion con `npm install --prod`
   - APROX menos de 1mb de peso en paquetes
6. Ejecutar servidor con `npm run start:deploy`

## Requisitos

1. Tener instalado node js

## Scripts

Algunos scripts que pueden ser utilizados en modo desarrollo

| Comando                | Descripcion                                 |
| ---------------------- | ------------------------------------------- |
| `npm install`          | Instala todas las dependencias              |
| `npm install --prod`   | Instala solo las dependencias de Producción |
| `rm -rf node_modules`  | Limpia las dependencias                     |
| `npm run test`         | Ejecuta las pruebas                         |
| `npm start`            | Iniciar app                                 |
| `npm run dev`          | Iniciar app en modo desarrollo              |
| `npm run build`        | Compilar la app a nodejs                    |
| `npm run start:deploy` | Iniciar la app compilada                    |

# Docker

Solo añadir lo minimo para que el server funcione, posee su propio archivo `.dockerignore` que contiene los mismos elementos de `.gitignore` para evitar archivos no relevantes.

- Construir

```sh
# Detectando el archivo `dockerfile` automaticamente
docker build --tag imagen_custom:nombre_tag .
# Lo mismo pero especificando el nombre del dockerfile
docker build --tag imagen_custom:nombre_tag -f dockerfile.produccion .
```

- Eliminar cache de construccion

```sh
# primero elimina la imagen creada
docker image rm imagen_custom:nombre_tag
# elimina la data en caché
docker buildx prune
```
