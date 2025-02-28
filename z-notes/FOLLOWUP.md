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