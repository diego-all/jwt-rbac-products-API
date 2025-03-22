# Follow Up


### PAYASADA

ValidToken (model) vs ValidateToken (handler)

func (t *Token) ValidToken(plainText string) (bool, error) {}
func (app *application) ValidateToken(w, r){}

Para CustoToken:

- ValidToken() es llamado desde **routes.go** /test-validate-token
- ValidateToken() es llamado desde **routes.go** mux.Post("/validate-token", app.ValidateToken)

**SON DE PRUEBA**

Realmente el protagonista es: func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {


### Token info (Database <> Response)

"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"

	cmd/api/handlersCustomTokenFull.txt 
	GetDataForUpdateHandlerToken()  // Forma artesanal

	InsertReturning() // Utilizando RETURNING con postgres

Hubo un tema de modificacion de punteros



### Relacion entre Token y User

type User struct {
	ID        int; Email string; FirstName string; LastName string; Password string; CreatedAt time.Time; UpdatedAt time.Time; Token Token
}

**No entiendo la estructura del response.** ¿Para que retornar en 2 veces el token? De igual forma es mala practica retornar tanta informacion.

El atributo Token Token \json:"token_at"`dentro de la estructuraUser` está ahí para asociar la información del token JWT directamente con el usuario. Sin embargo, esta práctica no es común ni recomendable en la mayoría de los casos.

¿Cual es el porque?

La intención podría ser:

Incluir el token en la respuesta del usuario cuando se autentica, para que el frontend pueda almacenarlo y usarlo en futuras solicitudes.
Asociar un token persistente con el usuario en la base de datos, lo cual puede ser útil si manejas sesiones activas en lugar de JWT puramente stateless.


### Retomando llamado de Handler (Valid)

Principio de responsabilidad unica en token.go

El nombre mas adecuado para validToken es validateToken().
PEndiente realizar el cambio.
QUedo funcionando bien.


### Visibilidad desde paquete models del paquete main.


Actualmente toda la API en main y modelos por aparte.

El Objetivo es poder invocar el infoLog y el errorLog que estan declarados en el paquete main desde el paquete models. como se muestra a continuacion:
	// app.infoLog.Println("API listening on port", app.config.port)  NO FUNCION!!!


	Dejar los modelos quietos
	Utilizar loggers
	Crear paquete config

	De no ser posible validar interfaces.

Exportar las interfaces. NO

	func New(dbPool *sql.DB, infoLog, errorLog *log.Logger) Models {
		return Models{
			Product:  Product{},
			User:     User{},
			Token:    Token{},
			JWTToken: JWTToken{},
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		}

Paquete de utilidades para logs

Definir un paquete logger donde almacenes instancias globales de infoLog y errorLog, y puedas usarlas en cualquier paquete.
Es decir un paquete global para que todos los psquetes lo usen.


https://github.com/uber-go/zap

Considerar el tema del flag debugging al final


### Arquitectura de capas

¿Dónde ubicar el módulo de configuración?
proyecto sigue una estructura modular y organizada
Path recomendado: internal/config/config.go
Esto mantendrá la configuración separada de la lógica de negocio y facilitará su mantenimiento.

¿Donde deberia ir alojada la logica de negocio? 

En un proyecto en Golang siguiendo buenas prácticas, la lógica de negocio debería estar separada en paquetes organizados dentro de la carpeta internal/.


**Por ahora** Sientase libre de colocar la logica de negocio como prefiera.
Depende de la arquitectura. Capas, MVC, etc ...

Recomendacion:

1. internal/services  → Aquí puedes alojar la lógica de negocio principal.

internal/services/product_service.go
internal/services/user_service.go
internal/services/auth_service.go

**Son necesarios los servicios.**   ==> Microservicios

Cada servicio contendría funciones específicas relacionadas con las entidades del negocio.


2. internal/repository/ → Para la capa de acceso a datos.

internal/repository/ → Para la capa de acceso a datos.
internal/repository/product_repository.go
internal/repository/user_repository.go

**otocse saca el repositorio**

3. internal/models/ → Definición de estructuras y modelos.

internal/models/product.go
internal/models/user.go

Contiene las estructuras que representan las entidades del dominio de la aplicación.

4. internal/logger/ → Ya lo tienes, y es correcto mantenerlo allí.

5. cmd/api/ → Solo debería contener la lógica para inicializar la aplicación y definir rutas.

**Clean Architecture (Uncle Bob) Separar la logica de negocio de la infraestructura, lo que facilita la implemenetación de medidas de seguridad en cada capa.**

scaffolding + API 
modelos (gemini)

microservicio autenticacion

**OBJETIVO:**

**DEFINIR UNA ARQUITECTURA**




## Modulo Config





