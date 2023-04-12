   FROM postgres:latest

   # Set the working directory
   WORKDIR /docker-entrypoint-initdb.d

   # Copy the SQL script
   COPY initial.sql /docker-entrypoint-initdb.d/
   