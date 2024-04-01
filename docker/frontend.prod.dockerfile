FROM node:21.7.1-alpine3.18 as builder

WORKDIR /app
COPY --chown=node:node ./frontend/ /app/
RUN npm ci
RUN npm run build
EXPOSE 4173

# FROM node:21.7.1-alpine3.18 as runner
# WORKDIR /app
# COPY --from=builder /app/dist /app/dist
# COPY --from=builder /app/package.json /app/package.json
# COPY --from=builder /app/package-lock.json /app/package-lock.json
# RUN npm ci
CMD [ "/bin/sh", "-c", "npm run serve"]
