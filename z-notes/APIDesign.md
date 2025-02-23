
# Control de versiones


    Example:

    https://api.moontravels/v1/customers/


https://www.youtube.com/watch?v=SdsaZ-t1QwA
https://github.com/aluna1997/ApiMoonTravelsSwagger

## Insertar un nuevo cliente en la BD

- URL: https://api.moontravels/v1/customers/
- Método: POST
- Parámetros: 
    token autenticacion (header)
    datos personales (body)

- Response:

    200: idCustomer
    400-500: Mensaje de error


## Obtener un cliente en la BD

- URL: https://api.moontravels/v1/customers/idCustomer
- Método: GET
- Parámetros: 
    token autenticacion (header)
    idCustomer (url)*

- Response:

    200: idCliente
    400-500: Mensaje de error

 

## 


## 



##


##




mux.Route("/api/v1/admin", func(mux chi.Router) {
    mux.Use(app.AuthTokenMiddleware) // Middleware para endpoints protegidos

    // Endpoints de productos
    mux.Post("/products", app.CreateProduct)      // Crear un producto
    mux.Get("/products", app.AllProducts)         // Obtener todos los productos
    mux.Get("/products/{id}", app.GetProduct)     // Obtener un producto por ID
    mux.Put("/products/{id}", app.UpdateProduct)  // Actualizar un producto por ID
    mux.Delete("/products/{id}", app.DeleteProduct) // Eliminar un producto por ID
})

// Endpoints públicos de autenticación
mux.Post("/api/v1/users/signup", app.SignUp)  // Registro de usuario
mux.Post("/api/v1/users/login", app.Login)    // Login
mux.Post("/api/v1/users/logout", app.Logout)  // Logout

// Validación de tokens
mux.Post("/api/v1/auth/validate-token", app.ValidateToken) 
mux.Post("/api/v1/auth/login-jwt", app.LoginJWT)  
mux.Post("/api/v1/auth/login-asim-jwt", app.LoginAsimJWT)  

// Endpoint de salud del servidor
mux.Get("/api/v1/health", app.Health)



1. ¿Es necesario el prefijo /admin para rutas protegidas?

Respuesta: No es necesario.

El prefijo /admin puede ser útil si la API tiene una distinción clara entre usuarios regulares y administradores.
Si todos los endpoints de /products están protegidos, es suficiente con aplicar el middleware en esa ruta sin necesidad de un prefijo admin.
Mejor práctica: Aplicar seguridad con el middleware y evitar prefijos redundantes.


- Por seguridad podria removerse el prefijo admin.
Mejor práctica: Aplicar seguridad con el middleware y evitar prefijos redundantes.


    mux.Route("/api/v1/products", func(mux chi.Router) {
        mux.Use(app.AuthTokenMiddleware) // Middleware para proteger los endpoints

        mux.Post("/", app.CreateProduct)     
        mux.Get("/", app.AllProducts)       
        mux.Get("/{id}", app.GetProduct)     
        mux.Put("/{id}", app.UpdateProduct)  
        mux.Delete("/{id}", app.DeleteProduct)
    })

Ventaja:

Mantiene la URL limpia y RESTful.
La seguridad sigue siendo aplicada por el middleware.


2. ¿Es necesario usar /users/ en los endpoints públicos de autenticación?
Respuesta: No es obligatorio, pero sí puede ser útil.

En una API con múltiples entidades (usuarios, productos, pedidos), usar /users/signup o /users/login indica claramente que se refiere a usuarios.
Sin embargo, si el contexto de la API es solo autenticación, se puede omitir y usar solo /signup y /login en /api/v1/.

Opciones recomendadas:

Opción 1 (más explícita, útil en APIs grandes con múltiples recursos) **MICROSERVICIO**

    mux.Post("/api/v1/users/signup", app.SignUp)  
    mux.Post("/api/v1/users/login", app.Login)    
    mux.Post("/api/v1/users/logout", app.Logout)  

Opción 2 (más limpia y directa, útil si la API solo maneja autenticación y no otras entidades de usuarios)  **MONOLITO**

    mux.Post("/api/v1/signup", app.SignUp)  
    mux.Post("/api/v1/login", app.Login)    
    mux.Post("/api/v1/logout", app.Logout)  

🔹 Ventaja:

Evita redundancia en APIs enfocadas en autenticación.
Usa /users/ solo si hay más operaciones relacionadas con usuarios (ej. perfil, recuperación de contraseña).


3. ¿Es necesario el prefijo /auth/ para validación de tokens?

Respuesta: No es obligatorio, pero mejora la organización.

/auth/ se usa para agrupar funciones relacionadas con autenticación.
Si solo hay un endpoint de validación, puede ir en /api/v1/validate-token directamente.
Opciones recomendadas:

Opción 1 (mantener /auth/ para agrupar)

    mux.Post("/api/v1/auth/validate-token", app.ValidateToken)  
    mux.Post("/api/v1/auth/login-jwt", app.LoginJWT)  
    mux.Post("/api/v1/auth/login-asim-jwt", app.LoginAsimJWT)  


Opción 2 (más simple y directa, sin /auth/)

    mux.Post("/api/v1/validate-token", app.ValidateToken)  
    mux.Post("/api/v1/login-jwt", app.LoginJWT)  
    mux.Post("/api/v1/login-asim-jwt", app.LoginAsimJWT)  

🔹 Ventaja:

Si hay varios endpoints de autenticación, mantener /auth/ es útil.
Si solo hay validación de token, se puede simplificar sin /auth/.


## References


https://cloud.google.com/endpoints/docs/frameworks/java/handling-api-versioning?hl=es-419   (Cloud endpoints)

Tambien se puede con swaggo para golang a partir de anotaciones en los handlers.

https://docsify.js.org/#/





eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzM5NjgxMzQxLCJuYmYiOjE3Mzk1OTQ5NDEsImlhdCI6MTczOTU5NDk0MSwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.Lo0fokrzg-85qPp7kFJx4ByvPEO_H14aoGn3BsebpokyJp2gI3CgJmjB8vfCL-ilW50D81lE_x34AwrNXPltTA