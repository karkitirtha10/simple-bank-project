# docs 

1. get base url in gin framework
    ``` go 
    baseURL := "http://" + context.Request.clientIdHost
    ```
2. golang migrate package
    1. migrate up command
        ``` sh
        migrate -path db/migrations -database "postgresql://pgsuperuser:Admin@1@localhost:5432/simplebank?sslmode=disable" -verbose up
         ```
    2. create migration command
        ``` sh
        migrate create -ext sql -dir pkg/common/db/migration $(name)    
        ```