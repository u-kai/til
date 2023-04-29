#!/usr/bin/python3

import os
import sys

ret = os.fork()

if ret == 0 :
    print("親プロセス")
    print("子プロセス:pid={},親プロセス:pid = {}".format(os.getpid(),os.getppid()))
    exit()

if ret > 0 :
    print("子プロセス")
    print("子プロセス:pid={},親プロセス:pid = {}".format(ret,os.getpid()))
    exit()

sys.exit(1)