#!/bin/bash
false &
wait $!
echo $?
