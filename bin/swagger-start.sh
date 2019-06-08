#!/bin/bash

# 8000 ポートが使用可能かの確認
lsof -i :8000 | grep "*:irdmi (LISTEN)" &> /dev/null

# 上のコマンドが正常に実行される -> 何かのアプリ起動してる
if [ $? = 0 ]; then
  echo "8000 番ポートはすでに使用されています"
  exit 1
fi

# swagger の起動
docker run -d --rm --name swagger -p 8000:8080 swaggerapi/swagger-editor 1> /dev/null

# 正常に起動したら、ブラウザで開く
if [ $? = 0 ]; then
  echo "localhost:8000 で起動中..."
  open http://localhost:8000
else
  echo "swagger の起動に失敗しました"
  exit 1
fi
