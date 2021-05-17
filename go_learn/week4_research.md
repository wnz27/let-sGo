- [X] [Develop Your Design Philosophy](https://www.ardanlabs.com/blog/2017/01/develop-your-design-philosophy.html)
  > A design philosophy is important because it should be driving everything you do and every decision you make when writing code and structuring a project. It forces you to ask:
    - Who is your audience?
    - What are your priorities?
    - When do you take exceptions to the rules?
    - How do things work?
    - Why you are making the decisions you make?
    
    > With that understanding you are capable of:
    - Reasoning about tradeoffs and costs
    - Determining when and why you do things
    - Developing best practices and guidelines
    - Smelling out good code from bad
    - Appreciating others opinions
    - Participating in healthy debates
    - Refactoring your philosophy as your learn more
- [X] [Design Philosophy On Packaging](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html)
  - Design Philosophies
    - Purpose
        - Packages must be named with the intent to describe what it provides.
        - Packages must not become a dumping ground of disparate concerns.
    - Usability
        - Packages must be intuitive and simple to use.
        - Packages must respect their impact on resources and performance.
        - Packages must protect the user’s application from cascading changes.
        - Packages must prevent the need for type assertions to the concrete.
        - Packages must reduce, minimize and simplify its code base.
    - Portability
        - Packages must aspire for the highest level of portability.
        - Packages must reduce setting policies when it’s reasonable and practical.
        - Packages must not become a single point of dependency.
- [X] [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)
  Project Structure
    - Kit Projects
      > Think of the Kit project as a company’s standard library, so there should only be one. 
      The packages that belong to the Kit project need to be designed with the highest levels of portability in mind. 
      These packages should be usable across multiple Application projects and provide a very specific but foundational domain of functionality. 
      To this end, the Kit project is not allowed to have a vendor folder. 
      If any of packages are dependent on 3rd party packages, they must always build against the latest version of those dependences.
      ```
        github.com/ardanlabs/kit
        ├── CONTRIBUTORS
        ├── LICENSE
        ├── README.md
        ├── cfg/
        ├── examples/
        ├── log/
        ├── pool/
        ├── tcp/
        ├── timezone/
        ├── udp/
        └── web/
      ```
    - Application Projects
    > Each Application project contains three root level folders. These are cmd/, internal/ and vendor/. There is also a platform/ folder inside of the internal/ folder, which has different design constraints from the other packages that live inside of internal/.
    ```
    github.com/servi-io/api
    ├── cmd/
    │   ├── servi/
    │   │   ├── cmdupdate/
    │   │   ├── cmdquery/
    │   │   └── servi.go
    │   └── servid/
    │       ├── routes/
    │       │   └── handlers/
    │       ├── tests/
    │       └── servid.go
    ├── internal/
    │   ├── attachments/
    │   ├── locations/
    │   ├── orders/
    │   │   ├── customers/
    │   │   ├── items/
    │   │   ├── tags/
    │   │   └── orders.go
    │   ├── registrations/
    │   └── platform/
    │       ├── crypto/
    │       ├── mongo/
    │       └── json/
    └── vendor/
    ├── github.com/
    │   ├── ardanlabs/
    │   ├── golang/
    │   ├── prometheus/
    └── golang.org/
    ```
- [ ] [golang-standards-project-layout](https://github.com/golang-standards/project-layout)
https://github.com/golang-standards/project-layout/blob/master/README_zh.md
- [ ] [浅析VO、DTO、DO、PO的概念、区别和用处](https://www.cnblogs.com/zxf330301/p/6534643.html)
https://blog.csdn.net/k6T9Q8XKs6iIkZPPIFq/article/details/109192475?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522160561008419724839224387%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=160561008419724839224387&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_v2~rank_v28-6-109192475.first_rank_ecpm_v3_pc_rank_v2&utm_term=阿里技术专家详解DDD系列&spm=1018.2118.3001.4449
https://blog.csdn.net/chikuai9995/article/details/100723540?biz_id=102&utm_term=阿里技术专家详解DDD系列&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-0-100723540&spm=1018.2118.3001.4449
https://blog.csdn.net/Taobaojishu/article/details/101444324?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522160561008419724838528569%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=160561008419724838528569&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_click~default-1-101444324.first_rank_ecpm_v3_pc_rank_v2&utm_term=阿里技术专家详解DDD系列&spm=1018.2118.3001.4449
https://blog.csdn.net/taobaojishu/article/details/106152641
- [ ] [ Google API 的错误模型](https://cloud.google.com/apis/design/errors)
https://kb.cnblogs.com/page/520743/
https://zhuanlan.zhihu.com/p/105466656
https://zhuanlan.zhihu.com/p/105648986
https://zhuanlan.zhihu.com/p/106634373
https://zhuanlan.zhihu.com/p/107347593
https://zhuanlan.zhihu.com/p/109048532
https://zhuanlan.zhihu.com/p/110252394
https://www.jianshu.com/p/dfa427762975
https://www.citerus.se/go-ddd/
https://www.citerus.se/part-2-domain-driven-design-in-go/
https://www.citerus.se/part-3-domain-driven-design-in-go/
https://www.jianshu.com/p/dfa427762975
https://www.jianshu.com/p/5732b69bd1a1
https://www.cnblogs.com/qixuejia/p/10789612.html
https://www.cnblogs.com/qixuejia/p/4390086.html
https://www.cnblogs.com/qixuejia/p/10789621.html
https://zhuanlan.zhihu.com/p/46603988
https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
https://blog.csdn.net/taobaojishu/article/details/106152641
https://apisyouwonthate.com/blog/creating-good-api-errors-in-rest-graphql-and-grpc
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
https://www.youtube.com/watch?v=oL6JBUk6tj0
https://github.com/zitryss/go-sample
https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md
https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f
https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
https://medium.com/wtf-dial/wtf-dial-domain-model-9655cd523182
https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
https://hackernoon.com/trying-clean-architecture-on-golang-2-44d615bf8fdf
https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/
https://github.com/katzien/go-structure-examples
- [ ] [Ashley McNamara + Brian Ketelsen. Go best practices](https://www.youtube.com/watch?v=MzTcsI6tn-0)
https://www.appsdeveloperblog.com/dto-to-entity-and-entity-to-dto-conversion/
https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/
https://github.com/google/wire/blob/master/docs/best-practices.md
https://github.com/google/wire/blob/master/docs/guide.md
https://blog.golang.org/wire
https://github.com/google/wire
https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html
https://blog.golang.org/examples
https://blog.golang.org/subtests
https://blog.golang.org/cover
https://blog.golang.org/module-compatibility
https://blog.golang.org/v2-go-modules
https://blog.golang.org/publishing-go-modules
https://blog.golang.org/module-mirror-launch
https://blog.golang.org/migrating-to-go-modules
https://blog.golang.org/using-go-modules
https://blog.golang.org/modules2019
https://blog.codecentric.de/en/2017/08/gomock-tutorial/
https://pkg.go.dev/github.com/golang/mock/gomock
https://medium.com/better-programming/a-gomock-quick-start-guide-71bee4b3a6f1

