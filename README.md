# go-multiaddr-dns-doh

> Resolve /dns4, /dns6, and /dnsaddr multiaddrs using DNS over HTTPS, json interface.

```sh
> madns -doh /dnsaddr/ipfs.io/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip4/104.236.151.122/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip6/2604:a880:1:20::1d9:6001/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip6/fc3d:9a4e:3c96:2fd2:1afa:18fe:8dd2:b602/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/dns4/jupiter.i.ipfs.io/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/dns6/jupiter.i.ipfs.io/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
```


In more detail:

```sh
> madns -doh /dns6/example.net
/ip6/2001:db8::a3
/ip6/2001:db8::a4
...

> madns -doh /dns4/example.net/tcp/443/wss
/ip4/192.0.2.1/tcp/443/wss
/ip4/192.0.2.2/tcp/443/wss

# No-op if it's not a dns-ish address.

> madns -doh /ip4/127.0.0.1/tcp/8080
/ip4/127.0.0.1/tcp/8080

# /dnsaddr resolves by looking up TXT records.

> dig +short TXT _dnsaddr.example.net
"dnsaddr=/ip6/2001:db8::a3/tcp/443/wss/ipfs/Qmfoo"
"dnsaddr=/ip6/2001:db8::a4/tcp/443/wss/ipfs/Qmbar"
"dnsaddr=/ip4/192.0.2.1/tcp/443/wss/ipfs/Qmfoo"
"dnsaddr=/ip4/192.0.2.2/tcp/443/wss/ipfs/Qmbar"
...

# /dnsaddr returns addrs which encapsulate whatever /dnsaddr encapsulates too.

> madns -doh example.net/ipfs/Qmfoo
info: changing query to /dnsaddr/example.net/ipfs/Qmfoo
/ip6/2001:db8::a3/tcp/443/wss/ipfs/Qmfoo
/ip4/192.0.2.1/tcp/443/wss/ipfs/Qmfoo