### Run Application
``` shell
go run main.go
```

### Test Application
``` shell
curl -X GET http://localhost:8888/go-demo
curl -X GET http://localhost:8888/go-demo/info
curl -X GET http://localhost:8888/go-demo/health
curl -X GET http://localhost:8888/go-demo/version
```

### Build Image Application
``` shell
docker build -t go-front:1.0.0 . --no-cache --platform linux/amd64
```

### Docker Compose Up
``` shell
docker-compose up
[+] Running 3/3
 ⠿ Container go-front-frontend-2  Recreated                                                                                                        0.1s
 ⠿ Container go-front-frontend-3  Recreated                                                                                                        0.1s
 ⠿ Container go-front-frontend-1  Recreated                                                                                                        0.2s
Attaching to go-front-frontend-1, go-front-frontend-2, go-front-frontend-3
go-front-frontend-1  | time="2023-04-02 21:04:24.964" level=info msg="********** Go Front Started **********"
go-front-frontend-3  | time="2023-04-02 21:04:25.170" level=info msg="********** Go Front Started **********"
go-front-frontend-2  | time="2023-04-02 21:04:25.378" level=info msg="********** Go Front Started **********"
go-front-frontend-1  | time="2023-04-02 21:04:48.241" level=info msg=" | 2.0.0 | 03cf1dba75f0 | 2023.04.02 21:04:48.241120 | "
go-front-frontend-3  | time="2023-04-02 21:04:52.327" level=info msg=" | 2.0.0 | 49ae4c0dff91 | 2023.04.02 21:04:52.327695 | "
go-front-frontend-2  | time="2023-04-02 21:04:55.235" level=info msg=" | 2.0.0 | e4ae3456ee20 | 2023.04.02 21:04:55.235740 | "
```

``` shell
docker ps
CONTAINER ID   IMAGE            COMMAND   CREATED          STATUS          PORTS                     NAMES
03cf1dba75f0   go-front:1.0.0   "./app"   45 seconds ago   Up 44 seconds   0.0.0.0:50792->8888/tcp   go-front-frontend-1
49ae4c0dff91   go-front:1.0.0   "./app"   45 seconds ago   Up 44 seconds   0.0.0.0:50793->8888/tcp   go-front-frontend-3
e4ae3456ee20   go-front:1.0.0   "./app"   45 seconds ago   Up 44 seconds   0.0.0.0:50794->8888/tcp   go-front-frontend-2
```

### Apply Deployment
``` shell
kubectl apply -f deployment.yaml
```

### Test Loop
``` shell
while true; do curl -X GET http://localhost:8888/go-demo/info; echo "" && sleep 1; done
```