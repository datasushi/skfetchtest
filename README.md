# Reproduction of Behaviour

In order to run this start the API with:
```Bash
go run api.go
```
Next run:
```Bash
npm run dev
```
Next visit localhost:3000

As a result the terminal in which go was run should show the following dumps:
```Bash
PUT / HTTP/1.1
Host: localhost:8080
Connection: close
Accept: */*
Accept-Encoding: gzip,deflate,br
Connection: close
Content-Length: 18
Content-Type: text/plain;charset=UTF-8
User-Agent: node-fetch

{"payload":"none"}
OPTIONS / HTTP/1.1
Host: localhost:8080
Accept: */*
Accept-Encoding: gzip, deflate, br
Accept-Language: en
Access-Control-Request-Method: PUT
Connection: keep-alive
Origin: http://localhost:3000
Referer: http://localhost:3000/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36


PUT / HTTP/1.1
Host: localhost:8080
Accept: */*
Accept-Encoding: gzip, deflate, br
Accept-Language: en
Connection: keep-alive
Content-Length: 18
Content-Type: text/plain;charset=UTF-8
Origin: http://localhost:3000
Referer: http://localhost:3000/
Sec-Ch-Ua: "Chromium";v="92", " Not A;Brand";v="99", "Google Chrome";v="92"
Sec-Ch-Ua-Mobile: ?0
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36

{"payload":"none"}
```
