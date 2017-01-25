# インストール手順
macの場合brewが入っている場合、下記コマンド一発で導入可能  

    brew install go

すごく簡単。。。

# GOPATHの設定
go getでgithubなどに公開されているパッケージを使用する場合、専用フォルダを作成しpathを通しておく必要がある。  
また作成するプログラムもGOPATH配下で作成する必要がある。  
そのため、下記コマンドでGOPATHの設定を行っておく

    mkdir ~/.go
    echo 'export GOPATH=$HOME/.go:$HOME/Documents/develop/projects/go' >> ~/.bashrc
    echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc

上記コマンドを実行後terminalを再起動する。  
## 解説
go getコマンドで取得したパッケージは$HOME/.goフォルダへ格納される。  
goプログラムを作成する場所を$HOME/Documents/develop/projects/goとし、フォルダ分けを行っている。  

# vscodeでのgo開発環境
下記サイト参考  
macの場合証明書発行が必要になってくるので注意  
http://dev.classmethod.jp/go/visual-studio-code-golang-debug/  
https://github.com/derekparker/delve/blob/master/Documentation/installation/osx/install.md  


# vimでのgo開発環境整備
vim-pluginでfatih/vim-goというpluginを導入する。
.vimrcに「Plug 'fatih/vim-go'」を記述しインストールする。
インストールが完了したらvim-goの各種ツールをインストールするため、vimコマンドで下記を入力

    :GoInstallBinaries

neocompleteを使用する場合はlua入りのvimが必要になってくるので注意
