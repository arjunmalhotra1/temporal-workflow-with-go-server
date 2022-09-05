# temporal-workflow-with-go-server

## Bring the Temporal cluster up
  ```
  git clone https://github.com/temporalio/docker-compose.git
  cd docker-compose
  docker-compose up
  ```

## Run the worker
In a separate terminal window.
```
go run worker/main.go
```

## Run the golang web server
In a separate terminal window.
```
go run main.go
```
Hit the endpoint from the browser using
```
localhost:8083/