#!/bin/bash

stat_tail="-stat.csv"

# runtime of vrfhash/vrfprove with other mining algorithms -> not too much overhead
runtime_dir="../eval/runtime-comparison"
echo "Runtime comparison..."
go test -bench "(BenchmarkSHA256D|BenchmarkScrypt|BenchmarkCryptoNight|BenchmarkCompute|BenchmarkProve)" -count=10 > $runtime_dir
benchstat -csv $runtime_dir > $runtime_dir$stat_tail
echo "Okay. See" $runtime_dir$stat_tail "\n"

# runtime breakdown of vrfhash
vrfhash_breakdown_dir="../eval/vrf-runtime-breakdown"
echo "Runtime breakdown..."
go test -bench "(BenchmarkH1|BenchmarkScalarMult|BenchmarkH2)" -count=10 > $vrfhash_breakdown_dir
benchstat -csv $vrfhash_breakdown_dir > $vrfhash_breakdown_dir$stat_tail
echo "Okay. See" $vrfhash_breakdown_dir$stat_tail "\n"
