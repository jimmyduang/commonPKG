package initPKG

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	Ginurl       string `yaml:"ginurl"`
	Ginport      string `yaml:"ginport"`
	Logurl       string `yaml:"logurl"`
	Logport      string `yaml:"logport"`
	Tokenurl     string `yaml:"tokenurl"`
	Tokenport    string `yaml:"tokenport"`
	Grpcurl      string `yaml:"grpcurl"`
	Grpcport     string `yaml:"grpcport"`
	Mysqlhost    string `yaml:"mysqlhost"`
	Mysqlport    string `yaml:"mysqlport"`
	Mysqluser    string `yaml:"mysqluser"`
	Mysqlpass    string `yaml:"mysqlpass"`
	Mysqldb      string `yaml:"mysqldb"`
	Mysqlcharset string `yaml:"mysqlcharset"`
	Redisstatus  string `yaml:"redisstatus"`
	Redisnetwork string `yaml:"redisnetwork"`
	Redisaddr    string `yaml:"redisaddr"`
	Redisport    string `yaml:"redisport"`
	Redisprefix  string `yaml:"redisprefix"`
	Redispwd     string `yaml:"redispwd"`
	Redisdb      int    `yaml:"redisdb"`
	Redisbug     string `yaml:"redisbug"`
	Encodeurl    string `yaml:"encodeurl"`
	Decodeurl    string `yaml:"decodeurl"`
	Webcode      string `yaml:"webcode"`
	JwtSecret    string `yaml:"jwtsecret"`
}

func (c *Conf) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

var Config Conf
var db *sql.DB //数据库句柄指针
var Mysql_connStr = ""

// SALT 密钥
var SALT = ""

func init() {
	Config.GetConf()
	Mysql_connStr = Config.Mysqluser + ":" + Config.Mysqlpass + "@tcp(" + Config.Mysqlhost + ":" + Config.Mysqlport + ")/" + Config.Mysqldb + "?charset=" + Config.Mysqlcharset
	SALT = Config.JwtSecret
}
