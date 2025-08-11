FROM nginx:stable

ARG APP_UID
ARG APP_GID

ENV APP_UID=${APP_UID}
ENV APP_GID=${APP_GID}

RUN groupadd --gid ${APP_GID} testusers && \
    useradd --uid ${APP_UID} --gid ${APP_GID} --shell /bin/false user-${APP_UID}

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

RUN chmod +x /entrypoint.sh && \
    chmod 644 /etc/nginx/ssl/certs/server.crt && \
    chmod 640 /etc/nginx/ssl/certs/server.key && \
    chown -R ${APP_UID}:${APP_GID} /etc/nginx/ssl /usr/share/nginx/html /var/log/nginx /var/cache/nginx /var/run /run

EXPOSE 80 443
# USER ${APP_UID}:${APP_GID}

ENTRYPOINT ["/entrypoint.sh"]
