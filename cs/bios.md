- BIOS は周辺装置を制御するためのプログラム
- BIOS がないと周辺装置の回路図を調べてプログラムを全て作らないといけなくなる
- コンピュータを起動した時に一番初めに起動するプログラム
- 最初はハードウェアが正常に動くか BIOS がチェックする
- その次にブートストラップロードの起動

---

- BIOS はマザーボードの中の ROM に格納されている小さなプログラム
- マザーボードとは部品同士を連携させる役割がある

  - BIOS は必要最低限のものしかないので、マザーボードに置くことができる

- OS は巨大なソフトなのでハードディスクにある
- BIOS は起動したらハードディスクにある ブートストラップローダを起動する
- ブートストラップローダーが HD にある OS をメモリに読み込んで OS を起動する
- アプリケーション->OS->BIOS->ハードウェアという層がある
  - BIOS にハードウェアの操作をお願いする感じなんかな？