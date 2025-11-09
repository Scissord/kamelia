docker run --rm --network=nurdaulet_kong-net \
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-db" \
  -e "KONG_PG_USER=kong" \
  -e "KONG_PG_PASSWORD=kongpass" \
  kong:latest kong migrations bootstrap