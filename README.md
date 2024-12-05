# short-url-4go
short-url的Go版本

🙏 特别感谢Mehul Gohil和Arron两位优秀的创作者，这个项目分别从两位的https://github.com/mehulgohil/shorti.fy和https://github.com/ArronYR/short-url获得了灵感！

 🏝️ 项目版本

- { node: v18.20.5,  go: go1.22.5 }

🫧 设置环境变量

|      名称       |        值        |
| :-------------: | :--------------: |
|      PORT       |        80        |
|     ORIGIN      | http://localhost |
|     DB_HOST     |    127.0.0.1     |
|     DB_PORT     |       3306       |
|   DB_USERNAME   |       root       |
|   DB_PASSWORD   |       root       |
|     DB_NAME     |    short_url     |
| CACHE_MAX_ITEMS |       1200       |
| CACHE_LIFETIME  |       60s        |
|   ACCESS_LOG    |       true       |
|      TOKEN      |  53ROYinHId9qke  |
|   API_SECRET    |  1FIsiEpxQo5l7H  |

🍉 以下是随机的 50 个不同类型的网站链接，每个链接按行分隔（测试用）：

```
https://www.google.com  
https://www.youtube.com  
https://www.amazon.com  
https://www.wikipedia.org  
https://www.reddit.com  
https://www.netflix.com  
https://www.facebook.com  
https://www.twitter.com  
https://www.instagram.com  
https://www.linkedin.com  
https://www.apple.com  
https://www.microsoft.com  
https://www.github.com  
https://www.stackoverflow.com  
https://www.khanacademy.org  
https://www.coursera.org  
https://www.medium.com  
https://www.bbc.com  
https://www.cnn.com  
https://www.nytimes.com  
https://www.aliexpress.com  
https://www.taobao.com  
https://www.jd.com  
https://www.imdb.com  
https://www.soundcloud.com  
https://www.spotify.com  
https://www.twitch.tv  
https://www.pinterest.com  
https://www.deviantart.com  
https://www.quora.com  
https://www.etsy.com  
https://www.adobe.com  
https://www.paypal.com  
https://www.dropbox.com  
https://www.weibo.com  
https://www.zhihu.com  
https://www.tiktok.com  
https://www.salesforce.com  
https://www.samsung.com  
https://www.nike.com  
https://www.hulu.com  
https://www.airbnb.com  
https://www.booking.com  
https://www.tripadvisor.com  
https://www.uber.com  
https://www.zoom.us  
https://www.slack.com  
https://www.upwork.com  
https://www.fiverr.com  
https://www.wix.com  
https://www.shopify.com  
```

# 参考

🙏 感谢以下开源项目：

- https://github.com/go-gorm/gorm
- https://github.com/bluele/gcache
- https://github.com/ArronYR/short-url
- https://github.com/mehulgohil/shorti.fy
- https://github.com/auth0-samples/auth0-golang-web-app
- https://github.com/herusdianto/gorm_crud_example
- https://github.com/uber-go/zap
- https://github.com/go-gorm/gorm
- https://github.com/umijs/umi
- https://github.com/alibaba/hooks

🙏 感谢以下作者，他们写了很多优秀的文章

- gorm官方文档  https://gorm.io/zh_CN/docs/query.html
- material-ui官方文档  https://mui.org.cn/material-ui/getting-started/、https://mui.com/x/react-data-grid/getting-started/
- iris官方文档  https://docs.iris-go.com/iris
- 关于集成auth0的几个示例  https://developer.auth0.com/resources/code-samples/full-stack/hello-world/basic-access-control/spa/react-typescript/standard-library-golang、
  https://auth0.com/docs/quickstart/webapp/golang/interactive、
  https://auth0.com/docs/quickstart/webapp/golang、https://manage.auth0.com/、https://auth0.com/docs/quickstart/webapp/golang/interactive、https://developer.auth0.com/resources/code-samples
- actix官方文档  https://actix.rs/docs/getting-started
- react教程  https://www.runoob.com/react/react-jsx.html、https://developer.mozilla.org/zh-CN/docs/Learn/Tools_and_testing/Understanding_client-side_tools/Introducing_complete_toolchain、https://zh-hans.react.dev/learn/writing-markup-with-jsx
- 在React中使用Mock生成模拟数据你学会了吗  https://juejin.cn/post/7098610717493837860
- es6教程  https://es6.ruanyifeng.com/#docs/module
- js教程  https://www.runoob.com/js/js-type-conversion.html、https://www.w3school.com.cn/js/js_arrow_function.asp、https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Functions/Arrow_functions、https://zh.javascript.info/promise-chaining
- 关于es6通过import导入，什么时候需要花括号  https://blog.csdn.net/qq_51427204/article/details/122453283
- JSX 语法详解  https://juejin.cn/post/7030258791987806222、https://zh-hans.react.dev/learn/writing-markup-with-jsx
- ts教程  https://ts.xcatliu.com/introduction/what-is-typescript.html、 https://ts.xcatliu.com/、https://typescript.p6p.net/typescript-tutorial/declare.html
- 详解 TS 中的泛型  https://juejin.cn/post/7133810035171262501
- ts中的特殊符号 ?. ?: 等代表的含义与使用  https://blog.csdn.net/qq_41619796/article/details/129833416
- 使用 index.ts 文件重新导出模块：提高 TypeScript 项目的可读性、可维护性和可重用性  https://juejin.cn/post/7221004205271646245
- dynamodb相关  https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/AccessingDynamoDB.html#Tools.CLI、https://aws.amazon.com/cn/cli/
- 了解 JavaScript Promise  https://dev.nodejs.cn/learn/understanding-javascript-promises/
- Promise，async/await  https://zh.javascript.info/promise-basics
- 终于搞懂类型声明文件.d.ts和declare了，原来用处如此大  https://sduck.tech/archives/284
- namespace 命名空间  https://www.arryblog.com/vip/ts/module-namespace.html#_1%E3%80%81%E5%9F%BA%E7%A1%80%E7%94%A8%E6%B3%95
- 使用 index.ts 文件重新导出模块：提高 TypeScript 项目的可读性、可维护性和可重用性  https://juejin.cn/post/7221004205271646245
- ahook https://ahooks.js.org/zh-CN/hooks/use-request/index
- JavaScript Promise用法示例  https://www.leidazhifu.com/10828
- return  https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Statements/return
- Typography排版  http://121.4.220.50/components/typography-cn
- JS 和 JSX 、TS 和 TSX 的区别  https://juejin.cn/post/7226187111330431031
- 扩展语法 (...)   https://mdn.org.cn/en-US/docs/Web/JavaScript/Reference/Operators/Spread_syntax
- 计算属性名  https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Object_initializer
- State：组件的记忆  https://zh-hans.react.dev/learn/state-a-components-memory
- React Hooks 是什么  https://cloud.tencent.com/developer/article/1906643
- svg路径在线预览 https://uutool.cn/svg-path/
- 科普：Native App、Web App与Hybrid App  https://www.woshipm.com/pd/321844.html
- 说说你对SPA（单页应用）的理解?  https://github.com/febobo/web-interview/issues/3
- ant design官方文档  https://pro.ant.design/docs/request/
- Umi js官方文档  https://v3.umijs.org/zh-CN/docs
- ahooks官方文档  https://ahooks.js.org/zh-CN/hooks/use-request/basic
- JavaScript Promise用法示例  https://www.leidazhifu.com/10828
- HTML 和 JSX 的区别  https://www.freecodecamp.org/chinese/news/html-vs-jsx-whats-the-difference/
- go redis官方文档  https://redis.uptrace.dev/zh/guide/go-redis.html#dial-tcp-i-o-timeout

