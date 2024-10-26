# worker-evaluacion
## Configuración de Variables de Entorno para Docker Compose

Para ejecutar el contenedor, se deben configurar las siguientes variables de entorno en el archivo `docker-compose.yml` o al ejecutar `docker run`:

| Variable               | Descripción                                  | Valor por Defecto                |
|------------------------|----------------------------------------------|----------------------------------|
| `MAX_WORKERS`          | Número máximo de workers                     | `4`                              |
| `MAX_QUEUE`            | Tamaño máximo de la cola                     | `20`                             |
| `PORT`                 | Puerto en el que escucha la aplicación       | `:8081`                          |
| `MINIO_ENDPOINT`       | Dirección y puerto del servidor MinIO        | `host.docker.internal:9000`      |
| `MINIO_ACCESS_KEY`     | Clave de acceso de MinIO                     | `wDeP32IOLrUwrlmFLwOc`           |
| `MINIO_SECRET_KEY`     | Clave secreta de MinIO                       | `VzwSovPOwHeAe68asfwzBMjgCpbNzLvCXiUpNpCi` |
| `MINIO_USE_SSL`        | Define si MinIO utiliza SSL (`true` o `false`) | `false`                        |
| `RABBITMQ_USER`        | Usuario de RabbitMQ                          | `admin`                          |
| `RABBITMQ_PASS`        | Contraseña de RabbitMQ                       | `admin`                          |
| `RABBITMQ_HOST`        | Dirección del servidor RabbitMQ              | `host.docker.internal`           |
| `RABBITMQ_PORT`        | Puerto del servidor RabbitMQ                 | `5673`                           |
| `RABBITMQ_QUEUE_NAME_EVA` | Nombre de la cola en RabbitMQ              | `notification`                   |

### Ejemplo de Comando

Para ejecutar la imagen `nombre_que_se_le_de_a_la_imagen_del_repo` con las variables anteriores, usa el siguiente comando:

```bash
docker run -p 8081:8081 \
  -e MAX_WORKERS=4 \
  -e MAX_QUEUE=20 \
  -e PORT=:8081 \
  -e MINIO_ENDPOINT=host.docker.internal:9000 \
  -e MINIO_ACCESS_KEY=wDeP32IOLrUwrlmFLwOc \
  -e MINIO_SECRET_KEY=VzwSovPOwHeAe68asfwzBMjgCpbNzLvCXiUpNpCi \
  -e MINIO_USE_SSL=false \
  -e RABBITMQ_USER=admin \
  -e RABBITMQ_PASS=admin \
  -e RABBITMQ_HOST=host.docker.internal \
  -e RABBITMQ_PORT=5673 \
  -e RABBITMQ_QUEUE_NAME_EVA=notification \
  nombre_que_se_le_de_a_la_imagen_del_repo
