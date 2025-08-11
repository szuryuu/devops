#!/bin/bash

echo "Running as user: $(id)"

if [ ! -f /etc/nginx/ssl/certs/server.crt ]; then
    echo "SSL certificate not found, generating self-signed cert..."
    bash ./ssl/generate-ssl.sh
fi

nginx -t
if [ $? -ne 0 ]; then
    echo "Nginx configuration error"
    exit 1
fi

echo $(ls -l /run/nginx.pid)
echo $(ls -ld /run)

exec nginx -g 'daemon off;'
