# kty-cli
F島県A大学のGo言語の課プロでやった共同開発。</br>
コードはその一部です。

### やったこと
コマンドライン上から、DMを飛ばす。

### ぼくがやったこと
- サブコマンド、オプションの実装(テスト) -> cli.go </br>
- SlackのDMのクライアント実装 -> slack.go

### こんな感じになった
```
$ ./cli send -s slack -m こんにちは、tatatakky -u tatatakky
```
- -s : slack、twitterの送りたいほうのどちらかを指定。
- -all : 強制的にslack、twitterのどちらも指定。
- -m : 送りたいメッセージを書く。
- -u : 送りたいユーザを指定。
