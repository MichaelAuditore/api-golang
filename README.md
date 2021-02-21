# My Api Golang

**!!! Importante !!!**

mi variable $GOPATH tiene este valor **$HOME/work**

mi variable $GOROOT tienes este valor **$HOME/work/src** -> ruta donde se encuentra mi proyecto actualmente

GO version 1.16

Una API de Go extremadamente fácil con seis endpoints:

* '/' **GET** el endpoint base da un mensaje de bienvenida a la API
* '/tickets' **GET** endpoint para obtener todos los tickets creados en la API.
* '/tickets/{id}' **GET** endpoint para obtener un ticket en especifico basandose en un parametro **id** que corresponde al id del ticket a buscar.
* '/tickets' **POST** endpoint para crear un ticket
* '/tickets' **PUT** endpoint para actualizar un ticket
* '/tickets/{id}' **DELETE** endpoint para eliminar un ticket

Cada ticket creado se guardara en una colección de MongoDB. Para trabajar en este proyecto necesite las siguientes herramientas:


* Golang y sus dependencias
* Mongo DB como la base de datos
* Mongo Express (optional) para my interfaz de administracion de la DB

Esta aplicación contiene un archivo **dockerfile** por lo que es posible generar pruebas locales a traves de docker.

Compose es una herramienta para definir y ejecutar aplicaciones Docker de varios contenedores. Con Compose, usa un archivo YAML para configurar los servicios de su aplicación. Luego, con un solo comando, crea e inicia todos los servicios desde su configuración.

# [docker-compose.yml](docker-compose.yml)

* servicio go-app

    A continuación, definiremos los servicios necesarios para nuestra aplicación. Comenzamos con el servicio go-app que acabamos de crear arriba, ¿recuerdas ese Dockerfile?
    El servicio go-app usa una compilación de imágenes del Dockerfile dentro del directorio go-app.
    Le damos un nombre personalizado en nuestro caso go_app.
    Luego le indicamos que el servicio depende de otro servicio llamado mongo que será nuestro servicio de base de datos.
    Finalmente pasamos algunas variables de entorno para usar dentro del contenedor. En nuestro caso PORT.
    Por último, vinculamos el contenedor y la máquina host al puerto expuesto 8000.
    Finalmente definimos un volumen que mapea nuestro directorio fuente al directorio de trabajo dentro del contenedor, de esa manera los cambios en nuestro código fuente desde nuestra máquina host se reflejarán dentro del contenedor.

* Servicio Mongo

    Primero definimos la imagen que se utilizará en nuestro servicio mongo, en nuestro caso una imagen oficial con la etiqueta bionic.
    Luego le damos un nombre personalizado mongo que es el predeterminado, podría ser cualquier cosa.
    Luego vinculamos el contenedor y la máquina host al puerto expuesto 27017.

    Luego definimos un volumen que mapea un directorio local mongodata al directorio de datos dentro del contenedor, de esa manera tendremos almacenamiento persistente incluso si nuestro contenedor falla.

    Por último, agregamos un reinicio: bandera siempre para reiniciar siempre el contenedor si se detiene.

* Mongo Express (opcional)

    Una interfaz de administración de MongoDB basada en la web que permitirá administrar nuestra base de datos de mongo en el navegador.

    Primero definimos la imagen, en nuestro caso la imagen oficial de mongo-express de docker hub.

    Le damos un nombre de contenedor personalizado

    Agregamos un indicador depend_on para que el contenedor solo se ejecute si se está ejecutando el servicio mongo.

    Luego vinculamos el contenedor y la máquina host al puerto expuesto 8081.

    Por último, agregamos un reinicio: bandera siempre para reiniciar siempre el contenedor si se detiene.

# [DockerFile](go-app/Dockerfile)
He comentado bien el Dockerfile pero explicaré algunos bits.

Un Dockerfile es un documento de texto que contiene todos los comandos
el usuario puede llamar a la línea de comando para ensamblar una imagen.

Comenzamos tirando de la imagen base go. La etiqueta alpine indica la versión alpine que es un poco liviana, aunque no es del todo óptima para la producción, pero debería funcionar bien en el desarrollo, no voy a optimizar mucho por el momento.
Luego instale git, que es una dependencia para obtener algunos paquetes.
Luego configuramos el directorio de trabajo dentro del contenedor; en mi caso, es / opt / go-app donde vivirán nuestros archivos dentro del contenedor.
Luego copiamos la fuente del directorio actual al directorio de trabajo establecido anteriormente. Nuestro código fuente también contiene go.mod y go.sum, que son archivos de dependencia
Después de eso, obtenemos las dependencias del servicio de la aplicación usando go mod download
Necesitaremos Live Reload para nuestra aplicación, por lo que instalamos air https://github.com/cosmtrek/air, que es una utilidad de recarga en vivo para las aplicaciones Go.
Finalmente pasamos un ENTRYPOINT que define el comando que se ejecutará cuando se inicie el contenedor, en nuestro caso es la utilidad de recarga Live.

# Pasos para la ejecución

Ahora que sabemos que instrucciones se han especificado en el archivo **[docker-compose.yml](docker-compose.yml)** ejecutaremos el siguiente comando:

1. Clona el repositorio.
2. Ingresa al directorio del proyecto.
2. Ejecuta el comando **docker-compose up** para construir y ejecutar el servicio docker y probar localmente.

Cuando puedas visualizar en consola la siguiente información:

![img](https://res.cloudinary.com/practicaldev/image/fetch/s--m1aOPDqO--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/0st1ju76w1xn6at05f7k.png)

Ya puedes acceder a la aplicación usando las siguientes URLs:

Acceder a la API usando la siguiente url **localhost:8000**
Acceder al administrador de DB MongoDB  **localhost:8081**

Para hacer peticiones a la API debes realizarlo mediante Postman o Insomnia.

# Nota:

Si no has podido realizar los pasos anteriores. puedes probar localmente esta aplicación sin docker.

## Requisitos
* Necesitas tener instalado MongoDB en tu PC.
* Necesitas tener instalado Go v1.16 en tu PC.

## Post-requisitos
* Ejecute el script run.sh

URI http://localhost:8000

# Ejemplos de Requests

## GetAllTickets GET

![GetAllTickets](https://imgur.com/XCAZJRW)

## CreateTicket POST

![CreateTicket](https://imgur.com/Fvl9kk4)

## UpdateTicket PUT

![UpdateTicket](https://imgur.com/Gye0YNK)

## DeleteTicket DELETE

![DeleteTicket](https://imgur.com/ua86EnH)

## GetTicketById GET

![getTicketById](https://imgur.com/V7Oeu9p)