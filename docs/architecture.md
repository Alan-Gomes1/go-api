# Arquitetura do Projeto

A Go API utiliza uma arquitetura baseada no padrão **MVC (Model-View-Controller)**, adaptada para o contexto de uma API RESTful. Essa escolha visa organizar o código de forma clara, modular e escalável.

## Visão Geral da Arquitetura

A estrutura do projeto é dividida nas seguintes camadas:

- **Model:** Representa os dados e a lógica de negócio da aplicação. No nosso caso, a camada de `model` lida com a estrutura de dados do usuário, validações e a comunicação com o banco de dados.
- **View:** Responsável pela apresentação dos dados. Em uma API RESTful, a "View" é a representação dos dados em formatos como JSON. A camada de `view` em nosso projeto converte os objetos de domínio para os formatos de resposta JSON.
- **Controller:** Atua como um intermediário entre o `Model` e a `View`. Ele recebe as requisições HTTP, aciona a lógica de negócio no `Model` e retorna a resposta formatada pela `View`.

## Benefícios da Arquitetura MVC

A adoção do padrão MVC traz diversos benefícios para o desenvolvimento e a manutenção da API:

1.  **Organização e Clareza:** A separação de responsabilidades torna o código mais fácil de entender e dar manutenção. Cada camada tem um propósito bem definido.
2.  **Reutilização de Código:** A lógica de negócio no `Model` pode ser reutilizada por diferentes `Controllers`, evitando duplicação de código.
3.  **Desenvolvimento Paralelo:** Como as camadas são desacopladas, desenvolvedores podem trabalhar em paralelo no `Model`, na `View` e no `Controller` sem grandes conflitos.
4.  **Facilidade de Testes:** A separação de responsabilidades facilita a criação de testes unitários para cada camada de forma isolada.

## Motivos da Escolha

A escolha da arquitetura MVC para este projeto foi motivada pelos seguintes fatores:

-   **Simplicidade e Eficiência:** O MVC é um padrão bem conhecido e relativamente simples de implementar, o que acelera o desenvolvimento.
-   **Escalabilidade:** A estrutura modular permite que a aplicação cresça de forma organizada, facilitando a adição de novas funcionalidades.
-   **Manutenibilidade:** Com o código bem organizado, a correção de bugs e a implementação de melhorias se tornam mais simples e seguras.
-   **Clareza para a Equipe:** O uso de um padrão de arquitetura conhecido facilita a colaboração e a integração de novos desenvolvedores ao projeto.

## Detalhamento das Camadas

Além da estrutura MVC principal, o projeto aprofunda a separação de responsabilidades com camadas adicionais dentro do `Model`, como `Service` e `Repository`.

### Camada de Serviço (Service)

A camada de **Serviço** é responsável por orquestrar a lógica de negócio da aplicação. Ela atua como uma ponte entre o `Controller` e a camada de acesso a dados (`Repository`).

**Responsabilidades:**

-   Implementar as regras de negócio complexas.
-   Coordenar operações que envolvem múltiplos repositórios ou fontes de dados.
-   Realizar validações de negócio antes de persistir os dados.
-   Garantir a consistência das transações.

**Benefícios:**

-   **Desacoplamento:** Isola a lógica de negócio dos detalhes de implementação do banco de dados e da apresentação (controller).
-   **Reutilização:** A mesma lógica de serviço pode ser consumida por diferentes `Controllers` ou até mesmo por outros serviços.
-   **Testabilidade:** Facilita a criação de testes unitários para a lógica de negócio, pois os serviços podem ser testados de forma isolada, utilizando *mocks* para os repositórios.

### Camada de Repositório (Repository)

A camada de **Repositório** implementa o padrão *Repository Pattern*, que tem como principal objetivo abstrair a lógica de acesso e persistência de dados.

**Responsabilidades:**

-   Encapsular as queries e a lógica de comunicação com o banco de dados (ex: MongoDB, PostgreSQL).
-   Fornecer uma interface clara e consistente para a manipulação de entidades de domínio (CRUD - Create, Read, Update, Delete).
-   Mapear os dados entre os objetos de domínio da aplicação e a estrutura de dados do banco.

**Benefícios:**

-   **Abstração do Banco de Dados:** A camada de serviço e o restante da aplicação não precisam saber qual banco de dados está sendo utilizado. Isso facilita a troca do SGBD no futuro, se necessário.
-   **Centralização do Acesso a Dados:** Todo o código de acesso a dados fica centralizado em um único lugar, tornando a manutenção mais simples.
-   **Código Mais Limpo:** Os `Services` ficam mais limpos e focados na lógica de negócio, sem se preocupar com os detalhes da persistência.