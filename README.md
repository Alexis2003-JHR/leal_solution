# Levantamiento del Servicio

Este documento explica cómo levantar el servicio para realizar la prueba técnica, junto con información sobre documentación, pruebas y aclaraciones importantes.

## Levantamiento del Servicio

1. **Navegar a la carpeta `backend`:**

   Abre una terminal y dirígete al directorio `backend`:

   ```bash
   cd backend
   ```

2. **Iniciar los servicios con Docker Compose:**

   Dentro de la carpeta `backend`, ejecuta el siguiente comando para levantar los contenedores en modo "detached" (en segundo plano):

   ```bash
   docker compose up -d
   ```

   Esto levantará tanto el servicio de la API como el de la base de datos.

## Acceso al API

Una vez que los servicios estén levantados, el API estará accesible en la siguiente URL:

- **API Base URL:** [http://localhost:6060/api/v1](http://localhost:6060/api/v1)


## Migraciones
Las migraciones se realizaron mediante GORM. La librería gestiona automáticamente la migración de la base de datos en base a las estructuras ORM definidas en el código.

## Notas y Documentación

- **Swagger:**
  - La documentación interactiva de la API, generada con Swagger, se encuentra en el endpoint:
    - [http://localhost:6060/api/v1/docs/index.html](http://localhost:6060/api/v1/docs/index.html)

- **Documentación Adicional:**
  - En la carpeta `docs` se han incluido archivos que ayudaron a comprender los requerimientos del proyecto. Estos documentos incluyen análisis escritos, diagramas de secuencia y otros recursos. La mayoría están escritos en Markdown y algunos gráficos utilizan Mermaid, por lo que es recomendable usar un previsualizador compatible con Mermaid para verlos correctamente.

## Pruebas

- **Pruebas Unitarias:**
  - Las pruebas unitarias del proyecto se encuentran en la carpeta `/internal/core/services/service_test.go`.
  - Para ejecutar todas las pruebas unitarias con salida detallada, navega al directorio `backend` y ejecuta:

    ```bash
    go test ./internal/core/services -v
    ```

- **Pruebas de Carga:**
  - Si tienes instalada la herramienta `k6`, puedes ejecutar el archivo `load-test.js` para realizar una prueba de carga en el endpoint de transacciones. Esta prueba se realizó para evaluar el comportamiento de esta funcionalidad bajo condiciones de alta demanda y para verificar el manejo de la concurrencia.

## Aclaraciones

- **Archivos `.env`:**
  - Para facilitar el acceso a esta prueba técnica, se han dejado los archivos `.env` dentro del proyecto. 
  - **Nota:** En entornos productivos, no es recomendable dejar estos archivos en el repositorio, ya que contienen información sensible. Se debe gestionar la configuración de manera segura.
