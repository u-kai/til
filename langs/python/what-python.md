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
row1 = ["0","1","2"]
row2 = ["3","4","5"]
row3 = ["6","7","8"]

# リストのリストも可能
matrix = [
    row1,
    row2
]

print(matrix)

# よく使うメソッド
# リスト.appendでリストに要素を追加できる

matrix.append(row3)

print(matrix)

# リスト[index]で要素を取得できる
# indexは0から始まる要素の並び順

first = row1[0]
print(first)

print(row2[2])
print(matrix[2])
# 要素外にアクセスすると？
row3[3]


```

## 辞書型(他の言語ではハッシュテーブルとか HashMap とかとも呼ばれる)

- dict 型とも言う
- リストのように複数のデータを保持できるデータ構造
- リストと違うのはインデックスが番号ではなく，key と呼ばれる何かしらのデータ

  - key は等価性がある必要があり，格納されているデータは全て同じデータ型である必要がある

- リストよりも直感的にデータアクセスできる
  - リストはアクセスしたいデータが何番目にあるのか知っている必要があるが，辞書型は key という人間が認識しやすい値を覚えておけば良いため

```python

# dictは{}で囲み,key:valueの記述で書く
# 下の例はkeyがstr型(文字列型)のdict
d1 = {"name":"kai","age":20,"hobby":"it"}

# dict[key]で要素を取得できる
print(d1["name"])
# 要素外にアクセスすると？
print(d1["like"])


# 要素を指定して書き換えも可能

d1["age"] = 27

```

## for (繰り返し)

- python の for は近代的で，in で指定されているものから一つずつ取り出すような動きをする

```python

# pythonのrangeは0から与えられた引数までの連続的な値を返すもの(これは一機能の説明)
# なのでforと組み合わせると0~9までが一つずつ取り出される
for i range(9):
    print(i)

list = ["kai","Kai","KAI"]
# リストの中身を一つずつ取り出す
for data in list:
    print(data)

dict = {"name":"kai","age":20,"hobby":"it"}

# 辞書型.items()でkeyとvalueの組みを一つずつ取り出す
# 辞書型の要素には順番がないため，記述した順番通りに取り出されるかは言語によって異なる
for key,value in dict.items():
    print("key = ",key)
    print("value = ",value)

# 文字を一つずつ切り出して取り出す

for s in "hello world":
    # \nは改行文字
    print(s + "\n")

# 10と言う数字は複数データではないので値を取り出せない
for no in 10:

```

## while (繰り返し)

- while は与えられた条件式が真である限り繰り返しが続く

```
# このwhileは常に真であるため無限ループ
while True:
    print("loop!")

# Ctrl+cでプログラムを強制終了しよう
```

```
sum = 0
while sum < 10
    sum = sum + 1

```

# continue, break

- for や while で利用できるもの

- continue はそれより下の処理を行わずに繰り返しに戻る
- break はそれより下の処理を行わずに繰り返しを終える

```
not_contain_8 = []

for i in range(10):
    if i == 8:
        continue

    not_contain_8.append(i)

many_many_data = ["kai","Kai","KAI","書い","開","かい","カイ","海"]

for data in many_many_data:
    if data == "開":
        print(data+"さん")
        break
    print("data")

```

## None(データが存在しないことを示すもので他の言語では nil や null)

## 関数

- 1 つ以上の処理を記述して実行できるもの
- 複数の処理をまとめて管理しやすくしたり，可読性を上げたり，繰り返し利用できることを可能にする
- 関数は 0 個以上の引数を取ることができ，処理に利用できる

  - 引数とは関数の外部から引き込んでくる値
  - これがあることで，関数定義時には決定できないような値でも仮置きすることができ，柔軟に関数を定義できる

- 関数は処理後に値を返すことができる
  - return VALUE で VALUE を返すことができる

```python

# defで関数を定義する
def sum(start,end):
    # 複雑な処理を隠蔽して，名前をつけることで利用しやすくなる
    result = 0
    for i in range(start,end+1):
        result += i

    # 計算した値を返す
    return result

# 少ない記述で繰り返し実行できる
sum_range_1_50 = sum(1,50)
print(sum_range_1_50)
# 利用者はわかりやすい名前で関数にアクセスできる
sum_range_100_1000 =um(100,1000)
print(sum_range_100_1000)



# 一行だけだが，普通の記述に名前がつくので可読性が上がる
def add_double_quote(string):
    # 引数で受け取った文字列をダブルクォーとで囲む
    return '"'+ string+ '"'

data = "hello"
print(add_double_quote(data))

# returnがないと何も返さない
def many_print(data,num):
    for _ in range(num):
        print(data)

many_print("hello",100)
```

## クラス

- データの属性を複数持つことができ，何かの概念を表現したいときに利用する
  - 読みやすくなり,管理もしやすくなる
- データの属性だけではなく，複数の振る舞いもメソッドとして持つことができる
- オブジェクト指向というプログラムの考え方では，クラスを入念に設計して，プログラムで表現したい物に近づけ，クラス同士を強調動作させることで可読性を上げたり，管理性をあげる考え
- うまく設計して，クラスの責務をあるべきクラスに閉じ込めることでプログラムの追加や修正をするときにどのクラスを修正すれば良いかわかる
- また，プログラムも処理の連続ではなく，クラスに対して何かを呼び出すような文章に近くなり，プログラムが読みやすくなる

```python
# クラスの定義
class Human():
    # 初期化時に呼ばれるメソッド(クラスに紐づいている関数のこと)
    # selfはこのクラス自体のこと(クラスを利用するときのおまじないだと思えば良い)
    def __init__(self,name,age,hobby):
        # 自身の属性にデータを入れ込んでいる
        self.name = name
        self.age = age
        self.hobby = hobby

    # 人間って挨拶するよね
    def say_hello(self):
        print("Hello, My name is",self.name)

    def who_is_the_senior(self,other):
        if self.age > other :
            return self
        if other > self.age:
            return other
        return self,other


# クラスを利用するとき
kai = Human("kai",20,"it")
hiroki = Human("hiroki",28,"後輩いじり")


# クラスのインスタンス.属性名でそのクラスの属性を呼び出せる
# .を~の，と読み替えるとわかりやすい
print(kai.name)
print(kai.age)
print(kai.hobby)

kai.say_hello()
hiroki.say_hello()

print(kai.who_is_the_senior(hiroki).name)

class HomeRoom():
    def __init__(self,students,teacher,class_name):
        self.class_name = class_name
        self.students = students

        if teacher.class_name != class_name:
            print("Error")
        else :
            self.teacher = teacher

    def add_student(self,student):
        if student.class_name == self.class_name:
            self.students.append(student)

    def all_student(self,student):
        return self.students

    def do_class(self):
        self.teacher.teach()
        for s in self.students:
            print(s.name + ":" + s.reaction())

class Student():
    def __init__(self,name,age,class_name):
        self.name = name
        self.age = age
        self.class_name = class_name

    def reaction(self):
        return "はーい"

class Teacher():
    def __init__(self,name,class_name,expert):
        self.name = name
        self.class_name = class_name
        self.expert = expert

    def teach(self):
        print(self.expert+"の授業を始めますよ")


# どのような処理を実際に行っているかはわからないが,大まかに何が立っ性されているのかはわかりやすい
# 抽象化や詳細の隠蔽によって複雑度が各クラスに閉じ込められている
student1 = Student("kai",20,"A")
student2 = Student("hiroki",28,"B")
student3 = Student("nishimachi",25,"A")
teacher = Teacher("jun","A","it")

home_room = HomeRoom(
    [student1,student2,student3],teacher,"A"
)

home_room.do_class()

```

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
