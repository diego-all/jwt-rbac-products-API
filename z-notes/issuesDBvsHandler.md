# ISSUES

Inconsistencia de valores de la response en login.


    root@pho3nix:/home/diegoall/CORRECTION/jwt-rbac-products-API/cmd/client# curl --key client.key         --cert client.pem         -k         -X POST         -H "Content-Type: application/json"         -d '{
                "email": "diego@diego.com",
                "password": "password"
                }'         https://localhost:9090/users/login
    {"error":false,"message":"logged in","data":{"token":{"id":0,"user_id":1,"token":"POTODIH3RSZVRO6WOKQAN56EZ4","expiry":"2025-02-24T02:37:50.047008944-05:00","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"user":{"id":1,"email":"diego@diego.com","first_name":"diego","last_name":"last_name_placeholder","password":"$2a$12$ZTTPWrzmiT5wSz9gl2FrJuiP4wwoXNKriRbSnZuWCvN/ZgxutLYjG","created_at":"2025-01-15T03:17:54.942686Z","updated_at":"2025-01-15T03:17:54.942686Z","token_at":{"id":0,"token":"","expiry":"0001-01-01T00:00:00Z","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}}}}root@pho3nix:/home/diegoall/CORRECTION/jwt-rbac-products-API/cmd/client#

- Error en Login, valores del response son diferentes en la base de datos.
Esto es debido a que el valor de created_at y updated_at son calculados directamente en la base de datos.
Solucion: 

1. Calcularlos con Go antes de insertarlos.

2. Postgres permite utilizar RETURNING al hacer el insert. Se captura y se asigna al response.

        stmt = `INSERT INTO tokens (user_id, email, token, token_hash, created_at, updated_at, expiry) 
                VALUES ($1, $2, $3, $4, $5, $6, $7) 
                RETURNING id, created_at, updated_at`


        INSERT INTO tokens (user_id, email, token, token_hash, created_at, updated_at, expiry) 
        VALUES 
            (1, 
            'diego@diego.com', 
            'RK25UBTFWWVXGVYECNOBKPBFHM', 
            E'\\x3a8ccd0d8e4a9e8d04b69a83d9634570446557f962c0f78b4ea6a488f3', 
            '2025-02-23 22:51:50.638287+00', 
            '2025-02-22 17:51:50.643585', 
            '2025-02-22 17:51:50.643586') 
        RETURNING id, created_at, updated_at;

3. Realizar una consulta a la base de datos luego de Insertar. (Ineficiente)


**¿Cual es la mejor práctica?**

    Here are a few tips for maximizing the usefulness of created_at and updated_at columns:

    Always use the same type of timestamp columns (`TIMESTAMP` or `DATETIME`) consistently across your entire schema.
    Avoid manual updates to these columns as it defeats their purpose and may introduce errors or inconsistencies.
    Ensure your server’s time zone is correctly set since TIMESTAMP columns are affected by the time zone.

> Automatic timestamps with PostgreSQL




Es mejor desde la base de datos!


- Chatgpt

La mejor práctica para asignar los valores de created_at y updated_at depende del contexto y de las necesidades específicas de tu aplicación. Ambos enfoques 

— hacerlo directamente en la base de datos o manejarlo desde la aplicación— tienen ventajas y desventajas, pero generalmente se recomienda manejarlos en la aplicación, por las siguientes razones:


1. Manejo en la aplicación (desde el código)

Ventajas:

Control total sobre los valores: Puedes tener un control más preciso sobre cuándo y cómo se asignan estas fechas, ya que están directamente en el código de la aplicación. Esto te permite asegurarte de que los valores sean correctos antes de enviarlos a la base de datos.
Sin dependencia de la base de datos: Si, en el futuro, decides cambiar de base de datos o realizar ajustes en la lógica de la base de datos, no tendrías que depender de triggers o funciones específicas de la base de datos que generen estos valores.
Mayor flexibilidad: Si, por ejemplo, necesitas asignar diferentes valores para created_at o updated_at bajo ciertas condiciones (como dependiendo de una configuración de la aplicación), manejarlo en el código te da esa flexibilidad.
Desventajas:

Necesitas gestionar las fechas en el código: Esto puede aumentar ligeramente la complejidad del código, ya que deberás asegurarte de que estas fechas se asignen correctamente en cada operación que involucre la creación o actualización de registros.
Cómo implementarlo: Puedes asignar los valores de created_at y updated_at en tu código 



- 

- 

- 

- 