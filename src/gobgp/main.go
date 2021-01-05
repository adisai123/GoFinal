package main

import (
	"context"
	"fmt"
	"strings"
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
	addpeer("192.168.43.186", s, 60000)
	d1 := &api.DefinedSet{
		DefinedType: api.DefinedType_PREFIX,
		Name:        "d1",
		Prefixes: []*api.Prefix{
			&api.Prefix{
				IpPrefix:      "25.25.0.0/16",
				MaskLengthMax: 24,
				MaskLengthMin: 16,
			},
		},
	}
	d2 := &api.DefinedSet{
		DefinedType: api.DefinedType_AS_PATH,
		Name:        "d2",
		List:        []string{"61000"},
	}
	s.AddDefinedSet(context.Background(), &api.AddDefinedSetRequest{DefinedSet: d1})
	s.AddDefinedSet(context.Background(), &api.AddDefinedSetRequest{DefinedSet: d2})
	// s1 := &api.Statement{
	// 	Name: "s1",
	// 	Conditions: &api.Conditions{
	// 		PrefixSet: &api.MatchSet{
	// 			Name:      "d1",
	// 			MatchType: api.MatchType_INVERT,
	// 		},
	// 	},
	// 	Actions: &api.Actions{
	// 		RouteAction: api.RouteAction_ACCEPT,
	// 	},
	// }
	s2 := &api.Statement{
		Name: "s2",
		Conditions: &api.Conditions{
			PrefixSet: &api.MatchSet{
				Name:      "d1",
				MatchType: api.MatchType_ANY,
			},
			AsPathSet: &api.MatchSet{
				Name:      "d2",
				MatchType: api.MatchType_ANY,
			},
		},
		Actions: &api.Actions{
			RouteAction: api.RouteAction_ACCEPT,
		},
	}
	p1 := &api.Policy{
		Name:       "p3",
		Statements: []*api.Statement{s2},
	}
	s.AddPolicy(context.Background(), &api.AddPolicyRequest{Policy: p1})
	s.AddPolicyAssignment(context.Background(), &api.AddPolicyAssignmentRequest{
		Assignment: &api.PolicyAssignment{
			Name:          "global",
			Direction:     api.PolicyDirection_IMPORT,
			Policies:      []*api.Policy{p1},
			DefaultAction: api.RouteAction_REJECT,
		},
	})

	if err := s.MonitorTable(context.Background(), &api.MonitorTableRequest{}, func(p *api.Path) {
		fmt.Printf("Received update from %v", p)
		rsp, _ := s.GetTable(context.Background(), &api.GetTableRequest{
			TableType: api.TableType_GLOBAL,
		})
		fmt.Printf("****** table := %v", rsp)
	}); err != nil {
		fmt.Printf("Unable to monitor BGP Path, error: %v", err)
	}
	// deletion of exiting policy
	var policies []*api.Policy
	s.ListPolicyAssignment(context.Background(), &api.ListPolicyAssignmentRequest{
		Direction: api.PolicyDirection_IMPORT,
		Name:      "global",
	}, func(pa *api.PolicyAssignment) {
		// s.DeletePolicyAssignment(context.Background(), &api.DeletePolicyAssignmentRequest{
		// 	All:        false,
		// 	Assignment: pa,
		// })
		for _, policy := range pa.Policies {
			fmt.Println(policy)
			if !strings.HasPrefix(policy.Name, "p") {
				s.DeletePolicy(context.Background(), &api.DeletePolicyRequest{
					All:                true,
					PreserveStatements: false,
					Policy:             policy,
				})
			} else {
				policies = append(policies, policy)

			}
		}
		s.SetPolicyAssignment(context.Background(), &api.SetPolicyAssignmentRequest{
			Assignment: &api.PolicyAssignment{
				Name:          "global",
				Direction:     api.PolicyDirection_IMPORT,
				Policies:      policies,
				DefaultAction: api.RouteAction_REJECT,
			},
		})
	})

	// s.ListDefinedSet(context.Background(), &api.ListDefinedSetRequest{
	// 	DefinedType: api.DefinedType_PREFIX,
	// }, func(d *api.DefinedSet) {
	// 	fmt.Println("deleteing defineset ", d)
	// 	s.DeleteDefinedSet(context.Background(), &api.DeleteDefinedSetRequest{
	// 		All:        true,
	// 		DefinedSet: d,
	// 	})
	// })
	// s.ListDefinedSet(context.Background(), &api.ListDefinedSetRequest{
	// 	DefinedType: api.DefinedType_AS_PATH,
	// }, func(d *api.DefinedSet) {
	// 	fmt.Println("deleteing defineset ", d)
	// 	s.DeleteDefinedSet(context.Background(), &api.DeleteDefinedSetRequest{
	// 		All:        true,
	// 		DefinedSet: d,
	// 	})
	// })

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
