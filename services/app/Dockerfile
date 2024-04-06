FROM node:20.10.0-alpine

WORKDIR /app

# Copy package.json and package-lock.json
COPY . /app/

# Run the install 
RUN npm ci

# Build the application
RUN npm run build

# Install NGINX
RUN apk add --no-cache nginx

# Copy NGINX configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Copy static dist files to NGINX default directory
COPY dist /usr/share/nginx/html

# Expose port 80 for NGINX
EXPOSE 80

# Start NGINX server
CMD ["nginx", "-g", "daemon off;"]