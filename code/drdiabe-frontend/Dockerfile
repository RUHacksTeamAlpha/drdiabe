# FROM node:latest
# COPY package*.json ./
# RUN yarn install
# COPY . .
# RUN yarn build
# EXPOSE 8080
# CMD ["yarn", "start"]

# FROM node:latest
# COPY package*.json ./
# RUN yarn install
# COPY nuxt.config.js ./
# COPY .nuxt ./
# # RUN yarn install
# # COPY . .
# # RUN yarn build
# EXPOSE 8080
# CMD ["yarn", "start"]

# Dockerfile
FROM node:latest

# create destination directory
RUN mkdir -p /usr/src/nuxt-app
WORKDIR /usr/src/nuxt-app

# copy the app, note .dockerignore
COPY . /usr/src/nuxt-app/
RUN yarn install
COPY .nuxt .

EXPOSE 8080

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=8080

CMD [ "yarn", "start" ]