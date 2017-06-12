package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/grpclog"

	pb "github.com/zhangyongjiang/grpc-go-test/blockchain"
)

var (
	port      = flag.Int("port", 9090, "The server port")
	chainFile = flag.String("chain_info_file", "testdata/chaininfo.json", "A json file containing bc info")
)

type blockChainServer struct {
	savedChaininfo *pb.Chaininfo
}

// GetChaininfo returns the chaininfo.
func (s *blockChainServer) GetChaininfo(ctx context.Context, em *pb.EmptyMsg) (*pb.Chaininfo, error) {
	return s.savedChaininfo, nil
}

func (s *blockChainServer) GetTransaction(ctx context.Context, em *pb.MsgInput) (*pb.Transaction, error) {
	var t = new(pb.Transaction)
	t.HeaderSignature = "xxxxxxxxxxxxx"
	t.Id = em.Data
	return t, nil
}

func (s *blockChainServer) CreateTransaction(ctx context.Context, em *pb.Transaction) (*pb.Transaction, error) {
	return em, nil
}

func (s *blockChainServer) GetUnconfirmedTransactionList(ctx context.Context, em *pb.EmptyMsg) (*pb.TransactionList, error) {
	var txs = new(pb.TransactionList)
	return txs, nil
}

// loadChaininfo loads chain from a JSON file.
func (s *blockChainServer) loadChaininfo(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		grpclog.Fatalf("Failed to load default chaininfo: %v", err)
	}
	if err := json.Unmarshal(file, &s.savedChaininfo); err != nil {
		grpclog.Fatalf("Failed to load default chaininfo: %v", err)
	}
}

func (s *blockChainServer) GetBlockByHash(ctx context.Context, em *pb.MsgInput) (*pb.Block, error) {
	var t = new(pb.Block)
	t.HeaderSignature = "head sig here"
	return t, nil
}

func (s *blockChainServer) GetBlockByHeight(ctx context.Context, em *pb.MsgInput) (*pb.Block, error) {
	var t = new(pb.Block)
	t.HeaderSignature = "head sig here. block by height"
	return t, nil
}

func (s *blockChainServer) GetAddress(ctx context.Context, em *pb.MsgInput) (*pb.Address, error) {
	var t = new(pb.Address)
	t.Id = em.Data
	return t, nil
}

func (s *blockChainServer) GetAddressBalance(ctx context.Context, em *pb.MsgInput) (*pb.Address, error) {
	var t = new(pb.Address)
	t.Id = em.Data
	return t, nil
}

func newServer() *blockChainServer {
	s := new(blockChainServer)
	s.loadChaininfo(*chainFile)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBlockChainServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
