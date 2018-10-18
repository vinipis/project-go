#!/usr/bin/python
# -*- coding: utf-8 -*-

f = open('sample.txt', 'r')
fields = {}

for line in f.readlines():
    line = line.strip()
    (name, value) = line.split('^')
    fields[name] = value

f.close
print fields
