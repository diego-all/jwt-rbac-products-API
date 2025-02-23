docker-compose build
docker-compose up --build



El modelo de postgres y sqlite son muy similares solo se cambio el string de conexion y ya.


Hay un error con [UPDATE]



PUT http://localhost:9090/products/update/12

{
"Name": "Limpiador y Blanqueador random",
"Description": "Limpiador y Blanqueador random",
"Price": 4000,
"category_id": 12
}

    {    "error": true,    "message": "ERROR: could not determine data type of parameter $4 (SQLSTATE 42P18)"}


Despues de agregar uno nuevo el indice aumenta entonces hay una inconsistencia con el param que se envia.  (SOLVED)



update products set name = $1, description = $2, price = $3, updated_at = $5, where id = $6;



 10 | Limpiador y Limpido            | Limpiador y Blanqueador puro                    | 4000.00 | 2024-11-09 03:37:34.173944 | 2024-11-09 02:05:36.761629
 13 | Limpiador y Limpido            | Limpiador y Blanqueador puro                    | 4000.00 | 2024-11-09 02:04:22.493882 | 2024-11-09 02:08:10.351493
 14 | Limpiador y Limpido            | Limpiador y Blanqueador puro                    | 4000.00 | 2024-11-09 02:20:14.089344 | 2024-11-09 02:22:35.50275
 15 | garbanzo                       | garbanzo                                        | 2000.00 | 2024-11-09 02:23:30.16326  | 2024-11-09 02:23:30.16326



utilizando pq el timestamp es mas pequeño


    curl --key client.key --cert client.pem -k https://localhost:9090/health (Funciona)
    
    curl -X POST --header "Content-Type: application/json" --header "Accept: */*" "http://localhost:8080/api/rest/v1/category/p1/{id}?id=${a}&name=${b}&typecode=${c}

    HTTP_PROXY=http://127.1:8080 arjun -u http://127.1:5000/