FROM openjdk:8

# Create dirs
RUN mkdir -p /opt/liquibase &&\
    mkdir -p /opt/jdbcdrivers/lib &&\
    mkdir -p /home/rundir &&\
    mkdir /scripts

# Add liquibase
ADD https://repo1.maven.org/maven2/org/liquibase/liquibase-core/3.6.3/liquibase-core-3.6.3-bin.tar.gz /opt/liquibase/liquibase-core-bin.tar.gz
WORKDIR /opt/liquibase
RUN tar -xzf liquibase-core-bin.tar.gz &&\
    rm liquibase-core-bin.tar.gz &&\
    chmod +x /opt/liquibase/liquibase &&\
    ln -s /opt/liquibase/liquibase /usr/local/bin/liquibase

WORKDIR /

# Add JDBC drivers
ENV DRIVER_VERSION 42.2.6
ENV DRIVER_JAR postgresql-${DRIVER_VERSION}.jar
ADD https://repo1.maven.org/maven2/org/postgresql/postgresql/${DRIVER_VERSION}/${DRIVER_JAR} /opt/jdbcdrivers/lib/${DRIVER_JAR}
RUN chmod 644 /opt/jdbcdrivers/lib/${DRIVER_JAR}

# Install JQ
RUN apt-get update &&\
    apt-get install -y jq

# Copy changelogs
COPY ./schema/changelogs /changelogs
COPY ./schema/scripts /scripts

WORKDIR /home/rundir

CMD cd /home/rundir && /scripts/apply_schema.sh