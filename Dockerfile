FROM traefik:2.9.6

# COPY *.yml *.mmdb go.* ./geoip.go /plugins/go/src/github.com/GiGInnovationLabs/traefikgeoip2/
# COPY vendor/ /plugins/go/src/github.com/GiGInnovationLabs/traefikgeoip2/vendor/

# COPY GeoLite2-City.mmdb /var/lib/traefikgeoip2/

COPY . plugins-local/src/github.com/forestvpn/traefikgeoip2


