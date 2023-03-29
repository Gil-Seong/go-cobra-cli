#!/bin/sh

chmod 755 ./go-cli
mkdir /etc/ktctl
cp ./go-cli /etc/ktctl
ln -s /etc/ktctl/go-cli /usr/bin/ktctl
