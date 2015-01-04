# 1. For Server

### 1.1 Global json
    [
      "tokenStr1":{
        "username": "username1",
        "password": "password1",
        "signintime": "2014-12-31 23:59:59",
        "pushtime": "2014-12-31 23:59:59",
        "pushtoken":"pushtokenstr1"
      },
      "tokenStr2":{
        "username": "username2",
        "password": "password2",
        "signintime": "2014-12-31 23:59:59",
        "pushtime": "2014-12-31 23:59:59",
        "pushtoken":"pushtokenstr2"
      }
      ...
    ]

### 1.2 taskJson

    [
        "tokenStr1":{
            "task": "--- task1",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:59",
            "donetime":null,
            "updatetime":"2014-12-31 23:59:57",
            "status":"doing"
        },
        "tokenStr2":{
            "task": "--- task2",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:59",
            "updatetime":"2014-12-31 23:59:57",
            "donetime":null,
            "status":"doing"
        }
        ...
    ]


### 1.3 taskHistory
    [
        "tokenStr1":{
            "task": "--- task1",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:59",
            "donetime":null
        },
        "tokenStr2":{
            "task": "--- task2",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:59",
            "donetime":null
        }
        ...
    ]

# 2. For Client

### 2.1 Global json
    
    {
        "usertoken":"tokenStr",
        "pushtime":"2014-12-31 23:59:57",
        "pushtoken":"2014-12-31 23:59:59",
        "username":"username",
        "password":"password"
    }

### 2.2 taskJson

    [
        "taskToken1":{
            "task":"---task1...",
            "updatetime":"2014-12-31 23:59:57",
            "status":"doing",
        }
    ]
    