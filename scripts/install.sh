#!/bin/bash

set -e

ASSETS_PATH=/usr/local/etc/drummachine/sounds
mkdir -p $ASSETS_PATH
cp assets/sounds/*.sf2 $ASSETS_PATH/

GO111MODULE=on
INSTALL_PATH=/usr/local/bin/
echo Installing in $INSTALL_PATH
go build ./cmd/drummachine
mv ./drummachine $INSTALL_PATH
chmod +x $INSTALL_PATH/drummachine