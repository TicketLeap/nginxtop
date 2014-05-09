#!/bin/bash
tail -f /var/log/nginx/access.log | nodejs nginxtop.js -n 10
