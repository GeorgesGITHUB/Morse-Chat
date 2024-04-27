# Note on the multi-stage build process:
#   By using multi-stage builds, it results in smaller Docker images and more efficient deployments.
#   Larger node base image only used when building and smaller one only used during runtime
#   Ensuring that final image only contains necessary files to run the application, 
#   Development dependencies or build artifacts, useful during build, are discarded
#   Result: Smaller image size, performant, more secure 

# Stage 1: Build Environment
# Use official Node.js image as base for best compatibility
FROM node:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to take advantage of Docker layer caching
COPY package*.json ./

# Install dependencies, Docker layer caching is used if dependencies haven't changed
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Production Environment
# Serve the application with a lightweight Node.js image
FROM node:alpine

# Set the working directory in the container
WORKDIR /app

# Copy the build output from the previous stage
COPY --from=build /app/dist .

# Install serve to run the built application
RUN npm install -g serve

# Expose port 5173 that React app listens to, allowing external access to app when container running.
EXPOSE 5173

# Command to Start server and serve the built React app. 
# -s flag specifies directory to serve
# -l flag specifies port to listen on.
CMD ["serve", "-s", ".", "-l", "5173"]
