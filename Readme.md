
# Documentação do Projeto: Sistema de Leilões

Este projeto implementa um sistema de leilões, onde é possível criar, monitorar e fechar leilões automaticamente, considerando uma duração pré-estabelecida. O sistema interage com o MongoDB para armazenar os dados dos leilões.

## Funcionalidades

- **Criar Leilão**: Permite que um usuário crie um leilão com um produto, categoria, descrição e condição do produto.
- **Fechar Leilão**: O leilão é automaticamente fechado após um tempo configurado (por exemplo, 1 hora).
- **Monitoramento de Leilões Vencidos**: O sistema verifica periodicamente se algum leilão foi vencido e, se necessário, fecha o leilão.

## Como Rodar o Projeto

### Requisitos

- **Go (Golang)** versão 1.18 ou superior
- **MongoDB** (local ou Docker)
- **Variáveis de ambiente** configuradas corretamente (exemplo: `AUCTION_DURATION`)

### 1. Configuração do Ambiente

Primeiro, você deve instalar o Go em sua máquina. Para verificar se o Go está instalado corretamente, use:

```bash
go version
```

Além disso, você precisará do MongoDB. Você pode rodá-lo localmente ou via Docker. Aqui está um exemplo de como rodar o MongoDB usando Docker:

```bash
docker run --name mongodb -d -p 27017:27017 mongo
```

Isso fará o MongoDB rodar na sua máquina local na porta `27017`.

### 2. Clonando o Repositório

Clone o repositório do projeto para sua máquina:

```bash
git clone https://github.com/seu-usuario/auction-go.git
cd auction-go
```

### 3. Configuração das Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto ou defina as variáveis de ambiente manualmente. O arquivo `.env` deve conter pelo menos a variável `AUCTION_DURATION`, que define a duração do leilão em segundos.

Exemplo de arquivo `.env`:

```env
AUCTION_DURATION=3600  # 1 hora
```

Se não definir a variável, o sistema usará o valor padrão de 1 hora.

### 4. Instalando as Dependências

Instale as dependências do projeto executando:

```bash
go mod tidy
```

Isso irá baixar as bibliotecas necessárias para o funcionamento do projeto.

### 5. Rodando o Projeto

Para rodar o projeto, basta executar o comando:

```bash
go run main.go
```

Isso iniciará o servidor e começará a monitorar os leilões.

### 6. Testando o Projeto

Se desejar rodar os testes automatizados para garantir que tudo está funcionando corretamente, use o comando:

```bash
go test ./...
```

Isso executará os testes de integração e valida se as funcionalidades de criação e fechamento de leilões estão funcionando como esperado.

### 7. Arquitetura do Sistema

O sistema é dividido em várias camadas:

- **Camada de Entidade (`auction_entity`)**: Contém os modelos e lógicas de validação dos leilões.
- **Camada de Repositório (`auction_repository`)**: Responsável pela interação com o banco de dados MongoDB.
- **Camada de Serviço**: Contém as funções para gerenciar o ciclo de vida dos leilões, como criação e monitoramento de leilões vencidos.

### 8. Considerações Finais

Este sistema permite a criação de leilões dinâmicos, monitoramento de seu status e fechamento automático com base no tempo configurado. É altamente configurável através de variáveis de ambiente e é de fácil integração com o MongoDB.

