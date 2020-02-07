#!/bin/zsh

ab_config="-n 100 -c 10"
declare pools=( ["moo"]="cow" ["woof"]="dog")

echo "${pools[moo]}"

# ab $ab_config http://xmr-eu2.nanopool.org:14433/