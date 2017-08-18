 # :boom: Demiourgo 

Demiurgo is a generator of template for CEHDUN .


## Table of Contents

:construction: construction

## Purpose

Generador of the template ```design``` with a series of parameters to can create a API with [GOA](https://github.com/goadesign/goa)


## Use

:construction: construction


### Install goagen:

```
go get -u github.com/goadesign/goa/...
```

### Generate the Demiurgo

Position in:

```
$ cd CEHDUN/core
```

Execute goagen:

```
$ $GOBIN/goagen bootstrap -d github.com/merakiVE/CEHDUN/core/design
```

### Install core

```
$ go install github.com/merakiVE/CEHDUN/core
```

### Run Core

```
$ $GOBIN/core
```

Now in another terminal create the cli to consult:

```
$ cd github.com/merakiVE/CEHDUN/core/tool/cehdun-cli
```

And execute:

```
$ go build
```

And test of resource:

```
$ ./cehdun-cli connect syndesi --payload '{"name": "Victor", "password": "11", "typeDB": "postgresql"}' --dump
```

### Response success:

```
2017/08/18 16:44:30 [INFO] started id=q2epxp2c POST=http://localhost:8080/syndesi/connect
2017/08/18 16:44:30 [INFO] request headers Content-Type=application/json User-Agent=CEHDUN-cli/0
2017/08/18 16:44:30 [INFO] request body={"name":"Victor","password":"11","typeDB":"postgresql"}
2017/08/18 16:44:30 [INFO] completed id=q2epxp2c status=200 time=1.675399ms
2017/08/18 16:44:30 [INFO] response headers Content-Type=application/json Date=Fri, 18 Aug 2017 20:44:30 GMT Content-Length=25
2017/08/18 16:44:30 [INFO] response body={"message":"Success :)"}
```

### Response Error:


Execute:

```
$ ./cehdun-cli connect syndesi --payload '{"name": "Hugo", "password": "11", "typeDB": "postgresql"}' --dump
```

Response:

```
2017/08/18 16:46:11 [INFO] started id=HQBfP4re POST=http://localhost:8080/syndesi/connect
2017/08/18 16:46:11 [INFO] request headers Content-Type=application/json User-Agent=CEHDUN-cli/0
2017/08/18 16:46:11 [INFO] request body={"name":"Hugo","password":"11","typeDB":"postgresql"}
2017/08/18 16:46:11 [INFO] completed id=HQBfP4re status=200 time=1.586726ms
2017/08/18 16:46:11 [INFO] response headers Content-Type=application/json Date=Fri, 18 Aug 2017 20:46:11 GMT Content-Length=31
2017/08/18 16:46:11 [INFO] response body={"message":"Error mi pana :("}
```

You can also try the resource with curl or another tool


