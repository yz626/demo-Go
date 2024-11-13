# Wire依赖注入框架

依赖注入的例子：
```Go
type userStore struct {
    db *gorm.DB
}

func NewStore(db *gorm.DB) *userStore {
    return &userStore{db: db}
}

func (u *userStore) Create(ctx context.Context, user *model.UserM) error {
    return u.db.Create(&user).Error
}
```
```Go
var db = NewDB()
store := NewStore(db)
store.Create(ctx, user)
```
这段代码就正在使用依赖注入。通过注入的方式，可以轻松的实现解耦，让代码更易维护。
构造`store`时,`*userStore`依赖`*gorm.DB`，我们使用构造函数 `NewStore` 创建`*userStore`
对象，并且将它的依赖对象`*gorm.DB`通过函数参数的形式注入进来。

## Wire安装
- 获取Wire依赖包
```
go get -u github.com/google/wire
```
- 导入Wire包
```Go
import "github.com/google/wire"
```
- 安装Wire工具包
```go
go install github.com/google/wire/cmd/wire
```

## Wire使用介绍
### 核心概念
- providers(提供者):可导出的Go函数，返回值是注入的依赖对象。比如`main.go`文件中的
`NewEvent`、`NewGreeter`、`NewMessage`,都是providers。
    - 在providers中，通过wire.NewSet(...)函数，可以组合多个providers，生成一个注入器。

- injectors(注入器):是一个函数，可以按照依赖顺序调用providers，注入依赖对象。比如`wire_gen.go`
文件中的`InitializeEvent`函数。

### 使用方式
- `wire.Build()`函数。

  函数原型`func Build(...interface{}) string`。