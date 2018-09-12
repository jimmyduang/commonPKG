/**
* set 的是通过 hash table 实现的
* hash table 会随着添加或者删除自动的调整大小。
* 需要注意的是调整 hash table 大小时候需要同步(获取写锁)会阻塞其他读写操作
 */
package redisClient

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
* 添加数据到集合当中去，如果不存在，则创建这个key
* @key	string	集合的key
* @val	string	集合的值，多个用空格隔开
* return int   	影响的条数
 */

func (this *RedisPool) SortSetAddString(key, val string, score int64) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("ZADD", key, score, val))
	if err != nil {
		fmt.Println("redis命令写入SORTSET错误->", err.Error())
	}
	return res
}

/**
* 向redis中写入一个key—val类型的字符串
* @key 	string 	参数的主键
* @val 	string 	需要写入缓存的值
* @Ex	int		超时时间（秒）
				参数为0时，永远不过期

* return 返回插入行数
*/
func (this *RedisPool) SortSetAddStringAndOutTime(key, val string, score int64, EX int) int {
	var result = 0

	key = this.prefix + ":" + key
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()

	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	result, err = redis.Int(c.Do("ZADD", key, score, val))

	if EX > 0 && result > 0 {
		this.KeyExpire(key, EX, 0)
	}

	if err != nil {
		if debug {
			fmt.Println("SortSetAddStringAndOutTime->" + err.Error())
		}
	}
	return result
}

/**
* 返回有序集合中的排名
* @param key string 主健
* @param val  string 元素
 */
func (this *RedisPool) SortSetRank(key, val string) int {
	res := -1
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("ZRANK", key, val))
	if err != nil {
		if debug {
			fmt.Println("redis命令查询SORTSET错误->", err.Error())
		}
	} else {
		if res >= 0 {
			res = res
		} else {
			res = -1
		}
	}
	return res
}

/**
* 删除指定分值区间的元素
 */
func (this *RedisPool) RemoveSortSetByScore(key string, min, max int64) int {
	res := -1
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("ZREMRANGEBYSCORE", key, min, max))

	if err != nil {
		if debug {
			fmt.Println("redis命令删除SORTSET错误->", err.Error())
		}
	} else {
		if res >= 0 {
			res = res
		} else {
			res = -1
		}
	}
	return res
}

/**
* 获取指定分值区间内的所有元素列表
 */
func (this *RedisPool) GetSortSetRangeByScore(key string, min, max int64) []string {
	res := []string{}
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("ZRANGEBYSCORE", key, min, max))
	if err != nil {
		if debug {
			fmt.Println("redis命令查询SORTSET指定分值区间元素错误->", err.Error())
		}
	}
	return res
}

/**
* 返回有序集合中的基数(即数目)
* @param key string 主健
 */
func (this *RedisPool) GetSortSetCount(key string) int {
	res := -1
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()

	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("ZCARD", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令查询SORTSET错误->", err.Error())
		}
	} else {
		if res >= 0 {
			res = res
		} else {
			res = -1
		}
	}
	return res
}

/**
* 获取指定区间内的(按照分值的递增排序,分值相同则按照字典逆序)元素列表
 */
func (this *RedisPool) GetSortSetRangeByZrange(key string, min, max int) []string {
	res := []string{}
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("ZRANGE", key, min, max))
	if err != nil {
		if debug {
			fmt.Println("redis命令查询SORTSET指定区间元素错误->", err.Error())
		}
	}
	return res
}
