#!/bin/bash

# Start ollama server in the background
ollama serve &

# Wait for the server to be ready
timeout=30
counter=0
echo "Waiting for Ollama server to be ready..."
while ! curl -s http://localhost:11434/api/version >/dev/null; do
    sleep 1
    counter=$((counter + 1))
    if [ $counter -ge $timeout ]; then
        echo "Timeout waiting for Ollama server"
        exit 1
    fi
done

echo "Ollama server is ready. Pulling model..."

# Pull the model
ollama pull deepseek-r1:1.5b

# Keep the container running by waiting for the background process
wait $! 