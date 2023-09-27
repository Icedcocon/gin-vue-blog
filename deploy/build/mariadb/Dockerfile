FROM  mariadb:11

# 定义工作目录
ENV WORK_PATH /usr/local/work
# 定义被容器自动执行的目录（MariaDB不起作用）
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d
# 定义要执行的 shell 文件
ENV RUN_SHELL run.sh
# 定义被Mariadb容器自动执行的文件
ENV MARIADB_AUTO_RUN ${AUTO_RUN_DIR}/${RUN_SHELL}

COPY ./50-client.cnf /etc/mysql/mariadb.conf.d
COPY ./50-mariadb-clients.cnf /etc/mysql/mariadb.conf.d/
COPY ./50-server.cnf /etc/mysql/mariadb.conf.d/
# 把数据库初始化数据的文件复制到工作目录
COPY ./gvb.sql ${WORK_PATH}/
# 把执行初始化的脚本放到 /docker-entrypoint-initdb.d 目录下
COPY ./${RUN_SHELL} ${AUTO_RUN_DIR}/

# 给执行文件添加可执行权限
RUN chmod a+x ${AUTO_RUN_DIR}/${RUN_SHELL}