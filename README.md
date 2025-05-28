# GoNeoway

# GoNeoway

Serviço em Go para importação, higienização e persistência de dados em PostgreSQL via Docker.

---

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

```bash
docker compose up --build
```

## 📁 Estrutura do Projeto

.
├── cmd/app/ # main.go - ponto de entrada
├── internal/db/ # conexão com o banco de dados
├── pkg/ # módulos reutilizáveis (parser, validações, etc.)
├── docker-compose.yml # define os serviços (app e banco)
├── Dockerfile # build da aplicação Go
├── .env # variáveis de ambiente
├── go.mod / go.sum # dependências do projeto
├── README.md # instruções e documentação

### Documentações e guias usados:

https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
https://github.com/jhonyzam/curso_golang
https://github.com/golang-standards/project-layout
