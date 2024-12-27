# Devcontainer

Devcontainer de laravel php, si quieres crear el proyecto de 0 desde el directorio de trabajo, corre esto 1 vez.

-   Creacion de template + copia de app

```sh
# Template folder
mkdir -p /workspaces
# Create app and copy to workspace
cd /workspaces
laravel new laravel_app
cp -r laravel_app/. /app
# Clean and go back to folder
rm -rf laravel_app
cd -
```

-   Creacion desde composer

```sh
composer create-project --prefer-dist laravel/laravel laravel_app
```

-   Instalacion de nvm

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | zsh
source ~/.zshrc
```

## Referencias

1. Mover archivos a directorios [stack-overflow](https://stackoverflow.com/questions/20192070/how-to-move-all-files-including-hidden-files-into-parent-directory-via)

2. usando las preinstalaciones de otras imagenes en dockerfile [stack-overflow](https://stackoverflow.com/questions/76109982/installing-specific-version-of-nodejs-and-npm-on-alpine-docker-image)

> Es por este motivo que existe el `dockerfile.alt`
