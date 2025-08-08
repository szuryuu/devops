FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl bash cron ca-certificates coreutils gawk && \
    rm -rf /var/lib/apt/lists/*

COPY . .

# Permissions
RUN chmod +x ./entrypoint.sh ./scripts/*.sh

RUN echo "* * * * * root /bin/bash /scripts/disk-monitor.sh >> /var/log/cron.log 2>&1" \
    > /etc/cron.d/high-usage-check && \
    chmod 0644 /etc/cron.d/high-usage-check && \
    crontab /etc/cron.d/high-usage-check && \
    touch /var/log/cron.log

CMD ["./entrypoint.sh"]
