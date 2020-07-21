module github.com/micro/examples

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/gosimple/slug v1.9.0
	github.com/lithammer/shortuuid/v3 v3.0.4
	github.com/micro/go-micro/v2 v2.9.1-0.20200720090451-a3a7434f2cd9
	github.com/micro/micro/v2 v2.9.2-0.20200721144551-2451560dd2c9
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)

replace	google.golang.org/grpc => google.golang.org/grpc v1.26.0
