Follow this tutorial for linux / ubuntu 

https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04

1. First update package and install mysql server

sudo apt update

sudo apt install mysql-server

2. Check if service start or not

sudo systemctl start mysql.service

3. Configure mysql

sudo mysql 

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';

exit

mysql_secure_installation

// change password or reset user, privilege

mysql -u root -p

ALTER USER 'root'@'localhost' IDENTIFIED WITH auth_socket; // this mean next time to connect mysql, should use 'sudo mysql'

ALTER USER 'sammy'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password'; // for connect to mysql which not use sudo anymore

4. Create a new user for application only

sudo mysql

CREATE USER 'username'@'host' IDENTIFIED BY 'password';

5. grant privileges

GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* TO 'username'@'localhost' WITH GRANT OPTION;

6. Now can connect from main.go: 

go run .

7. allow remote access with ubuntu firewall 

sudo ufw enable

sudo ufw allow mysql

8. Install Workbench

sudo apt install mysql-workbench-community

9. Launch workbench

mysql-workbench

