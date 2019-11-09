package redis

// 单次执行命令的R函数,执行完命令自动关闭连接
func R(command string, args ...interface{}) (reply interface{}, err error) {
	conn := GetRedisConn()
	defer func() {
		conn.Close()
	}()

	var params []interface{}
	for _, i := range args {
		params = append(params, i)
	}
	return conn.Do(command, params...)

}

// pipeline 串行命令，减少网络开销
