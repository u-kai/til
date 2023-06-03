# 低凝集 - バラバラになったものたち -

- 凝集度とはモジュール内におけるデータとロジックの関係性の強さを表す指標

  - モジュールはクラス，パッケージ(言語によって異なる概念,go だとディレクトリ),レイヤー(開発者が決めるもの)など様々な粒度がある

- 高凝集な構造は変更に強く，望ましい構造
- 低凝集な構造は壊れやすく，変更が困難

## static メソッドの誤用

```python
# bad
class OrderManager():
    def add(money_amount,money_amount):
        return money_amount + money_amount

class MoneyData():
    def __init__(self):
        self.amount = 0

money_data = MoneyData()
money_data.amount = OrderManager.add(100,200)
```

- 上のコードはデータ部分(MoneyData)とロジック部分(OrderManager)が別々に定義されていて低凝縮

## static メソッドはインスタンス変数を使えない

- static メソッドを持ち出した時点でデータとデータを操作するロジックが乖離するためどうしても低凝集になる

## インスタンス変数を使う構造に作り替える

```python
# good

class MoneyData():
    def __init__(self,value):
        self.amount = value

    def add(self,other):
        return self.amount + other.amount

money_data = MoneyData(100)
money_data2 = MoneyData(200)
money_data.add(money_data2)

```

## インスタンスメソッドのフリをした static メソッドに注意

```python
# bad
class PaymentManager():
    def __init__(self,discount_rate):
        self.discount_rate = discount_rate

    def add(self,money_data,money_data2):
        return money_data.amount + money_data2.amount
```

- 上の add メソッドははインスタンス変数を全く使っていないので static メソッドと変わりなく，低凝集
- どれが実質的に static メソッドなのか確かめる方法として，static キーワードを入れても問題ないかを確認する

  - python だと，self を抜いてみる

- インスタンス変数を使っていればその時点でエラーが出るが，エラーが出ない場合はインスタンス変数を使っていないことがわかるため，実質 static メソッドと言うことがわかる

## どうして static メソッドがつかわれていしまうのか

- C 言語などの手続き型言語ではデータとロジックが別々になるように設計されるため，データとロジックが別々のものになることが多い
- そのなごり?

## どう言う時に static メソッドを使うべきか

- 凝縮度に影響がない場合に使える
  - ログ出力やフォーマット変換用のメソッドなど
    - これらはデータを持たないため，static メソッドにすることで凝縮度に影響がない
    - ログやフォーマット変換はドメインロジックではない(やりたいことの核ではなく，横断的な関心ごと)のためいろんなものに適用される，そのため static メソッドにしても凝縮度に影響がない
    - 例:MoneyData とその Log 出力は凝縮してなくて良い
      - 逆に MoneyData の中に Log があるの凝縮ではなく，結合
- ファクトリメソッドに使うのは良い

## 初期化ロジックの分散

```typescript
// bad
class GiftPoint {
  private point: number;
  constructor(point: number) {
    if (point < GiftPoint.MIN_POINT) {
      throw new Error("ポイントは0以上でなければなりません");
    }
    this.point = point;
  }
  add(point: number) {
    if (point < GiftPoint.MIN_POINT) {
      throw new Error("ポイントは0以上でなければなりません");
    }
    this.point += point;
  }
}
```

- 高凝集のように見えるが,実は標準会員の入会ポイントやプレミアム会員の入会ポイントといったものがあり，その度にコンストラクタが公開されているため,様々な利用をされてしまう
  - 入会ポイントを変更したい時に全てのソースをチェックする必要がある
- このような場合はコンストラクタを非公開にして代わりにファクトリメソッドを static メソッドで作る
  - ただし python ではコンストラクタを非公開にすることは難しいのでファクトリメソッドを利用するように勤めるしかないのかな？

```typescript
// good
class GiftPoint {
  static MIN_POINT = 0;
  static PREMIUM_MEMBER_POINT = 1000;
  static STANDARD_MEMBER_POINT = 500;
  private point: number;
  // private にしてコンストラクタを非公開にする
  private constructor(point: number) {
    if (point < GiftPoint.MIN_POINT) {
      throw new Error("ポイントは0以上でなければなりません");
    }
    this.point = point;
  }

  // ファクトリメソッドをstatic メソッドで作る
  static forPremiumMember(): GiftPoint {
    return new GiftPoint(GiftPoint.PREMIUM_MEMBER_POINT);
  }
  static forStandardMember(): GiftPoint {
    return new GiftPoint(GiftPoint.STANDARD_MEMBER_POINT);
  }
  // 省略
}
const premium = GiftPoint.forPremiumMember();
const standard = GiftPoint.forStandardMember();
const point = new GiftPoint(100); // コンパイル(tsだとトランスパイル)エラー
```

```python
# good
# __init__はclass名()で呼ばれてしまうため，非公開にすることはできないが,ファクトリメソッドを作ること自体が意味のないことになるわけではない

class GiftPoint:
    MIN_POINT = 0
    PREMIUM_MEMBER_POINT = 1000
    STANDARD_MEMBER_POINT = 500

    def __init__(self, point):
        if point < GiftPoint.MIN_POINT:
            raise Exception("ポイントは0以上でなければなりません")
        self.point = point

    # selfをとっていないためstaticメソッド
    def for_premium_member():
        return GiftPoint(GiftPoint.PREMIUM_MEMBER_POINT)

    def for_standard_member():
        return GiftPoint(GiftPoint.STANDARD_MEMBER_POINT)
    # 省略

premium_member_point = GiftPoint.for_premium_member()
standard_member_point = GiftPoint.for_standard_member()

# コンストラクタを呼び出せてはしまう．ファクトリメソッドがある時はファクトリメソッドを使うように勤めるしかないのかな？
point = GiftPoint(100000000)
```

## 生成ロジックが増えすぎたらファクトリクラスを検討すること

- 生成ロジックが多すぎると，そのクラスの責務がよくわからなくなる
- 責務をはっきりさせるために，そのクラスの責務と生成する責務に分けて，生成する責務を追いやったクラスを別で定義することも視野に入れる

## 共通処理クラス(Common,Util)

- 消費税計算などの共通処理の置き場として用意されたクラス
- さまざまなロジックが雑多に置かれがち
- 原因の一つは Common や Util といった共通を匂わせる名前
- 根本的な原因は共通化や再利用性に関して理解が不足しているところ
- OOP らしく，あるべき構造に処理を押し込むこと

## 横断的関心ごと

- ログ出力処理やエラー検出処理はアプリケーションのどんな処理でも必要
- このようにさまざまなユースケースでも必要となる基盤的処理を横断的関心ごととよび，このような処理は static メソッドとして定義することや共通処理としてまとめることは OK
- 例
  - ログ出力
  - エラー検出
  - デバッグ
  - 例外処理
    - 関心ごとによって処理は変わらんか？
    - 例外処理は関心ごとによって処理が変わるので，横断的関心ごとではないのでは？
  - キャッシュ
    - よくわからん
  - 同期処理
    - よくわからん
  - 分散処理
    - よくわからん

## 結果を返すために引数を使わないこと

```python
class ActorManager():
    def __init__(self):
        pass

    def shift(self,location,shift_x,shift_y):
        location.x += shift_x
        location.y += shift_y

location = Location(0,0)
manager = ActorManager()
manager.shift(location,10,10)
print(location.x,location.y) # 10,10
```

- 上のように出力として用いている引数を出力引数と呼ぶ

  - データ操作対象は Location,操作ロジックは ActorManager で分離されており低凝集

- ロジックが分散されることはもちろん，引数が知らない間に書き換わっているので，デバッグが困難になる

```python
# good
class Location():
    def __init__(self,x,y):
        self.x = x
        self.y = y
        pass

    def shift(self,shift_x,shift_y):
        self.x += shift_x
        self.y += shift_y


location = Location(0,0)
location.shift(10,10)
print(location.x,location.y) # 10,10
```

## 多すぎる引数

- 多すぎる引数は低凝集になりがち
- 引数の量が多いと言うことは，すなわちそれだけ処理させたい内容が膨らむことになる
  - 処理内容が増えるとロジックが複雑化したり，重複コードが増えたりする

## プリミティブ型執着

- bool,int,str などのプログラミング言語で用意されている基本データ型をプリミティブ型という
- クラスでラップせずにこれらの値を利用することをプリミティブ型執着という
- プリミティブ型だけで実装すると，そのデータにあるべき振る舞いが重複コードやロジックとして分散しがちになる

- データがただ存在しているだけと言うことはほとんどなく，データを使って計算したり，データを判断して制御を変えたりするので，プリミティブ型だけで実装しようとすると，データのアリかと，データを使って制御するロジックのありかがバラバラになって低凝集になる

## 意味のある単位ごとにクラス化する

- 引数が多すぎる事態に陥らないためには概念的に意味のあるクラスを作ることが肝心
- 引数が多い場合はデータを引数として扱うのではなく，そのデータをインスタンス変数として持つクラスへ設計変更するのが良い

```python
class MagicPoint():
    def __init__(self,current_amount,original_max_amount,max_increments):
        self.max_increments = max_increments
        self.original_max_amount = original_max_amount
        self.current_amount = current_amount

    def current(self):
        return self.current_amount

    def max(self):
        amount = self.original_max_amount
        for each in self.max_increments:
            amount += each
        return amount

    def recover(self,recover_amount):
        self.current_amount += recover_amount
        if self.current_amount > self.max():
            self.current_amount = self.max()

    def consume(self,consume_amount):
        # 消費処理
        pass

```

- 疑問
- 初期化のコストが上がっているが，それぞれのメソッドの引数は少なくなっている？
  - ファクトリを使うのか？
  - そもそもの解決策はクラス化ではなく，意味のある単位ごとにクラス化することだと思うが，これがそのアプローチになっているのか？

## メソッドチェイン

```python
def equip_armor(party,member_id,new_armor):
    if party.members[member_id].equipments.can_change:
        party.members[member_id].armor = new_armor
```

- 上のようなコードをメソッドチェインを使っていると言うことは，内部のロジックが外部に漏れていると言うこと
- デメテル法則
  - 利用するオブジェクツの内部を知るべきでない，とするもので，知らない人には話しかけるなとも要約される
  - メソッドチェインで内部詳細を渡り歩くつくりはまさにデメテルの法則に反している

## 尋ねるな，命じろ

- 他のオブジェクトの内部状態を尋ねたり，その状態に応じて呼び出し側が判断したりするのではなく，呼び出し側はただメソッドで命じるだけで，命令された側で適切な判断や制御をするように設計する
- これにはインスタンス変数を private にして，外部からアクセスできなくさせる
- インスタンス変数に対する制御はメソッドとして外部から命じる形にする
- そして命令された側が詳細な判断や制御を担う作りにする

```python
def equip_armor(party,member_id,new_armor):
    party.change_armor(member_id,new_armor)
```
