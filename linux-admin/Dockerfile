FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl bash cron ca-certificates coreutils gawk && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN chmod +x ./entrypoint.sh ./scripts/*.sh
RUN printf "* * * * * root /bin/bash /scripts/disk-monitor.sh >> /var/log/cron.log 2>&1\n" > /etc/cron.d/high-usage-check && \
    chmod 0644 /etc/cron.d/high-usage-check
RUN crontab /etc/cron.d/high-usage-check
RUN touch /var/log/cron.log

EXPOSE 80

ENTRYPOINT ["/bin/bash", "-c", "./entrypoint.sh & cron && tail -f /var/log/cron.log"]
CMD ["echo", "Running Bash\n", "Jawa"]
