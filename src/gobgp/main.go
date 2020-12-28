package main

import (
	"context"
	"fmt"
	"time"

	api "github.com/osrg/gobgp/api"
	gobgp "github.com/osrg/gobgp/pkg/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	s := gobgp.NewBgpServer(gobgp.GrpcListenAddress(":50051"), gobgp.GrpcOption(grpcOpts))
	go s.Serve()

	// global configuration
	if err := s.StartBgp(context.Background(), &api.StartBgpRequest{
		Global: &api.Global{
			As:         65000,
			RouterId:   "192.168.43.108",
			ListenPort: 179, // gobgp won't listen on tcp:179
		},
	}); err != nil {
		log.Fatal(err)
	}

	// add peer
	addpeer("192.168.43.143", s, 65000)
	d1 := &api.DefinedSet{
		DefinedType: api.DefinedType_PREFIX,
		Name:        "d1",
		Prefixes: []*api.Prefix{
			&api.Prefix{
				IpPrefix:      "20.1.0.0/24",
				MaskLengthMax: 24,
			},
		},
	}
	d2 := &api.DefinedSet{
		DefinedType: api.DefinedType_AS_PATH,
		Name:        "d2",
		List:        []string{"^65000"},
	}
	err := s.AddDefinedSet(context.Background(), &api.AddDefinedSetRequest{DefinedSet: d1})
	s.AddDefinedSet(context.Background(), &api.AddDefinedSetRequest{DefinedSet: d2})
	s1 := &api.Statement{
		Name: "s1",
		Conditions: &api.Conditions{
			PrefixSet: &api.MatchSet{
				Name:      "d1",
				MatchType: api.MatchType_INVERT,
			},
			AsPathSet: &api.MatchSet{
				Name:      "d2",
				MatchType: api.MatchType_INVERT,
			},
		},
		Actions: &api.Actions{
			RouteAction: api.RouteAction_REJECT,
		},
	}
	p2 := &api.Policy{
		Name:       "p2",
		Statements: []*api.Statement{s1},
	}
	err = s.AddPolicy(context.Background(), &api.AddPolicyRequest{Policy: p2})
	err = s.AddPolicyAssignment(context.Background(), &api.AddPolicyAssignmentRequest{
		Assignment: &api.PolicyAssignment{
			Name:      "global",
			Direction: api.PolicyDirection_IMPORT,
			Policies:  []*api.Policy{p2},
		},
	})
	if err = s.MonitorTable(context.Background(), &api.MonitorTableRequest{}, func(p *api.Path) {
		fmt.Printf("Received update from %v", p)
		rsp, _ := s.GetTable(context.Background(), &api.GetTableRequest{
			TableType: api.TableType_GLOBAL,
		})
		fmt.Printf("****** table := %v", rsp)
	}); err != nil {
		fmt.Printf("Unable to monitor BGP Path, error: %v", err)
	}
	//s.GetTable(context.Background(), api.AddTabl)

	time.Sleep(time.Minute * 300)
}

func addpeer(ip string, s *gobgp.BgpServer, as uint32) {
	n := &api.Peer{
		Conf: &api.PeerConf{
			NeighborAddress: ip,
			PeerAs:          as,
		},
	}
	if err := s.AddPeer(context.Background(), &api.AddPeerRequest{
		Peer: n,
	}); err != nil {
		log.Fatal(err)
	}
}
