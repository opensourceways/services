module github.com/micro/services

go 1.14

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/SlyMarbo/rss v1.0.1
	github.com/davecgh/go-spew v1.1.1
	github.com/getkin/kin-openapi v0.26.0
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/gosimple/slug v1.9.0
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/micro/dev v0.0.0-20201117163752-d3cfc9788dfa
	github.com/micro/micro/v3 v3.0.5-0.20210202220323-c0f344d21cc6
	github.com/miekg/dns v1.1.31 // indirect
	github.com/stoewer/go-strcase v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/ulikunitz/xz v0.5.8 // indirect
	golang.org/x/crypto v0.0.0-20201002094018-c90954cbb977
	golang.org/x/net v0.0.0-20200930145003-4acb6c075d10
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	google.golang.org/genproto v0.0.0-20201001141541-efaab9d3c4f7 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	googlemaps.github.io/maps v1.3.1
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.9
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
