FROM nginx:stable

RUN groupadd -g 1001 testusers && \
    useradd -u 1001 -g testusers user-1001

RUN mkdir -p /etc/nginx/sites-available &&\
    mkdir -p /etc/nginx/snippets &&\
    mkdir -p /etc/nginx/ssl/certs &&\
    mkdir -p /var/log/nginx &&\
    mkdir -p /ssl/certs

COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY conf/sites-available /etc/nginx/sites-available
COPY conf/snippets /etc/nginx/snippets
COPY html /usr/share/nginx/html
COPY ssl/certs /etc/nginx/ssl/certs
COPY ssl/generate-ssl.sh /ssl/generate-ssl.sh
COPY entrypoint.sh /entrypoint.sh

RUN chown -R user-1001:testusers /etc/nginx/ssl/certs && \
    chmod 600 /etc/nginx/ssl/certs/server.key && \
    chmod 644 /etc/nginx/ssl/certs/server.crt && \
    chown -R user-1001:testusers /usr/share/nginx/html && \
    chown -R user-1001:testusers /var/log/nginx && \
    chmod +x /entrypoint.sh

EXPOSE 80 443

ENTRYPOINT ["/entrypoint.sh"]
