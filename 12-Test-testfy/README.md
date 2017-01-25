# 概要
testパッケージであるtestfyを勉強する。  
[リポジトリ]  
https://github.com/stretchr/testify

# testfyのインストール
下記コマンドでtestfyをインストールする。

    go get github.com/stretchr/testify/assert

# 実行方法
プロジェクトのカレントディレクトリで下記コマンドを実行する。

    go test ./...

このコマンドで再帰的にフォルダを探索しテストが走る。  
