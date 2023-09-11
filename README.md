![Coverage](https://github.com/OmnGabriel/Api-Crud-Characters/blob/badge/badge.svg)

# Api Crud Characters
-----------------------------------
Como instalar o projeto:

Primeiro passo: Baixe o Docker
para ter acesso ao banco de dados e poder utilizar a aplicação de forma correta, é necessario baixar o docker para poder subir o container apos baixar o docker e ja ter sua branch local do projeto rode os seguintes comandos: 

    docker-compose build 

  após subir o container rode:

    docker-compose up

  Com o container rodando, instale o golang: 
  
  https://go.dev/

  Com o golang baixado, o projeto deve rodar normalmente, teste seu funcionamento com o comando:
		
    GIN_MODE=release go test -v -v -coverpkg=./controllers
	
  Se tudo foi feito corretamente ele ira rodar todos os testes.
