## Criando um Dockerfile
### Exemplo Aplicação

1. O código abaixo é de um servidor http escrito em Go, ele está disponível no arquivo `main.go` 
~~~ go
package main

import (
	"fmt"
	"net/http"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Olá, %s!", name)
}

func main() {
	http.HandleFunc("/greeting", greetingHandler)
	fmt.Println("listening on :8000")
	http.ListenAndServe(":8000", nil)
}
~~~

2. Crie um Dockerfile para a aplicação.

3. Compile a imagem com o comando abaixo:
~~~ Batch
docker build -t minha-primeira-aplicacao .
~~~

4. Execute o servidor com o comando abaixo:
~~~ Batch
docker run -p 8000:8000 -t minha-primeira-aplicacao
~~~
5. Abra http://localhost:8000/greeting?name=SeuNome no seu navegador. Você deve ver uma saída como esta:
~~~
Olá, SeuNome!
~~~
