# 継承について考察

- 継承をされているサブクラスは定義されてから使われるまでずっと，そのサブクラス
  - ずっと依存関係を持つことになる
    - そのため，共通メソッドを変更，機能追加しにくくなる
      - 引数とか返り値とかそういう話だと思う

```javascript
class Super {
  constructor(id, name) {
    this.id = id;
    this.name = name;
  }
  super = () => {
    return this.id + this.name;
  };
}

class Sub extends Super {
  constructor(id, name, data) {
    super(id, name);
    this.data = data;
  }
  super = () => {
    return super.super() + this.data;
  };
}
```
