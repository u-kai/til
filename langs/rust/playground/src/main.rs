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
    println!("can not complie case : {}", generator.gen_original());
}
