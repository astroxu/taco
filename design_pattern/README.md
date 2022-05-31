# design_pattern
> 设计模式：https://github.com/sevenelevenlee/go-patterns

> 一切设计模式都是灵活应用struct的组合模式，以及go隐形继承接口的特性
go中的interface就是一些方法装饰, 而struct并不依赖于接口

## 创建类
----

### [建造者模式(Builder Pattern)](./builder_pattern)
> 将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示
> 
> 当一个类的构造函数参数个数超过4个，而且这些参数有些是可选的参数，考虑使用构造者模式。

#### [工厂方法模式(Factory Method)](./factory_method_pattern)
> 使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化

#### [对象池模式(Object Pool)](./object_pool_pattern)
> 根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本

#### [单例模式(Singleton)](./singleton_pattern)
> 单例模式是最简单的设计模式之一, 保证一个类仅有一个实例, 并提供一个全局的访问接口

#### [生成器模式(Generator)](./generator_pattern)
> 生成器模式可以允许使用者在生成要使用的下一个值时与生成器并行运行

#### [抽象工厂模式(Abstract Factory)](./abstract_factory_pattern)
> 提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类

#### [原型模式(Prototype)](./prototype_pattern)
> 复制一个已存在的实例

## 结构类
---

#### [装饰模式(Decorator)](./decorator_pattern)
> 装饰模式使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能

#### [代理模式(Proxy)](./proxy_pattern)
> 代理模式用于延迟处理操作或者在进行实际操作前后对真实对象进行其它处理。

#### [适配器模式(Adapter)](./apapter_pattern)
> 将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作

#### [组合模式(Composite)](./composite_pattern)
> 组合模式有助于表达数据结构, 将对象组合成树形结构以表示"部分-整体"的层次结构, 常用于树状的结构

#### [享元模式(Flyweight)](./flyweight_pattern)
> 把多个实例对象共同需要的数据，独立出一个享元，从而减少对象数量和节省内存

#### [外观模式(Facade)](./facade_pattern)
> 外观模式在客户端和现有系统之间加入一个外观对象, 为子系统提供一个统一的接入接口, 类似与委托

