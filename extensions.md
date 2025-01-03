# Extensiones

Tomando como referencia la configuracion de `dotnet_app`, las configuraciones de vscode se agregan dentro de la llave `customizations.vscode`.

- `settings` configuracion similar al archivo de vscode `settings.json` esto aplica al devcontainer
- `extensions` Extensiones que se instalan en el devcontainer

```json
{
	"name": "C# (.NET)",
	"image": "mcr.microsoft.com/devcontainers/dotnet:1-8.0",
	"customizations": {
		"vscode": {
			"settings": {
				"editor.formatOnSave": true
			},
			"extensions": ["ms-dotnettools.csdevkit", "humao.rest-client"]
		}
	}
}
```

## General

Extensiones de uso general que pueden ser utilidades o temas, recomendado usar algunos fuera del devcontainer

- `formulahendry.code-runner` [CodeRunner](https://marketplace.visualstudio.com/items?itemName=formulahendry.code-runner)
- `IgorSbitnev.error-gutters` [ErrorGutters](https://marketplace.visualstudio.com/items?itemName=IgorSbitnev.error-gutters)
- `EliverLara.andromeda` [Andromeda](https://marketplace.visualstudio.com/items?itemName=EliverLara.andromeda)
- `shalldie.background` [background](https://marketplace.visualstudio.com/items?itemName=shalldie.background)
- `aaron-bond.better-comments` [Better Comments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)
- `adpyke.codesnap` [CodeSnap](https://marketplace.visualstudio.com/items?itemName=adpyke.codesnap)
- `inu1255.easy-snippet` [Easy Snippet](https://marketplace.visualstudio.com/items?itemName=inu1255.easy-snippet)
- `usernamehw.errorlens` [Error Lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens)
- `bierner.markdown-mermaid` [Markdown Mermaid](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-mermaid)
- `PKief.material-icon-theme` [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme)
- `quicktype.quicktype` [Paste JSON as Code](https://marketplace.visualstudio.com/items?itemName=quicktype.quicktype)
- `johnpapa.vscode-peacock` [Peacock](https://marketplace.visualstudio.com/items?itemName=johnpapa.vscode-peacock)
- `esbenp.prettier-vscode` [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
- `icrawl.discord-vscode` [Discord Pressence](https://marketplace.visualstudio.com/items?itemName=icrawl.discord-vscode)

## Backend

Extensiones que son utiles para el desarrollo de api en general

- `rohinivsenthil.postcode` [Postcode](https://marketplace.visualstudio.com/items?itemName=rohinivsenthil.postcode)
- `humao.rest-client` [REST CLIENT](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
- `ckolkman.vscode-postgres` [PostgreSQL](https://marketplace.visualstudio.com/items?itemName=ckolkman.vscode-postgres)
- `zxh404.vscode-proto3` [Proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)

## Utils

- `tomoki1207.pdf` [PDF Viewer](https://marketplace.visualstudio.com/items?itemName=tomoki1207.pdf)
  - previsualizacion de pdfs en vscode
- `GrapeCity.gc-excelviewer` [Excel Viewer](https://marketplace.visualstudio.com/items?itemName=GrapeCity.gc-excelviewer)
  - para visualizar archivos excel
- `cweijan.vscode-database-client2` [Database Client](https://marketplace.visualstudio.com/items?itemName=cweijan.vscode-database-client2)
  - visualizacion de bases de datos
- `alexcvzz.vscode-sqlite` [Sqlite](https://marketplace.visualstudio.com/items?itemName=alexcvzz.vscode-sqlite)
  - Herramientas de sqlite
- `AykutSarac.jsoncrack-vscode` [Json Crack](https://marketplace.visualstudio.com/items?itemName=AykutSarac.jsoncrack-vscode)
  - visualizacion de json en modo visual
- `kisstkondoros.vscode-gutter-preview` [Image Preview](https://marketplace.visualstudio.com/items?itemName=kisstkondoros.vscode-gutter-preview)
  - visualizacion de previews de imagenes en el editor de codigo

## Docker

Extensiones especificas de docker , devcontainer y k8s.

- `jeff-hykin.better-dockerfile-syntax` [Better DockerFile Syntax](https://marketplace.visualstudio.com/items?itemName=jeff-hykin.better-dockerfile-syntax)
- `ms-vscode-remote.remote-containers`[Devcontainers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- `ms-azuretools.vscode-docker` [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
- `ms-kubernetes-tools.vscode-kubernetes-tools` [Kubernetes](https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools)
