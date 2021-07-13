module github.com/julianlee107/user

go 1.15

require (
	github.com/asim/go-micro/v3 v3.5.2
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/micro/v3 v3.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.1.1 // indirect
	gorm.io/gorm v1.21.11 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/julianlee107/user => github.com/julianlee107/user v0.0.0-20210709080748-fa92d0fe43cb
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
