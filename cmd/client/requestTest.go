package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Cargar los certificados del cliente
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		fmt.Println("Error cargando certificados:", err)
		return
	}

	// Configurar el pool de certificados
	certPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile("client.pem")
	if err != nil {
		fmt.Println("Error leyendo client.pem:", err)
		return
	}
	certPool.AppendCertsFromPEM(caCert)

	// Configurar el transporte HTTP con TLS
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            certPool,
			InsecureSkipVerify: true, // Permitir conexiones inseguras (equivalente a -k en curl)
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", "https://localhost:9090/admin/products/all", nil)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return
	}

	// Agregar encabezado de autorización
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzM5NTkxMTA4LCJuYmYiOjE3Mzk1MDQ3MDgsImlhdCI6MTczOTUwNDcwOCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.1VY014JlOLAjcjD1h6jNMc2pCLeJDbU0CKvl9ofSIHJMfHIjGAvOSdLAcSX7Nsl0kLJW1ZpOW0ehfRIYz-hl7g")

	// Ejecutar la solicitud
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al ejecutar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	fmt.Println("Código de estado:", resp.Status)
	fmt.Println("Respuesta:", string(body))
}
