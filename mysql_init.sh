#! /bin/bash
/etc/init.d/mysql start

mysql -uroot -D mysql << EOF
set password for 'root'@'localhost' = password('123456');
flush privileges;
exit
EOF

