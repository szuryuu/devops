#!/bin/bash

if [ ! -f /etc/nginx/ssl/certs/server.crt ]; then
    echo "SSL certificate not found, generating self-signed cert..."
    bash ./ssl/generate-ssl.sh
fi

nginx -g 'daemon off;'
