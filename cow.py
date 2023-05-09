#!/usr/bin/python3


import mmap
import os
import subprocess
import sys

ALLOC_SIZE = 100 * 1024 * 1024
PAGE_SIZE = 4 * 1024


def access(data):
    for i in range(0, ALLOC_SIZE, PAGE_SIZE):
        data[i] = 0


def show_meminfo(msg, process):
    print(msg)
    print("free コマンドの実行結果")
    subprocess.run("free")
    print("{}のメモリ関連情報".format(process))
    subprocess.run(["ps", "-orss,maj_flt,min_flt", str(os.getpid())])
    print()


data = mmap.mmap(-1, ALLOC_SIZE, mmap.MAP_PRIVATE)
access(data)
show_meminfo("子プロセス生成前", "親プロセス")
pid = os.fork()

if pid < 0:
    print("fork failed")
elif pid == 0:
    show_meminfo("子プロセス生成後", "子プロセス")
    access(data)
    show_meminfo("子プロセスによるメモリアクセス後", "子プロセス")
    sys.exit(0)
os.wait()
