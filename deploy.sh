#!/usr/bin/env bash
case $1 in
        server)
            cd web_server
            sh autobuild.sh
        ;;
        vue)
            cd web_client
            npm run build:prod
            scp -r dist root@118.89.108.25:/root/vue-admin/
        ;;
        all)
            cd web_client
            npm run build:prod
            scp -r dist root@118.89.108.25:/root/vue-admin/
            cd ..
            cd web_server
            sh autobuild.sh
        ;;
        *)
                echo "$0 {server|vue|all}"
                exit 4
        ;;
esac
