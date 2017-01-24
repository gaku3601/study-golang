# 概要
DBマイグレーションツールであるgooseを勉強する。  
[リポジトリ]  
https://bitbucket.org/liamstask/goose/src

# 前設定
docker上にpostgreSQLコンテナを立ち上げ、testという名前でDBを作成する。  

    docker-compose up -d
    docker-compose exec postgres bash
    psql -U postgres
    create database test;

# gooseの設定
gooseのインストールを行う。
     
    go get bitbucket.org/liamstask/goose/cmd/goose

dbフォルダ内にdbconf.ymlというファイルを作成し、postgreSQLとの接続内容を記述する。

    development:
      driver: postgres
      open: user=postgres dbname=test host=localhost port=5432 sslmode=disable

これで設定は完了です。  
試しに接続できるか試します。

    gaku:09-DBConnect-goose gaku$ goose status
    goose: status for environment 'development'
    Applied At                  Migration
    =======================================

このように出力されれば設定は完了です。

# gooseを使用してみる
Usersというテーブルを作成することを想定し、gooseを使用してみます。

    goose create CreateTableTestm sql

このコマンドを実行するとdb/migrationフォルダ内にファイルが作成されるので、そのファイルを下記のように編集。

    -- +goose Up
    -- SQL in section 'Up' is executed when this migration is applied
    create table Users (
      key            char(008)     primary key,
      data1          int8,
      data2          int8,
      data3          int8
    );

    -- +goose Down
    -- SQL section 'Down' is executed when this migration is rolled back
    DROP TABLE Users;

最後に下記コマンドを実行してあげれば、migrationファイルに記載した内容でtableが作成される。

    goose up

ちなみに、goose downコマンドを使用するとdownに記載した内容が実行される。
