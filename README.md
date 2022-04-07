# UTByDocker

基于docker容器实现CI

mysql_init.sh 脚本中可source相关的sql文件

ci_noti.sh 脚本中可以添加聊天机器人的webhook 和 请求方式和参数 需要自定义修改。
这部分可以用gitlab自带的webhook功能，不过，需要起一个服务做转发到机器人webhook。

gitlab-ci.yml 文件中的tags 需要修改成gitlab中注册的runner-tag

如果需要在gitlab-ci中使用ssh git clone，参考官方文档和下面的exmaple
 ```yaml
before_script:
    - 'command -v ssh-agent >/dev/null || ( apt-get update -y && apt-get install openssh-client -y )'
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY" | tr -d '\r' | ssh-add -
    - mkdir -p ~/.ssh
    - chmod 700 ~/.ssh
    - ssh-keyscan git.garena.com >> ~/.ssh/known_hosts
    - chmod 644 ~/.ssh/known_hosts
    - git config --global --add user.email "$GITLAB_USER_EMAIL"
    - git config --global --add user.name "$GITLAB_USER_NAME"
    - git config --global --add url.ssh://gitlab@git.garena.com:2222.insteadOf "https://git.garena.com"
 ```
