# Bem-vindo à Documentação da Go API

Esta é a documentação de referência para a Go API, um projeto de CRUD de usuários para demonstrar como criar APIs em Go.

## Autenticação

Todas as requisições à API precisam de um token de autenticação no cabeçalho `Authorization`.

```
Authorization: Bearer SEU_TOKEN_AQUI
```

---

## Endpoints

### Buscar Usuário por ID

`GET /users/{id}`

Este endpoint retorna um usuário específico baseado no ID fornecido.

#### Parâmetros de URL

| Parâmetro | Tipo   | Descrição                         |
| --------- | ------ | --------------------------------- |
| `id`      | `uuid` | **Obrigatório.** O ID do usuário. |

#### Exemplo de Requisição

```bash
curl -X GET https://sua-api.com/users/be03247d-d53a-44ed-a4f8 \
  -H "Authorization: Bearer SEU_TOKEN_AQUI"
```

#### Exemplo de Resposta

```json
{
  "id": "be03247d-d53a-44ed-a4f8",
  "name": "Alan Gomes",
  "email": "alan.gomes@example.com",
  "age": 25
}
```

!!! note "Nota sobre Rate Limit"
Este endpoint tem um limite de 100 requisições por minuto.
