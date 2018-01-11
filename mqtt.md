````conf
# 禁止匿名访问
allow_anonymous false
# 账号密码文件
password_file /etc/mosquitto/pwfile
# 配置topic和用户关系
acl_file /etc/mosquitto/aclfile

# 服务绑定的IP地址
# bind_address

# 新增用户 choi 并设置密码为 lanseyujie [注意: 使用 -c 参数为创建一个文件，这将会覆盖已有文件！]
sudo mosquitto_passwd -c /etc/mosquitto/pwfile choi

sudo /etc/mosquitto/aclfile

user choi
topic write testtopic/#

user choi
topic read testtopic/#

# 订阅主题
mosquitto_sub -h 127.0.0.1 -t testtopic -u choi -P lanseyujie

# 发布消息
mosquitto_pub -h 127.0.0.1 -t testtopic -u choi -P lanseyujie -m "Hello MQTT"
````