FROM node:22-alpine

# quicktype is pinned: a different version may change the generated code.
ARG QUICKTYPE_VERSION=23.2.6

RUN npm install -g "quicktype@${QUICKTYPE_VERSION}" \
 && npm cache clean --force

COPY --chmod=755 gen.sh /usr/local/bin/gen

WORKDIR /work
ENTRYPOINT ["gen"]
