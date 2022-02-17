# 2022-01-06

Go net abstraction is pretty heavy. Hard to see what’s going on in sockets
Python thinner abstraction

DNS intro:
- Decentralized because administration totally centralized is hard. Don’t want one single person to update the whole list. Delegated recursively
- .com handled by whoever, google handled by whoever, subdomain handled by whoever
- Why have multiple ips per hostname? Load balancing
- Old copies of hostnames.txt were pretty okay
- but say we have Policy: nothing out of date for more than a minute
- If you have a large amount of data, query -> data, not data -> query
- All DNS records on earth, not too bad. Queryable
- System designer choice: fault tolerance / throughput vs consistency
- DNS is more for throughput not consistency
- Need to massively distribute to get all these queries out
- Taking a lot of traffic. Also malicious traffic. Want to stay up. Hardware fails, can’t stay up. Want a system where if something breaks, it’s okay

DNS query sequence:
- your computer -> local caching DNS server (called recursive resolver) -> root server to get org NS record for .org -> local caching DNS server -> .org TLD server for wikiepdia.org
- what kind of record will caching server get from root server? NS record for .com
- additional section has glue records so like if there's a different TLD they'll give you IP for those too

cool trick for DNS:
when you know you're going to make an update, make your TTL super low for a few minutes. make sure it propagates all the way

```
nc -l -u 127.0.0.1 8888
nc -u 1.27.0.0.1 8888
```

socket and sendto used correctly syscall
use recvfrom to get the thing back correctly
good skill to read an rfc


# 2021-01-13

transport layer responsibilities:
- port mapping
- processing multiplexing
- reliable delivery
    - retry dropped packets / no dropped packets (TCP)
    - integrity. verify we don't have flipped bits
    - gets to process in order (TCP)
    - de duplication (TCP)
- flow control (TCP)
- congestion control (TCP)
- connections (TCP)

error checking: 16 bit internet checker

reliability:
- lots of retrying. who should do retries? the host who sent
- acks
- timeout if no ack from recipient after some time. (other way to do this is to have recipient keep track of what has and hasn't arrived). how to decide? keep track of mean and variance of roundtrip time. mean + 4 * variance is timeout.
- packets are dropped on routers where there isn't enough space in the buffer. (newest packets dropped usually)
- there are cumulative acknowledgements. time it for once about every roundtrip time. bc many in flight during one single roundtrip. if one packet isn't gotten though, all following packets retried though.
- 3 acks in a row says hey i didn't get something

end to end principle: if something can be done between two hosts, do it between hosts. hard to update routers. routers shouldn't have state either.

tcp slow start:
- start slow, pessimistically. can have one un-acked packet in flight. as successes continue, more in flight.
- what happens if you don't do this? then everyone pushes a lot forward, overrunning a resource. nobody lets off the gas.
- so we want old connections more than new connections. handshake is costly, but slow start is much more costly
- use connection pools

window size: clamps how far congestion control can scale. it's about how much the recipient can handle

backpressure: you shouldn't just give up or let overflow. you should be able to send back downstream to say "give me this much instead". backpressure lets you adapt to things increasing and decreasing. want some way to say more or less pls


# 2022-01-17

maybe 30 years was reasonable to do things other IP. but now this specific network layer protocol is what runs the show.

bradfield: deep but not deep enough to set up your wifi printer

origins:
- myth about internet: was made for resilience against nuclear attack. helped to get funding routed that way
- let's connect independent networks together. do whatever you want internally. with this money, connect to the other ones
- ivan sutherland was in u of u. graphics researchers there too. teapot there
- satellite link used to connect to hawaii and norway
- email was a big reason. bulletin board. telnet. remote terminal operation. mostly about collaboration and resource sharing across institutions. email documents. level up from a fax. executables and so on. primarily universities

how to do routing:
- packet going to come into a router (edge router). router has many links out to like maybe 30 peers. google for example has 426 peers!
- autonomous systems are discrete things that need to connect at the edges. ones at the edges connect to peer
- important decision at edge: where to route to?
- router is a dedicated physical device. not general purpose. prefix matching as fast as possible. MIPS machines were used as routers. now longest prefix match lookup accelerated with hardware

a lot of stuff on the web about the internet is wrong now lol. like 20 years out of date

flag day story (look up IP flag day):
- NCP was prior routing protocol. like 200 machines on arpanet. wasn't super well designed. wasn't enough thought given to higher throughput. IP was replacement for NCP. everyone had to switch to IP. "in the morning, everyone has to go to IP." 
- "would like someone to do abrupt changes. NCP -> IP, sweden road side change"

some IPV4 stuff:
`A 8 bit AS 24 bit subnet` 256 subnets, 16m size
`B 16 bit AS 16 bit subnet` 65k subnets, 65k size
`C 24 bit AS 8 bit subnet` 16m subnets, 256 size

IPV6:
- first 64 bits is subnet id. next 64 bits is for within the subnet
- facebook hit 50% ipv6 traffic a little while ago

we need hacks for ipv4:
- ips are leased out to customers. router uses DHCP to get ip address. ISP can also tell you dns server to use. will tell you lease length also. i
- no reason to expect fixed ipv6 address either tho. still dynamic

NAT
- need mechanism for knowing that some port maps to a private ip address and port
- 2^16 max number of ports

can't bind to localhost uhhhh. 0.0.0.0 can 
cisco router does whatever 
doing protocols. mostly about writing good software now.


# 2022-01-24

only a few weird tricks in network security
security is a domain in itself as is network security
could spend a whole career on one part of this and be well employed and do interesting work
security is pretty competitive


## problems vs 1 weird tricks

problem: endpoint verification (how do i know you are who you say you are?)
solution: certificates, signature, certificate authorities. certificates to prove you are who you say you are. certificate contains signature(s). cryptographic trick where you can only generate private key

problem: integrity (did someone modify it?)
solution: cryptographic hashes / MAC (message auth codes)

problem: confidentiality
solution: symmetric key encryption + small amount of asymmetric 

problem: availability
solution: idk

problem: opsec
solution: paranoia

problem: privacy
solution: VPN, TOR

## misc

all handshakes messages are verified with hashes

mid-ish way through you know the rest of the data needs to be sent. maybe before final master key. would be nice to prevent where someone who's captured every MAC in a way that would be useful to them. cert is sent so early 

ban printers because opsec risk

## symmetric

AES (symmetric) jumbles things up. quite fast. cpu architectures have support also.
- what is it if you don't have starting key?
- going to look at plaintext and guess what key was? hard bc key expansion makes every round highly entropic. all sub steps super easy.
- lot of work to come up with the system
- are expanded keys dependent on previous rounds at all? not dependent on previous round
- not that hard to write out. it's a weekend project
- TLS probably uses AES 192 or something

## public key

problem: how do you get the same key / do key exchange?
solution: public key encryption

diffi helman and RSA created around the same time. same solution to same problem
recently found out before either of these groups, someone at MI6 figured it out first

stuff you gotta get:
- be able to get a bunch of big primes
- test big primes pretty quick
- generate N = p * q (both primes)
- want to take a message x = E(m), m = D(x); x cyphertext, m message
- given N, you cannot find p and q

extremely computationally expensive. constantly taking stuff down to mod N. remaining very heavy. not feasible to share large message this way. symmetry key crypto needed bc this is so computationally intensive

there's a nonce don't worry about it pal


# 2022-01-27

neat trick: idempotency key. include a key on each request
endpoint idempotency nice
want to do only the first one
