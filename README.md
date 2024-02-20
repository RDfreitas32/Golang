# CRUD Básico em Golang

Este projeto implementa um CRUD (Create, Read, Update, Delete) básico para usuários em Golang. Utiliza o pacote Gorilla Mux para lidar com as rotas HTTP e realiza interações com um banco de dados.

## Estrutura do Projeto

O projeto contém os seguintes arquivos e diretórios principais:

- **main.go**: O arquivo principal que contém a lógica para configurar o servidor HTTP e definir as rotas CRUD.
- **servidor/**: O diretório que contém os arquivos com as funções de manipulação das rotas HTTP para os diferentes métodos CRUD.
- **README.md**: Este arquivo, fornecendo informações sobre o projeto.

## Funcionalidades

O projeto oferece as seguintes funcionalidades CRUD para os usuários:

- **Create (POST)**: Adiciona um novo usuário ao banco de dados.
- **Read (GET)**: Retorna informações de todos os usuários ou de um usuário específico.
- **Update (PUT)**: Atualiza as informações de um usuário existente no banco de dados.
- **Delete (DELETE)**: Remove um usuário do banco de dados.
