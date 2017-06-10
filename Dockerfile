FROM scratch
MAINTAINER Jon Borgonia "jon@gomagames.com"
EXPOSE 80
COPY bin/main /
COPY static /static
CMD ["/main"]
