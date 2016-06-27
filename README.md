# siapool

[![Join the chat at https://gitter.im/robvanmieghem/siapool](https://badges.gitter.im/robvanmieghem/siapool.svg)](https://gitter.im/robvanmieghem/siapool?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Status

Early development phase, completely useless at the moment.

The intention is to make a p2pool for SIA. In a first phase the pool interface and blockgeneration will be created. This will result in a fully functional but centralized pplns pool. The sharechain and p2peer protocol will be added in phase 2.

## Connect your miner

Direct your miner to the pool using the following host: `<yourpayoutaddress>@<poolhost>:<poolport>`

Example using gominer:
```
gominer -H "1e80b18e7cdd92c3a03f307c5f453bb5a26784dfce054063b4976c8784b3a98f55ecf5f59627@siapool.tech:9985"
```

Passing this host will work on most other miners as well but siapool rejects requests for the same payout address if the previous request happened less then 5 seconds before, unless the miner submitted a share off course.

## Share difficulty

The pool adjusts the difficulty to have an accepted share every 30 seconds by average. The length of the sharechain is 8640, representing 3 days of work.

## Payout logic

Each share contains a generation transaction that pays to the previous n shares, where n is the number 8640 (= 3 days of shares).

The block reward and the transaction fees are combined and apportioned according to these rules:

A subsidy of 0.5% is sent to the miner that solved the block in order to discourage not sharing solutions that qualify as a block. (A miner with the aim to harm others could withhold the block, thereby preventing anybody from getting paid. He can NOT redirect the payout to himself.) The remaining 99.5% is distributed evenly to miners based on work done recently.

A node can choose to keep a fee for operating the node.

In the event that a share qualifies as a block, this generation transaction is exposed to the Sia network and takes effect, transferring each miner its payout.

## Support development

If you really want to, you can support the siapool development:

SIA: 1e80b18e7cdd92c3a03f307c5f453bb5a26784dfce054063b4976c8784b3a98f55ecf5f59627
