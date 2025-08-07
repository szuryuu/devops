#!/bin/bash

openssl genrsa -out ./certs/server.key 2048
openssl req -new -key ./certs/server.key -out ./certs/server.csr \
        -subj "/C=ID/ST=YOGYAKARTA/L=SLEMAN/O=BOTIKA/CN=localhost"
openssl x509 -req -days 365 -in ./certs/server.csr -signkey ./certs/server.key -out ./certs/server.crt
