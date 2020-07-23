module github.com/farit2000/shippy/shippy-cli-consignment

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/farit2000/shippy/shippy-service-consignment v0.0.0-20200720113700-c59c24fd29ff
	github.com/micro/go-micro/v2 v2.9.1
)
