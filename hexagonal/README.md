# Arquitetura Hexagonal

Trata-se de um modelo de 

## Adapters

A camada de adaptadores é composta por duas partes, `Input` e `Output`.

### Input 

A camada de input, concentra todas as formas de **entrada** de dados para a aplicação constituindo assim, um adaptador de entrada.

- Controllers
- Consumers/Listeners
- Models/DTOs
    - Request
    - Response
- Mappers e conversores 

### Output

A camada de output, concentra todas as formas de **saída** de dados para a aplicação constituindo assim, um adaptador de saída. 

- REST Clients (Ex. saída para consumir dados do ViaCEP)
- Producers
- Repositories
- Models/DTOs
    - Request
    - Response
- Mappers e conversores 

## Application

A camanda Application concentra todas as regras de negócio da aplicação em serviços (services), além do domínio (Domain) que é onde temos todas as informações a serem tratadas pela aplicação.

- Services
- Ports (Use Cases)
- Domain


## Configurations

Concentra as configurações referente a aplicação, por exemplo, variáveis de ambiente, secrets, portas HTTP, etc.


## Domínio

Esta camada contém a loǵica da aplicação, não deve depender de nenhuma outra camada a não ser dela mesma. Responsável por definir os modelos e estruturas de dados que representam as entidades e conceitos de negócio. 

Deve ser fechada e agnóstica a qualquer tipo de framework. Qualquer necessidade específica para atender requests, responses ou configurações de banco de dados *(tags json, bson, xml, sql, etc)* devem ser criadas em outro objeto a partir do domínio.


#### Domínio UserDomain

Abaixo temos um exemplo de um domínio `UserDomain`.

```go
type UserDomain struct {
    Id Int
    Name string
    Email string
    Password string
}
```

#### Request UserRequest

Abaixo temos o exemplo de um request `UserRequest`, que trata-se de um objeto gerado a partir do domínio `UserDomain` acima para atender a requisição para registro de dados de um usuário.

```go
type UserRequest struct {
    Name string     `json:"name" binding:"required,min=4,max=100"`
    Email string    `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8,containsany=!@#$%*"`
}
```

O propósito deste isolamento é garantir que caso seja necessária alguma alteração *(Exemplo, mudança de banco de dados de SQL para NoSQL ou de Http para Grpc)*, o domínio não seja afetado, uma vez que ele não depende de nenhuma outra camada e representa a verdade e essencia da aplicação.


## Tome couidado, use com moderação

A arquitetura hexagonal pode ser complexa, difícil de implementar e manter, portanto, use com moderação. Para projetos menores que não vão sofrer alterações e evoluções com frequência, pode ser melhor usar uma arquitetura mais simples e rápida.
