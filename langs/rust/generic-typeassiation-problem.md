# Generic と TypeAssiation 関連で生じた問題

-   下のような一つのトレイト中の一つのメソッドにしか Generic を使わないにも関わらず，トレイト全体に Generic が利用されており，さらに，そのトレイトを実装する struct がトレイトで指定されている Generic な型の情報がなく，トレイトの Generic は struct のメソッドを利用する時に外部から差し込むケースでうまくいかなくなった

```rust
trait Convertor {
    fn convert(&self, str: &str) -> String;
}

struct ConcreteConvertor {}
impl Convertor for ConcreteConvertor {
    fn convert(&self, str: &str) -> String {
        format!("!!{}!!", str)
    }
}

trait Generator<T>
where
    T: Convertor,
{
    fn gen(&self, c: T) -> String;
    fn gen_original(&self) -> String;
}

struct ConcriteGenerator {
    data: String,
}

impl<T> Generator<T> for ConcriteGenerator
where
    T: Convertor,
{
    fn gen(&self, c: T) -> String {
        c.convert(&self.data)
    }
    fn gen_original(&self) -> String {
        self.data.clone()
    }
}
fn main() {
    let generator = ConcriteGenerator {
        data: "Error".to_string(),
    };
    let convertor = ConcreteConvertor {};
    println!("can complie case : {}", generator.gen(convertor));
    // type assiated error
    println!("can not complie case : {}", generator.gen_original());
}
```

-   回避するために trait ではなく，必要な関数だけに Generic を追いやることで上記のケースはクリア

```rust
trait Convertor {
    fn convert(&self, str: &str) -> String;
}

struct ConcreteConvertor {}
impl Convertor for ConcreteConvertor {
    fn convert(&self, str: &str) -> String {
        format!("!!{}!!", str)
    }
}

trait Generator {
    fn gen<T: Convertor>(&self, c: T) -> String;
    fn gen_original(&self) -> String;
}

struct ConcriteGenerator {
    data: String,
}

impl Generator for ConcriteGenerator {
    fn gen<T: Convertor>(&self, c: T) -> String {
        c.convert(&self.data)
    }
    fn gen_original(&self) -> String {
        self.data.clone()
    }
}
fn main() {
    let generator = ConcriteGenerator {
        data: "Error".to_string(),
    };
    let convertor = ConcreteConvertor {};
    println!("can complie case : {}", generator.gen(convertor));
    println!("can complie case : {}", generator.gen_original());
}

```

-   ちなみに，ある struct について具体的に trait を実装する際に，Generic 部分が具体的にわかっていても実装時　(実際に使っているのではなく，定義している時ということ)には具体的な型を入れることができず，Generic のままにしなくてはいけないかも
    -   もしかしたら型定義をしっかり書けばいけるのか？
    -   それとも関連型を使えばいけるのか？
    -   普通にできてもおかしくなさそうだけどな

```rust
trait Convertor {
    fn convert(&self, str: &str) -> String;
}

struct ConcreteConvertor {}
impl Convertor for ConcreteConvertor {
    fn convert(&self, str: &str) -> String {
        format!("!!{}!!", str)
    }
}

trait Generator {
    fn gen<T: Convertor>(&self, c: T) -> String;
    fn gen_original(&self) -> String;
}

struct ConcriteGenerator {
    data: String,
}

impl Generator for ConcriteGenerator {
    // error!!
    // ConcriteGeneratorがConcreteConvertorを利用する場合であってもGenericにするべし？
    fn gen(&self, c: ConcreteConvertor) -> String {
        c.convert(&self.data)
    }
    fn gen_original(&self) -> String {
        self.data.clone()
    }
}

```
