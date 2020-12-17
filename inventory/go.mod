module inventory

go 1.15

require (
	github.com/bradfitz/slice v0.0.0-20180809154707-2b758aa73013
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/micro/v3 v3.0.0
	github.com/stretchr/testify v1.6.1
	go4.org v0.0.0-20201209231011-d4a079459e60 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.8
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
