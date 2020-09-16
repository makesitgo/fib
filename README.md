# fib

when running `go test -bench=.`:

```cmd
goos: darwin
goarch: amd64
pkg: github.com/makesitgo/fib
BenchmarkFib/Using_Go/fib_5-12   	41454134	        25.0 ns/op
BenchmarkFib/Using_Go/fib_10-12  	 3977953	       301 ns/op
BenchmarkFib/Using_Poplar/fib_5-12         	  189816	      5967 ns/op
BenchmarkFib/Using_Poplar/fib_10-12        	  189952	      6260 ns/op
PASS
ok  	github.com/makesitgo/fib	8.673s
```
