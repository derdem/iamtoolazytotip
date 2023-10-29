FROM --platform=arm64 node:21.1-alpine3.17

WORKDIR /app

COPY frontend ./
RUN npm install

CMD ["npm", "start"]
