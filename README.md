#  :point_right: CEHDUN :wink:
Es Conjunto estandarizado de herramientas de desarrollo para la union e interconexion de Neurones
Ademas es un componente del proyecto MerakiVE, que tiene como fin generar micro servicios para interconectarse con el CORE y sus principales funciones son:


- **Generar Micro Servicios.**
- **Tramites (opcional).**
- **Gestionar AGENTES.**
- **Gestionar RUTINAS.**
- **Gestionar SUB RUTINAS.**


## Caracteristicas:

Construccion   :no_entry_sign: :construction:  :muscle:


## Tecnologias:
 
- Back-end:
  - [Go](https://golang.org/)
  - [OAuth2](https://oauth.net/2/)

- Front-end:
  - [React](https://facebook.github.io/react/) o [Angular](https://angularjs.org/)
  - [Materialize](http://materializecss.com/)


## Diagramas

Comportamiento del CEHDUN:

![Image of CEHDUN](https://github.com/merakive/cehdun/blob/master/diagrams/cehdun.png)




## Build the enviroment develepment

**Install with Package Management For Go**

Install glide:

```
$ curl https://glide.sh/get | sh
```

**Install dependencies**

Execute in the terminal:

```
cd $GOPATH/src/github.com/merakiVE/CEHDUN
```

Execute glide

```
glide install
```

You can start to develop

Run the server:

```
go run main.go
```

**Install depends of mode global**

Execute in the terminal:

```
$ go get -u github.com/kataras/iris
$ go get -u github.com/merakiVE/CVDI
$ go get -u github.com/hostelix/aranGO
$ go get -u github.com/iris-contrib/middleware/cors
$ go get -u github.com/fatih/structs
$ go get -u github.com/fatih/structtag
$ go get -u github.com/go-playground/locales/en
$ go get -u github.com/go-playground/universal-translator
$ go get -u github.com/spf13/viper
$ go get -u gopkg.in/go-playground/validator.v9
$ go get -u gopkg.in/jeevatkm/go-model.v1
```

**Install ArangoDB**

The next example is how install arango in debian 8

First add the repository key to apt like this:

```
curl -O https://download.arangodb.com/arangodb32/Debian_8.0/Release.key
sudo apt-key add - < Release.key
```

Use apt-get to install arangodb:

```
echo 'deb https://download.arangodb.com/arangodb32/Debian_8.0/ /' | sudo tee /etc/apt/sources.list.d/arangodb.list
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install arangodb3=3.2.4
```


**Create the file cehdun.conf**

Create the file that content the arangoDB config, in the root the project with the following content:

```
{
  "DATABASE" : {
    "DB_HOST" : "http://127.0.0.1",
    "DB_NAME" : "meraki",
    "DB_USER" : "your_user",
    "DB_PASSWORD" : "your_password",
    "DB_PORT": "8529"
  },
  "PATH_KEYS": "keys",
  "PUBLIC_KEY_PATH" : "",
  "PRIVATE_KEY_PATH" : ""
}
```

**To finish execute:**

```
$ go run main.go
```



## ¿Cómo puedo contribuir? 
Solo debes leer el archivo `contributing.md`, que encontraras en [este enlace](https://github.com/merakive/cehdun/blob/master/.github/CONTRIBUTING.md)

