#!/bin/bash

docker-compose run --rm db mysql -h db -u root 1> /dev/null << EOF
CREATE DATABASE examin;
EOF

if [ $? = 0 ]; then
  echo "データベースが正常に作成されました."
else
  echo "データベースの作成に失敗しました."
  exit 1
fi
