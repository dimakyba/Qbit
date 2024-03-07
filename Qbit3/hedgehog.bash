#!/bin/bash

x=0
y=0
dx=0
dy=1

while IFS= read -r c
do
    if [[ -z "$c" ]]; then
        break
    fi
    if [[ "$c" == "R" ]]; then
        temp=$dx
        dx=$dy
        dy=$((-$temp))
        elif [[ "$c" == "L" ]]; then
        temp=$dx
        dx=$((-$dy))
        dy=$temp
    else
        x=$((x+dx))
        y=$((y+dy))
    fi
done

echo "$x $y"
