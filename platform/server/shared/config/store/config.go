/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package store

import (
	"errors"
	"fmt"
	"github.com/cloudwego/cwgo/platform/server/shared/consts"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Type  string `mapstructure:"type"`
	Mysql Mysql  `mapstructure:"mysql"`
	Mongo Mongo  `mapstructure:"mongo"`
	Redis Redis  `mapstructure:"redis"`
}

func (c Config) GetStoreType() consts.StoreType {
	return consts.StoreTypeMapToNum[c.Type]
}

func (c Config) NewMysqlDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.Mysql.GetDsn()), &gorm.Config{
		PrepareStmt: true,
	})
}

func (c Config) NewRedisClient() (redis.UniversalClient, error) {
	if c.Redis.StandAlone.Addr != "" {
		return redis.NewClient(&redis.Options{
			Addr:     c.Redis.StandAlone.Addr,
			Password: c.Redis.StandAlone.Password,
			DB:       c.Redis.StandAlone.Db,
		}), nil
	}
	if c.Redis.Cluster.Addrs != nil {
		addrs := make([]string, len(c.Redis.Cluster.Addrs))
		for i, addr := range c.Redis.Cluster.Addrs {
			addrs[i] = fmt.Sprintf("%s:%s", addr.Ip, addr.Port)
		}
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    addrs,
			Password: c.Redis.Cluster.Password,
		}), nil
	}

	return nil, errors.New("config not found")
}

type Mysql struct {
	Addr     string `mapstructure:"addr"`
	Port     string `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
}

func (m Mysql) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai",
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.Db,
		m.Charset)
}

type Mongo struct {
	Addr         string `mapstructure:"addr"`
	Port         string `mapstructure:"port"`
	DatabaseName string `mapstructure:"databaseName"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
}

func (m Mongo) GetAddr() string {
	return fmt.Sprintf("mongodb://%s:%s", m.Addr, m.Port)
}

type Redis struct {
	StandAlone RedisStandAlone `mapstructure:"standalone"`
	Cluster    RedisCluster    `mapstructure:"cluster"`
}

type RedisStandAlone struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type RedisCluster struct {
	MasterNum int `mapstructure:"masterNum"`
	Addrs     []*struct {
		Ip   string `mapstructure:"ip"`
		Port string `mapstructure:"port"`
	} `mapstructure:"addrs"`
	Password string `mapstructure:"password"`
}
