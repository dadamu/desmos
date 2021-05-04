#!/usr/bin/env bash

IBCDIR=$1
ACCOUNTNUM=$2
ibc0=$3
ibc1=$4

# Init account tx in ibc
for (( i = 0; i < $ACCOUNTNUM; i++ ))
do
    desmos tx bank send test0-$i desmos1punhxfyxvnwup70mc6cz9cmuxu74ars7x6hgjm 1desmos --home "$IBCDIR/ibc0" \
    --keyring-backend "test" --node $ibc0 --chain-id ibc0 --broadcast-mode async --yes &>/dev/null

    desmos tx bank send test1-$i desmos1punhxfyxvnwup70mc6cz9cmuxu74ars7x6hgjm 1desmos --home "$IBCDIR/ibc1" \
    --keyring-backend "test" --node $ibc1 --chain-id ibc1 --broadcast-mode async --yes &>/dev/null
done