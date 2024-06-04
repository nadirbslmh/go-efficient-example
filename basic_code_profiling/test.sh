# analyze CPU and memory profile
go test -bench='.' -cpuprofile='cpu.prof'

# CPU profiling
go tool pprof cpu.prof

# Commands for pprof interactive console
top15 -cum
granularity=lines
hide=runtime
hide= 
list funcName
web
