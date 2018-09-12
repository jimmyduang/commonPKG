/**
* set 的是通过 hash table 实现的
* hash table 会随着添加或者删除自动的调整大小。
* 需要注意的是调整 hash table 大小时候需要同步(获取写锁)会阻塞其他读写操作
 */
package redisClient

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

/**
* 添加数据到集合当中去，如果不存在，则创建这个key
* @key	string	集合的key
* @val	string	集合的值，多个用空格隔开
* return int   	影响的条数
 */

func (this *RedisPool) SetAddString(key string, val string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	res, err = redis.Int(c.Do("SADD", key, val))
	if err != nil {
		if debug {
			fmt.Println("redis命令写入SET错误->", err.Error())
		}
	}
	return res
}

/**
* 从集合中删除指定的值
* @param key string 主健
* @param item string 需要删除的值
 */
func (this *RedisPool) SetDelItem(key string, item string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("SREM", key, item))
	if err != nil {
		if debug {
			fmt.Println("redis命令删除SET错误->", err.Error())
		}
	}
	return res
}

/**
* 计算集合数量
* @key	string	集合的key
* @param int 返回集合中数据的数量
 */
func (this *RedisPool) SetCountKey(key string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("SCARD", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令统计SET错误->", err.Error())
		}
	}
	return res
}

/**
* 判断集合中是否存在项目
* @key	string	集合的key
* @item	string	需要判断的项
* @param int 如果存在返回1，不存在返回0
 */
func (this *RedisPool) SetExistItem(key, item string) int {
	res := 0
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Int(c.Do("SISMEMBER", key, item))
	if err != nil {
		if debug {
			fmt.Println("redis判断是否存在set命令错误->", err.Error())
		}
	}
	return res
}

/**
* 从集合中读取结果集
* @param key string 主健
* @return []string 返回结果集切片
 */
func (this *RedisPool) SetReadItems(key string) []string {
	res := []string{}
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("SMEMBERS", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令读取SET错误->", err.Error())
		}
	} else {
		if debug {
			fmt.Println("redis cmd->SMEMBERS", key)
		}
	}
	return res
}

/**
* 计算2个集合之间的差集
* 主健1对比主健2的差集
* @param key string 主健1
* @param key string 主健2
* @return []string 返回结果集切片
* 举例:
	key1 >> 1 2 3
	key2 >> 1
	则本函数得到的结果是: 2 3
*/
func (this *RedisPool) SetDiffItems(key1, key2 string) []string {
	res := []string{}
	//return res
	key1 = this.prefix + ":" + key1
	key2 = this.prefix + ":" + key2

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("SDIFF", key1, key2))
	if err != nil {
		if debug {
			fmt.Println("redis命令比较SET差集错误->", err.Error())
		}
	}
	return res
}

/**
* 计算2个集合之间的交集
* @param key string 主健1
* @param key string 主健2
* @return []string 返回结果集切片
* 举例:
	key1 >> 1 2 3
	key2 >> 1
	则本函数得到的结果是: 1
*/
func (this *RedisPool) SetInterItems(key1, key2 string) []string {
	res := []string{}
	//return res
	key1 = this.prefix + ":" + key1
	key2 = this.prefix + ":" + key2
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("SINTER", key1, key2))
	if err != nil {
		if debug {
			fmt.Println("redis命令比较SET交集错误->", err.Error())
		}
	}
	return res
}

/**
* 计算多个集合之间的并集
* @param key string 主健
* @param keys string 其他需要合并的并集，多个请使用英文逗号隔开
* @return []string 返回结果集切片
* 举例:
	key >> 1 2 3
	key2 >> 1
	key3 >> 4
	则本函数得到的结果是: 1 2 3 4
*/
func (this *RedisPool) SetSunionItems(key, keys string) []string {
	res := []string{}
	//return res
	key = this.prefix + ":" + key
	//切割多个key，分别加上前缀
	sList := strings.Split(keys, ",")
	keys = ""
	for _, v := range sList {
		keys = keys + " " + this.prefix + ":" + v
	}
	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("SUNION", key, keys))
	if err != nil {
		if debug {
			fmt.Println("redis命令比较SET并集错误->", err.Error())
		}
	}
	return res
}

/**
* 随机删除一条数据然后返回被删除的元素
* @return string  删除的元素
 */
func (this *RedisPool) SetDelRandomItem(key string) string {
	res := ""
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.String(c.Do("SPOP", key))
	if err != nil {
		if debug {
			fmt.Println("redis命令删除SET错误->", err.Error())
		}
	}
	return res
}

/**
* 随机返回集合的元素
* @return string  删除的元素
 */
func (this *RedisPool) SetRandomItem(key string, count int) []string {
	res := []string{}
	//return res
	key = this.prefix + ":" + key

	var err error
	// 从连接池里面获得一个连接
	c := this.getCon()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	res, err = redis.Strings(c.Do("SRANDMEMBER", key, count))
	if err != nil {
		if debug {
			fmt.Println("redis命令随机SET错误->", err.Error())
		}
	}
	return res
}
