#!/usr/bin/env bash

openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout nginx/certs/dev.key \
  -out nginx/certs/dev.crt \
  -subj "/CN=localhost"

