# send some requests
curl -XGET 'http://localhost:1323/check?w=thisisreallylongstring'
curl -XGET 'http://localhost:1323/check?w=civic'

# check heap
go tool pprof http://localhost:1323/debug/pprof/heap

# check memory allocation
go tool pprof http://localhost:1323/debug/pprof/allocs

# 30-sec CPU profile
go tool pprof http://localhost:1323/debug/pprof/profile?seconds=30