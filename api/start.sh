# !/bin/.sh
# MySQLサーバーが起動するまでループで待機する

until mysqladmin ping -h mysql --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "go app is started!"


