package price

import (
	fmt "fmt"
	"io"

	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Client interface {
	PriceServiceClient
	io.Closer
}

type client struct {
	PriceServiceClient
	*grpc.ClientConn
}

func Connect(gcfg cfg.GRPCServiceConfig) (Client, error) {
	// gcfg.FullServiceName() TODO: wtf is going on with registrator???
	conn, err := grpc.Dial(fmt.Sprintf("%s:9200", gcfg.ServiceName()), grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "Connect")
	}
	return client{
		NewPriceServiceClient(conn),
		conn,
	}, nil
}

const maxMsgSize = 100 * 1024 * 1024 // 100MB

var (
	MaxSendMsgSize     = grpc.MaxSendMsgSize(maxMsgSize)     // server side
	MaxRecvMsgSize     = grpc.MaxRecvMsgSize(maxMsgSize)     // server side
	MaxCallRecvMsgSize = grpc.MaxCallRecvMsgSize(maxMsgSize) // client side
	MaxCallSendMsgSize = grpc.MaxCallSendMsgSize(maxMsgSize) // client side
)
