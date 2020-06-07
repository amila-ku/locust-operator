FROM python:3-slim

# Update and install dependencies
RUN apt update && apt install -y build-essential

# Install locust
RUN pip install locustio==0.14.6

# Copy task files
COPY tasks /tasks

# Set script to be executable
RUN chmod 755 /tasks/main.py

# Expose the required Locust ports
EXPOSE 5557 5558 8089

# Start Locust 
CMD ["locust", "-f", "tasks/main.py" ] 