package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoRequest(t *testing.T) {
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Se esperaba POST, pero se recibió %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer mock_token" {
			t.Errorf("Token incorrecto: %s", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Success"}`))
	}))
	defer mockServer.Close()

	client := mockServer.Client()
	client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// 🔹 Llamada corregida a DoRequest
	resp, err := DoRequest(client, mockServer.URL, "mock_token")
	if err != nil {
		t.Fatalf("Error en la solicitud: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error al leer respuesta: %v", err)
	}

	expectedStatus := http.StatusOK
	if resp.StatusCode != expectedStatus {
		t.Errorf("Se esperaba %d, pero se recibió %d", expectedStatus, resp.StatusCode)
	}

	expectedBody := `{"message": "Success"}`
	if string(body) != expectedBody {
		t.Errorf("Respuesta incorrecta: %s", string(body))
	}
}

// DoRequest ejecuta una solicitud HTTP POST con el token de autorización y certificados
func DoRequest(client *http.Client, url, token string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return client.Do(req)
}
