package storage

/**
* @Author: super
* @Date: 2020-09-01 17:01
* @Description: 用于存储与elastic的映射
**/

const Mapping = `
{
    "mappings": {
        "properties": {
            "title": {
                "type": "text"
            },
            "url": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            },
            "content": {
                "type": "text"
            }
        }
    }
}`