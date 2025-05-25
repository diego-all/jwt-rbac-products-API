# JWT-RBAC-Product-API




### Start the project

**Build the API**

    go build ./cmd/api
    go build -o {{.LowerEntity}}API ./cmd/api/

    go run ./cmd/api/



## sqlite3

**Create sqlite database**

If you are using the Ubuntu Linux operating system, you can start the SQLite database by running the following command:

    sh script.sh

## Postgres



## Collections


    curl --key client.key --cert client.pem -k https://localhost:9090/health (Funciona)
    
    curl -X POST --header "Content-Type: application/json" --header "Accept: */*" "http://localhost:8080/api/rest/v1/category/p1/{id}?id=${a}&name=${b}&typecode=${c}

    curl --key client.key --cert client.pem -k https://localhost:9090/test-save-token

    curl --key client.key --cert client.pem -k https://localhost:9090/test-generate-token

    curl --key client.key --cert client.pem -k https://localhost:9090/test-generate-jwt-token

    curl --key client.key --cert client.pem -k https://localhost:9090/test-generate-asim-jwt-token

    curl --key client.key --cert client.pem -k https://localhost:9090/test-save-jwt-token   (EMULANDO UN LOGIN)


    curl --key client.key --cert client.pem -k \
        -X POST https://localhost:9090/admin/products \
        -H "Content-Type: application/json" \
        -d '{
            "Name": "agua",
            "Description": "agua",
            "Price": 2000,
            "category_id": 11
            }'


     curl --key client.key --cert client.pem -k \
    -X POST https://localhost:9090/admin/products \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzM5NjgwNDU0LCJuYmYiOjE3Mzk1OTQwNTQsImlhdCI6MTczOTU5NDA1NCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.AAde6TG9H_rzlyr8H1RpODqCRDtyDjk39FeXzoqSf2wVzIjXIrdqndT6TotzmTrQAKRnBsDaGvb_WsyA3OG31Q" \
    -d '{
        "Name": "agua",
        "Description": "agua",
        "Price": 2000,
        "category_id": 11
    }'
       


    curl --key client.key --cert client.pem -k \
        -X POST https://localhost:9090/products \
        -H "Content-Type: application/json" \
        -d '{
            "Name": "arcangelmarach",
            "Description": "arcangelmarach",
            "Price": 2000,
            "category_id": 11
            }' \
        -v

    curl --key client.key --cert client.pem -k \
        -X POST https://localhost:9090/admin/products \
        -H "Content-Type: application/json" \
        -d '{
            "Name": "arcangelmarach",
            "Description": "arcangelmarach",
            "Price": 2000,
            "category_id": 11
            }' \
        -v

curl --key client.key --cert client.pem -k \
    -X POST https://localhost:9090/admin/products \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer GFLZ2EDGDUUD3PMM6WJHBAJZO4" \
    -d '{
        "Name": "arcangelmarach",
        "Description": "arcangelmarach",
        "Price": 2000,
        "category_id": 30
    }'


    curl --key client.key \
        --cert client.pem \
        -k \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "email": "diego@diego.com",
            "password": "password"
            }' \
        https://localhost:9090/users/login

LOGIN ESTA FUNCIONANDO CON LOS PRIMEROS USUARIOS (diego@diego.com, you@there.com, you@trinis.com)


    curl --key client.key \
        --cert client.pem \
        -k \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "email": "diego@diego.com",
            "password": "password"
            }' \
        https://localhost:9090/users/login-asim-jwt



    curl --key client.key \
        --cert client.pem \
        -k \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "token": "BZJHTSNIA2KE3YUJS7MHVQZ6GI"
            }' \
        https://localhost:9090/users/logout


curl --key client.key --cert client.pem -k \
    -X POST https://localhost:9090/admin/foo \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer GFLZ2EDGDUUD3PMM6WJHBAJZO4" \
    -d '{
        "Message": "bar"
    }'

invalid authentication credentials



    docker exec -it products_pg_db psql -U postgres -d e_commerce




TLS/02-ms-go-HTTP-TLS/README.md:51:openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -subj "/C=AR/ST=CABA/L=CABA/O=ExampleOrg/OU=IT Department/CN=*"
TLS/02-ms-go-HTTP-TLS/README.md:180:    Issuer: C = AR, ST = CABA, L = CABA, O = ExampleOrg, OU = IT Department, CN = *
TLS/02-ms-go-HTTP-TLS/README.md:181:    Subject: C = AR, ST = CABA, L = CABA, O = ExampleOrg, OU = IT Department, CN = *



Indexes users by email column | NOT NULL UNIQUE


TOKENS  (Lógica de 8 funciones)
USERS (Lógica de 7 funciones)




curl --key client.key \
        --cert client.pem \
        -k \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "email": "diego@diego.com",
            "password": "password"
            }' \
        https://localhost:9090/users/login-jwt

/users/loginjwt



standard claims esta deprecado, puede usar jwt.RegisteredClaims, y para registar la fecha de expiracion usan jwt.NewNumericDate(time)


yo tenia unos comandos para ver la base de datos en real time salcedo



/test-generate-jwt-token

{
  "id": 0,
  "user_id": 1,
  "email": "diego@diego.com",
  "token": "",
  "expiry": "2025-02-01T23:53:43.751845443-05:00",
  "role": "trin",
  "SecretKey": "",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z",
  "expires_at": 0,
  "issued_at": 0
}




curl --key client.key --cert client.pem -k https://localhost:9090//test-validate-jwt-token



/test-validate-jwt-token

    curl --key client.key --cert client.pem -k https://localhost:9090/test-validate-jwt-token?token=SFKUE2EY63C3U7P42GW4T6B3IY  

    curl --key client.key --cert client.pem -k https://localhost:9090/test-validate-jwt-token?token=OVOHQZLJDOWH2SLU2FTD7ASMAI  (SI)

    curl --key client.key --cert client.pem -k "https://localhost:9090/test-validate-jwt-token?token=TU_TOKEN_AQUI"

    curl --key client.key --cert client.pem -k https://localhost:9090/test-validate-jwt-token

    curl --key client.key --cert client.pem -k "https://localhost:9090/test-validate-asim-jwt-token?token=TU TOKEN ACA"

ME APARECE EXPIRED EL JWT TOKEN PERO EL TOKEN SI APARECE VIGENTE.



    curl --key client.key \
        --cert client.pem \
        -k \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "email": "fulano@fulaaanox.com",
            "password": "passyyyywooooord",
            "firstname": "frankasa2",
            "lastname": "pasquines" 

            }' \
        https://localhost:9090/users/signup




curl --key client.key \
     --cert client.pem \
     -X POST "http://localhost:9090/admin/products" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer TU_JWT_AQUI"


curl -X POST "https://localhost:9090/admin/products" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDRUMjM6MjQ6MjAuOTkxOTcxNjM0LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg3MzY2NjAsImlhdCI6MTczODcyOTQ2MCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.4asw1aTA2D2yqWw399L0bXXwFk0mec7bGb2eSQGAFVc"



MALO
curl -X POST "https://localhost:9090/admin/products" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer DF76HCLYBHNQ5YODRSVUCVICJY"


CORRECTO
curl -X POST "https://localhost:9090/admin/products" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer DF76HCLYBHNQ5YODRSVUCVICJY" \
     -d '{"name":"Producto de prueba","price":10.5,"stock":20}'


curl -X GET "https://localhost:9090/products/get/4" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Authorization: Bearer Producto de trin"




curl -X GET "https://localhost:9090/admin/products/get/4" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDRUMjM6MjQ6MjAuOTkxOTcxNjM0LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg3MzY2NjAsImlhdCI6MTczODcyOTQ2MCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.4asw1aTA2D2yqWw399L0bXXwFk0mec7bGb2eSQGAFVc"



curl -X GET "https://localhost:9090/products/get/4" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDRUMjM6MjQ6MjAuOTkxOTcxNjM0LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg3MzY2NjAsImlhdCI6MTczODcyOTQ2MCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.4asw1aTA2D2yqWw399L0bXXwFk0mec7bGb2eSQGAFVc"

curl -X POST "https://localhost:9090/admin/products/all" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzM5OTM3Njg5LCJuYmYiOjE3Mzk4NTEyODksImlhdCI6MTczOTg1MTI4OSwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.jb3hRgquOlZ4r6DKT_h4f8sX8TXJU5uqbrP59-We-PXdJauCihe0VYGqBlTXOQBT2EZxRbEjRf7cJ4sdpbBJWQ" \
     -H "Content-Type: application/json" \
     -d '{}'

curl --key client.key --cert client.pem -k \
     -X POST https://localhost:9090/admin/products/all \
     -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzQwMDI2MjQ2LCJuYmYiOjE3Mzk5Mzk4NDYsImlhdCI6MTczOTkzOTg0NiwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.PwHSgmdbU7k5TuBu6vfKD3BFUtfm4Ln4shjV5QUcCaJnhJCOyTWVw9wbiIBkTyzwsZyRVQL-4b4U8bpT3b069A" \
     -H "Content-Type: application/json" \
     -d '{}'



curl --key client.key --cert client.pem -k \
     -X POST https://localhost:9090/admin/products/all \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDlUMTA6MzQ6MjQuMzU4NzQ5NTkxLTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3MzkxMTUyNjQsImlhdCI6MTczOTAyODg2NCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.tdoxJ_T4A9a16yjOsAuE9G_xX1YVc6LwFkBHsRdjcNI" \
     -H "Content-Type: application/json" \
     -d '{}'


curl -X POST "https://localhost:9090/admin/products/all" \
    --key client.key \
    --cert client.pem \
    -k \
    --http1.1 \
    -H "Authorization: Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwiZXhwaXJ5IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJyb2xlIjoidHJpbiIsImlzcyI6Imczbm90eXBlIiwic3ViIjoiZGllZ29AZGllZ28uY29tIiwiYXVkIjpbIm1pcy11c3VhcmlvcyJdLCJleHAiOjE3Mzk1MDU0MjIsIm5iZiI6MTczOTQxOTAyMiwiaWF0IjoxNzM5NDE5MDIyLCJjcmVhdGVkX2F0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJ1cGRhdGVkX2F0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoifQ.hc0rAkCMSZDDKgCRzEwCE-QllvnW4RefYEQpndvZA4bC1Q_UTpGBJ3gI0wnPdomjX-Da3GReegT_NJbx58OGWA"

curl -X POST "https://localhost:9090/admin/products/all" \
    --key client.key \
    --cert client.pem \
    -k \
    -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzQwMDI2MjQ2LCJuYmYiOjE3Mzk5Mzk4NDYsImlhdCI6MTczOTkzOTg0NiwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.PwHSgmdbU7k5TuBu6vfKD3BFUtfm4Ln4shjV5QUcCaJnhJCOyTWVw9wbiIBkTyzwsZyRVQL-4b4U8bpT3b069A"


curl -X POST "https://localhost:9090/admin/products/all" \
    --key client.key \
    --cert client.pem \
    -k \
    -H "Authorization: Bearer IRPT27IH5KLBRMH5AJGNLS52QY"

RECORDAR QUE:

ValidateJWTToken retorna los claims en cambio
ValidToken retorna true



	fmt.Println("VALID IN VALIDATEJWTTOKEN", valid)

LA FECHA DE CREACION DEL TOKEN ESTA EN EL FUTURO Y LA DE EXPIRY SIGUE ESTANDO EN EL FUTURO.
VALIDAR LA CONDICION DE COMO SE GENERA EL TOKEN.

POR QUE SI SE COLOCA UN TOKEN DE 24 HORAS ? hay una diferencia de 5?



    ALTER TABLE public.tokens 
    ALTER COLUMN expiry 
    SET DATA TYPE timestamp WITH time zone;

> Se pueden utilizar los verbos HTTP asi vayan con token.

    curl -X GET "https://localhost:9090/admin/products/get/1" \
        --key client.key \
        --cert client.pem \
        -k \
        -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzM5NjgwNDU0LCJuYmYiOjE3Mzk1OTQwNTQsImlhdCI6MTczOTU5NDA1NCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.AAde6TG9H_rzlyr8H1RpODqCRDtyDjk39FeXzoqSf2wVzIjXIrdqndT6TotzmTrQAKRnBsDaGvb_WsyA3OG31Q"



    APARECE CUANDO SE CAMBIA EL TIPO DE TOKEN.
    CREO QUE DEBE VALIDARSE EL TOKEN (ValidToken)

    2025/02/14 23:26:34 http2: panic serving 127.0.0.1:39438: runtime error: slice bounds out of range [-1:]



    curl -X POST "https://localhost:9090/admin/products/all" \
    --key client.key \
    --cert client.pem \
    -k \
    -H "Authorization: Bearer X57U7REZHRBNCPJLUB3YCLHQ2Q"


    