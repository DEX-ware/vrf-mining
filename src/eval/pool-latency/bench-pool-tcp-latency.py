#!/usr/bin/env python3

import csv
from tcp_latency import measure_latency
from statistics import mean

def main():
    row_dict_list = []
    with open('pools.csv') as f:
        reader = csv.DictReader(f, delimiter=',')
        for row in reader:
            row_dict = {}
            row_dict['coin'] = row['coin']
            row_dict['region'] = row['region']
            row_dict['host'] = row['host']
            row_dict['port'] = row['port']
            row_dict_list.append(row_dict)

    with open('latency.csv', 'w') as f:
        f.write('coin,region,host,port\n')
        for row_dict in row_dict_list:
            # print(row_dict)
            res = measure_latency(host=row_dict['host'], port=row_dict['port'], runs=10, timeout=2.5)
            res = [i for i in res if i] # remove None elements
            # print(res)
            try:
                m = mean(res)
                f.write("%s,%s,%s,%s,%f\n" % (row_dict['coin'], row_dict['region'], row_dict['host'], row_dict['port'], m))
            except:
                pass


if __name__ == "__main__":
    main()