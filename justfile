default:
	@just --list

new year day:
	#!/usr/bin/env bash
	set -euxo pipefail
	y=20{{year}}
	d=day{{day}}
	mkdir -p $y/$d
	cp -R template/ $y/$d
	rm $y/$d/.gitignore
