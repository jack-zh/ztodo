# 1. For Server

### 1.1 Global json
    {
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
    }

### 1.2 worktaskJson

    [
        {
            "task": "--- task1",
            "token":"tokenstr1",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:59",
            "donetime":"2006-01-02 15:04:05",
            "updatetime":"2014-12-31 23:59:57",
            "status":"doing"
        },
        {
            "task": "--- task2",
            "token":"tokenstr2",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:59",
            "updatetime":"2014-12-31 23:59:57",
            "donetime":"2006-01-02 15:04:05",
            "status":"doing"
        }
        ...
    ]


### 1.3 taskHistory
    [
        {
            "token":"tokenstr1",
            "task": "--- task1",
            "createtime": "2014-12-31 23:59:56",
            "doingtime": "2014-12-31 23:59:58",
            "donetime":"2006-01-02 15:04:05"
        },
        {
            "token":"tokenstr2",
            "task": "--- task2",
            "createtime": "2014-12-31 23:59:57",
            "doingtime": "2014-12-31 23:59:59",
            "donetime":"2006-01-02 15:04:05"
        },
        ...
    ]

# 2. For Client

### 2.1 Global json
    
    {
        "usertoken":"tokenStr",
        "pushtime":"2014-12-31 23:59:57",
        "pushtoken":"tokenstr",
        "username":"username",
        "password":"password"
    }

### 2.2 worktaskJson

    [
        {
            "task":"---task1...",
            "token":""tokenstr1,            
            "createtime":"2014-12-31 23:59:57",
            "doingtime":"2014-12-31 23:59:58",
            "donetime":"2014-12-31 23:59:59",
            "updatetime":"2014-12-31 23:59:59",
            "status":"doing"
        },
        {
            "task":"---task2...",
            "token":""tokenstr2,
            "createtime":"2014-12-31 23:58:57",
            "doingtime":"2014-12-31 23:58:58",
            "donetime":"2014-12-31 23:58:59",
            "updatetime":"2014-12-31 23:59:59",
            "status":"doing"
        }
        ...
    ]

### 2.2 historytaskJson

    [
        {
            "task":"---task1...",
            "token":""tokenstr1,            
            "createtime":"2014-12-31 23:59:57",
            "doingtime":"2014-12-31 23:59:58",
            "donetime":"2014-12-31 23:59:59",
            "updatetime":"2014-12-31 23:59:59",
            "status":"doing"
        },
        {
            "task":"---task2...",
            "token":""tokenstr2,
            "createtime":"2014-12-31 23:58:57",
            "doingtime":"2014-12-31 23:58:58",
            "donetime":"2014-12-31 23:58:59",
            "updatetime":"2014-12-31 23:59:59",
            "status":"doing"
        }
        ...
    ]
