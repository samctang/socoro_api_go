# socoro_api
### start api
    go run main.go
### generate swagger spec
    swagger generate spec -o ./docs/swagger.json
### serve swagger
    swagger serve -F swagger ./docs/swagger.json