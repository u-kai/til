#!/bin/bash 
<<COMMENT 
demand-paging. py プロセス について 1 秒間 に 1 回 メモリ に関する 情報 を 出力 し ます。 
各行 の 先頭 には 情報 を 採取 し た 時刻 を 表示 し ます。 
その後 に 続く フィールドの 意味 は 以下 の 通り です。 
    第 1 フィールド: 獲得 済 メモリ 領域 の サイズ
    第 2 フィールド: 獲得 済 物理 メモリ の サイズ 
    第 3 フィールド: メジャー フォールト の 数 
    第 4フィールド: マイナー フォールト の 数 
COMMENT

PID=$(pgrep -f "demand-paging\.py") 
if [ -z "${PID}" ]; then
    echo "demand-paging. py プロセス が 存在 し ませ ん ので $0 より 先 に 起動 し て ください。" >&2 
    exit 1 
fi 
while true; do 
    DATE=$(date |tr -d '\n') 
# -h は ヘッダ を 出力し ない ため の オプション です。 
    INFO=$(ps -h -o vsz,rss,maj_flt,min_flt -p ${PID}) 
    if [ $? -ne 0 ]; then 
     echo "$DATE: demand-paging. py プロセス は 終了 し まし た。" >&2 
     exit 1 
    fi 
    echo "${DATE}:${INFO}"
    sleep 1 
done