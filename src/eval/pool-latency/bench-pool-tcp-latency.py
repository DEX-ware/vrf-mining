#!/usr/bin/env python3

import csv
from tcp_latency import measure_latency
from statistics import mean, stdev
import logging

logging.basicConfig(level=logging.DEBUG)


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

    logging.info(row_dict_list)

    with open('latency.csv', 'w') as f:
        fields = ['coin', 'region', 'host',
                  'port', 'latency_mean', 'latency_std']
        f_csv = csv.DictWriter(f, fieldnames=fields)
        f_csv.writeheader()

        for row_dict in row_dict_list:
            # print(row_dict)
            res = measure_latency(
                host=row_dict['host'], port=row_dict['port'], runs=10, timeout=2.5)
            # basic data processing
            res = [i for i in res if i]  # remove None elements
            if len(res) == 0:
                continue
            # mean and std
            row_dict['latency_mean'] = mean(res)
            row_dict['latency_std'] = stdev(res)
            logging.info(row_dict)
            # write
            f_csv.writerow(row_dict)


if __name__ == "__main__":
    main()
