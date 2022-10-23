## TESTING DATVIETVAC
### Structure
```
├── Dockerfile
├── README.md
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── main_test.go
├── pkg
│   ├── cloud
│   │   ├── aws.go
│   │   ├── client.go
│   │   └── s3.go
│   └── utils
│       └── utils.go
├── src
│   ├── api
│   │   └── server.go
│   ├── cors
│   │   └── cors.go
│   ├── handlers
│   │   └── upload.handler.go
│   └── middleware
│       └── middleware.go
├── tmp
│   ├── cache
│   │   ├── machine.json
│   │   ├── server.test.pem
│   │   ├── server.test.pem.crt
│   │   ├── server.test.pem.key
│   │   └── service-catalog-1_2_1_dev-1_27_91.pickle
│   ├── lib
│   ├── logs
│   │   ├── localstack_infra.err
│   │   └── localstack_infra.log
│   ├── state
│   │   └── startup_info.json
│   └── tmp
└── upload
```
### How to run
1. Clone repo
```sh
   git clone 
```
2. Run with docker-compose
```sh
    cd testdatvietvac
    docker-compose up
```
3. Service run on

| Service | URL |
|:-------:|:----:|
| API      | http://localhost:8080 |
| S3 local | http://localhost:4566 |

### API Description

- Api upload data : 
    - /user/batch
    - Method: POST
    - Params : data(json)

### Contact Me:

- email: ngocnghia128@gmail.com - Nghia Nguyen :fire: :fire: :fire:# testingdatvietvac

