FROM hyperledger/besu:23.7.2

USER root

RUN mkdir -p /tmp/besu/data
WORKDIR /tmp/besu
ADD ibftConfigFile.json /tmp/besu/ibftConfigFile.json
RUN besu operator generate-blockchain-config --config-file=ibftConfigFile.json --to=networkFiles --private-key-file-name=key
RUN cp ./networkFiles/keys/*/* ./data/

EXPOSE 8545 8546 8547 30303
ENTRYPOINT [ "besu" ]
CMD ["--data-path", "./data", "--genesis-file", "./networkFiles/genesis.json", "--rpc-http-enabled", "--rpc-http-api", "ETH,NET,IBFT,DEBUG", "--host-allowlist", "*", "--rpc-http-cors-origins", "all", "--revert-reason-enabled"]
