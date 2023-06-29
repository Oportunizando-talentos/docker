# Docker Compose
## Conceito
Docker-compose é uma aplicação separada para executar aplicações em contêiner de forma mais configurável. Sua configurações são escritas num documento `yaml`, nele pode ser configurado um ou mais contêineres com diferentes configurações, chamados de **Serviços**.
Em cada serviço pode ser configurado:
- imagens;
- Port-Binding;
- rede;
- variáveis de ambiente;
- sequência de inicalização;
- persistência de dados;
- gerenciamento de recursos;
e mais.

A vantagem de usar o docker-compose em vez do docker é facilitar a configuração de aplicações multi-container e suas configurações, com ele podemos fazer deploy de aplicações facilmente com apenas um comando, pois suas configurações ficam no arquivo.

### Exemplo


Este é o docker-compose.yml oficial da ferramenta **Wiki JS** 
~~~ yml
version: "3"
services:
  # serviço do banco de dados
  db:
    image: postgres:11-alpine
    # variáveis de ambiente para configuração do banco
    environment:
      POSTGRES_DB: wiki
      POSTGRES_PASSWORD: wikijsrocks
      POSTGRES_USER: wikijs
    # configuração de reinicialização do serviço
    restart: unless-stopped
    # configuração de persistencia de dados
    volumes:
      - db-data:/var/lib/postgresql/data

  # serviço da aplicação de fato
  wiki:
    image: ghcr.io/requarks/wiki:2
    # limitação de recursos para 0.5 cpus e 
    deploy:
      resources:
        limits:
          cpus: "0.5"
          memory: 200M
    # define a ordem de inicialização dos serviços
    depends_on:
      - db
    environment:
      DB_TYPE: postgres
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: wikijs
      DB_PASS: wikijsrocks
      DB_NAME: wiki
    restart: unless-stopped
    # port-binding da porta 3000 do serviço para 8080 do host
    ports:
      - "8080:3000"

# configuração do volume utilizado pelo banco
volumes:
  db-data:
~~~

## Prática

1. Primeiramente, instale o Docker Compose executando o comando abaixo:
~~~ Batch
sudo curl https://github.com/docker/compose/releases/download/latest/docker-compose-linux-x86_64 -o /usr/bin/docker-compose
~~~
2. Verifique a instalação com:
~~~
docker-compose version
~~~

Dentro da pasta `brazilian-states` há um Docker Compose que inicia um banco PostgreSQL com dados sobre os estados brasileiros. Na pasta também há uma aplicação Go que se conecta a um banco de dados e retorna dados sobre estados.
Seu trabalho é criar um Dockerfile para a aplicação e alterar o arquivo `docker-compose.yaml` para iniciar a aplicação após o banco de dados, e torná-la utilizável.

**Atenção:** A aplicação requer a variável de ambiente DSN para se conectar ao banco de dados e ela deve ter este valor:
~~~
user=postgres password=password dbname=brazilian_states host=localhost port=5432 sslmode=disable
~~~

Para subir o Docker Compose:
~~~
docker-compose up --build
~~~
