<div align="center">

<img width="100%" src="https://capsule-render.vercel.app/api?type=waving&color=00ADD8&height=200&section=header&text=GitHub%20User%20Activity%20CLI&fontSize=40&fontColor=fff&animation=twinkling&fontAlignY=40&desc=Go%20|%20CLI%20|%20GitHub%20API%20|%20Activity%20Tracker&descAlignY=60&descSize=18">

<p align="center">
  <i>🔍 A simple and efficient command-line application to fetch and display GitHub user activity built with Go.</i>
</p>

<p align="center">
  <i>🔍 Uma aplicação de linha de comando simples e eficiente para buscar e exibir atividades de usuários do GitHub construída com Go.</i>
</p>

---

### 📚 Practice Project | Projeto de Prática

<div align="center">

**Language:** Go (Golang) | **Linguagem:** Go (Golang)  
**Objective:** Learn Go fundamentals by consuming GitHub API | **Objetivo:** Aprender os fundamentos de Go consumindo a API do GitHub

</div>

### 🌟 Features | Funcionalidades

<div align="center">

| Feature | Description | Descrição |
|:---:|:---|:---|
| 👤 | User Activity | Atividade do Usuário |
| 🔄 | Push Events Tracking | Rastreamento de Push Events |
| 📅 | Formatted Timestamps | Timestamps Formatados |
| 📊 | Commit Count | Contagem de Commits |
| 🏪 | Repository Information | Informações do Repositório |
| ⚡ | GitHub API Integration | Integração com API do GitHub |
| 🌐 | HTTP Requests | Requisições HTTP |
| 📦 | JSON Data Processing | Processamento de Dados JSON |

</div>

### 🛠️ Technologies | Tecnologias

<div align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go,github&theme=dark" />
  </a>
</div>

### 🏗️ Architecture | Arquitetura

```
┌─────────────────────────────┐
│   GitHub User Activity CLI  │
│   (Command Line Interface)  │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│   GitHub API Endpoint       │
│   (api.github.com/users)    │
└─────────────────────────────┘
```

### 🔍 What it does | O que faz

Este aplicativo busca e exibe as atividades recentes de push de qualquer usuário do GitHub, mostrando:
- **Repository name** | **Nome do repositório**
- **Number of commits** | **Número de commits**
- **Timestamp** formatted in Brazilian format | **Timestamp** formatado no padrão brasileiro

### 🚀 Getting Started | Começando

#### Prerequisites | Pré-requisitos

- Go 1.25.1 or higher installed | Go 1.25.1 ou superior instalado
- Internet connection to access GitHub API | Conexão com a internet para acessar a API do GitHub
- Basic command line knowledge | Conhecimento básico de linha de comando

#### Installation | Instalação

```bash
# Clone the repository | Clone o repositório
git clone https://github.com/matheussricardoo/GitHubUserActivityCLI.git

# Navigate to project directory | Navegue até o diretório do projeto
cd GitHubUserActivityCLI

# Build the application | Compile a aplicação
go build -o github-activity.exe main.go

# Or run directly | Ou execute diretamente
go run main.go <username>
```

#### Usage | Uso

```bash
# Run with a GitHub username | Execute com um nome de usuário do GitHub
go run main.go matheussricardoo

# Or using the compiled version | Ou usando a versão compilada
./github-activity.exe matheussricardoo
```

#### Example Output | Exemplo de Saída

```
Looking for activities for: matheussricardoo
Connection successful!
Status: 200 OK

--- Atividade Recente de matheussricardoo ---

[02/10/2025 às 14:30] - Pushed 3 commits to matheussricardoo/GitHubUserActivityCLI
[01/10/2025 às 09:15] - Pushed 5 commits to matheussricardoo/TaskTracker
[30/09/2025 às 16:45] - Pushed 5 commits to matheussricardoo/CommitToLearn
```

### 📁 Project Structure | Estrutura do Projeto

```
GitHubUserActivityCLI/
├── main.go              # Main application file with GitHub API integration
├── go.mod              # Go module definition
├── LICENSE             # MIT License
└── README.md           # Project documentation
```

### 📦 GitHub API Response Structure | Estrutura da Resposta da API do GitHub

```json
{
  "type": "PushEvent",
  "repo": {
    "name": "username/repository"
  },
  "payload": {
    "size": 3
  },
  "created_at": "2025-10-02T14:30:00Z"
}
```

### 💡 Key Features Implementation | Implementação dos Recursos Principais

- **HTTP Client:** Makes requests to GitHub API | **Cliente HTTP:** Faz requisições para a API do GitHub
- **JSON Processing:** Parses GitHub API responses | **Processamento JSON:** Analisa respostas da API do GitHub
- **Time Formatting:** Converts UTC to Brazilian format | **Formatação de Tempo:** Converte UTC para formato brasileiro
- **Error Handling:** Handles network and API errors gracefully | **Tratamento de Erros:** Trata erros de rede e API adequadamente
- **Push Event Filtering:** Shows only push events | **Filtragem de Push Events:** Mostra apenas eventos de push

### 🎯 Learning Objectives | Objetivos de Aprendizado

- [x] HTTP requests in Go | Requisições HTTP em Go
- [x] JSON unmarshaling with interface{} | Desserialização JSON com interface{}
- [x] Time parsing and formatting | Análise e formatação de tempo
- [x] Command-line arguments | Argumentos de linha de comando
- [x] Error handling with defer | Tratamento de erros com defer
- [x] GitHub API integration | Integração com API do GitHub
- [x] Type assertions in Go | Asserções de tipo em Go

### 📋 Usage Commands | Comandos de Uso

| Command | Description | Descrição |
|:---|:---|:---|
| `go run main.go <username>` | Fetch user activity | Buscar atividade do usuário |
| `go build -o app.exe main.go` | Build executable | Compilar executável |
| `./app.exe <username>` | Run compiled version | Executar versão compilada |

### 👤 Author | Autor

<div align="center">
  <a href="https://github.com/matheussricardoo" target="_blank">
    <img src="https://skillicons.dev/icons?i=github" alt="GitHub"/>
  </a>
  <a href="https://www.linkedin.com/in/matheus-ricardo-426452266/" target="_blank">
    <img src="https://skillicons.dev/icons?i=linkedin" alt="LinkedIn"/>
  </a>
</div>

### 📜 License | Licença

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

<img width="100%" src="https://capsule-render.vercel.app/api?type=waving&color=00ADD8&height=120&section=footer"/>

</div>
