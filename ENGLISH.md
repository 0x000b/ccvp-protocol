<h1 align="center">

<img src="./assets/CCVP.png" width=250>


CPF-CNPJ Validation Protocol
</h2>

- [Português](./README.md)
- [English](./ENGLISH.md)

CCVP is a protocol that aims to guide CPF and CNPJ validation on the Internet.

## Structure

### Blocks

The CCVP has 7 blocks:

- NOD --> Number of Document
- TOD --> Type of Document
- STT --> Status Code
- TOR --> Type of Request
- TYP --> Allowed Types
- Types Length --> Size of TYP Block
- NOD Length --> Size of NOD Block

Each of the values within the block has its own value-based meaning:

- NOD
    - CPF
    - CNPJ
- TOD
    - 0 --> CPF
    - 1 --> CNPJ
- TOR
    - 0 --> Verify
    - 1 --> Allowed Types

- TYP
    - Allowed Types
- TYP Length
    - Size of TYP Block
- NOD Length
    - Size of NOD Block

### Message Size

Each message (Request or Response) must contain a correct size to allow correct communication. Each type of request made by the client can contain two sizes, depending on the document required (CPF or CNPJ) and the type of request (Allowed types or Validation).

|   Block      |  Size   |
|--------------|------------|
|   TOR        |     1      |
|   TOD        |     1      |
|   NOD [CPF]  |     11     |
|   NOD [CNPJ] |     14     |
|   STT        |     3      |
|   NOD Length |     2      |
|   TYP Length |     1      |
|   TYP        |     8      |

### Request

    |     |                 |     |                     |
    | TOR |   NOD Length    | TOD |         NOD         |
    |     |                 |     |                     |

### Response
    Validation Request Case

    |     |                 |                    |
    | STT |   NOD Length    |        NOD         |
    |     |                 |                    |

    Allowed Types Case

    |     |                 |                    |
    | STT |   TYP Length    |        TYP         |
    |     |                 |                    |

### Example

    Request:

    |     |                 |     |                             |
    |  0  |       14        |  1  |         12345678878         |
    |     |                 |     |                             |

    Response:

    |     |                 |                            |
    | 301 |       14        |        12345678878         |
    |     |                 |                            |
## Status Code

The STT is one of the blocks that comes from the response message communication, it can contain warning and success messages, informing the client of its status.

| Code  | Meaning |
| ------------- | ------------- |
| 101  | Wrong Type of Document |
| 102  | Wrong Request |     
| 103  | Wrong Size |     
| 301  | Accepted |
| 302  | Rejected |     
| 303  | Received |  

## Server and Client 

In order to exemplify the protocol, in this project there is a server and a client to demonstrate the protocol.

- [Server](./server/README.md)
- [Client](./client/README.md)