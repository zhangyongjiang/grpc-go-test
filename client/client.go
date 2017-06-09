package main

import (
	"flag"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "github.com/zhangyongjiang/grpc-go-test/blockchain"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	serverAddr         = flag.String("server_addr", "127.0.0.1:9090", "The server address in the format of host:port")
)

// printChaininfo gets the chaininfo.
func printChaininfo(client pb.BlockChainClient, em *pb.EmptyMsg) {
	grpclog.Printf("Getting chaininfo")
	chaininfo, err := client.GetChaininfo(context.Background(), em)
	if err != nil {
		grpclog.Fatalf("%v.GetChaininfo(_) = _, %v: ", client, err)
	}
	grpclog.Println(chaininfo)
}

func printTransaction(client pb.BlockChainClient, em *pb.EmptyMsg) {
	tran, err := client.GetTransaction(context.Background(), em)
	if err != nil {
                grpclog.Fatalf("%v.GetTransaction(_) = _, %v: ", client, err)
        }
        grpclog.Println(tran)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewBlockChainClient(conn)

	printChaininfo(client, &pb.EmptyMsg{})
	printTransaction(client, &pb.EmptyMsg{})

}
