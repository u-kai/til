#!/usr/bin/python3

import subprocess

size = 10000000

print("メモリ確保前のシステム全体のメモリ使用量を表示する")
subprocess.run("free")
array = [0]*size

print("メモリ獲得後のシステム全体のメモリ空き容量を表示する")
subprocess.run("free")