# 概要
簡易的なAPIサーバを作成します。

# インストール手順
下記コマンドでGo-Json-Restをインストールする。

    go get github.com/ant0ine/go-json-rest/rest

# 実行内容
get,postをcurlコマンドで行う。

    //Getコマンド
    curl -i http://127.0.0.1:8080/
    curl -i http://127.0.0.1:8080/ababa
    //Postコマンド
    curl -i -H 'Content-Type: application/json' -d '{"Name":"gaku","Password":"1234"}' http://127.0.0.1:8080/

これらのコマンドを実行すると、それぞれに見合ったjsonデータが返却される。
