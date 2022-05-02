#!/bin/bash
cd cmd
go build
mv terraform-provider-junos-vsrx ../
cd ../
mv terraform-provider-junos-vsrx terraform-provider-junos-vsrx_v$1
zip terraform-provider-junos-vsrx_$1_darwin_amd64.zip terraform-provider-junos-vsrx_v$1

rm terraform-provider-junos-vsrx_v$1

cd cmd
GOOS=linux go build
mv terraform-provider-junos-vsrx ../
cd ../
mv terraform-provider-junos-vsrx terraform-provider-junos-vsrx_v$1
zip terraform-provider-junos-vsrx_$1_linux_amd64.zip terraform-provider-junos-vsrx_v$1

rm terraform-provider-junos-vsrx_v$1

shasum -a 256 *.zip *.json > terraform-provider-junos-vsrx_$1_SHA256SUMS
gpg --detach-sign terraform-provider-junos-vsrx_$1_SHA256SUMS


