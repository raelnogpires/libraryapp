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

<details>
  <summary><strong>Detalhes</strong></summary>

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

</details>

### Login
Realiza o login com email e senha de uma pessoa usuária já cadastrada.  
Corpo da resposta contém um objeto JSON com o token de autenticação gerado e o `id` do usuário.

<details>
  <summary><strong>Detalhes</strong></summary>

#### URL
```sh
 POST http://localhost:5000/api/v1/login
```

#### Parâmetros

##### Body
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| email         | string   | Email da pessoa usuária. **Obrigatório**.             |
| password      | string   | Senha da pessoa usuária. **Obrigatório**.             |

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 200        | token generated with success       |
| 400        | invalid credentials                |
| 404        | user not found                     |

#### Exemplo
Requisição:
```json
 {
   "email": "ilovebooks@lispector.com",
   "password": "1d5as15d"
 }
```

Resposta:
```json
 {
   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjIsImV4cCI6MTY1MzY5MDQ2MCwiaWF0IjoxNjUzNjY4ODYwLCJpc3MiOiJsaWJyYXJ5In0.kp7TQZmRHz-5ENJMa9KZ0mRVg35Xd8sm08WKmX_MFXg",
   "user_id": 1
 }
```

</details>

### Deletar seu usuário
Deleta o usuário que está logado.

<details>
  <summary><strong>Detalhes</strong></summary>

#### URL
```sh
 DELETE http://localhost:5000/api/v1/user/me
```

#### Autorização
Requer token de autenticação no campo `Token` do `Bearer Token` .

#### Parâmetros
Nenhum.

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 204        | user deleted with success          |
| 400        | invalid token                      |
| 404        | token not found                    |
| 404        | user not found                     |

#### Exemplo
Requisição:  
![deleteMe](./readme_utils/deleteMe.png)

</details>

### Buscar todos os livros
Busca todos os livros cadastrados.  
Corpo da resposta contém um array JSON com os livros e suas informações.

<details>
  <summary><strong>Detalhes</strong></summary>

#### URL
```sh
 GET http://localhost:5000/api/v1/books
```

#### Autorização
Requer token de autenticação no campo `Token` do `Bearer Token` .

#### Parâmetros
Nenhum.

#### Campos da resposta
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| id            | int      | ID do livro.                                          |
| name          | string   | Nome do livro                                         |
| description   | string   | Sinopse do livro.                                     |
| category_id   | int      | ID da categoria que o livro pertence.                 |
| category_name | string   | Nome da categoria que o livro pertence.               |
| author_id     | int      | ID da pessoa autora.                                  |
| author_name   | string   | Nome da pessoa autora.                                |
| img_url       | string   | URL da imagem do livro.                               |

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 200        | books returned with success        |
| 400        | invalid token                      |
| 404        | token not found                    |

#### Exemplo
Requisição:  
![getAllBooks](./readme_utils/getAllBooks.png)

Resposta:
```json
 [
   {
     "id": 1,
     "name": "Sentimento do mundo",
     "description": "O Drummond de Sentimento do mundo oscila entre diversos polos: cidade x interior, atualidade x memórias, eu x mundo. Perfeita depuração dos livros anteriores, este é um verdadeiro marco.",
     "category_id": 6,
     "category_name": "Poesia",
     "author_id": 2,
     "author_name": "Carlos Drummond de Andrade",
     "img_url": "https://images-na.ssl-images-amazon.com/images/I/41Q6T14Y0EL._SX324_BO1,204,203,200_.jpg"
   },
   {
     "id": 2,
     "name": "Crime e Castigo",
     "description": "Crime e Castigo é a obra mais célebre de Fyodor Dostoevsky. Neste livro, Raskólnikov, um jovem estudante, pobre e desesperado, perambula pelas ruas de São Petersburgo até cometer um crime que tentará justificar por uma teoria: grandes homens, como César ou Napoleão, foram assassinos absolvidos pela História.",
     "category_id": 5,
     "category_name": "Literatura Russa",
     "author_id": 4,
     "author_name": "Fyodor Dostoevsky",
     "img_url": "https://images-na.ssl-images-amazon.com/images/I/517DdyXpc5L._SX348_BO1,204,203,200_.jpg"
   },
 ]
```

</details>

### Buscar livro
Buscar um livro específico pelo seu ID.  
Corpo da resposta contém um objeto JSON as com informações do livro.

<details>
  <summary><strong>Detalhes</strong></summary>

#### URL
```sh
 GET http://localhost:5000/api/v1/books/{id}
```

#### Autorização
Requer token de autenticação no campo `Token` do `Bearer Token` .

#### Parâmetros

##### Path
| **Campo** | **Tipo** | **Descrição**              |
|:----------|:---------|:---------------------------|
| id        | int      | ID do livro a ser buscado. |

#### Campos da resposta
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| id            | int      | ID do livro.                                          |
| name          | string   | Nome do livro                                         |
| description   | string   | Sinopse do livro.                                     |
| category_id   | int      | ID da categoria que o livro pertence.                 |
| category_name | string   | Nome da categoria que o livro pertence.               |
| author_id     | int      | ID da pessoa autora.                                  |
| author_name   | string   | Nome da pessoa autora.                                |
| img_url       | string   | URL da imagem do livro.                               |

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 200        | book returned with success         |
| 400        | invalid token                      |
| 404        | token not found                    |
| 404        | book not found                     |

#### Exemplo

Requisição:  
![getBookById](./readme_utils/getBookById.png)

Resposta:  
```json
 {
   "id": 2,
   "name": "Crime e Castigo",
   "description": "Crime e Castigo é a obra mais célebre de Fyodor Dostoevsky. Neste livro, Raskólnikov, um jovem estudante, pobre e desesperado, perambula pelas ruas de São Petersburgo até cometer um crime que tentará justificar por uma teoria: grandes homens, como César ou Napoleão, foram assassinos absolvidos pela História.",
   "category_id": 5,
   "category_name": "Literatura Russa",
   "author_id": 4,
   "author_name": "Fyodor Dostoevsky",
   "img_url": "https://images-na.ssl-images-amazon.com/images/I/517DdyXpc5L._SX348_BO1,204,203,200_.jpg"
 }
```

</details>

### Criar um livro
Cria um novo livro.  
Corpo da resposta contém um objeto JSON com as informações do livro criado.

<details>
  <summary><strong>Detalhes</strong></summary>

#### URL
```sh
 POST http://localhost:5000/api/v1/books
```

#### Autorização
Requer token de autenticação no campo `Token` do `Bearer Token` .

#### Parâmetros

##### Body
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| name          | string   | Nome do livro. **Obrigatório**                        |
| description   | string   | Sinopse do livro. **Obrigatório**                     |
| category_id   | int      | ID da categoria que o livro pertence. **Obrigatório** |
| author_id     | int      | ID da pessoa autora. **Obrigatório**                  |
| img_url       | string   | URL da imagem do livro. **Obrigatório**               |

#### Campos da resposta
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| id            | int      | ID do livro.                                          |
| name          | string   | Nome do livro                                         |
| description   | string   | Sinopse do livro.                                     |
| category_id   | int      | ID da categoria que o livro pertence.                 |
| author_id     | int      | ID da pessoa autora.                                  |
| img_url       | string   | URL da imagem do livro.                               |

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 201        | book created with success          |
| 400        | invalid request body               |
| 400        | invalid token                      |
| 404        | token not found                    |

#### Exemplo
Requisição:
```json
 {
   "name": "O idiota",
   "description": "Publicado originalmente em 1868, este é um desses livros em que o leitor reconhece de imediato a marca do gênio. Nele, o autor russo constrói um dos personagens mais impressionantes de toda a literatura mundial ― o humanista e epilético príncipe Míchkin, mescla de Cristo e Dom Quixote, cuja compaixão sem limites vai se chocar com o desregramento mundano de Rogójin e a beleza enlouquecedora de Nastácia Filíppovna.",
   "category_id": 5,
   "author_id": 4,
   "img_url": "https://images-na.ssl-images-amazon.com/images/I/51EuSosoqJL._SX346_BO1,204,203,200_.jpg"
 }
```

Resposta:
```json
 {
   "id": 3,
   "name": "O idiota",
   "description": "Publicado originalmente em 1868, este é um desses livros em que o leitor reconhece de imediato a marca do gênio. Nele, o autor russo constrói um dos personagens mais impressionantes de toda a literatura mundial ― o humanista e epilético príncipe Míchkin, mescla de Cristo e Dom Quixote, cuja compaixão sem limites vai se chocar com o desregramento mundano de Rogójin e a beleza enlouquecedora de Nastácia Filíppovna.",
   "category_id": 5,
   "author_id": 4,
   "img_url": "https://images-na.ssl-images-amazon.com/images/I/51EuSosoqJL._SX346_BO1,204,203,200_.jpg"
 }
```

</details>

### Editar um livro
Edita um livro especificado pelo seu ID.  
Corpo da resposta contém um objeto JSON com as informações do livro editado.

<details>
  <summary><strong>Detalhes</strong></summary>

#### URL
```sh
 PUT http://localhost:5000/api/v1/books/{id}
```

#### Autorização
Requer token de autenticação no campo `Token` do `Bearer Token` .

#### Parâmetros

#### Path
| **Campo** | **Tipo** | **Descrição**              |
|:----------|:---------|:---------------------------|
| id        | int      | ID do livro a ser editado. |

##### Body
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| name          | string   | Nome do livro. **Obrigatório**                        |
| description   | string   | Sinopse do livro. **Obrigatório**                     |
| category_id   | int      | ID da categoria que o livro pertence. **Obrigatório** |
| author_id     | int      | ID da pessoa autora. **Obrigatório**                  |
| img_url       | string   | URL da imagem do livro. **Obrigatório**               |

#### Campos da resposta
| **Parâmetro** | **Tipo** | **Descrição**                                         |
|:--------------|:---------|:------------------------------------------------------|
| id            | int      | ID do livro.                                          |
| name          | string   | Nome do livro                                         |
| description   | string   | Sinopse do livro.                                     |
| category_id   | int      | ID da categoria que o livro pertence.                 |
| author_id     | int      | ID da pessoa autora.                                  |
| img_url       | string   | URL da imagem do livro.                               |

#### Códigos de status da resposta
| **Código** | **Descrição**                      |
|:-----------|:-----------------------------------|
| 201        | book edited with success           |
| 400        | invalid request body               |
| 400        | invalid token                      |
| 404        | token not found                    |

Requisição:
```json
 {
   "name": "The idiot",
   "description": "A book written by Fyodor Dostoevsky.",
   "category_id": 5,
   "author_id": 4,
   "img_url": "https://images-na.ssl-images-amazon.com/images/I/51EuSosoqJL._SX346_BO1,204,203,200_.jpg"
 }
```

Reposta:
```json
 {
   "id": 3,
   "name": "The idiot",
   "description": "A book written by Fyodor Dostoevsky.",
   "category_id": 5,
   "author_id": 4,
   "img_url": "https://images-na.ssl-images-amazon.com/images/I/51EuSosoqJL._SX346_BO1,204,203,200_.jpg"
 }
```

</details>