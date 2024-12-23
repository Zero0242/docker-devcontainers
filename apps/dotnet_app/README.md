# Instrucciones para ejecutar el DevContainer

1. **Instalar Docker**: Asegúrate de tener Docker instalado en tu máquina. Puedes descargarlo desde [aquí](https://www.docker.com/products/docker-desktop).

2. **Instalar Visual Studio Code**: Descarga e instala Visual Studio Code desde [aquí](https://code.visualstudio.com/).

3. **Instalar la extensión de Dev Containers**: Abre Visual Studio Code y ve a la extensión de Dev Containers. Instálala desde el marketplace de extensiones.

4. **Abrir el proyecto en un contenedor de desarrollo**:

   - Abre Visual Studio Code.
   - Ve a `Archivo` > `Abrir Carpeta` y selecciona la carpeta del proyecto `/Users/pedro/Desktop/devcontainer/dotnet_app`.
   - Presiona `F1` para abrir la paleta de comandos y escribe `Dev Containers: Reopen in Container`.

5. **Esperar a que se construya el contenedor**: Visual Studio Code descargará y construirá el contenedor de desarrollo. Este proceso puede tardar unos minutos.

6. **Ejecutar la aplicación**:
   - Una vez que el contenedor esté listo, abre una terminal integrada en Visual Studio Code.
   - Ejecuta el siguiente comando para iniciar la aplicación:
     ```sh
     dotnet run
     ```

¡Listo! Ahora deberías tener tu aplicación .NET corriendo dentro de un DevContainer.
