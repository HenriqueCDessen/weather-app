# ☁️ Weather App - Go + Cloud Run

Este é um serviço web simples desenvolvido em **Go** que recebe um **CEP brasileiro** e retorna a **temperatura atual** da cidade correspondente, formatada em:

- 🌡️ Celsius
- 🌡️ Fahrenheit
- 🌡️ Kelvin

O serviço utiliza:

- 🔎 [ViaCEP](https://viacep.com.br) para descobrir a cidade/estado a partir do CEP
- 🌤️ [WeatherAPI](https://www.weatherapi.com/) para obter a temperatura atual
- 🚀 Deploy no **Google Cloud Run**
- 🐳 Empacotado com **Docker**

---

# Request google cloud run 
```
https://weather-api-775625234960.southamerica-east1.run.app/weather/14030430 
```

## 📦 Exemplo de uso

### Requisição:

```http
GET /weather/01001000
Resposta:

{
  "temp_C": 27.3,
  "temp_F": 81.14,
  "temp_K": 300.3
}
```

✅ Regras de negócio
```
-Cenário código HTTP	Resposta
-CEP válido e encontrado	200	Temperaturas
-CEP inválido (ex: menos de 8 dígitos)	422	invalid zipcode
-CEP válido mas inexistente	404	can not find zipcode
```
🚀 Como rodar localmente (com Docker)
Crie seu .env (ou defina a variável):


WEATHER_API_KEY=your_weather_api_key
Execute com Docker Compose:
```
docker compose up --build
```
Acesse:
```
http://localhost:8080/weather/01001000
```
🛠️ Build manual

```
docker build -t weather-api .
docker run -p 8080:8080 -e WEATHER_API_KEY=your_weatherapi_key weather-api
```

🧪 Testes automatizados
Os testes estão localizados em test/weather_test.go.

Execute com:
```
go test ./...
```

📁 Estrutura do projeto

```
weather-app/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── handler/
│   └── weather.go
├── service/
│   ├── zipcode_service.go
│   └── weather_service.go
├── model/
│   └── weather.go
├── internal/
│   └── http_client.go
├── test/
│   └── weather_test.go
└── README.md
```

📌 Tecnologias
Go 1.22+
Docker / Docker Compose
Google Cloud Run
WeatherAPI
ViaCEP

👨‍💻 Autor
Henrique Dessen

Projeto criado para fins de estudo com foco em integração de APIs e deploy serverless na nuvem ☁️
