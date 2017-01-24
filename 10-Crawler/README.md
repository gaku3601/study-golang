# 概要
クローラー技術を習得します。  
使用するリポジトリはこちらになります。  
[リポジトリ]  
https://github.com/PuerkitoBio/goquery

# 設定
クローラーを行う上で、goqueryというライブラリを使用するため下記コマンドでインストールします。  
  
    go get github.com/PuerkitoBio/goquery

また、utf8以外のページをクロールする場合、文字コード変換ロジックが必要になってくる。  
そのためencoding/japaneseというライブラリもインストールしておく。

    go get -v golang.org/x/text/encoding/japanese

# 実行内容
今回は2chからスレッド名だけ取得するロジックを記述した。  
実行方法はいつも通り  

    go run Crawler.go

このコマンドを実行するとクロールされ、スレッド名をlogに出力する。
