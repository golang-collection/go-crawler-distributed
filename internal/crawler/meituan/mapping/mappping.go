package mapping

/**
* @Author: super
* @Date: 2021-01-05 15:15
* @Description:
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