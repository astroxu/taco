# design_pattern
> 设计模式：https://github.com/sevenelevenlee/go-patterns

> 一切设计模式都是灵活应用struct的组合模式，以及go隐形继承接口的特性
go中的interface就是一些方法装饰, 而struct并不依赖于接口

## 创建类
----

### [建造者模式(builder pattern)](./builder_pattern)
> 将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示
> 
> 当一个类的构造函数参数个数超过4个，而且这些参数有些是可选的参数，考虑使用构造者模式。




