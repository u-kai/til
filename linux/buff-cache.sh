#!/bin/bash
#

free 

echo "make 1G"

dd if=/dev/zero of=testfile bs=1M count=1K

echo "ページキャッシュ獲得後のシステム全体のメモリ使用量を表示する"

free

echo "ファイル削除後、つまりページキャッシュ削除後のシステム全体のメモリ使用量を表示する"

rm testfile