FROM harbor.apusic.com/apusic/centos:centos7-amd64
ADD  ginDemo/  /opt/ginDemo
WORKDIR   /opt/ginDemo
ENTRYPOINT ["/opt/ginDemo/gin-demo"]
CMD ["--contextPath=gin"]