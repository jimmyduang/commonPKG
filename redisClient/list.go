package redisClient

/**
* list表:从左到右的链表
* 可以实现先进先出,也可以实现先进后出
* 可以插入重复值
 */

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	//	"strings"
)

/**
* 将一个或多个值 value 插入到列表 key 的表头
* @params   key    链表
* @params   value   插入的值,可以是单个,也可以是多个
 */
func (this *RedisPool) LPushList(key, value string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err = redis.Int(c.Do("LPUSH", key, value))
	if err != nil {
		if debug {
			fmt.Println("redis命令写入SET错误->", err.Error())
		}
	}
	return res
}

/**
* 将一个或多个值 value 插入到列表 key 的表尾
* @params   key    链表
* @params   value   插入的值,可以是单个,也可以是多个
 */
func (this *RedisPool) RPushList(key, value string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err = redis.Int(c.Do("RPUSH", key, value))
	if err != nil {
		if debug {
			fmt.Println("redis命令写入SET错误->", err.Error())
		}
	}
	return res
}

/**
* 计算list中元素数量
 */
func (this *RedisPool) LenList(key string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("LLEN", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令统计SET错误->", err.Error())
		}
	}
	return res
}

/**
* 从表头删除元素并返回该元素
 */
func (this *RedisPool) RPopList(key string) string {
	res := ""
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.String(c.Do("RPOP", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令删除SET错误->", err.Error())
		}
	}
	return res
}

/**
* 从表尾删除元素并返回该元素
 */
func (this *RedisPool) LPopList(key string) string {
	res := ""
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.String(c.Do("LPOP", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令删除SET错误->", err.Error())
			fmt.Println("redis错误->", err)
		}
	}
	return res
}
