#!/usr/bin/python3

import os
import sys

ret = os.fork()

if ret == 0:
    print("child process = {} parent process = {}".format(os.getpid(),os.getppid()))
    os.execv("/bin/echo",["echo","pid=hello from {}".format(os.getpid())])
    exit()

if ret > 0:
    print("child process = {} parent process = {}".format(ret,os.getpid()))
    exit()
    
sys.exit(1)