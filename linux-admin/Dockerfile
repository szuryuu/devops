FROM debian:bookworm-slim

# Update the package lists and install necessary packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl bash cron ca-certificates coreutils gawk && \
    rm -rf /var/lib/apt/lists/*

COPY . .
RUN chmod +x ./entrypoint.sh
RUN chmod +x ./scripts/*.sh

# Expose the port your application listens on
EXPOSE 80

ENTRYPOINT [ "./entrypoint.sh" ]
CMD ["echo", "Running Bash\n", "Jawa"]
