# GearBarter

### Exchange your gear for other gear

As of right now, this project is mainly to develop and demonstrate my development skills.  This project is the backen API and will receive a mobile front end in the near future (Android) as well as a web front end later on (React)

## How to install on your machine 
### Current Environment

**Postgres** 	9.6

**Go** 		1.9

### Installation Process
Install go 1.9+

Install postgres version 9.6

Create database credentials file
```
{
    "username": "briand787b",
    "password": "password123",
    "database": "gearbarter"
}
```
This file is not checked into source control on purpose.  I am providing it here so that the structure is defined with dummy data.  The `database` field holds the name of the database

With go installed, you can use it to create the certificates for ssl by running this command in the auth directory
```
go run /usr/local/go/src/crypto/tls/generte_cert.go -host=localhost
```

You can generate the rsa files for jwt authentication by running these commands in the auth directory as well.
```
openssl genrsa -out app.rsa <keysize>
openssl rsa -in app.rsa -pubout > app.rsa.pub
```

