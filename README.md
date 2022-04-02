# UTByDocker

基于docker容器实现CI

mysql_init.sh 脚本中可source相关的sql文件

ci_noti.sh 脚本中可以添加聊天机器人的webhook 和 请求方式和参数 需要自定义修改。
这部分可以用gitlab自带的webhook功能，不过，需要起一个服务做转发到机器人webhook。

gitlab-ci.yml 文件中的tags 需要修改成gitlab中注册的runner-tag
