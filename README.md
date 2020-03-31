# 个人博客

## 命令行工具
    -syncdb     生成数据表
    -register   注册账号（注意保存生成的userkey）

## config 文件
    {
      "debug": "true",  // debug 模式
      "host": ":8081",  // 监听端口
      "title": "",      // 网站标题
      "record_number": "",        // 备案号
      "onetimePassword": "true",  // 是否使用动态密钥
      "db_host": "", // mysql ip
      "db_port": "", // mysql 端口
      "db_user": "", // mysql 用户名
      "db_pass": "", // mysql 密码
      "db_name": "", // mysql 数据库名
      "redis_addr": "", // redis ip
      "redis_port": "", // redis 端口
      "redis_pass": "", // redis 密码
      "serve_name": ""  // 服务域名
    }
    
## UserKey config 文件
    {
        "user_key": "" // 注册时生成的密钥
    }