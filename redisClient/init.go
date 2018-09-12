/**
* redis相关操作的处理
* @author:james
* @date:2015-09-09修改
* 包含功能
	字符串的读取和写入
	整形的读取和写入
	浮点的读取和写入
	布尔的读取和写入
	hash读取和写入（golang处理是map）
	set集合的读取和写入
用例：
	redisClient.InitRedis("gc", "tcp", "192.168.1.211", "6379", true)
	d := map[string]string{}
	d["a"] = "1"
	d["b"] = "2"
	d["c"] = "3"
	//map写入hash
	redisClient.HashWrite("myhash", d, 0)
	//读取hash中的field
	s := redisClient.HashReadField("myhash", "b")
	fmt.Println("s=>", s)
	//删除一个field
	redisClient.HashDelField("myhash", "c")
	//读取hash到map
	d2 := redisClient.HashReadAllMap("myhash")
	fmt.Println("d2", d2)

*/
package redisClient

import (
	"commonPKG/initPKG"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 连接池结构体
type RedisPool struct {
	pool      *redis.Pool
	prefix    string
	redis_pwd string
}

//redis调试开关
var debug = false
var Config initPKG.Conf

//被外部调用的结构体对象
var Redis *RedisPool

//加载包就会执行
func init() {
	Config.GetConf()
	//初始化redis
	redis_prefix := Config.Redisprefix
	redis_network := Config.Redisnetwork
	redis_addr := Config.Redisaddr
	redis_port := Config.Redisport
	redis_pwd := Config.Redispwd
	redis_db := Config.Redisdb
	redis_bug := Config.Redisbug
	if redis_bug == "on" {
		debug = true
	}
	// if db_err != nil {
	// 	redis_db = 0
	// }

	Redis = initRedis(redis_prefix, redis_network, redis_addr, redis_port, redis_pwd, redis_db)
}

/**
*检查redis连接是否正常
* 正常的话，将返回nil
* @param k string  redis中主健使用的key
* @param n string	定义redis连接类型，一般为tcp
* @param ip string  redis服务器的ip地址
* @param p string redis服务器的端口
* @param w string redis认证密码
 */
func initRedis(k, n, ip, p, w string, db int) *RedisPool {

	opt_conn_time := redis.DialConnectTimeout(time.Second * 15)
	opt_conn_pwd := redis.DialPassword(w)
	opt_conn__db := redis.DialDatabase(db)

	pool := &redis.Pool{
		MaxIdle:   80,    //最大空闲
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			host := ip + ":" + p
			c, err := redis.Dial(n, host, opt_conn_time, opt_conn_pwd, opt_conn__db)
			if err != nil {
				if debug {
					fmt.Println("不能连接到redis->", n, ip, p, err.Error())
				}
			} else {
				//设置访问权限
				c.Send("AUTH", w)
				if debug {
					fmt.Println("成功连接到redis->", n, ip, p)
				}
			}
			return c, err
		},
		IdleTimeout: time.Second * 15, //设置空闲15秒超时

	}
	return &RedisPool{pool, k, w}
}

/**
* 返回一个可以操作的链接
 */
func (this *RedisPool) getCon() redis.Conn {
	c := this.pool.Get()
	if len(this.redis_pwd) > 1 {
		c.Send("AUTH", this.redis_pwd)
	}
	if debug {
		fmt.Println("开启连接数->", this.pool.ActiveCount())
	}
	return c
}
