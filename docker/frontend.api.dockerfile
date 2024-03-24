FROM node:21.7.1-alpine3.18 as builder

WORKDIR /app
COPY ./frontend/ /app/
RUN npm ci --production
RUN npm run build

FROM node:21.7.1-alpine3.18 as runner
COPY --from=builder /app/dist /app/dist
CMD [ "/bin/sh", "-c", "npm run preview"]
