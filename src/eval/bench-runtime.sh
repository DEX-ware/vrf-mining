#!/bin/bash

cd ../vrf

# runtime of vrfhash/vrfprove with other mining algorithms -> not too much overhead
runtime_dir="../eval/runtime-comparison"
rm -rf $runtime_dir*
mkdir -p $runtime_dir
echo "Runtime comparison..."
go test -bench "(BenchmarkSHA256D|BenchmarkScrypt|BenchmarkCryptoNight|BenchmarkCompute|BenchmarkProve)" -count=10 -cpuprofile=$runtime_dir/cpu.out -benchmem -memprofile=$runtime_dir/mem.out -trace $runtime_dir/trace.out > $runtime_dir/raw.log
benchstat -csv $runtime_dir/raw.log > $runtime_dir/stat.csv
echo "Okay. See" $runtime_dir/stat.csv "\n"

# runtime breakdown of vrfhash
vrfhash_breakdown_dir="../eval/vrf-runtime-breakdown"
rm -rf $vrfhash_breakdown_dir*
mkdir -p $vrfhash_breakdown_dir
echo "Runtime breakdown..."
go test -bench "(BenchmarkH1|BenchmarkScalarMult|BenchmarkH2)" -count=10 -cpuprofile=$vrfhash_breakdown_dir/cpu.out -benchmem -memprofile=$vrfhash_breakdown_dir/mem.out -trace $vrfhash_breakdown_dir/trace.out > $vrfhash_breakdown_dir/raw.log
benchstat -csv $vrfhash_breakdown_dir/raw.log > $vrfhash_breakdown_dir/stat.csv
echo "Okay. See" $vrfhash_breakdown_dir/stat.csv "\n"

cd ../eval
