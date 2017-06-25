#!/bin/sh

./db_load \
	--infile f1.xlsx \
	--host 127.0.0.1 \
	--port 3306 \
	--user user \
	--password pass \
	--database population \
	$@
