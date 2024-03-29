# SSL-VPN

- TLS プロトコルを利用した VPN 技術
- 主に 3 つの動作方式がある
  - リバースプロキシ方式
    - 専用モジュールが不要
    - ブラウザで動作できるアプリケーションに限定される
  - ポートフォワーディング方式
    - 専用モジュールが必要
      - 専用モジュールが PC の hosts ファイルを書き換え，AP サーバのホスト名に対応する IP アドレスとしてループバックアドレスを登録する
      - その後モジュールは AP 通信ようポートを宛先とする通信を待ち受ける
    - VPN 装置が PC から受信するパケットと，VPN 装置が AP サーバに送信するパケットのそれぞれの宛先ポート番号を比較すると前者は 443 で後者は AP 通信ようのポート番号であり，ポート番号が変化していることからこの名前がつけられている
    - ポートフォワーディング方式では AP サーバが決まればポートフォワーディングするポート番号が決まるので，このマッピングを VPN 装置に事前に登録しておき，実行時に独自ヘッダなどで判断し，参照する仕組みになっている
    - ポート番号が実行時に変化しないアプリケーションに限定される
  - L2 フォワーディング方式
    - 専用モジュールが必要
    - アプリケーションに制限がない
    - VPN 装置の役割はデフォルトゲートウェイであったり，L2 スイッチであったりする
    - L2 フォワーディングは仮想的に見ると同じ L2 ネットワークに VPN 装置と PC が存在する仕組みになっている
    - PC が VPN 装置に接続すると，専用モジュールの働きにより，PC 内部に仮想 NIC が構築される
      - さらに VPN 装置から VPN 接続用ネットワークの IP アドレスが動的に割り当てられる
      - 仮想 NIC を経由する通信はすべて専用モジュールが受け取る
      - 専用モジュールは受け取ったイーサネットフレーム全体を暗号化して VPN 装置に転送する仕組みになっている
      - 具体的な手順は以下
        - 仮想 NIC を指定し，AP サーバ宛に AP 通信を行う
        - 仮想 NIC を経由する通信は専用モジュールが受け取る
        - VPN 装置に転送するために，イーサネットフレーム全体を TLS でカプセル化し，暗号化する
        - VPN 装置宛に TLS 通信を行う
        - TLS の複合とカプセル化解除を行なってイーサネットフレームを取り出し，AP サーバに転送する
    - この方式はリバースプロキシ方式やポートフォワーディング方式と異なり， IP に限らずさまざまな L3 パケットを転送することができる
    - さらにポートフォワーディング方式と異なり，実行時にポート番号が変化する通信も転送することができる

## SoftEther
