InfluxDB Proxy INSTALL
======
-----------
Usage
------------

```
git clone http://gitlab.chinac.com/OpsUltra/influx-proxy
go build 
./influx-proxy -c cfg.json
```

---------------
Configuration(示例:详细解释)
---------------

## influxdb-proxy 代理三台influxdb ##
三台influxdb所在节点假设为:test1,test2,test3
其中一台influxdb-proxy的配置如下(三台influxdb-proxy冗余配置一模一样)
```
{
    "backends" :{
        "test1":{
            "zone":"default",默认区域，用来做跨机房同步，这里默认为default
            "url": "http://10.21.1.220:9086", influxdb-proxy所在的节点test1,对外提供的url地址用于插入查询
            "db": "testDB", influxdb对外服务的db名称
            "interval": 1000,
            "timeout": 10000,
            "timeoutquery":600000,
            "maxrowlimit":10000,
            "checkinterval":1000,
            "rewriteinterval":10000
        },
        "test2" :{
            "zone":"default",
            "url": "http://10.21.1.221:9086",
            "db": "testDB",
            "interval": 1000,
            "timeout": 10000,
            "timeoutquery":600000,
            "maxrowlimit":10000,
            "checkinterval":1000,
            "rewriteinterval":10000
        },
        "test3" :{
            "zone":"default",
            "url": "http://10.21.1.222:9086",
            "db": "testDB",
            "interval": 1000,
            "timeout": 10000,
            "timeoutquery":600000,
            "maxrowlimit":10000,
            "checkinterval":1000,
            "rewriteinterval":10000
        }
    },
    "node" : {
        "listenaddr": ":8080",influxdb-proxy监听对外提供服务端口
        "db": "testDB", influxdb-proxy监听的influxdb的数据库，必须一一对应起来
        "zone": "default",influxdb-proxy所在的区域，必须和需要代理的influxdb在同一个区域
        "interval":10,
        "idletimeout":10,
        "writetracing":0,
        "querytracing":0
    },
    "measurements" : {
        "cpu": ["test1","test2"],influxdb-proxy所监听的表(cpu),对应到后端哪个influxdb节点的名称
        "weather": ["test1","test2"],
    }
}
```

## haproxy 配置代理influxdb-proxy##
示例配置如下:
```
listen influxdb_proxy
    bind 10.21.1.225:8080
    mode tcp
    balance  source
    option  tcpka
    option  tcplog
    server test1 10.21.1.220:8080 check inter 1s rise 3 fall 5
    server test1 10.21.1.221:8080 check inter 1s rise 3 fall 5
    server test1 10.21.1.222:8080 check inter 1s rise 3 fall 5

```
