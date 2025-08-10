#!/bin/bash

if [ ! -f /etc/nginx/ssl/certs/server.crt ]; then
    echo "SSL certificate not found, generating self-signed cert..."
    bash ./ssl/generate-ssl.sh
fi

nginx -t

if [ $? -eq 0 ]; then
    echo "Nginx configuration is valid"
    nginx -g 'daemon off;'
else
    echo "Nginx configuration error"
    exit 1
fi
