# toronto-time-db

In this assignment we have created main.go file in different functions are there which fetch the Toronto city time from External API.
Also, it adds entries into created mysql database table for each run.

For mysql driver, we used 'github.com/go-sql-driver/mysql'

We have also dockerized the Go code file and we created Dockerfile for Go code, Dockerfile.mysql for mysql.
Created docker-compose.yml which contains services for both Database and Go application.
Also created sql script to set up database and table

Docker repo --> docker pull narveersaharan/torontotimedb:v0.2