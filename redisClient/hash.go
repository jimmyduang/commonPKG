package redisClient

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
* 向redis中写入一个hash类型的map
* @key 	string 	参数的主键
* @hash 	map 	需要写入缓存的值
* @Ex	int		超时时间（秒）
				参数为0时，永远不过期
*/
func (this *RedisPool) HashWrite(key string, hash map[string]string, EX int) {
	//return
	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	for k, v := range hash {
		_, err = c.Do("HSET", key, k, v)
		if EX > 0 {
			this.KeyExpire(key, EX, 0)
		}
	}
	if err != nil {
		if debug {
			fmt.Println("redis命令写入字符串错误->", err.Error())
		}
	}
}

/**
* 从hash数据读取读取一个key的field的vales字符串
* @param key string 主健
* @param field string 字段
* @return string 结果值
 */
func (this *RedisPool) HashReadField(key string, field string) string {
	//return ""
	key = this.prefix + ":" + key
	var err error
	res := ""

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.String(c.Do("HGET", key, field))
	if err != nil {
		if debug {
			fmt.Println("redis命令读取hash错误->", err.Error())
		}
	}
	return res
}

/**
* 从hash数据读取删除一个key的field
* @param key string 主健
* @param field string 字段
* @return string 结果值
 */
func (this *RedisPool) HashDelField(key string, field string) {
	//return
	key = this.prefix + ":" + key
	var err error

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	_, err = c.Do("HDEL", key, field)
	if err != nil {
		if debug {

			fmt.Println("redis命令读取hash错误->", err.Error())
		}
	}
}

/**
* 从hash数据读取所有的fields
* @param key string 主健
* @return []string 结果值
 */
func (this *RedisPool) HashReadFields(key string) []string {
	res := []string{}
	//return res
	key = this.prefix + ":" + key
	var err error

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("HKEYS", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令读取hash切片数组错误->", err.Error())
		}
	}
	return res
}

/**
* 从hash数据读取所有的values
* @param key string 主健
* @return []string 结果值
 */
func (this *RedisPool) HashReadValues(key string) []string {

	res := []string{}
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err = redis.Strings(c.Do("HVALS", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令读取hash切片数组错误->", err.Error())
		}
	}
	return res
}

/**
* 更新hash一个字段的值,如果该字段不存在,则新建这个字段
* 如果存在这个字段,则新值覆盖旧值
* @key    string   hash的key值
* @field    string   hash的key值的字段
* @value    string   hash的key值的字段值
* 返回值:-1 -> 返回错误, 1 -> 新增一个字段, 0 -> 覆盖旧值
 */
func (this *RedisPool) HashWriteField(key, field, value string, EX int) int {
	res := -1
	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err = redis.Int(c.Do("HSET", key, field, value))
	if EX > 0 {
		this.KeyExpire(key, EX, 0)
	}
	if err != nil {
		fmt.Println("redis命令写入字符串错误->", err.Error())
	}
	return res
}

/**
* 从hash数据读取所有的key和vals
* @param key string 主健
* @return map 结果值
 */
func (this *RedisPool) HashReadAllMap(key string) map[string]string {
	m := map[string]string{}
	//return m
	key = this.prefix + ":" + key

	var err error
	res := []string{}
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err = redis.Strings(c.Do("HGETALL", key))
	lens := len(res)
	if err != nil {
		if debug {
			fmt.Println("redis命令读取hash切片数组错误->", err.Error())
		}
	} else {
		for i := 0; i < lens; i = i + 2 {
			t1 := res[i]
			t2 := res[i+1]
			m[t1] = t2
		}
	}
	return m
}

/**
* 将名称为key的hash中field的value增加integer
* @param key string 主健
* @param field string 字段
* @param integer string 加减数
 */
func (this *RedisPool) HashHincrby(key string, field string, integer string) {
	key = this.prefix + ":" + key
	var err error

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	_, err = c.Do("HINCRBY", key, field, integer)
	if err != nil {
		if debug {
			fmt.Println("redis命令读取hash错误->", err.Error())
		}
	}
}
