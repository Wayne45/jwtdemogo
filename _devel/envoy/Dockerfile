FROM envoyproxy/envoy:v1.24.1

ARG envoy_user_id=1001
ENV ENVOY_UID=$envoy_user_id
ENV ENVOY_GID=$envoy_user_id

RUN mkdir -p /app/logs/go-envoy && chown envoy:envoy /app/logs