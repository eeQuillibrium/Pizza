package grpckitchen

import (
	"context"
	"log"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
	"google.golang.org/grpc"
)

type Kitchen struct {
	gRPCClient nikita_kitchen1.KitchenClient
	conn       *grpc.ClientConn
	port       int
}

func New(
	port int,
	conn *grpc.ClientConn,
) *Kitchen {
	gRPCClient := nikita_kitchen1.NewKitchenClient(conn)
	return &Kitchen{
		gRPCClient,
		conn,
		port,
	}
}

func (k *Kitchen) SendOrder(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Print("try to proceeds req with kitchen...")
	_, err := k.gRPCClient.SendOrder(ctx, in)
	if err != nil {
		log.Fatalf("sendorder grpc req err: %v", err)
	}
	log.Print("empty order was gotten!")
	return &nikita_kitchen1.EmptyOrderResp{}, nil
}

func (k *Kitchen) Stop() {
	k.conn.Close()
}
