package logger

import (
	"log"
	"os"
)

// ROVERT
// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	DebugLog *log.Logger
)

//Las interfaces son variables, complementar.

func Init() {
	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLog = log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile) // Oportunity
}

// Es necesario guardar los logs?
// zap Permite cambiar dinámicamente el nivel de logs

// Por ahora no se sabe si es seguro.
// No justifica enviar a produccion una API rota.
//

// logger.Init("DEBUG") // Cambia "DEBUG" por "INFO" o "ERROR" según lo que necesites

// ✔ Puedes cambiar el nivel de logs dinámicamente (logger.Init("INFO")).
// ✔ Solo imprime Debug() si el nivel está en DEBUG.
// ✔ Sigue usando log.New() pero de forma más flexible.
