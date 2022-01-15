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
