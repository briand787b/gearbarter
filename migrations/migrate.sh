#!/bin/bash

#echo "please provide password\n:  "
#read -s PASSWORD

psql -d gearbarter -U gearbarter -f migration.sql
