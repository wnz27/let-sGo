# 1、redis-benchmark
## 本机信息
```
  型号名称：	MacBook Pro
  型号标识符：	MacBookPro14,1
  处理器名称：	Intel Core i5
  处理器速度：	2.3 GHz
  处理器数目：	1
  核总数：	2
  L2 缓存（每个核）：	256 KB
  L3 缓存：	4 MB
  超线程技术：	已启用
  内存：	8 GB
  Boot ROM 版本：	205.0.0.0.0
  SMC 版本（系统）：	2.43f7
  序列号（系统）：	FVFY85W9HV29
  硬件 UUID：	4172700F-05C3-5B6A-A2DA-04E0C3A3A10F
```
## redis-benchmark help
键入： `benchmark --help` ：
```shell
Usage: redis-benchmark [-h <host>] [-p <port>] [-c <clients>] [-n <requests>] [-k <boolean>]

 -h <hostname>      Server hostname (default 127.0.0.1)
 -p <port>          Server port (default 6379)
 -s <socket>        Server socket (overrides host and port)
 -a <password>      Password for Redis Auth
 --user <username>  Used to send ACL style 'AUTH username pass'. Needs -a.
 -c <clients>       Number of parallel connections (default 50)
 -n <requests>      Total number of requests (default 100000)
 -d <size>          Data size of SET/GET value in bytes (default 3)
 --dbnum <db>       SELECT the specified db number (default 0)
 --threads <num>    Enable multi-thread mode.
 --cluster          Enable cluster mode.
 --enable-tracking  Send CLIENT TRACKING on before starting benchmark.
 -k <boolean>       1=keep alive 0=reconnect (default 1)
 -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD,
                    random members and scores for ZADD.
  Using this option the benchmark will expand the string __rand_int__
  inside an argument with a 12 digits number in the specified range
  from 0 to keyspacelen-1. The substitution changes every time a command
  is executed. Default tests use this to hit random keys in the
  specified range.
 -P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
 -q                 Quiet. Just show query/sec values
 --precision        Number of decimal places to display in latency output (default 0)
 --csv              Output in CSV format
 -l                 Loop. Run the tests forever
 -t <tests>         Only run the comma separated list of tests. The test
                    names are the same as the ones produced as output.
 -I                 Idle mode. Just open N idle connections and wait.
 --help             Output this help and exit.
 --version          Output version and exit.

Examples:

 Run the benchmark with the default configuration against 127.0.0.1:6379:
   $ redis-benchmark

 Use 20 parallel clients, for a total of 100k requests, against 192.168.1.1:
   $ redis-benchmark -h 192.168.1.1 -p 6379 -n 100000 -c 20

 Fill 127.0.0.1:6379 with about 1 million keys only using the SET test:
   $ redis-benchmark -t set -n 1000000 -r 100000000

 Benchmark 127.0.0.1:6379 for a few commands producing CSV output:
   $ redis-benchmark -t ping,set,get -n 100000 --csv

 Benchmark a specific command line:
   $ redis-benchmark -r 10000 -n 10000 eval 'return redis.call("ping")' 0

 Fill a list with 10000 random elements:
   $ redis-benchmark -r 10000 -n 10000 lpush mylist __rand_int__

 On user specified command lines __rand_int__ is replaced with a random integer
 with a range of values selected by the -r option.
```
### 单机 -c 1  单客户端10000 请求 -n 10000 值大小10字节 -d 10
`redis-benchmark -t set,get -d 10 -c 1 -n 10000`
```shell
====== SET ======
  10000 requests completed in 11.94 seconds
  1 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.599 milliseconds (cumulative count 1)
50.000% <= 1.103 milliseconds (cumulative count 5170)
75.000% <= 1.231 milliseconds (cumulative count 7570)
87.500% <= 1.351 milliseconds (cumulative count 8775)
93.750% <= 1.543 milliseconds (cumulative count 9387)
96.875% <= 1.959 milliseconds (cumulative count 9691)
98.438% <= 2.535 milliseconds (cumulative count 9845)
99.219% <= 3.063 milliseconds (cumulative count 9922)
99.609% <= 3.767 milliseconds (cumulative count 9961)
99.805% <= 4.735 milliseconds (cumulative count 9981)
99.902% <= 6.447 milliseconds (cumulative count 9991)
99.951% <= 9.975 milliseconds (cumulative count 9996)
99.976% <= 12.375 milliseconds (cumulative count 9998)
99.988% <= 12.407 milliseconds (cumulative count 9999)
99.994% <= 72.127 milliseconds (cumulative count 10000)
100.000% <= 72.127 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 0.607 milliseconds (cumulative count 1)
0.150% <= 0.703 milliseconds (cumulative count 15)
2.150% <= 0.807 milliseconds (cumulative count 215)
9.880% <= 0.903 milliseconds (cumulative count 988)
29.970% <= 1.007 milliseconds (cumulative count 2997)
51.700% <= 1.103 milliseconds (cumulative count 5170)
72.010% <= 1.207 milliseconds (cumulative count 7201)
84.310% <= 1.303 milliseconds (cumulative count 8431)
90.390% <= 1.407 milliseconds (cumulative count 9039)
93.170% <= 1.503 milliseconds (cumulative count 9317)
94.650% <= 1.607 milliseconds (cumulative count 9465)
95.660% <= 1.703 milliseconds (cumulative count 9566)
96.220% <= 1.807 milliseconds (cumulative count 9622)
96.670% <= 1.903 milliseconds (cumulative count 9667)
97.120% <= 2.007 milliseconds (cumulative count 9712)
97.430% <= 2.103 milliseconds (cumulative count 9743)
99.280% <= 3.103 milliseconds (cumulative count 9928)
99.690% <= 4.103 milliseconds (cumulative count 9969)
99.840% <= 5.103 milliseconds (cumulative count 9984)
99.880% <= 6.103 milliseconds (cumulative count 9988)
99.910% <= 7.103 milliseconds (cumulative count 9991)
99.940% <= 9.103 milliseconds (cumulative count 9994)
99.970% <= 10.103 milliseconds (cumulative count 9997)
99.990% <= 13.103 milliseconds (cumulative count 9999)
100.000% <= 72.127 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 837.31 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.176     0.592     1.103     1.647     2.887    72.127
====== GET ======
  10000 requests completed in 12.69 seconds
  1 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.559 milliseconds (cumulative count 1)
50.000% <= 1.175 milliseconds (cumulative count 5013)
75.000% <= 1.319 milliseconds (cumulative count 7531)
87.500% <= 1.455 milliseconds (cumulative count 8774)
93.750% <= 1.647 milliseconds (cumulative count 9379)
96.875% <= 2.055 milliseconds (cumulative count 9688)
98.438% <= 2.591 milliseconds (cumulative count 9845)
99.219% <= 2.975 milliseconds (cumulative count 9923)
99.609% <= 3.399 milliseconds (cumulative count 9961)
99.805% <= 4.543 milliseconds (cumulative count 9981)
99.902% <= 11.351 milliseconds (cumulative count 9991)
99.951% <= 14.215 milliseconds (cumulative count 9996)
99.976% <= 16.463 milliseconds (cumulative count 9998)
99.988% <= 20.559 milliseconds (cumulative count 9999)
99.994% <= 50.271 milliseconds (cumulative count 10000)
100.000% <= 50.271 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 0.607 milliseconds (cumulative count 1)
0.270% <= 0.703 milliseconds (cumulative count 27)
1.860% <= 0.807 milliseconds (cumulative count 186)
6.600% <= 0.903 milliseconds (cumulative count 660)
18.340% <= 1.007 milliseconds (cumulative count 1834)
35.640% <= 1.103 milliseconds (cumulative count 3564)
56.650% <= 1.207 milliseconds (cumulative count 5665)
72.980% <= 1.303 milliseconds (cumulative count 7298)
84.500% <= 1.407 milliseconds (cumulative count 8450)
89.910% <= 1.503 milliseconds (cumulative count 8991)
92.880% <= 1.607 milliseconds (cumulative count 9288)
94.590% <= 1.703 milliseconds (cumulative count 9459)
95.500% <= 1.807 milliseconds (cumulative count 9550)
96.100% <= 1.903 milliseconds (cumulative count 9610)
96.580% <= 2.007 milliseconds (cumulative count 9658)
97.100% <= 2.103 milliseconds (cumulative count 9710)
99.400% <= 3.103 milliseconds (cumulative count 9940)
99.780% <= 4.103 milliseconds (cumulative count 9978)
99.830% <= 5.103 milliseconds (cumulative count 9983)
99.840% <= 6.103 milliseconds (cumulative count 9984)
99.870% <= 7.103 milliseconds (cumulative count 9987)
99.880% <= 8.103 milliseconds (cumulative count 9988)
99.890% <= 9.103 milliseconds (cumulative count 9989)
99.900% <= 10.103 milliseconds (cumulative count 9990)
99.910% <= 12.103 milliseconds (cumulative count 9991)
99.930% <= 13.103 milliseconds (cumulative count 9993)
99.950% <= 14.103 milliseconds (cumulative count 9995)
99.960% <= 15.103 milliseconds (cumulative count 9996)
99.970% <= 16.103 milliseconds (cumulative count 9997)
99.980% <= 17.103 milliseconds (cumulative count 9998)
99.990% <= 21.103 milliseconds (cumulative count 9999)
100.000% <= 51.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 788.21 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.251     0.552     1.175     1.743     2.895    50.271
```
### 单机 -c 1  单客户端10000 请求 -n 10000 值大小20字节 -d 10
```shell
====== SET ======
  10000 requests completed in 13.22 seconds
  1 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
  
Summary:
  throughput summary: 756.49 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.304     0.600     1.159     1.767     2.871   629.759

====== GET ======
  10000 requests completed in 12.89 seconds
  1 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 775.86 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.271     0.640     1.167     1.895     3.615    62.655
```

### 单机 -c 1  单客户端10000 请求 -n 10000 值大小50字节 -d 10
```shell
====== SET ======
  10000 requests completed in 12.88 seconds
  1 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 776.40 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.270     0.608     1.215     1.807     2.863    44.959
        
====== GET ======
  10000 requests completed in 13.26 seconds
  1 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 753.92 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.308     0.560     1.103     2.407     4.183    92.223

```
### 单机 -c 1  单客户端10000 请求 -n 10000 值大小100字节 -d 10
```shell
====== SET ======
  10000 requests completed in 13.95 seconds
  1 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
  
Summary:
  throughput summary: 716.79 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.376     0.528     1.135     2.607     4.751    79.807
     
====== GET ======
  10000 requests completed in 15.78 seconds
  1 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 633.79 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.559     0.528     1.159     2.927     8.567   199.423

```
## 单机 -c 1  单客户端10000 请求 -n 10000 值大小200字节 -d 10
```shell
====== SET ======
  10000 requests completed in 14.53 seconds
  1 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
  
Summary:
  throughput summary: 688.18 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.434     0.536     1.159     2.695     5.495    70.207
        

====== GET ======
  10000 requests completed in 15.72 seconds
  1 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
  
Summary:
  throughput summary: 636.13 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.552     0.632     1.247     3.071     5.559   127.039
```



# 2、 [内存占用测试](./redis_learn_test.go)
