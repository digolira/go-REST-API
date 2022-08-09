FROM mysql

#MySQL docker container executes script files from /docker-entrypoint-initdb.d/ folder in the ascending order of file names
COPY DDL_CreateTables.sql /docker-entrypoint-initdb.d/
COPY DML_InsertFakeData.sql /docker-entrypoint-initdb.d/


EXPOSE 3306