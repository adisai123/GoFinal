package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	api "github.com/osrg/gobgp/api"
	gobgp "github.com/osrg/gobgp/pkg/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	s := gobgp.NewBgpServer()
	go s.Serve()

	// global configuration
	if err := s.StartBgp(context.Background(), &api.StartBgpRequest{
		Global: &api.Global{
			As:         65003,
			RouterId:   "192.168.43.1",
			ListenPort: -1, // gobgp won't listen on tcp:179
		},
	}); err != nil {
		log.Fatal(err)
	}

	// add peer
	addpeer("192.168.43.143", s, 65002)
	addpeer("192.168.43.62", s, 65001)
	// monitor the change of the peer state
	// if err := s.MonitorPeer(context.Background(),
	// 	&api.MonitorPeerRequest{},
	// 	func(p *api.Peer) {
	// 		log.Info(p)
	// 	}); err != nil {
	// 	log.Fatal(err)
	// }

	// add routes
	// nlri, _ := ptypes.MarshalAny(&api.IPAddressPrefix{
	// 	Prefix:    "10.33.0.1",
	// 	PrefixLen: 20,
	// })

	 a1, _ := ptypes.MarshalAny(&api.OriginAttribute{
	// 	Origin: 0,
	// })
	// a2, _ := ptypes.MarshalAny(&api.NextHopAttribute{
	// 	NextHop: "192.168.43.5",
	// })
	// a3, _ := ptypes.MarshalAny(&api.AsPathAttribute{
	// 	Segments: []*api.AsSegment{
	// 		{
	// 			Type:    2,
	// 			Numbers: []uint32{65003},
	// 		},
	// 	},
	// })
	attrs := []*any.Any{a1, a2}
	//attrs := []*any.Any{a1, a2}
	// v4Family := &api.Family{Afi: api.Family_AFI_IP, Safi: api.Family_SAFI_UNICAST}

	// v6Family := &api.Family{
	// 	Afi:  api.Family_AFI_IP6,
	// 	Safi: api.Family_SAFI_UNICAST,
	// }

	// // add v6 route
	// nlri, _ = ptypes.MarshalAny(&api.IPAddressPrefix{
	// 	PrefixLen: 64,
	// 	Prefix:    "2001:db8:1::",
	// })
	// v6Attrs, _ := ptypes.MarshalAny(&api.MpReachNLRIAttribute{
	// 	Family:   v6Family,
	// 	NextHops: []string{"2001:db8::1"},
	// 	Nlris:    []*any.Any{nlri},
	// })

	// c, _ := ptypes.MarshalAny(&api.CommunitiesAttribute{
	// 	Communities: []uint32{100, 200},
	// })

	// _, err = s.AddPath(context.Background(), &api.AddPathRequest{
	// 	Path: &api.Path{
	// 		Family: v6Family,
	// 		Nlri:   nlri,
	// 		Pattrs: []*any.Any{a1, v6Attrs, c},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// s.ListPath(context.Background(), &api.ListPathRequest{Family: v4Family}, func(p *api.Destination) {
	// 	log.Info(p)
	// })
	// time.Sleep(time.Second * 10)
	// _, err := s.AddPath(context.Background(), &api.AddPathRequest{
	// 	Path: &api.Path{
	// 		Family: v4Family,
	// 		Nlri:   nlri,
	// 		Pattrs: attrs,
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	/**policy := getStatement("import")
	s.AddPolicy(context.Background(), &api.AddPolicyRequest{
		Policy:                  policy,
		ReferExistingStatements: true,
	})
	// do something useful here instead of exiting*/
	time.Sleep(time.Minute * 300)
}

func getStatement(name string) *api.Policy {
	stmts := make([]*api.Statement, 0, 1)
	stmts = append(stmts, &api.Statement{Name: name, Actions: &api.Actions{RouteAction: api.RouteAction_ACCEPT, Nexthop: &api.NexthopAction{Self: true}}, Conditions: &api.Conditions{RouteType: api.Conditions_ROUTE_TYPE_LOCAL, AsPathSet: &api.MatchSet{MatchType: api.MatchType_ANY}}})
	policy := &api.Policy{
		Name:       name,
		Statements: stmts,
	}
	return policy
}

func addpeer(ip string, s *gobgp.BgpServer, as uint32) {
	n := &api.Peer{
		Conf: &api.PeerConf{
			NeighborAddress: ip,
			PeerAs:          as,
		},
		ApplyPolicy: &api.ApplyPolicy{
			ImportPolicy: &api.PolicyAssignment{
				DefaultAction: api.RouteAction_ACCEPT,
			},
			ExportPolicy: &api.PolicyAssignment{
				DefaultAction: api.RouteAction_REJECT,
			},
		},
	}

	if err := s.AddPeer(context.Background(), &api.AddPeerRequest{
		Peer: n,
	}); err != nil {
		log.Fatal(err)
	}
}
