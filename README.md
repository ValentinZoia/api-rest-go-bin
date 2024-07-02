# Directorio Profesional en Go utilizando gin.

En este tópico aprenderemos a crear un directorio profesional para un proyecto en Go utilizando el framework Gin.

Pero antes introduciré varios conceptos.

## Qué es Gin?

Gin es un framework web escrito en Go (Golang) que se utiliza para construir aplicaciones y servicios web. Es conocido por su rendimiento, simplicidad y facilidad de uso. El componente central del framework Gin es el `gin.Engine`, que maneja las solicitudes HTTP y las rutas dentro de la aplicación.

### Características principales de Gin:

1. **Rendimiento**: Gin es extremadamente rápido debido a su diseño minimalista y eficiente.
2. **Simplicidad**: La API de Gin es intuitiva y fácil de aprender, lo que permite a los desarrolladores crear aplicaciones rápidamente.
3. **Enrutamiento**: Gin proporciona un enrutador de solicitudes HTTP que es sencillo de configurar y manejar.
4. **Middleware**: Permite el uso de middleware para gestionar tareas comunes como la autenticación, el registro y el manejo de errores.
5. **Gestión de errores**: Gin facilita la captura y gestión de errores a lo largo del ciclo de vida de la solicitud.
6. **Soporte para JSON**: Incluye soporte nativo para trabajar con datos JSON, lo que es muy útil para las API RESTful.

### Instalación
```
go get -u github.com/gin-gonic/gin
```

## Introducción a los Handlers

**Handlers**, en el contexto de un servidor web, son funciones que se encargan de manejar las solicitudes HTTP entrantes. En Go, utilizando el framework Gin, los handlers son funciones que responden a rutas específicas. Su propósito principal es procesar las solicitudes HTTP, realizar operaciones necesarias (como acceder a una base de datos o procesar datos), y devolver una respuesta HTTP adecuada.
## Directorio Profesional
```
myapp/
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── handler.go
│   ├── router/
│   │   └── router.go
│   ├── service/
│   │   └── service.go
│   └── model/
│       └── model.go
├── pkg/
│   └── somepackage/
│       └── somepackage.go
├── config/
│   └── config.go
├── scripts/
│   └── setup.sh
├── .env
├── .gitignore
├── go.mod
└── go.sum

```

### Descripción de las carpetas y archivos:

- **cmd/**: Contiene los puntos de entrada de la aplicación. En este caso, tenemos un subdirectorio `myapp` que contiene el archivo `main.go`.
- **internal/**: Contiene el código interno de la aplicación. Este código no debería ser importado por otras aplicaciones.
    - **handler/**: Contiene los manejadores (handlers) de las rutas.
    - **router/**: Define las rutas de la aplicación.
    - **service/**: Contiene la lógica de negocio.
    - **model/**: Contiene las definiciones de los modelos/datos.
- **pkg/**: Contiene código que puede ser utilizado por múltiples proyectos.
- **config/**: Contiene archivos de configuración y utilidades de configuración.
- **scripts/**: Contiene scripts útiles, como scripts de configuración o despliegue.
- **.env**: Archivo de variables de entorno.
- **.gitignore**: Archivo para definir qué archivos/directorios deben ser ignorados por Git.
- **go.mod** y **go.sum**: Archivos de gestión de dependencias de Go.

### Contenido de los archivos:

**main.go (cmd/myapp/main.go)**:

### `main.go`

**Propósito**: Este es el punto de entrada de la aplicación. Aquí es donde se inicializa el servidor web y se define la configuración básica de la aplicación.

**Funciones**:

- **`main`**: Esta función se encarga de iniciar el servidor web. Llama a la función `SetupRouter` del paquete `router` para obtener una instancia del router configurado y luego inicia el servidor en el puerto especificado.

```go
package main

import (
    "myapp/internal/router"
    "log"
)

func main() {
    r := router.SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
```

### **router.go (internal/router/router.go)**:

### `router.go`

**Propósito**: Este archivo define las rutas de la aplicación y asocia cada ruta con su correspondiente handler. Básicamente, configura cómo las solicitudes HTTP serán enrutadas y manejadas.

**Funciones**:

- **`SetupRouter`**: Esta función configura el router de Gin. Define las rutas y asigna los handlers correspondientes a cada una de ellas.

```go
package router

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler.Initial)
	r.GET("/ping", handler.Ping)
	return r
}

```

### **handler.go (internal/handler/handler.go)**:

### `handler.go`

**Propósito**: Contiene las funciones handlers que procesan las solicitudes HTTP para rutas específicas.

**Funciones**:

- **`Ping`**: Este handler responde a la solicitud GET en la ruta `/ping`. Devuelve un mensaje JSON con la respuesta "pong".

```go
package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "pong",
		"message2": "Soy Crack",
	})
}

func Initial(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "Initial",
		"message2": "Mensaje de Bienvenida",
	})
}

```

### **config.go (config/config.go)**:

### `config.go`

**Propósito**: Este archivo maneja la configuración de la aplicación, utilizando la librería Viper para cargar y gestionar las variables de configuración.

**Funciones**:

- **`LoadConfig`**: Esta función utiliza Viper para cargar las variables de configuración desde un archivo `.env`.

**¿Qué es Viper y para qué sirve?**:

- **Viper**: Es una librería en Go diseñada para manejar la configuración de aplicaciones. Viper permite cargar configuraciones desde archivos, variables de entorno, etc., y facilita el acceso a estas configuraciones en la aplicación.
- **Propósito**: Simplifica la gestión de configuraciones y hace que sea fácil cambiar configuraciones sin modificar el código fuente.

```go
package config

import (
    "github.com/spf13/viper"
    "log"
)

func LoadConfig() {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }
}
```

### **.env**:

```makefile
PORT=8080
```

### **.gitignore**:

```bash
/vendor/
/bin/
*.log
.env
```

Con esta estructura, tu proyecto estará bien organizado y será más fácil de mantener y escalar.

### Correr el programa

Primero nos ubicamos en la carpeta raíz e inicializamos un modulo

```
go mod init myapp
```

Luego nos ubicamos en la carpeta donde tenemos el main.go y ejecutamos

```
go run main.go
```

Terminal:

```
[GIN-debug] Listening and serving HTTP on :8080
```

Si nos dirigimos a [http://localhost:8080/](http://localhost:8080/) veremos lo siguiente si todo salio bien.

![Captura de pantalla 2024-07-02 131521.png](Directorio%20Profesional%20en%20Go%20utilizando%20gin%20f6974d0e1ff24bb783264dd4ebeceb7e/7f8d5350-ec9d-49ba-af34-63198b38c442.png)

Y si nos dirigimos a [http://localhost:8080/ping](http://localhost:8080/ping) veremos lo siguiente si todo salio bien.

![Captura de pantalla 2024-07-02 131637.png](Directorio%20Profesional%20en%20Go%20utilizando%20gin%20f6974d0e1ff24bb783264dd4ebeceb7e/Captura_de_pantalla_2024-07-02_131637.png)

Y eso es todo, ya tenemos nuestra api web.

## Qué hago con service.go y model.go?

En un proyecto más grande y complejo, los archivos `service.go` y `model.go` desempeñan roles importantes en la organización del código. A continuación, se describe cómo se pueden utilizar estos archivos y qué tipo de contenido incluir en ellos.

### `service.go`

**Propósito**: Contiene la lógica de negocio de la aplicación. Aquí es donde se implementan las operaciones y procesos principales que la aplicación realiza. Los servicios suelen ser utilizados por los handlers para realizar acciones específicas.

**Ejemplo de contenido**:
Supongamos que estamos desarrollando una API para gestionar usuarios. Podríamos tener un servicio para gestionar las operaciones relacionadas con los usuarios, como crear, actualizar, obtener y eliminar usuarios.

```go
package service

import "myapp/internal/model"

// UserService define las operaciones relacionadas con los usuarios
type UserService struct{}

// NewUserService crea una nueva instancia de UserService
func NewUserService() *UserService {
    return &UserService{}
}

// CreateUser crea un nuevo usuario
func (s *UserService) CreateUser(user model.User) error {
    // Lógica para crear un usuario
    return nil
}

// GetUser obtiene un usuario por su ID
func (s *UserService) GetUser(id int) (*model.User, error) {
    // Lógica para obtener un usuario
    return &model.User{}, nil
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(user model.User) error {
    // Lógica para actualizar un usuario
    return nil
}

// DeleteUser elimina un usuario por su ID
func (s *UserService) DeleteUser(id int) error {
    // Lógica para eliminar un usuario
    return nil
}
```

### `model.go`

**Propósito**: Define las estructuras de datos (modelos) utilizadas en la aplicación. Estos modelos representan entidades del mundo real con las que la aplicación trabaja, como usuarios, productos, pedidos, etc.

**Ejemplo de contenido**:
Continuando con el ejemplo de gestión de usuarios, podríamos tener un modelo de usuario.

```go
package model

// User representa un usuario en la aplicación
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

### Integración de `service.go` y `model.go` en el proyecto

Una vez definidos los servicios y los modelos, los handlers pueden utilizarlos para procesar las solicitudes HTTP. Aquí tienes un ejemplo de cómo podrías modificar el handler para gestionar usuarios utilizando el servicio:

**Modificar `handler.go` para incluir operaciones de usuario**:

```go
goCopiar código
package handler

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/service"
    "myapp/internal/model"
    "net/http"
)

var userService = service.NewUserService()

func CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := userService.CreateUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := userService.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

```

**Modificar `router.go` para incluir rutas de usuario**:

```go
goCopiar código
package router

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/ping", handler.Ping)

    userGroup := r.Group("/users")
    {
        userGroup.POST("/", handler.CreateUser)
        userGroup.GET("/:id", handler.GetUser)
    }

    return r
}

```

### Resumen

- **`service.go`**: Contiene la lógica de negocio. Define servicios para manejar operaciones específicas relacionadas con la funcionalidad de la aplicación.
- **`model.go`**: Define las estructuras de datos (modelos) que representan las entidades con las que trabaja la aplicación.
