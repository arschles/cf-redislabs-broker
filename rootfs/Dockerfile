FROM quay.io/deis/base:0.2.0
COPY bin/broker /bin/broker
COPY conf/config.yml /conf/config.yml
EXPOSE 8080
CMD ["/bin/broker", "-c", "/conf/config.yml"]
