package rank

import (
	fmt "fmt"
	"io"

	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Client interface {
	RankServiceClient
	io.Closer
}

type client struct {
	RankServiceClient
	*grpc.ClientConn
}

func Connect(gcfg cfg.GRPCServiceConfig) (Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:9200", gcfg.ServiceName()), grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "Connect")
	}
	return client{
		NewRankServiceClient(conn),
		conn,
	}, nil
}

const maxMsgSize = 100 * 1024 * 1024 // 10MB

var (
	MaxSendMsgSize     = grpc.MaxSendMsgSize(maxMsgSize)     // server side
	MaxRecvMsgSize     = grpc.MaxRecvMsgSize(maxMsgSize)     // server side
	MaxCallRecvMsgSize = grpc.MaxCallRecvMsgSize(maxMsgSize) // client side
	MaxCallSendMsgSize = grpc.MaxCallSendMsgSize(maxMsgSize) // client side
)
