# Python とは

- 1991 年発のプログラミング言語(実は結構古い)
- version は 3 が主流で現在のほとんどの version は暗黙的に 3
- 最近流行りの言語
- 動的型付け言語
- 特徴は少ないコードで簡潔にプログラムが書けること,人気があること,多くの便利ライブラリやフレームワークが存在すること

  - 以下は多くのところで利用されている Java 言語と Python 言語で標準出力するときの違い

    ```java
    public class Main() {
        public static void main(String[] args){
            System.out.println("Hello, World!");
        }
    }
    ```

    ```python
    print("Hello, World!")
    ```

    - 明らかに記述数が少なく，わかりやすい

  - Python はライブラリやフレームワークが多いため，それらを利用することで少ない記述で多くのことができてしまう
    - ライブラリやフレームワークが多いということ自体は，言語の仕様と関係ないが，言語を選択するモチベーションの一つになる

- 特に機械学習やデータ分析のライブラリやフレームワークが多く，これらを行うのであれば現状の第一候補は Python になる

  - これが流行っている大きな要因の一つ

- 初心者がまずやるならこの言語！とよく言われている印象
- 記述を進めるにはインデントを揃える必要があり，だれでも同じように綺麗に書くことができる

## 基本構文／データ構造

- Python にかかわらずプログラミング言語にはそれぞれの書き方がある

  - ただし言語の基本機能はほとんどの際がなく，一度一つの言語を覚えてしまえば他の言語もすんなりできるようになることが多い

- 今回はプログラミング言語であればほとんど兼ね備えている Python の基本構文とデータ構造を学ぶ

- 一覧
  - 変数代入
  - if else(条件分岐)
  - for (繰り返し)
  - while (繰り返し)
  - リスト(コレクションのデータ構造)
  - 辞書型(コレクションのデータ構造で他の言語ではハッシュテーブルとか HashMap とかとも呼ばれる)
  - None(データが存在しないことを示すもので他の言語では nil や null)
  - 関数
  - クラス

## 変数代入

```python
a = 10

input_data = int(input())

print(a + input_data)

```

## if else(条件分岐)

```python
import datetime

today = datetime.date.today()

if today == 29 :
    print("肉の日やん")
elif today == 1:
    print("今月も頑張ろう")
else :
    print("帰りたい")

```

## リスト

- 何かを複数保持しておきたい時がある
- 例えば，行を表現する時とか，行列を考える時とか,複数のデータを考える必要がある時
- このような複数のデータを保持する一つのデータ構造がリスト

```python

# []の中にカンマ区切りで要素を入れる
row1 = ["0","1","0"]
row2 = ["1","1","1"]
row3 = ["0","1","0"]

# リストのリストも可能
matrix = [
    row1,
    row2
]

print(matrix)

# よく使うメソッド
# リスト.appendでリストに要素を追加できる

matrix.append(row3)


```

## 辞書型(他の言語ではハッシュテーブルとか HashMap とかとも呼ばれる)

## for (繰り返し)

- python の for は近代的で，in で指定されているものから一つずつ取り出すような動きをする

```python

```

## while (繰り返し)

## None(データが存在しないことを示すもので他の言語では nil や null)

## 関数

## クラス

## 書き方

- Python はインデントによって，スコープを切る
- スコープとはある機能の範囲のこと
- 例えば関数 hoge と fuga を定義する時や if を使う時は以下

```python
# def は関数の定義に必要なワード
# :はインデントを行う前に必要なワード

def hoge():
   # ここからhogeの定義になるのでインデントを下げてかつ，揃える
   print("hoge")
   print("hoge")

# hogeからfugaの定義に変わるのでインデントを戻す
def fuga():
   print("fuga")
   print("fuga")
   # ifの条件が真であればifの下のインデントが実行される
   if 1 + 1 == 2 :
      print("fuga")
   # この行はifのしたであるがインデントがfugaのインデントなのでifの条件にかかわらず実行される
   print("fuga")

```

- 以下の間違い探しをしてください

```python
def hoge():
print("hoge")

def fuga():
    print("hoge")
    def chome():
        print("chomechome")
    chome()
    if 1 + 1 == 2
        print("fuga")

```

## そもそもライブラリ/フレームワークとは？

- ライブラリとはプログラムファイルをまとめたもので，インストールすると自分のプログラム中でそのライブラリを利用できるもの

  ```python
  # requestsというライブラリを自分のプログラムの中にimport(Pythonの構文でライブラリをそのファイルで利用できるようにする機能)している
  # ただしこの前にコマンドラインでpip install requestsをして，インターネットに存在しているrequestsライブラリを自分のPCにインストールする必要がある
  import requests
  response = requests.get("https://www.google.com")
  html = response.text()
  print(html)
  ```

  - ライブラリのおかげで自分で多くの記述を書かなくとも，すごい人が書いてくれたプログラムを利用することができ，簡単にプログラムを描くことができる

- フレームワークとはライブラリとほぼ似ているが，特徴としてはフレームワークは自分がフレームワークに合わせて書く必要がある
  - ライブラリは自分のプログラムの中の好きなところで好きに利用できる
  - インストール方法や import 方法はライブラリと同じ

```python
from flask import Flask

app = Flask(__name__)

# このように書くことをflaskに強制させられている
@router.routes("/")
def index():
  return app.send_static_file("index.html")

app.run()
```

### 標準ライブラリについて

- 上で説明したライブラリは外部からインストールしたライブラリになるが，ほとんどの言語では標準ライブラリという，言語に最初から備わっているものもある

  - もちろん Python にも存在する

- 標準ライブラリは基本的な操作が可能なものが多い
- この標準ライブラリも import 構文を使って import する

```python
import os # osの機能を利用できる標準ライブラリ
import json # json のエンコード，デコードなどができるライブラリ
from unittest import unitttest # 単体テスト用のライブラリ
...

```

## import の使い方

- import ファイル名とだけすると，そのファイル全体を利用できる
- from はファイル名，import はそのファイルの機能名を指す
- 下はどちらも結果が一緒

```python
import os

# importだけでファイル全体を呼び出した時は
# importしたファイル名.機能名である機能を呼び出すことができる

os.get_env("HOME")

from os import get_env

# from ファイル名 import 機能名
# で機能名を選んで呼び出すことも可能
# こちらの方が，やることがプログラムを見ると明確にわかって良い
get_env("HOME")

...

```

## モジュール

- 多くのプログラミング言語はプログラムの部品ごとにファイルを分けることができる
- これによって，一枚のファイルに全てのプログラムを書く必要がなく，保守性を高めることができる
- たとえば LinuxOS のコード行数は 2000 万行といわれているが，それが一枚のファイルにあるわけではない
- https://github.com/aws/aws-sdk-go/tree/main/service は aws-sdk のディレクトリ構造を見ることができる

  - ここには aws サービス毎にディレクトリを分けており，自分が利用したいサービスのコードを読んだり，修正したりするときにすぐに目的のファイルに辿り着くことが可能になる
  - 逆にいうと，ファイルやディレクトリは誰が見てもわかるように分割すべき

- このようにファイルを分ける場合，分けたとしても使えなければ意味がない
- 別のファイルにある機能を呼び出すときに使うのが Python では import 構文
- 自分でファイルに作成した機能を呼び出す時は標準ライブラリや外部からインストールしたライブラリやフレームワークを import で呼び出すのと少し異なる
- まず，Python はモジュールとして認識されたファイルやその中の機能を import できる
- Python でのモジュールは\_\_init\_\_.py というファイル(中身はからで問題なし)が格納されているディレクトリのことを指す
- import する時は import するファイルのパスを考えて import を書く必要がある
- 例えば以下のディレクトリ構成で main から func.py の機能 func_func と db.py の db_func を import して利用する場合は...

```
├── src
│   ├── __init__.py
│   ├── functions1
│   │   ├── __init__.py
│   │   ├── database
│   │   │   ├── __init__.py
│   │   │   └── db.py
│   │   └── func.py
│   └── main.py

```

- 以下のようになる

```python
from src.database.db import db_func
from src.functions1 import func_func

func_func()
db_func()

```

- では，src の下に class1 ディレクトリを用意して中身が以下の class1.py の SampleClass を main で利用するにはどうすれば良いか？

```python
# src/class1/class1

class SampleClass():
    def __init__(self):
        print("これはサンプルです")

```
