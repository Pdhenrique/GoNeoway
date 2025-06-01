# GoNeoway

# GoNeoway

## Serviço em Go para importação, higienização e persistência de dados em PostgreSQL via Docker.

## 🐳 Executando com Docker Compose

### Requisitos

- Docker instalado
- Docker Compose v2+

### Passos para execução

1. Clone este repositório:

```bash
git clone https://github.com/Pdhenrique/GoNeoway.git
cd GoNeoway
```

2. Crie um arquivo .env:

```env
DB_HOST=db
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=mydb
```

3. Suba a aplicação com Docker Compose:

Arquivo txt é importado automaticamente ao inicializar a aplicação, com o banco rodando.

```bash
docker compose up --build
```

4. Importe o arquivo `goneoway.postman_collection.json` no postman
   Ou utilize estes curls

```bash
GET: curl --location 'http://localhost:8080/v1/clients/12345678900'

DELETE: curl --location --request DELETE 'http://localhost:8080/v1/clients/12345678900'

POST: curl --location 'http://localhost:8080/v1/clients' \
--header 'Content-Type: application/json' \
--data '{
  "cpf": "12345678900",
  "private": 1,
  "incompleto": 0,
  "data_ultima_compra": "2024-06-01T00:00:00Z",
  "ticket_medio": 341.45,
  "ticket_ultima_compra": 399.00,
  "loja_mais_frequente": "79379491000850",
  "loja_ultima_compra": "79379491000850"
}'

PUT: curl --location 'http://localhost:8080/v1/clients' \
--header 'Content-Type: application/json' \
--data '{
  "cpf": "12345678900",
  "private": 1,
  "incompleto": 0,
  "data_ultima_compra": "2024-06-01T00:00:00Z",
  "ticket_medio": 341.45,
  "ticket_ultima_compra": 399.00,
  "loja_mais_frequente": "79379491000850",
  "loja_ultima_compra": "79379491000850"
}'
```

## 📁 Estrutura do Projeto

```
.gitignore
base_teste_2.txt
base_teste.txt
cmd/app/main.go
db/init/create_clients_table.sql
docker-compose.yml
Dockerfile
domain/client.go
go.mod
internal/db/clientRepository.go
internal/db/connect.go
internal/http/client.go
internal/http/handler.go
internal/http/server.go
pkg/client/client.go
pkg/importer/service.go
pkg/parser/parser.go
pkg/sanitizer/sanitizer.go
README.md
```

### Documentações e guias usados:

https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
https://github.com/jhonyzam/curso_golang
https://github.com/golang-standards/project-layout
