# TARGET 

Beneficios de esta estructura:
✅ Separa la lógica de negocio de la infraestructura (HTTP handlers, base de datos, etc.).
✅ Facilita el mantenimiento y pruebas unitarias.
✅ Permite cambiar la base de datos o la capa de transporte sin modificar la lógica del negocio.

Si implementas esto, tu API será más modular, segura y escalable. 

La capa de transporte en una arquitectura de software es la que maneja la comunicación entre el cliente y el servidor. En el contexto de tu API en Golang, esta capa es la responsable de recibir las solicitudes HTTP, procesarlas y devolver respuestas al cliente.

En tu proyecto, la capa de transporte es el paquete cmd/api/

Aquí es donde defines los handlers, las rutas y la inicialización del servidor HTTP.

🔥 Beneficios de separar la capa de transporte
✅ Permite cambiar el protocolo (por ejemplo, migrar a gRPC sin tocar la lógica de negocio).
✅ Hace el código más modular y fácil de mantener.
✅ Separa responsabilidades (HTTP, negocio, y base de datos en capas distintas).


**Clean Architecture (Uncle Bob) Separar la logica de negocio de la infraestructura, lo que facilita la implemenetación de medidas de seguridad en cada capa.**

scaffolding + API 
modelos (gemini)

microservicio autenticacion

**OBJETIVO:**

**DEFINIR UNA ARQUITECTURA**



🔒 Medidas de Seguridad en Cada Capa
Para que la API sea segura, hay que considerar seguridad en varias capas:

1️⃣ Capa de Transporte (cmd/api/)
✅ Middleware de Autenticación → JWT, OAuth2.
✅ Rate Limiting → Evita ataques de fuerza bruta (ejemplo: github.com/ulule/limiter).
✅ CORS Policy → Solo permite orígenes confiables.
✅ Validación de Entrada → Sanitización de datos en handlers/.

2️⃣ Capa de Servicios (internal/services/)
✅ Reglas de Negocio → Controla qué acciones están permitidas.
✅ Protección contra Inyección SQL → Usa Consultas Preparadas con database/sql.
✅ Logs de Seguridad → Registro de eventos críticos (intentos de login fallidos, etc.).

3️⃣ Capa de Seguridad (internal/security/)
✅ Manejo Seguro de Contraseñas → Hashing con bcrypt (golang.org/x/crypto/bcrypt).
✅ Manejo de JWT Seguro → Firma con clave privada, expiración corta.
✅ Control de Accesos (RBAC o ABAC) → Definir permisos según roles.

4️⃣ Capa de Datos (internal/repository/)
✅ Principio de Menor Privilegio → Usuarios de BD con permisos mínimos.
✅ Evitar Exposición de Datos Sensibles → No exponer contraseñas, emails en logs.
✅ Cifrado de Datos Sensibles → Usa AES-GCM para información confidencial