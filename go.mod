module github.com/eloyekunle/world-bank-grpc

go 1.14

require (
	github.com/go-logr/logr v0.1.0
	github.com/golang/protobuf v1.3.5
	github.com/jkkitakita/wbdata-go v0.0.0-20200322142525-c603839aa304
	github.com/spf13/cobra v0.0.7
	google.golang.org/grpc v1.28.0
	k8s.io/klog/v2 v2.0.0-20200327143823-c91ea2cba2c8
)

replace github.com/jkkitakita/wbdata-go => github.com/eloyekunle/wbdata-go v0.0.0-20200330022426-7ae03a254516
