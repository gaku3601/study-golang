# 概要
docker上でgolangを動作させます。

# 実行方法
下記コマンドでdockerコンテナを立ち上げる。

    docker-compose build
    docker-compose up

golangのプログラムは06で作ったものなので、06と同じように下記コマンドでレスポンスがあるか検証した。

    curl -i -H 'Content-Type: application/json' -d '{"Name":"gaku","Password":"1234"}' http://127.0.0.18080/

# 補足
golangではなくdockerの問題なのだが$PWDを使用する際、パスにマルチバイト文字が含まれている場合エラーとなる。  
そのためvolumeマウント時などで使用する$PWDのパスにマルチバイト文字が含まれないよう注意する必要がある。  
