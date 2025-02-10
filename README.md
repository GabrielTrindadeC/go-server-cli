# Projeto de Linha de Comando: Go Server CLI

## Descrição

Projeto de linha de comando para trazer algumas informações em domínios.

Projeto com o intuito de estudar a stdlib do Go e por em pratica alguns conceitos da linguagem.

## Instalação

```sh
# Clone o repositório
git clone https://github.com/GabrielTrindadeC/go-server-cli

# Navegue até o diretório do projeto
cd go-server-cli

# Instale as dependências
go mod tidy

# Compile o projeto
go build go-server-cli
```

# Uso

## Sintaxe básica:

```sh
./go-server-cli [comando] --[flag] [valor]
```

Comandos:

### ip - Busca IP de um host

```sh
./go-server-cli ip --host https://www.google.com
```

### server - Mostra nome do servidor

```sh
./go-server-cli server --host https://www.github.com
```

### ping - Teste de ping

```sh
./go-server-cli ping --ip 8.8.8.8
```

### help - Ajuda geral

```sh
./go-server-cli help
```

# Comandos

- `help`: Exibe a ajuda
- `ip`: Busca o IP de um host
- `server`: Busca o nome do servidor
- `ping:`: Faz um ping em um ip
