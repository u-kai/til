# IPsec

- IP パケットをカプセル化して伝送する規格

## 通信モード

- 通信経路上でカプセル化する範囲を一部の区間とするのか，全区間にするのかに応じて通信モードを決定する
- 全区間の場合はトランスポートモード
  - 送信もとから宛先まで
- 一部の区間の場合はトンネルモード
  - 通信経路上のルータが IPsec の機能を持つケース

## セキュリティポリシ

- パケットの種類を識別する情報をセレクタといい，IP アドレス，ポート番号，上位層プロトコルなどがある
- セレクタに応じた動作は PROTECT(IPsec の処理を行う),BYPASS(IPsec の処理を行わず，通常の処理を行う)DISCARD(パケットを破棄する)の 3 種類がある
- この決定を下すための情報(セレクタと動作の組)をセキュリティポリシという
- ゲートウェイは複数のセキュリティポリシを持つことができ，SPD(Security Policy Database) に登録される
  - つまり，このセレクタのときはこの動作にする，みたいなポリシをセキュリティポリシって言ってそれを SPD に保存できるって感じ？

## SA

- IPsec で使用されるメッセージ認証方式，暗号化方式，暗号鍵などは IKE(Internet Key Exchange)のやり取りを経てゲートウェイ間で合意される

  - この合意のことを SA(Security Association)と呼ぶ

- この合意に基づいてゲートウェイはそれぞれ，共通鍵方式に基づく暗号鍵を生成する
- そして SA を通過するすべてのトラフィックに対して同一の暗号化とメッセージ認証の処理を行う

## IKE SA

- IPsec で使用するパラメータを交換するために生成される
- IKE SA の通信は暗号化とメッセージ認証が行われるため安全
- IKE SA は双方向の SAQ
- ISAKMP SA とも呼ばれる

## IPsec SA(Child SA)

- IPsec SA は IPsec を用いたセキュアな通信を行うために生成される
- ESP が選択されると通信が暗号化される
- このセキュリティ処理に使用されるメッセージ認証方式，暗号化方式，暗号鍵は上り用と下り用の通信で異なるものが用いられる
- つまり，IPsec SA は単方向の SA であり，通信のたびに 2 個生成される
- IPsec のカプセル化に使用する IPsec プロトコル(AH,ESP)ごとに別々の IPsec SA を生成する

  - ESP と AH を併用する場合，上りくだりそれぞれに ESP と AH 用の SA を生成するので 4 つの IPsec SA が生成される

- IKEv2 では Child SA と呼ばれている

## ReKey

- IPsec では SA が生成されるたびに異なる暗号鍵が生成される
- 通信の暗号化は共通鍵暗号方式によって行われる
- SA の生存期間を定め生存期間が満了するたびに SA を作り替えることを ReKey と呼ぶ
- ReKey は生存期間が満了する前に後継の SA が作成され，通信が遮断されないように SA の移行が行われる
- 効率的に行いたい場合，IKE SA の寿命を IPsec SA よりも長く設定すると良い

  - そうすることで，再生性が必要になるのは IPsec SA だけになるため

- また IPsec SA を生成する際，PFS(Perfect Forward Security)を無効にして，Diffie-Hellman 鍵交換を行わずに IPsecSA の共通鍵を生成するとよい
- IKE v1 では IKE SA は生存期間が満ちて SA が破棄されても，IKE SA が必要とする通信が起こるまで新しい SA は作成されない
  - つまり事実上 ReKey はない
  - IKE v2 では IKE SA は IPsec SA と同様に ReKey が行われる

## SPI(Security Parameters Index)

- IPsec 通信を行う両端ホストが IPsec SA を識別するのに用いる 32 ビット長の番号
  - IKE v2 からは IKE SA を識別する SPI が規定されており，長さは 64 ビット
- 受信ホスト送信された IPsec ヘッダに格納されている SPI の値を取得し，その IPsec SA の暗号化鍵を使って複合している
- IPsec では IPsec SA 生成時に鍵を交換するだけでなく，互いが使用する暗号化方式やメッセージ認証方式も交換しており，これも SPI によって識別している

## IKE

- IPsec 鍵交換のためのプロトコル
- 送信もとと宛先のポート番号はともに UDP の 500 番が使用される
- 流れとしてはおおかに以下
  - SA を用いない通信で IKE SA の準備
  - IKE SA を用いた通信
  - IPsec SA を用いた通信

## IKE v1

- フェーズ 1 で IKE SA のパラメータを交換し，IKE SA を生成する

  - パラメータの交換と同時にエンティティ認証も行う
  - 最初のフェーズでは暗号化が行えないので,パラメータ交換には Diffie-Hellman 鍵交換を使う
  - メインモード
    - 相手の IP アドレスを元にエンティティ認証に必要な情報を選ぶ
      - 通常は事前共有鍵
    - イニシエータとレスポンダの両方が固定 IP アドレスを持っている必要がある
    - メインモードは 3 往復で行われる
  - アグレッシブモード
    - メインモードに比べて処理が簡素化されていて 1 往復半でやり取りが行われる
    - イニシエータは固定 IP アドレスでなくても良い

- いずれかのモードを用いて下記の処理が行われる

  - 暗号化やメッセージ認証のアルゴリズムの合意
    - IKE SA で使用される暗号化アルゴリズム，メッセージ認証アルゴリズムなどを交換する
    - ここで合意されたメッセージ認証用のハッシュアルゴリズムをもとに鍵生成のための疑似乱数関数が暗黙的に選定される
  - 暗号鍵の生成
    - Diffie-Hellman 鍵交換により，互いが同じ暗号鍵を生成し，共有できるよう，乱数を交換する
      - この鍵はフェーズ 2 の通信を暗号化するために用いられる
      - 安全な鍵交換ができるが，なりすましが行われると，第三者に暗号鍵が漏洩してしまうので，次のエンティティ認証が必要
  - エンティティ認証
    - 方式としては，事前鍵共有方式，デジタル署名方式，公開鍵暗号方式(RSA)など
    - 通常事前鍵共有方式が多い
      - これはイニシエータ，レスポンダーが事前に同じキー(PSK)を共有する方式
      - ただしこれは機器同士の話で，利用者の認証ではない
    - エンティティを認証するために，フェーズ 1 で交換した情報のハッシュ値を利用する
  - IKE SA の生成
    - パラメータを交換した後 IKE SA を生成する
    - 交換される主なパラメータは以下
      - 暗号化アルゴリズム
        - IKE SA で使用する暗号化のアルゴリズム
      - ハッシュアルゴリズム
        - IKE SA で使用するメッセージ認証のアルゴリズム，鍵生成のための疑似乱数関数の元となるハッシュアルゴリズム
      - Diffie-Hellman グループ
        - Diffie-Hellman 鍵交換に使う離散対数の計算方法，およびその計算用パラメータの値
      - Diffie-Hellman 鍵交換の公開値
        - Diffie-Hellman 鍵交換でやり取りする乱数で公開鍵とも呼ばれる
    - エンティティ認証の方式
    - ID
      - ゲートウェイの ID
      - 事前鍵共有方式の場合，メインモードの時は IP アドレスのみ指定できる
      - アグレッシブモードの時は IP アドレス以外のものであるため，FQDN やユーザ名を指定できる
      - デジタル署名方式の場合，第三者認証を受けた電子証明書を相手に送る
    - ライフタイム
      - IKE SA の生存期間

- フェーズ 2 で IPsec SA のパラメータを交換し，これを生成する
- フェーズ 2 の通信は IKE SA を用いているのでセキュリティ処理(暗号化，メッセージ認証)が行われる
- フェーズ 2 で交換される IPsec SA のパラメータは上り用と下り用の単方向の通信で別々のものがある
- IPsec SA を生成してから IPsec 通信を開始する
- フェーズ 2 の通信はクイックモードと呼ばれ，1 往復半で行われる
  - IPsec 通信に関する合意
    - カプセル化モード，通信モード，IPsec 通信でカプセル化するオリジナルパケットの IP アドレスの範囲などを交換する
      - カプセル化モードは AH のみか，ESP のみか，AH と ESP の併用かのどれか
      - 通信モードはトンネルモードかトランスポートモード
  - 暗号化やメッセージ認証のアルゴリズムの合意
    - IPsec SA で使用される暗号化アルゴリズム，メッセージ認証アルゴリズムなどを交換する
  - 暗号鍵の生成
    - フェーズ 2 の通信は暗号化されており安全なためこの時生成する暗号鍵はナンスを交換するだけの低負荷な方法で生成される
    - ただし，PFS を有効にした場合に新たに Diffie-Hellman 鍵交換が行われる
  - IPsec SA の生成
    - パラメータを交換した後，IPsec SA を生成する
- フェーズ 2 で交換される主なパラメータ

  - カプセル化モード
    - IPsec 通信のカプセル化に使用するプロトコル
  - 通信モード
    - トンネルモードまたはトランスポートモード
  - ID
    - IPsec 通信でカプセル化するオリジナルパケットの IP アドレス範囲
      - IKEv1 ではん ID を用いてこの情報を交換する
  - 暗号化アルゴリズム
    - IPsec SA で使用する暗号化アルゴリズム
  - 認証アルゴリズム
  - Deffie-Hellman グループ
  - Deffie-Hellman 鍵交換の公開値
  - ライフタイム

- PFS
  - IKE SA の暗号鍵が破られると，IPsec SA では安全なやり取りができなくなり，ReKey してもやり取りを安全にできなくなる
  - そこで PFS を有効化すると，IPsec SA を生成するたびに新たに Diffie-Hellman 鍵交換を行い，安全に SA を生成する

## XAUTH

- IKEv1 には IPsec 通信を行う利用者の認証を行う仕組みが規定されていなったため XAUTH が追加された
- XAUTH で認証するには，まずフェーズ 1 の中でゲートウェイは互いに XAUTH をサポートしていることをつたえあう
- フェーズ 1 が完了すると，XAUTH 　による認証が行われ，認証に成功するとフェーズ 2 に入る
- しかし VPN つうしんでは IKE に頼らずとも利用者認証の仕組みを備えている

## IKEv2

- v1 との大きな違いは以下

  - フェーズ 1 が IKE_SA_INIT 交換にほぼ対応し，v1 のフェーズ 2 が v2 の IKE_AUTH 交換にほぼ対応している
  - エンティティ認証を行うタイミングが，v1 ではフェーズ 1 であったのに対して，v2 では IKE_AUTH 交換になっている
  - IKEv1 フェーズ 1 はメインモードとアグレッシブモードの 2 種類あったが v2 ではその区別がなくなっている

- 手順
  - IKE_SA_INIT 交換
    - 暗号化やメッセージ認証のアルゴリズムの合意
    - 暗号鍵の生成
    - IKE SA の生成
  - IKE_AUTH 交換
    - エンティティ認証
    - IPsec 通信に関する合意
      - v2 ではオリジナルパケットの範囲をトラフィックセレクタを用いて交換する
    - 暗号化やメッセージ認証のアルゴリズムの合意
    - 暗号化鍵の生成
    - IPsec SA の生成
  - その他の交換
    - CREATE_CHILD_SA 交換
      - すでに生成された IKE SA を使用して,Child SA を生成するために用いる
        - ReKey に伴い新たな Child SA を生成するときに用いられる
    - INFORMATIONAL 交換
      - エラーやステータスの通知，ネットワーク設定情報の通知，SA の削除などを行うために用いられる

## IPsec プロトコル(AH,ESP)

- IPsec SA が確立された後，ホストは IPsec パケットをやり取りする
- 具体的には AH(Authentication Header)ヘッダまたは ESP(Encapsulating Security Payload)ヘッダを付加する
- IP 層から見ると AH または ESP は上位層プロトコルに相当し，IP ヘッダ中のプロトコル番号は AH が 51 番を使用し，ESP が 50 番を使用している
- AH または ESP が上位層プロトコルに提供するセキュリティサービスはいか
  - メッセージ認証
    - AH はメッセージダイジェストを認証データとして AH ヘッダの中に格納している
      - AH の認証範囲に外側 IP ヘッダが含まれるため，NAT 機器経由の通信で支障が出る
      - 認証範囲は外側 IP ヘッダから全て
    - ESP は暗号化を行ってからメッセージ認証を行う
      - 認証データは ESP 認証データの中に格納する
      - 認証範囲は ESP ヘッダ から ESP トレーラまで
  - 暗号化
    - ESP のみ
    - 使用する暗号化アルゴリズムがブロック暗号方式であるとき，固定調のブロックが作成できないときのパディングとして末尾ブロックを合わせる
    - このパディングに相当するのが ESP トレーラ
    - 暗号化範囲は IP ペイロードとトレラー
    - トンネルモードの場合はオリジナル IP ヘッダまで暗号化する
  - リプレイ攻撃の拒否
    - AH ヘッダ，ESP ヘッダにはシーケンス番号が格納されており，送信側にはパケットを送信するたびにこれをカウントアップする
    - 受信側はこの番号の順序を確認することでリプレイ攻撃に対処できる
- AH も ESP も IPsec SA 生成時に交換したパラメータに基づきメッセージ認証や暗号化を行っている
- SA は SPI と呼ばれる番号で識別されており，AH ヘッダと ESP ヘッダには SPI が格納されている

## IPsec で IP マルチキャストパケットを直接カプセル化できない

- IKE の制約
  - IKE の鍵交換で生成される暗号鍵は共通鍵であり，多数には適応できない
- IPsec 通信のリプレイ攻撃防御の制約
  - 多数のカウントアップの同期を取るのが難しい

## IPsec の主な利用形態

- 拠点間を接続する VPN 通信
  - IPsec のトンネルモードにより生成された仮想的な専用線を使うやり方
  - この形態の場合各拠点のグローバル IP アドレスは固定なので IKEv1 を利用する場合はメインモードが利用される
- リモートアクセスによる VPN 通信
  - 拠点の IP アドレスが固定でなかったり，IP アドレスを固定するのが難しい端末などに適応する
  - 仮想的な専用線を構築することによって，リモートアクセス端末を接続先拠点の内部ネットワーク接続する携帯
  - リモートアクセス端末には VPN 接続用ソフトをインストールしておくことでゲートウェイの役割の担うことができるようになる
  - L2TP over IPsec がよく用いられる
    - 理由としては以下の 2 つを同時に行いたいから
      - 利用者の認証
      - リモート端末に対する接続先拠点のプライベート IP アドレスの払い出し
  - IPsec の通信の中で VPN 専用プロトコルの通信が行われる
    - この VPN 専用プロトコルと L2TP がよく用いられ，リモートアクセス端末と VPN サーバ間は L2TP over IPsec 通信が行われる
    - このとき IPsec トランスポートモードを用いて L2TP パケットをカプセル化している
    - また L2TP パケットは PPP パケットをカプセル化し，PPP パケットはオリジナル IP パケットをカプセル化している
    - つまり，最終的には PPP が使用されている
    - PPP は利用者の認証，リモート端末に対する接続先拠点のプライベート IP アドレスの払い出しを行う機能を有しているため，VPN サーバは利用者認証に成功したリモートアクセうす端末に対して，プライベート IP アドレスを払い出すこともできる

## NAT 環境下の IPsec の利用携帯

- ESP ヘッダにはポート番号が存在しないので NAPT が機能しない
- 問題を解決するための手法は NAT とラバーサルか，VPN パススルー
- NAT とラバーサルを使用するには NAT とラバーサルに対応した VPN クライアントソフトを用いる
  - これは新 IP ヘッダと ESP ヘッダの間に UDP ヘッダを挿入する仕組み
  - NAPT のポート番号変換にはこの UDP に対して行われる
  - NAT とラバーサルを使うにはまずは必要かどうかを判定する
    - そのために IKE 通信の段階でイニシエータ，レスポンダはそれぞれ，自分の IP アドレスとポート番号の組みから生成されたハッシュ値を通知し合う
      - IKE 通信の送信元と宛先のポート番号は 500
    - このハッシュ値が受信された際に再計算した値と異なっていると，NAPT で書き換えられたと判断して送信元と宛先のポート番号を 4500 万にして NAT とラバーサルを使用することを通知する
    - このように IKE の段階でイニシエータとレスポンダは NAT とラバーサルを使用することを事前に折衝する
- NAT とラバーサルを経て実施される IPsec 通信ではトンネルの送信側で UDP ヘッダを挿入し，トンネルの受信側でその UDP ヘッダを除去する動作が加わる
  - それ以外は同じ