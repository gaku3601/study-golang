# インストール手順
macの場合brewが入っている場合、下記コマンド一発で導入可能  

    brew install go

すごく簡単。。。

# goでのパッケージ管理設定
go getでgithubなどに公開されているパッケージを使用する場合、専用フォルダを作成しpathを通しておく必要がある。

    mkdir ~/.go
    echo 'export GOPATH=$HOME/.go' >> ~/.bashrc
    echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc

上記コマンドを実行後terminalを再起動する。

# vimでのgo開発環境整備
vim-pluginでfatih/vim-goというpluginを導入する。
.vimrcに「Plug 'fatih/vim-go'」を記述しインストールする。
インストールが完了したらvim-goの各種ツールをインストールするため、vimコマンドで下記を入力

    :GoInstallBinaries

neocompleteを使用する場合はlua入りのvimが必要になってくるので注意
