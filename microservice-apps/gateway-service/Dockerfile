FROM node:lts-alpine

# Create app directory
WORKDIR /usr/src/app

COPY package*.json ./

# Install app dependencies
RUN npm install

# Bundle app source
COPY server.js .

COPY products.proto .

EXPOSE 3000

CMD [ "node", "server.js" ]