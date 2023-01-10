#FROM envoyproxy/envoy-alpine:v1.21.6
FROM envoyproxy/envoy:v1.24.1

ARG envoy_user_id=0
ENV ENVOY_UID=$envoy_user_id
ENV ENVOY_GID=$envoy_user_id

RUN mkdir -p /app/logs/go-envoy && chown envoy:envoy /app/logs

RUN apt update && apt install -y logrotate