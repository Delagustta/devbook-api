# 📦 Desafio Backend Pleno - API de Catálogo para Marketplace

Este projeto é uma API RESTful desenvolvida para gerenciar produtos e categorias em um marketplace. A aplicação utiliza **Java com Spring Boot**, **MongoDB** e serviços da **AWS** como **S3**, **SQS** e **Lambda** para garantir escalabilidade e alta performance.

## 🚀 Tecnologias Utilizadas

- **Java 17**
- **Spring Boot**
- **MongoDB**
- **AWS S3**
- **AWS SQS**
- **AWS Lambda**
- **Maven**

## 📁 Estrutura do Projeto

```
📦 src
├── 📂 main
│   ├── 📂 java
│   │   └── 📂 com
│   │       └── 📂 marketplace
│   │           ├── 📂 controller       # Controladores da API
│   │           ├── 📂 service          # Lógica de negócios
│   │           ├── 📂 repository       # Interfaces para acesso ao banco de dados
│   │           └── 📂 model            # Modelos de dados
│   └── 📂 resources
│       └── 📄 application.properties   # Configurações da aplicação
├── 📂 test
│   └── 📂 java
│       └── 📂 com
│           └── 📂 marketplace
│               ├── 📂 controller       # Testes dos controladores
│               ├── 📂 service          # Testes dos serviços
│               └── 📂 repository       # Testes dos repositórios
```

## ⚙️ Configuração do Ambiente

Antes de executar a aplicação, configure as seguintes variáveis de ambiente:

- `AWS_ACCESS_KEY_ID`: Chave de acesso da AWS.
- `AWS_SECRET_ACCESS_KEY`: Chave secreta da AWS.
- `AWS_REGION`: Região onde os serviços da AWS estão configurados.
- `MONGODB_URI`: URI de conexão com o MongoDB.

## 📩 Endpoints da API

### Produtos

- **POST** `/api/product`: Criar um novo produto.
- **GET** `/api/product`: Listar todos os produtos.
- **PUT** `/api/product/{id}`: Atualizar um produto existente.
- **DELETE** `/api/product/{id}`: Remover um produto.

### Categorias

- **POST** `/api/category`: Criar uma nova categoria.
- **GET** `/api/category`: Listar todas as categorias.
- **PUT** `/api/category/{id}`: Atualizar uma categoria existente.
- **DELETE** `/api/category/{id}`: Remover uma categoria.

## ☁️ Integração com AWS

- **S3**: Utilizado para armazenar arquivos JSON que representam o catálogo de produtos e categorias.
- **SQS**: Gerencia mensagens relacionadas a atualizações no catálogo, permitindo processamento assíncrono.
- **Lambda**: Processa mensagens da fila SQS e atualiza os arquivos no S3 conforme necessário.