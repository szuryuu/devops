FROM debian:bookworm-slim

# Update the package lists and install necessary packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl bash cron ca-certificates coreutils gawk && \
    rm -rf /var/lib/apt/lists/*

COPY . .
RUN chmod +x ./entrypoint.sh ./scripts/*.sh

# Buat cron job file
RUN printf "* * * * * root /bin/bash /scripts/disk-monitor.sh >> /var/log/cron.log 2>&1\n" > /etc/cron.d/high-usage-check && \
    chmod 0644 /etc/cron.d/high-usage-check

# Jika perlu, pasang ke crontab root
RUN crontab /etc/cron.d/high-usage-check

# Buat file log
RUN touch /var/log/cron.log

# Expose the port your application listens on
EXPOSE 80

ENTRYPOINT ["/bin/bash", "-c", "./entrypoint.sh & cron && tail -f /var/log/cron.log"]
# ENTRYPOINT [ "./entrypoint.sh" ]
CMD ["echo", "Running Bash\n", "Jawa"]
