# DevOps Internship Exam

This repository contains the solutions for a DevOps internship exam. The challenges cover a range of essential DevOps skills and concepts.

## Project Goals

### 1. Linux Administration: User Management
- **Goal:** Write a Bash script to automate the creation of 1000 Linux users (`user-1001` to `user-1100`).
- **Details:**
    - Assign each user to a specific group ("testusers").
    - Set the UID for each user to match their index number.
    - Send a POST request with user details (username, UID, group, GID) to a webhook for each user created.

### 2. Linux Administration: System Monitoring
- **Goal:** Create a cron job that runs a Bash script every minute to monitor filesystem usage.
- **Details:**
    - If any filesystem usage reaches 90% or more, the script must send an alert via a POST request to a webhook.
    - The alert payload should be a JSON object containing details about the affected filesystem (type, total size, used space, free space, and mount point).

### 3. Programming & POSIX Signals
- **Goal:** Develop an HTTP server that can handle POSIX signals for graceful process management.
- **Details:**
    - **SIGHUP:** Reload the configuration without downtime.
    - **SIGTERM:** Perform a graceful shutdown, allowing existing requests to complete.
    - **SIGUSR1:** Re-open log files (for log rotation purposes).

### 4. NGINX Configuration
- **Goal:** Configure NGINX to act as a reverse proxy and serve static content with specific security and performance features.
- **Details:**
    - Set up a reverse proxy for the HTTP server created in the previous task.
    - Serve static content (images, audio) from a specific directory.
    - Implement a secure SSL configuration.
    - Enforce rate limiting on requests and connections for all virtual hosts.
    - (Bonus) Automate the setup using an Infrastructure as Code (IaC) tool like Ansible.

### 5. Containerization (Docker)
- **Goal:** Containerize the application and NGINX configuration using Docker.
- **Details:**
    - Use `docker-compose` to manage both the application server (from task 3) and the NGINX server (from task 4) as interconnected services.

---
*The following are optional challenges that were also part of the exam:*

### 6. Programming: Concurrency & Benchmarking (Optional)
- **Goal:** Build a command-line application to benchmark an HTTP server by sending a high volume of concurrent requests.
- **Details:**
    - The tool should allow configuration of concurrency level, total number of requests, and benchmark duration via CLI flags.

### 7. Containerization for Benchmark Tool (Optional)
- **Goal:** Containerize the benchmark application created in the previous optional task.

### 8. Network Programming: TCP Server & Parsing (Optional)
- **Goal:** Create a TCP server that parses incoming raw text as an HTTP message.
- **Details:**
    - The server should respond with a JSON payload containing the parsed components of the HTTP request (Method, URI, Version, Headers, and Body).
