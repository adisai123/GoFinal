package main

import (
	"context"
	"time"

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
			As:         65000,
			RouterId:   "192.168.43.5",
			ListenPort: -1, // gobgp won't listen on tcp:179
		},
	}); err != nil {
		log.Fatal(err)
	}

	// add peer
	addpeer("192.168.43.143", s, 65000)
	//addpeer("192.168.43.62", s, 65001)
//	monitor the change of the peer state
d1 := &api.DefinedSet{
	DefinedType: api.DefinedType_PREFIX,
	Name:        "d1",
	Prefixes: []*api.Prefix{
		&api.Prefix{
			IpPrefix:      "10.1.0.0/24",
			MaskLengthMax: 24,
		},
	},
}
d2 := &api.DefinedSet{
	DefinedType : api.DefinedType_AS_PATH,
	Name:       "d2",
	List: string[]{"65000"}
}
s1 := &api.Statement{
	Name: "s1",
	Conditions: &api.Conditions{
		PrefixSet: &api.MatchSet{
			Name: "d1",
		},
		AsPathSet: &api.MatchSet{
			Name: "d2"
		}
	},
	Actions: &api.Actions{
		RouteAction: api.RouteAction_ACCEPT,
	},
}
p2 := &api.Policy{
	Name:       "p2",
	Statements: []*api.Statement{s1},
}
err = s.AddPolicy(context.Background(), &api.AddPolicyRequest{Policy: p2})
err = s.AddPolicyAssignment(context.Background(), &api.AddPolicyAssignmentRequest{
	Assignment: &api.PolicyAssignment{
		Name:          table.GLOBAL_RIB_NAME,
		Direction:     api.PolicyDirection_IMPORT,
		Policies:      []*api.Policy{p2},
		DefaultAction: api.RouteAction_ACCEPT,
	},
})	
// if err := s.MonitorPeer(context.Background(),
	// 	&api.MonitorPeerRequest{},
	// 	func(p *api.Peer) {
	// 		log.Info(p.)
	// 	}); err != nil {
	// 	log.Fatal(err)
	// }
	
	// add routes
	// nlri, _ := ptypes.MarshalAny(&api.IPAddressPrefix{
	// 	Prefix:    "10.33.0.1",
	// 	PrefixLen: 20,
	// })
	// a2, _ := ptypes.MarshalAny(&api.LocalPrefAttribute{
	// 	LocalPref: 1000,
	// })
	// a1, _ := ptypes.MarshalAny(&api.OriginAttribute{
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
	//attrs := []*any.Any{a2}
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
	// 		// 		Family: v4Family,
	// 		// 		Nlri:   nlri,
	// 		Pattrs: attrs,
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	//s.SetPolicies
	// do something useful here instead of exiting*/
	time.Sleep(time.Minute * 300)

}

func createPolicyDefinition(defName string, stmt ...&api.Statement) &api.Policy {
	pd := &api.Policy{
		Statements :
		Name:       defName,
		Statements: []config.Statement(stmt),
	}
	return pd
}

func createStatement(name, psname, nsname string, accept bool) &api.Statement {
	c := &api.Conditions{
		MatchPrefixSet: &api.MatchPrefixSet{
			PrefixSet: psname,
		},
		MatchNeighborSet: &api.MatchNeighborSet{
			NeighborSet: nsname,
		},
	}
	rd := &api.ROUTE_DISPOSITION_REJECT_ROUTE
	if accept {
		rd = &api.ROUTE_DISPOSITION_ACCEPT_ROUTE
	}
	api.GetTableRequest().
	a := &api.Actions{
		RouteAction: ,
	}
	s := &api.Statement{
		
		Name:   name,
		Conditions: c,
		Actions:    a,
	}
	return s
}

func addpeer(ip string, s *gobgp.BgpServer, as uint32) {
	n := &api.Peer{
		Conf: &api.PeerConf{
			NeighborAddress: ip,
			PeerAs:          as,
		},
		ApplyPolicy: &api.ApplyPolicy{
			ImportPolicy: &api.PolicyAssignment{
				Name:          "aditya",
				DefaultAction: api.RouteAction_ACCEPT,
			},
			ExportPolicy: &api.PolicyAssignment{
				DefaultAction: api.RouteAction_REJECT,
			},
		},
	}
	//	s.AddPolicyAssignment
	if err := s.AddPeer(context.Background(), &api.AddPeerRequest{
		Peer: n,
	}); err != nil {
		log.Fatal(err)
	}
}
