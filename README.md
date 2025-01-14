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