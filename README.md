# go-api

API RESTful desenvolvida em Go para gerenciamento de usuários, utilizando MongoDB como banco de dados principal. O projeto segue uma arquitetura modular, separando responsabilidades em camadas de controller, service e repository, além de utilizar validação, logging e tratamento de erros customizados.

## Tecnologias Utilizadas

- **Go** (Golang)
- **Gin** (framework web)
- **MongoDB** (banco de dados NoSQL)
- **go.mongodb.org/mongo-driver** (driver oficial MongoDB)
- **go-playground/validator** (validação de dados)
- **Uber Zap** (logger)
- **godotenv** (carregamento de variáveis de ambiente)

## Estrutura do Projeto

```
go-api/
├── main.go                        # Ponto de entrada da aplicação
├── init_dependencies.go           # Injeção de dependências
├── go.mod / go.sum                # Gerenciamento de dependências
├── src/
│   ├── configuration/             # Configurações (logger, validação, erros, banco)
│   ├── controller/                # Controllers e rotas
│   ├── model/                     # Domínio, repository, service, entidades
│   └── view/                      # Conversores de resposta
```

## Funcionalidades Principais

- **Criar usuário**: `POST /createUser`
- **Buscar usuário por ID**: `GET /user/:userId`
- **Buscar usuário por e-mail**: `GET /userByEmail/:userEmail`
- **Atualizar usuário**: `PUT /updateUser/:userId`
- **Deletar usuário**: `DELETE /deleteUser/:userId`

## Como rodar o projeto

1. **Clone o repositório:**

```bash
git clone https://github.com/Alan-Gomes1/go-api.git
cd go-api
```

2. **Configure as variáveis de ambiente:**

Renomeie o arquivo `.env-example` para `.env` na raiz do projeto e defina os valores reais para as variáveis

3. **Instale as dependências:**

```bash
go mod tidy
```

4. **Execute a aplicação:**

```bash
go run main.go init_dependencies.go
```

A API estará disponível em `http://localhost:8080`.

## Observações

- Certifique-se de que o MongoDB está rodando localmente ou ajuste a URL conforme necessário.
- O projeto utiliza boas práticas de separação de camadas e pode ser expandido facilmente para novas funcionalidades.

---

Desenvolvido por Alan Gomes
