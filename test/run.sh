#!/bin/bash

script_dir=$(dirname $0)
cd ${script_dir}

set -o xtrace

# GET /health
curl -s -X GET localhost:18000/health | jq .

# POST /register (Administrator)
curl -s -X POST localhost:18000/register --data @json/admin_register.json | jq .

# POST /login
curl -s -X POST localhost:18000/login --data @json/admin_login_wrong.json | jq .
curl -s -X POST localhost:18000/login --data @json/admin_login.json | jq .
token=$(curl -s -X POST localhost:18000/login --data @json/admin_login.json | jq ".access_token" | sed "s/\"//g")

for i in {1..5}; do
# POST /tasks
curl -s -X POST -H "Authorization: Bearer ${token}" localhost:18000/tasks --data @json/add_task.json | jq .
done

# GET /tasks
curl -s -X GET -H "Authorization: Bearer ${token}" localhost:18000/tasks | jq .

# GET /admin
curl -s -X GET -H "Authorization: Bearer ${token}" localhost:18000/admin | jq .

# AdministratorからUserに切り替えてもう一度実行

# POST /register (User)
curl -s -X POST localhost:18000/register --data @json/user_register.json | jq .

# POST /login
curl -s -X POST localhost:18000/login --data @json/user_login.json | jq .
token=$(curl -s -X POST localhost:18000/login --data @json/user_login.json | jq ".access_token" | sed "s/\"//g")

# GET /tasks
curl -s -X GET -H "Authorization: Bearer ${token}" localhost:18000/tasks | jq .

# GET /admin
curl -s -X GET -H "Authorization: Bearer ${token}" localhost:18000/admin | jq .
