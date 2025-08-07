#!/bin/bash

openssl genrsa -out ./ssl/certs/server.key 2048
openssl req -new -key ./ssl/certs/server.key -out ./ssl/certs/server.csr \
        -subj "/C=ID/ST=YOGYAKARTA/L=SLEMAN/O=BOTIKA/CN=localhost"
openssl x509 -req -days 365 -in ./ssl/certs/server.csr -signkey ./ssl/certs/server.key -out ./ssl/certs/server.crt
