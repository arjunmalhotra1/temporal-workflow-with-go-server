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
![Screen Shot 2022-09-05 at 4 11 53 PM](https://user-images.githubusercontent.com/43081882/188513091-0301bb84-646b-4342-b8e0-e7170f850521.png)




## Run the golang web server
In a separate terminal window.
```
go run main.go
```
![Screen Shot 2022-09-05 at 4 11 45 PM](https://user-images.githubusercontent.com/43081882/188513136-93e95794-7313-441d-b399-94d430622a2b.png)

Hit the endpoint from the browser using
```
localhost:8083/
```
<img width="466" alt="Screen Shot 2022-09-05 at 4 11 40 PM" src="https://user-images.githubusercontent.com/43081882/188513023-7ea51ab7-2529-44b9-922e-f84e357b0b39.png">
