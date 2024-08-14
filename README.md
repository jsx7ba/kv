![](https://github.com/jsx7ba/kv/actions/workflows/go.yml/badge.svg)

# A Key Value Service
Cleverly named `kv`. 

KV is a learning project, and is very much a work in progress.

## Stage 1
- [x] Protocol Buffers
- [x] GRPC 
- [x] REST
- [ ] Watch API
  - [x] Client needs to close when server dies
  - [x] Server needs to cancel the watch when client disconnects
  - [x] Http long polling
  - [x] Implement 'All' watch type
- [ ] More tests
- [ ] Improve server error handling so 404 can be distinguished from 500

## Stage 2
- [X] Different kv service implementations
- [ ] Distributed coordination (using raft)

## Stage 3
- [ ] Kubernetes operator

