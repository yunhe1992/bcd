# bcd
百度签名-bcd 拓展

**使用方法**
```
1.项目引入bce包："github.com/baidubce/bce-sdk-go/bce"
2.把bcd文件夹放到项目指定目录
3.程序里边调用,下边的import路径需要改成自己的
package lib

import (
	"self/config"
	"self/model/cdn/bcd"
	"self/model/cdn/bcd/api"
)

type Baidu struct {
}

func (this Baidu) DomainList() (interface{}, error) {
	client, _ := bcd.NewClient(config.Config.Cloud["baidu"].AccessKey, config.Config.Cloud["baidu"].SecretKey, "")
	paras := &api.ListDomainArgs{}
	paras.PageSize = 100
	paras.PageNo = 1
	result, err := client.ListDomain(nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

```