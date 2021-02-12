#!/bin/sh

/usr/local/bin/geth --password /root/geth.password \
  --unlock "0" --syncmode full \
  --rpc --rpcvhosts "*" --rpcaddr "0.0.0.0" --rpcport "8545" --rpcapi web3,eth,net,personal,miner,txpool --rpccorsdomain '*' \
  --ws --wsapi eth,net,web3,personal,txpool --wsaddr "0.0.0.0" --wsport "8546" --wsorigins '*' \
  --datadir /root/.ethereum --networkid "2018" --nodiscover \
  --mine --minerthreads 1 --gasprice "0" \
  --allow-insecure-unlock --nousb \
  $@
