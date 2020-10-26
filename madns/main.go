package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
	madoh "github.com/steamraven/go-multiaddr-dns-doh"
)

func main() {

	var resolver *madns.Resolver

	var doDoh = flag.Bool("doh", false, "Resolve using DNS over HTTPS")
	var dohServer = flag.String("doh-url", "", "URL to use for DNS over HTTPS queries")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprint(flag.CommandLine.Output(),
			"       madns /dnsaddr/example.com/ipfs/Qmfoobar\n"+
				"       madns /dns6/example.com\n"+
				"       madns /dns6/example.com/tcp/443/wss\n"+
				"       madns /dns4/example.com\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *dohServer != "" {
		_, err := url.Parse(*dohServer)
		if err != nil {
			fmt.Printf("Error parsing doh query url: %s\n", err)
			os.Exit(1)
		}
		resolver = &madns.Resolver{Backend: &madoh.DOHResolver{
			Host:   *dohServer,
			Client: http.DefaultClient,
		}}
	} else if *doDoh {
		fmt.Fprint(os.Stderr, "Using DNS over HTTPS\n")
		resolver = madoh.DefaultDOHResolver
	} else {
		resolver = madns.DefaultResolver
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	query := flag.Arg(0)
	if !strings.HasPrefix(query, "/") {
		query = "/dnsaddr/" + query
		fmt.Fprintf(os.Stderr, "madns: changing query to %s\n", query)
	}

	maddr, err := ma.NewMultiaddr(query)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	rmaddrs, err := resolver.Resolve(context.Background(), maddr)
	if err != nil {
		fmt.Printf("error: %s (result=%+v)\n", err, rmaddrs)
		os.Exit(1)
	}

	for _, r := range rmaddrs {
		fmt.Println(r.String())
	}
}
