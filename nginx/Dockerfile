FROM nginx:stable

RUN mkdir -p /etc/nginx/sites-available \
    && mkdir -p /etc/nginx/snippets \
    && mkdir -p /etc/nginx/ssl/certs \
    && mkdir -p /var/log/nginx

COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY conf/sites-available /etc/nginx/sites-available
COPY conf/snippets /etc/nginx/snippets
COPY html /usr/share/nginx/html
COPY ssl/certs /etc/nginx/ssl/certs
COPY entrypoint.sh /entrypoint.sh

RUN chmod 755 /usr/share/nginx/*
RUN chown -R nginx:nginx /usr/share/nginx/*

RUN chmod +x /entrypoint.sh

EXPOSE 80 443

ENTRYPOINT ["/entrypoint.sh"]
