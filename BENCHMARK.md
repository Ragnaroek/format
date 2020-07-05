
# Benchmark History

Run on:
- Processor: 2,6 GHz 6-Core Intel Core i7
- Memory: 32 GB 2400 MHz DDR4

## Engine refactoring for recursive handling + float fixes
BenchmarkFmtSimple-12       	18412630	        59.0 ns/op
BenchmarkFmtLong-12         	 5016384	       235 ns/op
BenchmarkFmtFloat-12        	 4191775	       293 ns/op
BenchmarkFormatSimple-12    	17161602	        68.5 ns/op
BenchmarkFormatLong-12      	 4237104	       286 ns/op
BenchmarkFormatFloat-12     	 2731093	       429 ns/op
==> FormatLong, FormatSimple profitted from the engine rewrite, Float performance got slower since
    a lot of edge cases were missing before the fix (but more ifs are need to handle this)
==> down to 16%, 22% for normal, 47% slower for float

## Floating point performance added
BenchmarkFmtSimple-12       	18428300	        61.2 ns/op
BenchmarkFmtLong-12         	 4870514	       250 ns/op
BenchmarkFmtFloat-12        	 3802862	       313 ns/op
BenchmarkFormatSimple-12    	16296456	        73.9 ns/op
BenchmarkFormatLong-12      	 3836730	       309 ns/op
BenchmarkFormatFloat-12     	 3657698	       326 ns/op

## Cached format graphs
BenchmarkFmtSimple-12       	18712560	        63.4 ns/op
BenchmarkFmtLong-12         	 4741552	       263 ns/op
BenchmarkFormatSimple-12    	14139633	        81.1 ns/op
BenchmarkFormatLong-12      	 3741561	       318 ns/op

## Baseline, no optimisation done
BenchmarkFmtSimple-12       	19969461	        61.2 ns/op
BenchmarkFmtLong-12         	 4812308	       250 ns/op
BenchmarkFormatSimple-12    	 6080833	       201 ns/op
BenchmarkFormatLong-12      	 1219668	       984 ns/op

