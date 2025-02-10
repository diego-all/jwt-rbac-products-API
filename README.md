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

    curl --key client.key --cert client.pem -k https://localhost:9090/test-save-jwt-token   (EMULANDO UN LOGIN)


    curl --key client.key --cert client.pem -k \
        -X POST https://localhost:9090/products \
        -H "Content-Type: application/json" \
        -d '{
            "Name": "trin",
            "Description": "trin",
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
        https://localhost:9090/users/loginjwt

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


curl -X POST "https://localhost:9090/admin/products" \
     --key client.key \
     --cert client.pem \
     -k \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer DARAOKFMBWLVZD7N6RTMBYXYWE"



curl -X GET "https://localhost:9090/products/get/4" \
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
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDRUMjM6MjQ6MjAuOTkxOTcxNjM0LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg3MzY2NjAsImlhdCI6MTczODcyOTQ2MCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.4asw1aTA2D2yqWw399L0bXXwFk0mec7bGb2eSQGAFVc" \
     -H "Content-Type: application/json" \
     -d '{}'

curl --key client.key --cert client.pem -k \
     -X POST https://localhost:9090/admin/products/all \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDVUMjI6NTI6MDEuMzAxMTM2MDgyLTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg5MDAzMjEsImlhdCI6MTczODgxMzkyMSwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.R5l6HMY59TREjbebPRqZksKTDB7W2gCnx0e4DPj-7YI" \
     -H "Content-Type: application/json" \
     -d '{}'



curl --key client.key --cert client.pem -k \
     -X POST https://localhost:9090/admin/products/all \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDlUMTA6MzQ6MjQuMzU4NzQ5NTkxLTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3MzkxMTUyNjQsImlhdCI6MTczOTAyODg2NCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.tdoxJ_T4A9a16yjOsAuE9G_xX1YVc6LwFkBHsRdjcNI" \
     -H "Content-Type: application/json" \
     -d '{}'



eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDlUMTA6MzQ6MjQuMzU4NzQ5NTkxLTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3MzkxMTUyNjQsImlhdCI6MTczOTAyODg2NCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.tdoxJ_T4A9a16yjOsAuE9G_xX1YVc6LwFkBHsRdjcNI


eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDhUMjE6NDY6NDAuOTIwOTI2MDM1LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3MzkwNjkyMDAsImlhdCI6MTczODk4MjgwMCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.pr2qVOSUQD1o6pigw9KI3Tv_PaysLCpasI8oaUIYYlY






eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDhUMjE6MTU6NTkuMjYyOTIzNDQ2LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3MzkwNjczNTksImlhdCI6MTczODk4MDk1OSwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.IiYDGqLZIj1dMZ7afXVPfb9RmIqo3koCZ-ZPUCd5TOI



eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDVUMjI6NTI6MDEuMzAxMTM2MDgyLTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg5MDAzMjEsImlhdCI6MTczODgxMzkyMSwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.R5l6HMY59TREjbebPRqZksKTDB7W2gCnx0e4DPj-7YI

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDZUMjM6MzM6MTEuNjAxNDMzNzA4LTA1OjAwIiwicm9sZSI6InRyaW4iLCJTZWNyZXRLZXkiOiIiLCJzdWIiOiJkaWVnb0BkaWVnby5jb20iLCJleHAiOjE3Mzg5ODkxOTEsImlhdCI6MTczODkwMjc5MSwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwaXJlc19hdCI6MCwiaXNzdWVkX2F0IjowfQ.RDJBBN4W0LCG4Hggy277bjZkX0h_eKmu6CYUt6hcWqs

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwidXNlcl9pZCI6MSwiZW1haWwiOiJkaWVnb0BkaWVnby5jb20iLCJ0b2tlbiI6IiIsImV4cGlyeSI6IjIwMjUtMDItMDZUMjM6MzQ6MTEuMTMxMjk3NzUtMDU6MDAiLCJyb2xlIjoidHJpbiIsIlNlY3JldEtleSI6IiIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImV4cCI6MTczODk4OTI1MSwiaWF0IjoxNzM4OTAyODUxLCJjcmVhdGVkX2F0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJ1cGRhdGVkX2F0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHBpcmVzX2F0IjowLCJpc3N1ZWRfYXQiOjB9.cmj6v-gcdhUeGHsigeNQP8k7XUIdWVs-3qOlyPT20FQ


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




