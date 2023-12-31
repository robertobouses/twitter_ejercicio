# twitter_ejercicio

Twitter
V1
Desarrollaremos una API dónde un usuario podrá crearse una cuenta, configurar su contraseña y a partir de ahí podrá publicar, ver sus propios tweets y ver también los tweets de las personas que sigue.

Bases de datos
Toda la información tiene que estar guardada en bases de datos postgres y tendremos las siguientes tablas:
    • account: Aquí irá la información de los usuarios, al registrarse, tenemos que generar una entrada en esa tabla.
    • tweet: Aquí irán todos los tweets de todos los usuarios. 
    • account_following: Esta tabla almacenará todas las relaciones de qué usuarios sigue a qué usuarios. Por cada usuario que sigue a otro, tenemos que tener una entrada, por ejemplo:
        ◦ User A sigue a User B
        ◦ User A sigue a User C
        ◦ User C sigue a User A

API
La API tendrá que tener las siguientes rutas:
    1. POST /account: Crear una cuenta nueva, se le pasará el email, nombre y una contraseña. Al guardar el usuario, generaremos un ID para ese nuevo user.
    2. GET /account/id: Devolverá la info del usuario del ID que recibimos en el path param
    3. POST /tweet: Un post que recibirá el tweet y el ID del usuario y lo guardará en bases de datos
    4. GET /tweet: devuelve todos los tweets
    5. POST /:id/follow: Y que reciba por body, un json donde reciba el userID, y lo que indica es que el user ID quiere seguir al userID que viene por el body
    6. GET /:id/timeline: Devuelve todos los tweets de los usuarios que el user ID sigue


Entregables:
    • Las SQL para crear las tablas en postgres
    • El código en un proyecto nuevo de github:
        ◦ Un commit por cada endpoint, y que los mensajes de los commits sean descriptivos (una buena lectura puede ser https://www.conventionalcommits.org/en/v1.0.0/ pero no te obsesiones),  por ejemplo: 
            ▪ Add create account endpoint
            ▪ Fix create account endpoint
Consejos
Hacer cada paso a la vez:
    • Crear el proyecto con el go mod init, crear un main que tenga una API que solo diga “hola”
    • Luego crear los endpoints por orden que está en el ejercicio