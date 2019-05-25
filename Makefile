# プロジェクトのビルド
build:
	go build -o examin_go

# コンテナの作成
setup:
	docker-compose build

# サーバの起動
run:
	docker-compose up
