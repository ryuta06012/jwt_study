#FROM mysql:5.7
FROM mysql/mysql-server:5.7.28

RUN mkdir /var/log/mysql

RUN touch /var/log/mysql/mysqld.log

CMD ["mysqld"]