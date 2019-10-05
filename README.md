# Prerequisites
- [Docker](https://docs.docker.com/install/)
- [Go](https://golang.org/doc/install)
- [Dep](https://golang.github.io/dep/docs/installation.html)
- Create Account in [Mail Trap](https://mailtrap.io)

# Steps to Install
### Get source code
`go get github.com/rishinair11/notification-server-go/src`
### Move to source code directory
`cd $GOPATH/src/github.com/rishinair11/notification-server-go`
### Build source code
`make`
### Build docker image
`make package`

# Steps to Run
```
docker run --name notification-server -p 5252:5252 --env USERNAME=<MAILTRAP_USERNAME> --env PASSWORD=<MAILTRAP_PASSWORD> --env HOST=smtp.mailtrap.io --env PORT=2525 --env FROM=<FROM_EMAIL> notification-server-go:latest
```

# Sample Request 

**URL** - http://localhost:5252/mail
**METHOD** - POST
**REQUEST BODY** -
```
{
	"emailID": "test@email.com",
	"subject": "test",
    "body": "hello"
}
```