package redisClient

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
* 将你的指令直接发送给redis
 */
func (this *RedisPool) Do(cmd string) interface{} {
	//	return
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err := redis.Bool(c.Do(cmd))
	if err != nil {
		if debug {
			fmt.Println("redis命令执行错误->", err.Error())
		}
	}
	return res
}

/**
* 检查一个key是否存在
* @key string 主键
* return int 返回值=1 表示存在
 */
func (this *RedisPool) KeyExists(key string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err := redis.Int(c.Do("EXISTS", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令判断主键是否存在错误->", err.Error())
		}
	}
	return res
}

/**
* 删除一个指定的key
* @key string 需要删除的键值,多个用空格隔开
* return int  返回删除的个数
 */

func (this *RedisPool) KeyDel(key string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err := redis.Int(c.Do("DEL", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令删除主键错误->", err.Error())
		}
	}
	return res
}

/**
* 设置一个key的过期时间
* @key string 需要删除的键值,多个用空格隔开
* @times	int	过期时间，单位秒
* @isPrefix	int	判断是否需要key前缀   需要=1

* return int  返回设置成功的个数
 */
func (this *RedisPool) KeyExpire(key string, times, isPrefix int) int {
	res := 0
	//return res
	if isPrefix == 1 {
		key = this.prefix + ":" + key
	}

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err := redis.Int(c.Do("EXPIRE", key, times))
	fmt.Printf(string(res))
	if err != nil {
		fmt.Println("redis命令设置过期时间错误->", err.Error())
	}

	return res
}
