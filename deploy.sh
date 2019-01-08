#!/usr/bin/env bash


cd web_client
npm run build:prod

scp -r dist root@118.89.108.25:/root/vue-admin/
cd ..

cd web_server
sh autobuild.sh