mysql 可用性
    操作失败
        select 1;
        ping
慢查询次数
    show global status where variable_name='slow_queries'    #如果没有开启慢查询的话  开启就行  

    mysql> show global variables where variable_name like "%slow%";  查看相关情况
+---------------------------+--------------------------------------+
| Variable_name             | Value                                |
+---------------------------+--------------------------------------+
| log_slow_admin_statements | OFF                                  |
| log_slow_slave_statements | OFF                                  |
| slow_launch_time          | 2                                    |
| slow_query_log            | OFF                                  |
| slow_query_log_file       | /var/lib/mysql/ecs-owncloud-slow.log |

容量:
    qps:
        show global status where variable_name='Queries'
    tps:
        
        insert, update, delete *
        com_insert  show global status where variable_name='com_insert';
        com_update
        com_delete
        com_select
        com_replace
    连接:
        show global status where variable_name='Threads_connected'
        show global variables where variable_name= 'max_connections';  # 当前配置最大的数量

    流量：
        show global status where variable_name='Bytes_received'
        show global status where variable_name='Bytes_sent'


// mysql连接信息 => mysql host, port, user