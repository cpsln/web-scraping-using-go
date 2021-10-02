### Step Setup:
This is a Guide for setup for this application.


1. Installing Go 1.14.4.
    <pre>
        $ wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
        $ sudo tar -xvf go1.14.4.linux-amd64.tar.gz
        $ sudo mv go /usr/local
        $ echo 'export GOROOT=/usr/local/go' >>~/.profile
        $ echo 'export GOPATH=$HOME/go_projects' >>~/.profile
        $ echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >>~/.profile
        $ source ~/.profile
    </pre>

2. Mysql installation and user creation:

    <pre>    
        $ sudo apt-get install mysql-server
        $ sudo mysql
    </pre>

3. Create user for ALL GRANT.

    <pre>        
        mysql> CREATE USER 'root'@'%' IDENTIFIED BY 'root';
        mysql> GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
        mysql> FLUSH PRIVILEGES;
        mysql> exit

        $ mysql -u root -p
        Enter Password

        mysql> CREATE DATABASE IF NOT EXISTS testdb;
        mysql>
        Create TABLE `product_details` (
            `id` int NOT NULL AUTO_INCREMENT,
            `product_name` varchar(250) DEFAULT NULL, 
            `image_url` varchar(350) DEFAULT NULL, 
            `description` varchar(1500) DEFAULT NULL,
            `price` varchar(10) DEFAULT NULL,
            `total_review` varchar(20) DEFAULT NULL, 
            PRIMARY KEY (id)
        );
    </pre>

3. Run command.
    <pre>
        $ go mod tidy
        $ go run main.go
    </pre>


### Note.
1. Send url in the body, It scrap data and store into db
    <pre>
        localhost:8000/scrap -> POST -> body{"url":""}
    </pre>

2. Get the scrap product data by id.
    <pre>
        localhost:8000/get?pid=1 ->GET
    </pre>

3. Update scrap product data.
    <pre>
        localhost:8000/update -> PUT -> body(send json as GET api response)
    </pre>