#!/usr/bin/python3

import datetime
import mmap
import time

ALLOC_SIZE = 100 * 1024 * 1024
ACCESS_SIZE = 10 * 1024 * 1024
PAGE_SIZE = 4 * 1024


def show_message(msg):
    print("{}: {}".format(datetime.datetime.now(), msg))


memregion = mmap.mmap(-1, ALLOC_SIZE, flags=mmap.MAP_PRIVATE)
show_message("Memory allocated, push enter")

input()

for i in range(0, ALLOC_SIZE, PAGE_SIZE):
    memregion[i] = 0
    if i % ACCESS_SIZE == 0 and i != 0:
        show_message("Accessed {} bytes".format(i // (1024 * 1024)))
        time.sleep(1)


show_message("Memory accessed All, push enter")
input()
