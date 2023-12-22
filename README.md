<h1 align="center">

<img src="./assets/CCVP.png" width=250>


CPF-CNPJ Validation Protocol
</h2>

- [Português](./README.md)
- [English](./ENGLISH.md)

CCVP é um protocolo que visa guiar a validação de CPF e CNPJ na Internet.

## Estrutura

### Blocos

O CCVP conta com 7 blocos:

- NOD --> Número do Documento
- TOD --> Tipo do Documento
- STT --> Código
- TOR --> Tipo de Requisição
- TYP --> Tipos Permitidos
- Types Length --> Tamanho da mensagem TYP
- NOD Length --> Tamanho da mensagem NOD

Cada um dos valores dentro do bloco tem seu significado baseado em valores:

- NOD
    - CPF
    - CNPJ
- TOD
    - 0 --> CPF
    - 1 --> CNPJ
- TOR
    - 0 --> Verificar
    - 1 --> Tipos Permitidos

- TYP
    - Tipos Permitidos
- TYP Length
    - Tamanho do Bloco de Tipos
- NOD Length
    - Tamanho do Bloco do Documento

### Tamanho das mensagens

Cada mensagem (Requisição ou Resposta) deve cotner um tamanho correto para permitir uma comunicação correta. Cada tipo de requisição feita pelo cliente pode conter dois tamanhos, a depender do documento requerido (CPF ou CNPJ) e do tipo de requisição (Tipos permitidos ou Validação)

|   Bloco      |  Tamanho   |
|--------------|------------|
|   TOR        |     1      |
|   TOD        |     1      |
|   NOD [CPF]  |     11     |
|   NOD [CNPJ] |     14     |
|   STT        |     3      |
|   NOD Length |     2      |
|   TYP Length |     1      |
|   TYP        |     8      |

### Requisição

    |     |                 |     |                     |
    | TOR |   NOD Length    | TOD |         NOD         |
    |     |                 |     |                     |

### Resposta

    Caso de Requisição de Validação

    |     |                 |                    |
    | STT |   NOD Length    |        NOD         |
    |     |                 |                    |

    Caso de Requisição de Tipos Permitidos

    |     |                 |                    |
    | STT |   TYP Length    |        TYP         |
    |     |                 |                    |

### Exemplo

    Requisição:

    |     |                 |     |                             |
    |  0  |       14        |  1  |         12345678878         |
    |     |                 |     |                             |

    Resposta:

    |     |                 |                            |
    | 301 |       14        |        12345678878         |
    |     |                 |                            |
## Códigos de Status

O STT é um dos blocos vindos da mensagem de resposta, compondo a principal informação dentre a
comunicação feita, pode conter mensagens de aviso e sucesso, informando ao cliente seu estado.

| Código  | Significado |
| ------------- | ------------- |
| 101  | Tipo de documento errado |
| 102  | Requisição errada |     
| 103  | Tamanho errado |     
| 301  | Aceito |
| 302  | Rejeitado |     
| 303  | Recebido |  

## Sevidor e Cliente

Para realizar a exemplificação do protocolo, neste projeto há um servidor e cliente para demonstração do protocolo.

- [Server](./server/README.md)
- [Client](./client/README.md)