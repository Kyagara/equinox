#!/bin/bash

echo "| Benchmark | ops | ns/op | bytes/op | allocs/op |"
echo "| --------- | --: | ----: | -------: | --------: |"

while read -r line; do
    if [[ $line =~ Benchmark.*-16 ]]; then
        benchmark=$(echo "$line" | awk '{print $1}' | awk '{sub(/Benchmark/,""); sub(/-16/,""); print};')
        ops=$(echo "$line" | awk '{print $2}')
        time=$(echo "$line" | awk '{print $3}')
        bytes=$(echo "$line" | awk '{print $5}')
        allocs=$(echo "$line" | awk '{print $7}')

        echo "| $benchmark | $ops | $time | $bytes | $allocs |"
    fi
done
