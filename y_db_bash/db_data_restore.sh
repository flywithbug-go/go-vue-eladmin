#!/usr/bin/env bash

#//远程拉去数据
#mongodump -h 118.89.108.25:27017 -u doc -p doc11121014a -d docmanager  -o ./db_backup

//同步到本地数据库
mongorestore ./db_backup/docmanager -d doc_manager --drop

