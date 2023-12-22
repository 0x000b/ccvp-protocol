# Servidor CCVP

- [Português](./README.md)
- [English](./ENGLISH.md)

O servidor CCVP é escrito em Golang e visa exemplificar os conceitos empregados pelo protocolo CCVP.

## Instalação

### Pré-requisitos

Para instalação e `building` do projeto, há a necessidade de instalar a linguagem [Go](https://go.dev/) em seu ambiente.

### Construção

Tendo os pré-requisitos, `clone` este projeto em sua área de trabalho e execute `cd` para a pasta do projeto e do servidor.

Dentro da pasta raiz do servidor, `execute`
```bash
go build cmd/server/main.go
```

Com o build completo, simplesmete execute o binário:

```bash
./main
```
