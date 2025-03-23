

Recordar que los parametros seguros se cogen del .env de lo contrario no sube la imagen con docker-compose up -d, dice que estan vacios.


Se puede reconstruir la imagen con docker-compose build

  docker-compose --env-file .env up -d


