# go-shorted
Source code for https://srtd.me website

You have to create the database and the tables before you can run the server.

To create the database, open your PostgreSQL shell and run the following command:

```CREATE DATABASE shorted;```

Now you need to "use" the database:

```USE shorted;```

To create the tables, run the following command:

```CREATE TABLE urls (url TEXT NOT NULL, short TEXT NOT NULL, created_at TIMESTAMP NOT NULL);```

Now you can run the server:

```go build```  
```./go-shorted```

The server should be running on port 8080.
To make it work as a daemon, you can just put it in ```screen``` or run it as a SystemD service.