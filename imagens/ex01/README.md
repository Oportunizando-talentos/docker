## Criando um Dockerfile
### Exemplo simples:
1. Crie um arquivo `Dockerfile` com o conteúdo abaixo:
~~~ Dockerfile
FROM ubuntu:22.04

RUN apt update && apt install sl -y

ENTRYPOINT ["/usr/games/sl"]
~~~ 
2. Compile a imagem com o comando:
~~~ Batch
docker build -t minha-primeira-imagem .
~~~
3. Execute a imagem compilada com o comando:
~~~ Batch
docker run -t minha-primeira-imagem
~~~
> **NOTA:** 
> `-t` no comando de build significa *tag* que é um nome de identificação da imagem (forma extendida `--tag`).
> Já no comando de run, `-t` significa terminal e serve para mostrar o terminal de saída do comando entrypoint.
{.is-warning}
