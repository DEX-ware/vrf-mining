#!/usr/bin/env python3

import csv

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
        print(coin)
        for dest in dests:
            print(dest)


if __name__ == "__main__":
    main()