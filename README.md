# LibraryAPI

LibraryAPI é uma REST API que simula o gerenciamento de uma biblioteca. É possível criar, editar, visualizar e deletar livros, autores e categorias. Porém, para criar, editar ou deletar, é necessário estar cadastrado na aplicação. Caso o usuário não esteja cadastrado, será possível apenas visualizar os livros.

## Tecnologias utilizadas
Essa API foi desenvolvida utilizando a linguagem [Go](https://go.dev/), em conjunto com o Object-Relational Mapper (ORM) [GORM](https://gorm.io/) e o banco de dados relacional [MySQL](https://www.mysql.com/).  
A autenticação nas rotas é feita com JSON Web Token [(JWT)](https://jwt.io/), concedido no login.  
As senhas dos usuários são criptografadas antes de serem armazenadas no banco de dados.

## Preparando o ambiente
Para executar essa aplicação, será necessário que o [Docker](https://docs.docker.com/engine/install/) e o [Docker Compose](https://docs.docker.com/compose/install/) estejam instalados em sua máquina.  
Abra o terminal e siga os seguintes passos:

1. Clone o repositório da aplicação:
```sh
 git clone git@github.com:raelnogpires/libraryapp.git
```
2. Entre no repositório:
```sh
 cd libraryapp
```
3. Entre no diretório `src/` e copie o arquivo `.env.example` para `.env` :
```sh
 cd src && cp .env.example .env
```

## Executando a aplicação
Com o terminal aberto na raiz da aplicação, execute:
```sh
 docker compose up
```
Aguarde o seguinte log aparecer:
```sh
 library_api | db running
 library_api | server running at port: 5000
```
Para encerrar a aplicação utilize `Ctrl + C` .

## Documentação

### Cadastro de um novo usuário
Cadastra uma nova pessoa usuária.  

#### URL
```sh
 POST http://localhost:5000/api/v1/register
```

#### Parâmetros

##### Body
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| username      | string   | Nome/apelido de usuário. **Obrigatório**.             |
| email         | string   | Email no formato "user@library.com". **Obrigatório**. |
| password      | string   | Senha do usuário. **Obrigatório**.                    |

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 201        | user registered with success       |
| 400        | invalid request body               |
| 400        | email already registered           |

#### Exemplo
Requisição:
```json
 {
   "username": "reader",
   "email": "ilovebooks@lispector.com",
   "password": "1d5as15d"
 }
```

Resposta:
```json
 {
   "message": "user registered with success",
 }
```