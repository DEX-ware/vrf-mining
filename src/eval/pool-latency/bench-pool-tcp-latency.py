#!/usr/bin/env python3

import csv
from tcp_latency import measure_latency
from statistics import mean

def main():
    coins = {}
    with open('pools.csv') as f:
        reader = csv.DictReader(f, delimiter=',')
        for row in reader:
            dest = {}
            dest['host'] = row['host']
            dest['port'] = row['port']
            coin = row['coin']
            if coin not in coins:
                coins[coin] = [dest]
            else:
                coins[coin].append(dest)

    for coin, dests in coins.items():
        # print(coin)
        if coin.startswith("#"):
            continue

        for dest in dests:
            # print(dest)
            res = measure_latency(host=dest['host'], port=dest['port'], runs=10, timeout=2.5)
            res = [i for i in res if i] # remove None elements
            # print(res)
            try:
                m = mean(res)
                print("%s,%f" % (coin, m))
            except:
                pass


if __name__ == "__main__":
    main()