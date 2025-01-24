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
            "Name": "arcangel",
            "Description": "arcangel",
            "Price": 2000,
            "category_id": 11
            }' \
        -v


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



    curl --key client.key \
        --cert client.pem \
        -k \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "token": "HP4J6HCCN4HYXRFFDH2ODM5DKA"
            }' \
        https://localhost:9090/users/logout




    docker exec -it products_pg_db psql -U postgres -d e_commerce




TLS/02-ms-go-HTTP-TLS/README.md:51:openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -subj "/C=AR/ST=CABA/L=CABA/O=ExampleOrg/OU=IT Department/CN=*"
TLS/02-ms-go-HTTP-TLS/README.md:180:    Issuer: C = AR, ST = CABA, L = CABA, O = ExampleOrg, OU = IT Department, CN = *
TLS/02-ms-go-HTTP-TLS/README.md:181:    Subject: C = AR, ST = CABA, L = CABA, O = ExampleOrg, OU = IT Department, CN = *



Indexes users by email column | NOT NULL UNIQUE


TOKENS  (Lógica de 8 funciones)
USERS (Lógica de 7 funciones)