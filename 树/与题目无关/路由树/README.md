**路由树**

[leetcode题目地址](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

为啥会写这个路由树，纯粹是因为有一个业务场景需要保存api与操作内容的映射关系，并且在请求过来时检索这个api接口是否有需要保存操作日志。  
之前的实现场景是把所有的映射关系保存到map中(写死在代码中)，请求过来的时候在map中获取。列如在map中有这样的{"/hello/name":"打招呼"}映射关系，那就在请求过来时通过map[request.path]获取数据。  
后面随着要记录的接口越来越多，这样方法明显不行了，每一次增加都需要直接修改代码，还要经历一次上线发布流程。并且如果路由是RESTful风格的例如"/hello/:name",那么上述的方法明显就不行了。  
所以去参考了GO语言的httprouter这个路由框架，通过路由树来保存这些路由，因为业务场景的原因就不需要保存Handler了。