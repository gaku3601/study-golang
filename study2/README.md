# Helloworld
Helloworldを出力してみます。

# build手順

    go build Hello.go

これでHelloという実行ファイルが作成される。  
下記コマンドで実行可能

    ./Hello

すごく簡単。。。  

# 補足
プログラム内にfmtというものをimportしている部分がある。  
少し気になったので調べてみた。  

## fmtとは

    fmtパッケージは、フォーマットI/Oを実装しており、
    C言語のprintfおよびscanfと似た関数を持ちます。
    フォーマットの「書式」はC言語から派生していますが、
    より単純化されています。

ただの入出力関連のパッケージだけっぽい？  
println関数を呼び出していると解釈。  
下記参照  
[fmtパッケージ](http://golang.jp/pkg/fmt "fmtパッケージ")
