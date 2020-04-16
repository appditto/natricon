FROM node:11.3-alpine

ENV NODE_ENV=production
ENV HOST 0.0.0.0

ENV APP_ROOT /src

RUN mkdir ${APP_ROOT}
WORKDIR ${APP_ROOT}
ADD . ${APP_ROOT}

RUN npm install
RUN npm run build

# Expose the app port
EXPOSE 3000

CMD [ "npm", "start" ]