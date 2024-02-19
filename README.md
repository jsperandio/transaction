# Transaction

API REST para gerencimento de conta e transação.

#### Com uso de libs:

* [Echo (Web Framework)](https://github.com/labstack/echo)
* [Validator (struct validation)](https://github.com/go-playground/validator)
* [Koanf (config load)](https://github.com/knadh/koanf)
* [PostgreSQL (Database)](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)


#### Endpoints


| Nome | Path | Method | Content-Type | Descrição |
| ------ | ------ | ------ | ------ | ------ |
| Saúde da app| /health | GET | application/json | Retorna Ok caso a app esteja saudável. |
| Swagger| /swagger | GET | application/json | Documentação Swaager. |
| Criar conta| /accounts | POST | application/json | Cria uma nova conta. |
| Buscar conta| /accounts/:accountId | GET | application/json | Retorna dados de uma conta. |
| Criar Transação | /transactions | POST | application/json | Insere uma transação na base de dados de acordo com o layout proposto. |

## Testes

### Executar test
```bash
$ make test
```

## Notas

Neste projeto foi usado gerador de mocks, a ferramenta: 
* [mockery (Mocking lib)](https://github.com/vektra/mockery)

Ferramenta de patch dinamico
* [monkey (monkeypatching)](https://github.com/bouk/monkey)