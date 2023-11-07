FROM node:18-alpine

WORKDIR /fractals/

COPY public ./public
COPY src ./src
COPY package.json ./package.json

RUN npm install

CMD ["npm", "start"]


