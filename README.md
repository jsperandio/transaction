# Transaction


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