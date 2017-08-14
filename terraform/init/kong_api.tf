resource "kong_api" "admin" {
  name               = "kong-admin"
  upstream_url       = "http://localhost:8001"
  hosts              = "localhost"
  strip_uri          = true
  retries            = 5
  preserve_host      = true
  https_only         = false
  http_if_terminated = true
}

resource "kong_consumer" "admin" {
  username  = "localz"
}

resource "kong_api_plugin" "admin_basic_auth" {
  api = "${kong_api.admin.id}"
  name = "basic-auth"
}

resource "kong_consumer_basic_auth_credential" "basic_auth_credential" {
  consumer = "${kong_consumer.admin.id}"
  username = "localz"
  password = "password"
}

