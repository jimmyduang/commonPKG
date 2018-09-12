package redisClient

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
* 向redis中写入一个key—val类型的字符串
* @key 	string 	参数的主键
* @val 	string 	需要写入缓存的值
* @Ex	int		超时时间（秒）
				参数为0时，永远不过期

* return	返回影响行数
*/
func (this *RedisPool) StringWrite(key string, val string, EX int) {

	key = this.prefix + ":" + key
	var err error

	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	_, err = c.Do("SET", key, val)

	if EX > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		fmt.Println("StringWrite->" + err.Error())
	}

}

/**
* 向redis中写入一个key—val类型的字符串
* @key 	string 	参数的主键
* @val 	string 	需要写入缓存的值
* @Ex	int		超时时间（秒）
				参数为0时，永远不过期

* return	返回影响行数
*/
func (this *RedisPool) NXStringWrite(key string, val string, EX int) int {
	var result = 0

	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	result, err = redis.Int(c.Do("SETNX", key, val))

	if EX > 0 && result > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		fmt.Println("NXStringWrite->" + err.Error())
	}
	return result
}

/**
* 读取一个key的字符串
 */
func (this *RedisPool) StringRead(key string) string {
	key = this.prefix + ":" + key
	res := ""
	//return res
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, _ = redis.String(c.Do("GET", key))
	return res
}

/**
* 写入一个int值到redis
* @param	key		健名
* @param	val		健值
* @param	Ex		超时时间（秒）
				参数为0时，永远不过期
* return	返回影响行数
*/
func (this *RedisPool) IntWrite(key string, val int, EX int) int {
	var result = 0

	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	result, err = redis.Int(c.Do("SETNX", key, val))

	if EX > 0 && result > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		fmt.Println("IntWrite->" + err.Error())
	}
	return result
}

func (this *RedisPool) IntRead(key string) int {
	key = this.prefix + ":" + key
	res := 0
	//return res
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, _ = redis.Int(c.Do("GET", key))
	return res
}

/**
* 写入一个bool值到redis
* @param	key		健名
* @param	val		健值
* @param	Ex		超时时间（秒）
				参数为0时，永远不过期
* return	返回影响行数
*/
func (this *RedisPool) BoolWrite(key string, val bool, EX int) int {
	var result = 0
	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	result, err = redis.Int(c.Do("SETNX", key, val))
	if EX > 0 && result > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		fmt.Println("BoolWrite->" + err.Error())
	}
	return result
}

func (this *RedisPool) BoolRead(key string) bool {
	key = this.prefix + ":" + key
	res := false
	//return res
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, _ = redis.Bool(c.Do("GET", key))
	return res
}

/**
* 写入一个byte值到redis
* @param	key		健名
* @param	val		健值
* @param	Ex		超时时间（秒）
				参数为0时，永远不过期
* return	返回插入是否成功
*/
func (this *RedisPool) BytesWrite(key string, val []byte, EX int) int {
	var result = 0
	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	result, err = redis.Int(c.Do("SETNX", key, val))
	if EX > 0 && result > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		fmt.Println("BytesWrite->" + err.Error())
	}
	return result
}

func (this *RedisPool) BytesRead(key string) []byte {
	key = this.prefix + ":" + key
	var res []byte
	//return res
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	if debug == true {
		fmt.Println("GET->" + key)
	}
	res, _ = redis.Bytes(c.Do("GET", key))
	return res
}

/**
* 写入一个float64值到redis
* @param	key		健名
* @param	val		健值
* @param	Ex		超时时间（秒）
				参数为0时，永远不过期
* return	返回影响行数
*/
func (this *RedisPool) Float64Write(key string, val float64, EX int) int {
	var result = 0
	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	result, err = redis.Int(c.Do("SETNX", key, val))
	if EX > 0 && result > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		fmt.Println("Float64Write->" + err.Error())
	}
	return result
}

func (this *RedisPool) Float64Read(key string) float64 {
	key = this.prefix + ":" + key
	var res float64
	//return res
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, _ = redis.Float64(c.Do("GET", key))
	return res
}
