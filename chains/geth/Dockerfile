FROM ethereum/client-go:v1.9.25

ADD geth.password /root/geth.password
ADD genesis.json  /root/genesis.json
ADD privatekey  /root/privatekey
ADD run.sh  /run.sh

RUN /usr/local/bin/geth --nousb --datadir /root/.ethereum init /root/genesis.json

RUN /usr/local/bin/geth --nousb account import /root/privatekey --password /root/geth.password

EXPOSE 8545

ENTRYPOINT ["/run.sh"]
CMD ["--verbosity", "5", "--vmdebug", "--shh"]
