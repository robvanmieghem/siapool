# siapool


[![Build Status](https://travis-ci.org/robvanmieghem/siapool.svg?branch=master)](https://travis-ci.org/robvanmieghem/siapool)
[![Join the chat at https://gitter.im/robvanmieghem/siapool](https://badges.gitter.im/robvanmieghem/siapool.svg)](https://gitter.im/robvanmieghem/siapool?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Status

Early development phase, completely useless at the moment.

The intention is to make a p2pool for SIA. In a first phase the pool interface and blockgeneration will be created. This will result in a fully functional but centralized pplns pool. The sharechain is currently just a list of accepted shared and the p2peer protocol will be added in phase 2.

## Connect your miner

Direct your miner to the pool using the following host: `<poolhost>:<poolport>/<yourpayoutaddress>`

Example using gominer:
```
gominer -H "siapool.tech:9985/1e80b18e7cdd92c3a03f307c5f453bb5a26784dfce054063b4976c8784b3a98f55ecf5f59627"
```

Passing this host will work on most other miners as well but siapool rejects requests for the same payout address if the previous request happened less then 5 seconds before, unless the miner submitted a share off course.

## Share difficulty

The pool has a fixed difficulty for a 1Gh/s miner to find two shares/day on average. The length of the sharechain is 8640 * 4 (= 4 days of shares if the all miners combined would find a share every 10 seconds).

## Payout logic

Each share contains a generation transaction that pays to the previous n shares, where n is the length of the sharechain.

The block reward and the transaction fees are combined and apportioned according to these rules:

A subsidy of 0.5% is sent to the miner that solved the block in order to discourage not sharing solutions that qualify as a block. (A miner with the aim to harm others could withhold the block, thereby preventing anybody from getting paid. He can NOT redirect the payout to himself.) The remaining 99.5% is distributed evenly to miners based on work done recently.

A node can choose to keep a fee for operating the node.

In the event that a share qualifies as a block, this generation transaction is exposed to the Sia network and takes effect, transferring each miner its payout.

## Architectural concept

Siapool needs a lot of information from the sia network to be able to construct the blocks for which it hands out headers to miners and needs to feed complete blocks to the sia network. Siad does not expose this information through it's api and siapool needs to react fast on new blocks. It's a lot more comfortable if siapool accesses the internal datastructures of siad directly to be able to serve it's miners up to date block headers and to submit custom made blocks to the sia network.

This left the option of implementing siapool as a siad module or vice versa, namely importing siad and launching the modules we require ourselves. The second option has been chosen to limit the impact on the sia project itself and to leave the pool landscape for sia mining open.

An additional benefit of embedding the necessary siad functionality is that there is only a single binary, there is no need to run a separate siad and to configure the pool and siad to work together.

## Support development

If you really want to, you can support the siapool development:

SIA: 1e80b18e7cdd92c3a03f307c5f453bb5a26784dfce054063b4976c8784b3a98f55ecf5f59627
