API REST em Go com Docker

Uma API RESTful moderna desenvolvida em Go (Golang) utilizando Docker para conteinerização e implantação simplificada.
🚀 Tecnologias Utilizadas

    Go (Golang) - Linguagem de programação principal

    Docker - Conteinerização da aplicação

    Docker Compose - Orquestração de containers

    Gin Framework - Framework web para Go

    PostgreSQL - Banco de dados relacional (se aplicável)

📋 Pré-requisitos

Antes de executar o projeto, certifique-se de ter instalado:

    Docker

    Docker Compose

    Git

🛠️ Instalação e Execução
1. Clonar o repositório
bash

git clone https://github.com/viniciusbackend/apirestgo-docker.git
cd apirestgo-docker

2. Executar com Docker Compose
bash

# Executar a aplicação em modo desenvolvimento
docker-compose up

# Executar em segundo plano
docker-compose up -d

# Parar a aplicação
docker-compose down

3. Build manual da imagem Docker
bash

# Build da imagem
docker build -t apirestgo-docker .

# Executar container
docker run -p 8080:8080 apirestgo-docker

📦 Estrutura do Projeto
text


🔧 Variáveis de Ambiente

O projeto utiliza as seguintes variáveis de ambiente (configuradas no docker-compose.yml):

    PORT - Porta da aplicação (padrão: 8080)

    DB_HOST - Host do banco de dados

    DB_PORT - Porta do banco de dados

    DB_USER - Usuário do banco

    DB_PASSWORD - Senha do banco

    DB_NAME - Nome do banco de dados

📡 Endpoints da API
Exemplo de endpoints disponíveis:
text

GET    /api/health       - Status da aplicação
GET    /api/users        - Listar usuários
POST   /api/users        - Criar usuário
GET    /api/users/:id    - Buscar usuário por ID
PUT    /api/users/:id    - Atualizar usuário
DELETE /api/users/:id    - Deletar usuário

🐛 Desenvolvimento
Executar localmente sem Docker:
bash

# Instalar dependências
go mod download

# Executar aplicação
go run main.go

Testar a aplicação:
bash

# Testar endpoints
curl http://localhost:8080/api/health

📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.
🤝 Contribuição

Contribuições são sempre bem-vindas! Por favor:

    Faça um fork do projeto

    Crie uma branch para sua feature (git checkout -b feature/AmazingFeature)

    Commit suas mudanças (git commit -m 'Add some AmazingFeature')

    Push para a branch (git push origin feature/AmazingFeature)

    Abra um Pull Request