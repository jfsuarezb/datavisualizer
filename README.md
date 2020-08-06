# Data Visualizer

This web app was created with the purpose of passing a developer test. It does, however, showcase my skills in full-stack development and languages like golang and javascript.

Since this app depends on a dgraph instance, in order to run this app, you first have to run the dgraph instance. If you have docker installed, run the following command:

```
docker run --rm -it -p 8001:8001 -p 8081:8081 -p 9080:9080 dgraph/standalone
``` 

Once the dgraph instance is running, you can use the following command to run the app:

```
go run .
```

Before running the app, make sure that you have a .env file that assigns the proper endpoints to the link variable names. Otherwise, the app will not find the links to the endpoints and, thus, it will fail. 

Once the app is running, the front end interface should be accesible with the following link:


http://localhost:8000


