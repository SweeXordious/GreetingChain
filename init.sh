#!/bin/bash

rm -r ~/.hellocli
rm -r ~/.hellod

hellod init helloNode --chain-id helloworld

hellocli config keyring-backend test

hellocli keys add me
hellocli keys add you

hellod add-genesis-account $(hellocli keys show me -a) 1000foo,10000000000000stake
hellod add-genesis-account $(hellocli keys show you -a) 1000foo,10000000000000stake

hellocli config chain-id helloworld
hellocli config output json
hellocli config indent true
hellocli config trust-node true

hellod gentx --name me --keyring-backend test
hellod collect-gentxs