# Reproduction of Behaviour

In order to run this start the API with:
```Bash
go run api.go
```
Next run:
```Bash
npm run dev
```
Next visit 
```Bash
localhost:3000/?token=jkaskjhsahkjsdkhjsadhjksad
```

As a result the terminal in which go was run should show the following dumps:
```Bash
PUT /test HTTP/1.1
Host: localhost:8080
Connection: close
Accept: */*
Accept-Encoding: gzip,deflate,br
Connection: close
Content-Length: 38
Content-Type: text/plain;charset=UTF-8
Cookie: userid=73523aab-281c-431c-a271-7e71120628a6
User-Agent: node-fetch

{"token":"jkaskjhsahkjsdkhjsadhjksad"}
```
