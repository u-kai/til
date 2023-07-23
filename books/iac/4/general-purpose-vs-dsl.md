General-Purpose Languages Versus DSLs for Infrastructure Most infrastructure DSLs are declarative languages rather than imperative languages. An internal DSL like Chef is an exception, although even Chef is primarily declarative.6 One of the biggest advantages of using a general-purpose programming language, such as JavaScript, Python, Ruby, or TypeScript, is the ecosystem of tools. These languages are very well supported by IDEs,7 with powerful productivity features like syntax highlighting and code refactoring. Testing support is an especially useful part of a programming language’s ecosystem. Many infrastructure testing tools exist, some of which are listed in “Verification: Making Assertions About Infrastructure Resources” and “Testing Server Code”. But few of these integrate with languages to support unit testing. As we’ll discuss in “Challenge: Tests for Declarative Code Often Have Low Value”, this may not be an issue for declarative code. But for code that produces more variable outputs, such as libraries and abstraction layers, unit testing is essential.
インフラストラクチャのための汎用言語とDSL
多くのインフラストラクチャDSLは、命令型の言語ではなく宣言型の言語です。Chefのような内部DSLは例外ですが、それでも主に宣言的です。ジェネラルパーパスのプログラミング言語（JavaScript、Python、Ruby、TypeScriptなど）を使用する最大の利点の1つは、その言語のエコシステムです。これらの言語は、シンタックスハイライトやコードのリファクタリングといった強力な生産性機能を備えたIDEで非常によくサポートされています。テストのサポートは、プログラミング言語のエコシステムの特に有用な部分です。多くのインフラストラクチャテストツールが存在し、その一部は「検証：インフラストラクチャリソースについての断定を行う」と「サーバーコードのテスト」にリストされています。しかし、これらのほとんどがユニットテストをサポートする