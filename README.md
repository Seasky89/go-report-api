# Go Report API

API REST simples desenvolvida em Go para gerar um relatório de usuários com total de posts.

A aplicação consome dados de uma API externa (JSONPlaceholder) e organiza a estrutura utilizando arquitetura em camadas (handler → service → store).

---

## 🚀 Funcionalidades

- Listar usuários
- Listar posts
- Gerar relatório com total de posts por usuário
- Armazenamento em memória
- Consumo de API externa
- Injeção manual de dependências

---

## 🏗 Arquitetura

O projeto segue uma organização em camadas:

- **Handlers** → Camada HTTP (Gin)
- **Services** → Regras de negócio
- **Store** → Abstração de dados (implementação em memória)
- **Bootstrap** → Inicialização da aplicação

As dependências são organizadas através de uma *composition root* (`App`), centralizando a criação de serviços e handlers.

---

## 📁 Estrutura do Projeto

```text
cmd/
internal/
├── app/
├── bootstrap/
├── handlers/
├── services/
├── store/
├── models/
├── routes/
└── external/
go.mod
README.md
```

---

## 🛠 Tecnologias utilizadas

- Go
- Gin (framework HTTP)
- JSONPlaceholder (fonte de dados externa)

---

## ▶️ Como executar

Clone o repositório:

```shell
git clone https://github.com/seu-usuario/go-report-api.git
```

Acesse a pasta do projeto:

```shell
cd go-report-api
```

Instale as dependências:

```shell
go mod tidy
```

Execute a aplicação:

```shell
go run ./cmd
```

Servidor disponível em:

```shell
http://localhost:8080
```

---

## 📌 Endpoints

### Gerar relatório completo
```shell
GET /report
```

### Listar usuários
```shell
GET /report/users
```

### Listar posts
```shell
GET /report/posts
```

---
