## Quick Start
Try on [Demo] page. Or Install on your local PC. Check [Guide](#install-on-your-local-pc) to how to install.

## Getting Started

### Technology Stack
* Golang(1.9.3 darwin/amd64)
* MongoDB
* RabbitMQ

### Install On your local PC

#### Download Sources

use git

```bash
git clone https://github.com/gauravbansal74/mlserver.git
```

```bash
cd mlserver
```


#### Download dependencies

use [Golang/Dep](https://github.com/golang/dep)

```bash
dep ensure
```

#### Update Configuration

update server configuration file as per your local configuration [.mlserver.yml](https://github.com/gauravbansal74/mlserver/blob/master/.mlserver.yml)

```yml
server: #server configuration
  port: 8081        #server port
  host: "0.0.0.0"   #server host
  debug: false      #server debug mode on/off
  name: "mlserver"  #server log app name
mongo:                  #mongo db server connection details
  database: "mlserver"  #mongo database name
  url: "127.0.0.1"      #mongo host URL (Please don't use mongo culster URL - For cloud hosted mongo we need to provide master/slave details)
rabbitmq:                     #rabbitMQ server connection details
  host: "127.0.0.1"           #rabbitMQ host URL
  port: 5672                  #rabbitMQ port
  username: "guest"           #rabbitMQ username
  password: "guest"           #rabbitMQ password
  exchange: "fileProcessing"  #rabbitMQ exchange name to push and process messages
datasource:                       #datasource folder to watch if datasource file is added or not
  folder: "./datasource_folder"   #folder path to watch for data source file
```

#### Server Components
ID  |Server Component|Details
----|----|----
1   |Rest Service|Rest API HTTP server which is communicating with Front-end Application
2   |Folder Watcher|Watcher to watch on datasource folder and if any new file added then push message to RabbitMQ
3   |DataSource File Process|Get Message from RabbitMQ, Process DataSource File data and Save into DB "visits" collection.

#### clean go binaries

```bash
 go clean
```


#### build server
```bash
 go build
```

#### Run Server Components in different terminal tabs or using TMUX
ID  |Server Component|command to run
----|----|----
1   |Rest Service|```./mlserver --config=./.mlserver.yml server ```
2   |Folder Watcher|```./mlserver --config=./.mlserver.yml watcher start```
3   |RabbitMQ Message Process|```./mlserver --config=./.mlserver.yml watcher process```


#### Run Server using *.Sh files in different terminal tabs or using TMUX
ID  |Server Component|command to run
----|----|----
1   |Rest Service|```./run.sh```
2   |Folder Watcher|```./run-watcher.sh```
3   |DataSource File Process|```./run-watcher-process.sh```


#### use Rest Service End Point
```bash
 http://localhost:8081 #host and port from .mlserver.yml file
```
