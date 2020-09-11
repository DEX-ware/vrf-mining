# VRF-based mining

This idea is co-developed by Runchao Han (runchao.han@monash.edu), Haoyu Lin (chris.haoyul@gmail.com) and Jiangshan Yu (jiangshan.yu@monash.edu). Please find [the writeup](https://github.com/DEX-ware/vrf-mining/blob/master/paper/main.pdf) for more detail.

We introduce VRF-based mining, a surprisingly simple and effective way of making pooled mining impossible. Instead of using hash functions, we use Verifiable Random Functions (VRFs) for proof-of-work-based consensus. As VRF binds the authorship with hashes, a pool operator should reveal his private key to outsource the mining process to miners, and miners can easily steal cryptocurrency in the pool operator’s wallet anonymously.

## Directories

- `/paper`: The Latex source code of the writeup
- `/src`: The code for implementation and performance evaluation of VRF

## Acknowledgement

We thank Mikerah, Omer Shlomovits, Silur, John Trump, Cheng Wang and anonymous reviewers for insightful discussion and feedback.