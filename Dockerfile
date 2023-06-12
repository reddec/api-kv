FROM scratch
# 8080 for API
EXPOSE 8080/tcp
VOLUME /data
ENV API_KV_BIND=":8080" \
    API_KV_DIR="/data" \
    API_KV_TOKEN="changeme"
ENTRYPOINT ["/api-kv"]
RUN ["serve"]
ADD api-kv /