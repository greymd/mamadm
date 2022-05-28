# これはなに？

フォロワー外のDMしちゃってごめんなさい😅
ママ活の勧誘で飛んできそうなメッセージを自動生成する、mamadm ってご存知ですか⁉️

# インストール方法

```
go get -u github.com/greymd/mamadm
```

# 使い方

```
$ mamadm -h
mamadm: Mama-katsu direct message generator.

Usage of mamadm:
  -V	バージョンを表示
  -h	ヘルプを表示
  -t int
    	0 = 文体をランダムに設定（デフォルト）, 1 = 風情ある文体にする, 2 = 通常の文体にする

Copyright (c) 2022 Yamada, Yasuhiro
Released under the MIT License.
https://github.com/greymd/mamadm
```

そのまま実行すると文言が出力される。

```
$ mamadm
お初にお目にかかります💗🐉
綺麗な美女の相手をするだけでお小遣いが得られるまま活の仲介をしているものです🥞
いきなりなんですけどmamaかつって気になったりしますか？🍄
いけてるメンズいなくて困ってます・・・。💸🐮📷
もし興味かあれば固定ツイート見に来てください🪶 👡
```

`-t` オプションの数字を 1 に指定すると、より文体に風情がでてくる。

```
$ mamadm -t 1
お初にお目にかかります
🐧
いきなりなんですが、mama★活興味ないですか？
🎯
いけてるメンズいなくて困ってます・・・。
🥞
固定ツイートからご連絡頂けましたらすぐにはじめれます！
🏐👏
```
