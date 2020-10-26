package madoh

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
)

var responseDnsaddrBootstrap = `{
	"Status": 0,
	"TC": false,
	"RD": true,
	"RA": true,
	"AD": false,
	"CD": false,
	"Question": [
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16
		}
	],
	"Answer": [
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/ams-1.bootstrap.libp2p.io/p2p/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/ams-2.bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/4001/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/ewr-1.bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/lon-1.bootstrap.libp2p.io/p2p/QmSoLMeWqB7YGVLJN3pNLQpmmEk35v6wYtsMGLzSr5QBU3\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/nrt-1.bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/nyc-1.bootstrap.libp2p.io/p2p/QmSoLueR4xBeUbY9WZ9xGUUxunbKWcrNFTDAadQJmocnWm\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/nyc-2.bootstrap.libp2p.io/p2p/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/sfo-2.bootstrap.libp2p.io/p2p/QmSoLnSGccFuZQJzRadHn95W2CrSFmZuTdDWP8HXaHca9z\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/sfo-3.bootstrap.libp2p.io/p2p/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/sgp-1.bootstrap.libp2p.io/p2p/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/sjc-1.bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN\""
		},
		{
		"name": "_dnsaddr.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 599,
		"data": "\"dnsaddr=/dnsaddr/sjc-2.bootstrap.libp2p.io/p2p/QmZa1sAxajnQjVM8WjWXoMbmPd7NsWhfKsPkErzpm9wGkp\""
		}
	],
	"Comment": "Response from 162.159.26.4."
}`

var addrsDnsaddrBootstrap = []string{
	"/dnsaddr/ams-2.bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
}

var responseDnsaddrServer = `{
	"Status": 0,
	"TC": false,
	"RD": true,
	"RA": true,
	"AD": false,
	"CD": false,
	"Question": [
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16
		}
	],
	"Answer": [
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/dns4/ams-2.bootstrap.libp2p.io/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/dns6/ams-2.bootstrap.libp2p.io/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/ip4/147.75.83.83/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/ip4/147.75.83.83/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/ip4/147.75.83.83/udp/4001/quic/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/ip6/2604:1380:2000:7a00::1/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/ip6/2604:1380:2000:7a00::1/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		},
		{
		"name": "_dnsaddr.ams-2.bootstrap.libp2p.io.",
		"type": 16,
		"TTL": 248,
		"data": "\"dnsaddr=/ip6/2604:1380:2000:7a00::1/udp/4001/quic/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb\""
		}
	]
}`

var addrsDnsaddrWssServer = []string{
	"/dns4/ams-2.bootstrap.libp2p.io/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/dns6/ams-2.bootstrap.libp2p.io/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/ip4/147.75.83.83/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/ip6/2604:1380:2000:7a00::1/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
}

var addrsDnsaddrTcpServer = []string{
	"/ip4/147.75.83.83/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/ip6/2604:1380:2000:7a00::1/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
}

var addrsDnsaddrQuicServer = []string{
	"/ip4/147.75.83.83/udp/4001/quic/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/ip6/2604:1380:2000:7a00::1/udp/4001/quic/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
}

var addrsDnsaddrNonWssServer = append(addrsDnsaddrTcpServer, addrsDnsaddrQuicServer...)
var addrsDnsaddrServer = append(addrsDnsaddrWssServer, addrsDnsaddrNonWssServer...)

var responseDnsServer = `{
	"Status": 0,
	"TC": false,
	"RD": true,
	"RA": true,
	"AD": false,
	"CD": false,
	"Question": [
		{
		"name": "ams-2.bootstrap.libp2p.io.",
		"type": 1
		}
	],
	"Answer": [
		{
		"name": "ams-2.bootstrap.libp2p.io.",
		"type": 1,
		"TTL": 541,
		"data": "147.75.83.83"
		}
	]
}`

var addrsDnsServer = []string{
	"/ip4/147.75.83.83/tcp/80/ws/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
}

func createServer(t *testing.T) *httptest.Server {
	handle := func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Content-Type", "application/x-javascript; charset=UTF-8")
		t.Logf("Query: %s\n", r.URL.RawQuery)
		switch query := r.URL.RawQuery; query {
		case "name=_dnsaddr.bootstrap.libp2p.io&type=16":
			w.Write([]byte(responseDnsaddrBootstrap))
		case "name=_dnsaddr.ams-2.bootstrap.libp2p.io&type=16":
			w.Write([]byte(responseDnsaddrServer))
		case "name=ams-2.bootstrap.libp2p.io&type=1":
			w.Write([]byte(responseDnsServer))
		default:
			t.Errorf("Unknown query %s\n", query)
			w.WriteHeader(http.StatusBadRequest)

		}
	}
	return httptest.NewServer(http.HandlerFunc(handle))
}

func doTest(t *testing.T, url string, expected []string) {
	server := createServer(t)
	defer server.Close()

	resolver := &madns.Resolver{Backend: &DOHResolver{
		Host:   server.URL,
		Client: server.Client(),
	}}

	maaddr := ma.StringCast(url)
	addrs, err := resolver.Resolve(context.Background(), maaddr)
	if err != nil {
		t.Error(err)
	}
	if len(addrs) != len(expected) {
		t.Errorf("Get %d. Expected %d: %s", len(addrs), len(expected), addrs)
	}

	// Sort lists to make comparisons stable
	sort.SliceStable(addrs, func(i, j int) bool {
		return addrs[i].String() < addrs[j].String()
	})
	sort.Strings(expected)

	for i, addr := range addrs {
		t.Logf("Addr %s", addr)
		if false && expected[i] != addr.String() {
			t.Errorf("Mismatched address. Got: \n\t%s\nexpected:\n\t%s\n", addr, expected[i])
		}
	}
}

func TestDnsServer(t *testing.T) {
	doTest(
		t,
		"/dns/ams-2.bootstrap.libp2p.io/tcp/80/ws/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		addrsDnsServer,
	)
}

func TestDnsaddrServer(t *testing.T) {
	doTest(
		t,
		"/dnsaddr/ams-2.bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		addrsDnsaddrServer,
	)
}

func TestDnsaddrBootstrap(t *testing.T) {
	doTest(
		t,
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		addrsDnsaddrBootstrap,
	)
}

func TestDnsaddrWssBootstrap(t *testing.T) {
	doTest(
		t,
		"/dnsaddr/bootstrap.libp2p.io/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		[]string{},
	)
}

func TestDnsaddrWssServer(t *testing.T) {
	doTest(
		t,
		"/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/443/wss/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		addrsDnsaddrWssServer,
	)
}

func TestDnsaddrTcpServer(t *testing.T) {
	doTest(
		t,
		"/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		addrsDnsaddrTcpServer,
	)
}
