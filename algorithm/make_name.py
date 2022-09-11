#!/usr/bin/env python3

import sys

if len(sys.argv) < 2:
    raise RuntimeError(f"usage: python3 {sys.argv[0]} <name>")

name = sys.argv[1]
parts = name.title().split()
print("".join(parts))
