# GoRestAPI

A simple Golang REST API that can do CRUD operations with Gorm, Echo and PostgreSQL

getUser: 
```console
curl localhost:1323/users/{id}
```

getAllUsers: 
```console
curl localhost:1323/users
```

createUser: 
```console
curl -X POST -H 'Content-Type:application/json' -d '{"first_name":"foo","last_name":"bar","age":"24","email":"foobar@foobar.com"}' localhost:1323/users
```

updateUser: 
```console
curl -X PUT -H 'Content-Type:application/json' -d '{"first_name":"foo","last_name":"bar","age":"24","email":"foobar@foobar.com"}' localhost:1323/users/{id}
```

deleteUser: 
```console
curl -X DELETE localhost:1323/users/{id}
```
