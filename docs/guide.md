# Gerando Documentação de API com MkDocs

Este guia irá ajudá-lo a criar uma documentação de API bonita e profissional usando MkDocs com o tema Material for MkDocs.

## 1. Pré-requisitos

Antes de começar, você precisa ter o Python e o pip (gerenciador de pacotes do Python) instalados.

Você pode instalar o MkDocs e o tema Material for MkDocs usando o pip:

```bash
pip install mkdocs mkdocs-material
```

## 2. Configuração do Projeto

Se você ainda não tem um projeto MkDocs, pode criar um com o seguinte comando:

```bash
mkdocs new .
```

Isso criará um arquivo `mkdocs.yml` e um diretório `docs` com um arquivo `index.md`.

## 3. Configurando o `mkdocs.yml`

Edite seu arquivo `mkdocs.yml` para configurar o tema Material e adicionar algumas funcionalidades úteis. Abaixo está um exemplo de configuração:

```yaml
site_name: Minha API
site_url: https://exemplo.com/

theme:
  name: material
  palette:
    # Paleta de cores (opcional)
    scheme: default
  features:
    # Navegação superior
    - navigation.tabs
    # Expansão automática da navegação à esquerda
    - navigation.expand

nav:
  - 'Início': 'index.md'
  - 'Guia da API': 'guide.md'

markdown_extensions:
  - pymdownx.highlight:
      anchor_linenums: true
  - pymdownx.superfences
  - admonition
```

## 4. Escrevendo a Documentação da API

Agora você pode começar a escrever sua documentação no diretório `docs`. Use Markdown para formatar seu conteúdo. Aqui estão algumas dicas:

### Estrutura

Organize sua documentação por endpoints. Por exemplo, você pode ter uma página para cada recurso da sua API (e.g., `usuarios.md`, `produtos.md`).

### Exemplos de Requisição e Resposta

Use blocos de código para mostrar exemplos de requisição e resposta. Com a extensão `pymdownx.superfences`, você pode aninhar blocos de código para mostrar, por exemplo, um `curl` e a resposta JSON.

````markdown
```json
{
  "id": 1,
  "nome": "Usuário Exemplo"
}
```
````

### Parâmetros

Use tabelas para descrever os parâmetros da sua API:

```markdown
| Parâmetro | Tipo   | Descrição                |
|-----------|--------|--------------------------|
| `id`      | `int`  | O ID do usuário.         |
| `nome`    | `string` | O nome do usuário.       |
```

### Admonitions

Use admonitions para destacar informações importantes:

```markdown
!!! note "Nota"
    Esta é uma nota importante.

!!! warning "Atenção"
    Este endpoint será descontinuado.
```

## 5. Documentação Automática com `mkdocstrings`

`mkdocstrings` é uma extensão poderosa que gera documentação a partir do seu código-fonte. Como você está trabalhando com Go, você pode usar o handler para Go.

### Instalação

```bash
pip install mkdocstrings
```

### Configuração

Adicione `mkdocstrings` ao seu `mkdocs.yml`:

```yaml
markdown_extensions:
  - pymdownx.highlight:
      anchor_linenums: true
  - pymdownx.superfences
  - admonition
  - mkdocstrings:
      handlers:
        go:
          options:
            # Caminho para o seu código Go
            source: ../
```

### Uso

Agora, em seus arquivos Markdown, você pode usar a seguinte sintaxe para inserir a documentação de um pacote ou função Go:

```markdown
::: meu.pacote.go
```

Isso irá renderizar a documentação para o pacote `meu.pacote.go` diretamente na sua página.

## 6. Visualizando e Compilando

Para visualizar sua documentação localmente, execute o seguinte comando na raiz do projeto:

```bash
mkdocs serve
```

Isso iniciará um servidor web local e você poderá ver sua documentação em `http://127.0.0.1:8000`.

Para compilar sua documentação em arquivos HTML estáticos, use:

```bash
mkdocs build
```

Os arquivos serão gerados em um diretório chamado `site`.

## 7. Publicando

Você pode hospedar sua documentação em qualquer lugar que sirva arquivos estáticos. O MkDocs tem um comando integrado para publicar no GitHub Pages:

```bash
mkdocs gh-deploy
```

Seguindo estes passos, você terá uma documentação de API profissional, bonita e fácil de manter.