# Tokens keys


Para generar un par de claves EC (Elliptic Curve) privadas y públicas en formato PEM, puedes usar openssl. Aquí te dejo los pasos para obtener ec_private.pem y ec_public.pem con el algoritmo ES256 (basado en la curva prime256v1).

## Generar la clave privada (ec_private.pem)

Ejecuta el siguiente comando en la terminal:

    openssl ecparam -name prime256v1 -genkey -noout -out ec_private.pem

Esto generará un archivo ec_private.pem que contiene la clave privada en formato PEM.




## Extraer la clave pública (ec_public.pem)

Ahora, genera la clave pública a partir de la privada:

    openssl ec -in ec_private.pem -pubout -out ec_public.pem

Esto guardará la clave pública en ec_public.pem, que se usará para verificar los tokens JWT.


## Verificar que las claves se generaron correctamente

Si quieres ver el contenido de las claves, puedes usar:

    cat ec_private.pem
    cat ec_public.pem

O para ver los detalles de la clave privada:

    openssl ec -in ec_private.pem -text -noout





