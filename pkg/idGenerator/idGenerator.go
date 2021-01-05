package idGenerator

/**
* @Author: super
* @Date: 2020-11-24 13:45
* @Description:
**/

import (
	"github.com/bwmarrin/snowflake"

	"strconv"
)

/**
* @Author: super
* @Date: 2020-09-09 22:04
* @Description: 雪花算法介绍：https://juejin.im/post/6844903562007314440
**/

var node *snowflake.Node

// InitSnowflake initiate Snowflake node singleton.
func InitSnowflake() error {
	// Get node number from env TIX_NODE_NO
	//key, ok := os.LookupEnv("TIX_NODE_NO")
	//if !ok {
	//	return fmt.Errorf("TIX_NODE_NO is not set in system environment")
	//}
	// Parse node number
	key := "1"
	nodeNo, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return err
	}
	// Create snowflake node
	n, err := snowflake.NewNode(nodeNo)
	if err != nil {
		return err
	}
	// Set node
	node = n
	return nil
}

// GenerateSnowflake generate Twitter Snowflake ID
func GenerateID() string {
	return node.Generate().String()
}
