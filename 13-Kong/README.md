# 概要
APIアグリゲータであるkongを起動し、goApiがレスポンスとして返ってくるか検証する。

# ネットワーク
docker-composeファイルにはネットワーク定義を記載している。  
詳細は以下  

    kong-database:192.168.1.2
    kong:192.168.1.3
    goAPI:192.168.1.4

# コンテナ起動
コンテナの起動はいつも通り下記コマンドで行う。  

    docker-compose up -d

これで、kong,kong-database,goAPIのコンテナが起動する。

# kongへのapi登録
今回はyahoo.co.jpとgoAPIを下記コマンドを実行し登録した。  

    curl -i -X POST   --url http://localhost:8001/apis/   --data 'name=yahoo'   --data 'upstream_url=http://www.yahoo.co.jp/'   --data 'strip_request_path=true'   --data 'request_path=/yahoo' --data 'request_host=www.yahoo.co.jp'
    curl -i -X POST   --url http://localhost:8001/apis/   --data 'name=goApi'   --data 'upstream_url=http://192.168.1.4:8080/'   --data 'strip_request_path=true'   --data 'request_path=/goApi' --data 'request_host=192.168.1.4'

これでWebブラウザ等でlocalhost:8000/yahooとlacalhost:8000/goApiへアクセスすると、それぞれのデータが返却されていることを確認できる。  
また、kongでlink指定しているので下記のようなこともできる  

    curl -i -X POST   --url http://localhost:8001/apis/   --data 'name=goApi2'   --data 'upstream_url=http://go:8080/'   --data 'strip_request_path=true'   --data 'request_path=/gogo' --data 'request_host=go'

192.168.1.4がkongコンテナでgoエイリアスでアクセス可能になっているため。  
localhost:8000/gogoを覗くと同じようにjsonが返却される。  
