FROM nginx:stable

ARG APP_UID
ARG APP_GID

RUN mkdir -p /etc/nginx/sites-available &&\
    mkdir -p /etc/nginx/snippets &&\
    mkdir -p /etc/nginx/ssl/certs &&\
    mkdir -p /var/log/nginx &&\
    mkdir -p /ssl/certs

COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY conf/sites-available /etc/nginx/sites-available
COPY html /usr/share/nginx/html
COPY ssl/certs /etc/nginx/ssl/certs
COPY ssl/generate-ssl.sh /ssl/generate-ssl.sh
COPY entrypoint.sh /entrypoint.sh

RUN chmod 600 /etc/nginx/ssl/certs/server.key && \
    chmod 644 /etc/nginx/ssl/certs/server.crt && \
    chmod +x /entrypoint.sh && \
    chown -R root:root /etc/nginx/ssl/certs && \
    chown -R ${APP_UID}:${APP_GID} /etc/nginx/ssl /usr/share/nginx/html /var/log/nginx

# USER ${APP_UID}:${APP_GID}
EXPOSE 80 443

ENTRYPOINT ["/entrypoint.sh"]
