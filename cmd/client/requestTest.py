import requests

# Ruta a los certificados
CERT_FILE = "client.pem"
KEY_FILE = "client.key"

# URL del endpoint
URL = "https://localhost:9090/admin/products/all"

# Token de autorización
TOKEN = "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImRpZWdvQGRpZWdvLmNvbSIsInRva2VuIjoiIiwidG9rZW5faGFzaCI6IiIsImV4cGlyeSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicm9sZSI6InRyaW4iLCJpc3MiOiJnM25vdHlwZSIsInN1YiI6ImRpZWdvQGRpZWdvLmNvbSIsImF1ZCI6WyJtaXMtdXN1YXJpb3MiXSwiZXhwIjoxNzM5NTkxMTA4LCJuYmYiOjE3Mzk1MDQ3MDgsImlhdCI6MTczOTUwNDcwOCwiY3JlYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwidXBkYXRlZF9hdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIn0.1VY014JlOLAjcjD1h6jNMc2pCLeJDbU0CKvl9ofSIHJMfHIjGAvOSdLAcSX7Nsl0kLJW1ZpOW0ehfRIYz-hl7g"

# Encabezados de la petición
HEADERS = {"Authorization": TOKEN}

try:
    # Realizar la solicitud POST con certificados
    response = requests.post(URL, headers=HEADERS, cert=(CERT_FILE, KEY_FILE), verify=False)
    
    # Imprimir código de estado y respuesta
    print("Código de estado:", response.status_code)
    print("Respuesta:", response.text)
except requests.exceptions.RequestException as e:
    print("Error al ejecutar la solicitud:", e)
