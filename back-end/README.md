# LibraryAPI

LibraryAPI é uma REST API que simula o gerenciamento de uma livraria. É possível criar, editar, visualizar e deletar livros, autores e categorias. Porém, para criar, editar ou deletar, é necessário estar cadastrado na aplicação. Caso o usuário não esteja cadastrado, será possível apenas visualizar os livros.

## Tecnologias utilizadas
Essa API foi desenvolvida utilizando a linguagem [Go](https://go.dev/), em conjunto com o Object-Relational Mapper (ORM) [GORM](https://gorm.io/) e o banco de dados relacional [MySQL](https://www.mysql.com/).  
A autenticação nas rotas é feita com JSON Web Token [(JWT)](https://jwt.io/), concedido no login.  
As senhas dos usuários são criptografadas antes de serem armazenadas no banco de dados.
