FROM locustio/locust:1.1

# Copy task files
COPY tasks /tasks

# Start Locust 
CMD ["locust", "-f", "/tasks/main.py" ] 
