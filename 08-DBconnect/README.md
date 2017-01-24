# 概要
golangからdocker上に立ち上げたpostgreSQLへアクセスしselectする方法を勉強する。

# 参考文献
こちらの記事を参考に作成しました。  
http://qiita.com/qube81/items/8d4a766280c25d0ef16c

# 実行方法
まず、postgreSQLフォルダへ遷移しdockerコンテナを立ち上げます。

    docker-compose up

立ち上げたコンテナへアクセスし、DBとテーブル、並びにデータをインサートします。

    docker-compose exec postgres bash
    
    psql -U postgres
    create database godbtest;
    CREATE TABLE jojo ( id serial, chapter int, data json );
    INSERT INTO jojo VALUES (1,1, '{ "title": "ファントムブラッド", "character": { "hero": "ジョナサン・ジョースター" } }');
    INSERT INTO jojo VALUES (2,2, '{ "title": "戦闘潮流", "character": { "hero": "ジョセフ・ジョースター" } }');
    INSERT INTO jojo VALUES (3,3, '{ "title": "スターダストクルセイダース", "character": { "hero": "空条承太郎" } }');

これでデータの準備は整いましたので、postgresコンテナからexitでログアウトします。  
最後にgoフォルダに移動し、実行コマンドを入力し、DBに格納されたデータが表示されることを確認しました。

    go run DBconnect.go


