Challenges with Testing Infrastructure Code Most of the teams I encounter that work with Infrastructure as Code struggle to implement the same level of automated testing and delivery for their infrastructure code as they have for their application code. And many teams without a background in Agile software engineering find it even more difficult. The premise of Infrastructure as Code is that we can apply software engineering practices such as Agile testing to infrastructure. But there are significant differences between infrastructure code and application code. So we need to adapt some of the techniques and mindsets from application testing to make them practical for infrastructure. The following are a few challenges that arise from the differences between infrastructure code and application code. Challenge: Tests for Declarative Code Often Have Low Value As mentioned in Chapter 4 (“Declarative Infrastructure Languages”), many infrastructure tools use declarative languages rather than imperative languages. Declarative code typically declares the desired state for some infrastructure, such as this code that defines a networking subnet: Example 8-1. subnet:
name: private_A
address_range: 192.168.0.0/16 A test for this would simply restate the code: assert:
subnet("private_A").exists
assert:
subnet("private_A").address_range is("192.168.0.0/16") A suite of low-level tests of declarative code can become a bookkeeping exercise. Every time you change the infrastructure code, you change the test to match. What value do these tests provide? Well, testing is about managing risks, so let’s consider what risks the preceding test can uncover: The infrastructure code was never applied.

The infrastructure code was applied, but the tool failed to apply it correctly, without returning an error. Someone changed the infrastructure code but forgot to change the test to match. The first risk may be a real one, but it doesn’t require a test for every single declaration. Assuming you have code that does multiple things on a server, a single test would be enough to reveal that, for whatever reason, the code wasn’t applied. The second risk boils down to protecting yourself against a bug in the tool you’re using. The tool developers should fix that bug or your team should switch to a more reliable tool. I’ve seen teams use tests like this in cases where they found a specific bug, and wanted to protect themselves against it. Testing for this is okay to cover a known issue, but it is wasteful to blanket your code with detailed tests just in case your tool has a bug. The last risk is circular logic. Removing the test would remove the risk it addresses, and also remove work for the team. Declarative Tests The Given, When, Then format is useful for writing tests.8 A declarative test omits the “When” part, having a format more like “Given a thing, Then it has these characteristics.” Tests written like this suggest that the code you’re testing doesn’t create variable outcomes. Declarative tests, like declarative code, have a place in many infrastructure codebases, but be aware that many tools and practices for testing dynamic code may not be appropriate. There are some situations when it’s useful to test declarative code. Two that come to mind are when the declarative code can create different results, and when you combine multiple declarations. Testing variable declarative code The previous example of declarative code is simple — the values are hardcoded, so the result of applying the code is clear. Variables introduce the possibility of creating different results, which may create risks that make testing more useful. Variables don’t always create variation that needs testing. What if we add some simple variables to the earlier example?
subnet:
name: ${MY_APP}-${MY_ENVIRONMENT}
address_range: ${SUBNET_IP_RANGE} There isn’t much risk in this code that isn’t already managed by the tool that applies it. If someone sets the variables to invalid values, the tool should fail with an error. The code becomes riskier when there are more possible outcomes. Let’s add some conditional code to the example: subnet:
  name: ${MY_APP}-${MY_ENVIRONMENT}
address_range: get_networking_subrange(
get_vpc(${MY_ENVIRONMENT}),
data_centers.howmany,
data_centers.howmany++
) This code has some logic that might be worth testing. It calls two functions, get_networking_subrange and get_vpc, either of which might fail or return a result that interacts in unexpected ways with the other function. The outcome of applying this code varies based on inputs and context, which makes it worth writing tests. Note Imagine that instead of calling these functions, you wrote the code to select a subset of the address range as a part of this declaration for your subnet. This is an example of mixing declarative and imperative code (as discussed in “Separate Declarative and Imperative Code”). The tests for the subnet code would need to include various edge cases of the imperative code — for example, what happens if the parent range is smaller than the range needed? If your declarative code is complex enough that it needs complex testing, it is a sign that you should pull some of the logic out of your declarations and into a library written in a procedural language. You can then write clearly separate tests for that function, and simplify the test for the subnet declaration.
Testing combinations of declarative code Another situation where testing is more valuable is when you have multiple declarations for infrastructure that combine into more complicated structures. For example, you may have code that defines multiple networking structures — an address block, load balancer, routing rules, and gateway. Each piece of code would probably be simple enough that tests would be unnecessary. But the combination of these produces an outcome that is worth testing — that someone can make a network connection from point A to point B. Testing that the tool created the things declared in code is usually less valuable than testing that they enable the outcomes you want.
インフラストラクチャコードのテストにおける課題
インフラストラクチャをコードとして扱うチームのほとんどが、アプリケーションコードと同じレベルの自動化されたテストとデリバリーを実装することに苦労しています。また、アジャイルソフトウェアエンジニアリングのバックグラウンドがないチームは、さらに困難と感じることが多いです。インフラストラクチャコードは、アジャイルなテストなどのソフトウェアエンジニアリングのプラクティスを適用できるという前提です。しかし、インフラストラクチャコードとアプリケーションコードには重要な違いがあります。そのため、アプリケーションのテストにおける一部の手法や考え方を適用し、インフラストラクチャ用に実用的にする必要があります。以下は、インフラストラクチャコードとアプリケーションコードの違いから生じるいくつかの課題です。

課題: 宣言的なコードのテストは通常低い価値しか持ちません
第4章の「宣言的インフラストラクチャ言語」で述べたように、多くのインフラストラクチャツールは、命令型言語ではなく宣言型言語を使用しています。宣言的コードは通常、ネットワーキングサブネットを定義する次のようなコードのように、特定のインフラストラクチャのための所望の状態を宣言します。

subnet:
  name: private_A
  address_range: 192.168.0.0/16

このため、テストでは通常、コードをそのまま記述します。

assert:
  subnet("private_A").exists
assert:
  subnet("private_A").address_range is("192.168.0.0/16")

宣言的なコードの低レベルなテストスイートは、煩雑な作業になる可能性があります。インフラストラクチャコードを変更するたびに、テストも変更する必要があります。これらのテストはどのような価値を提供するのでしょうか？テストはリスクを管理することに関連していますので、前述のテストが発見できるリスクを考えてみましょう。

- インフラストラクチャコードが適用されていなかった。
- インフラストラクチャコードが適用されたが、ツールが正しく適用されず、エラーも返されていなかった。
- 誰かがインフラストラクチャコードを変更したが、テストを変更し忘れた。

最初のリスクは実際に起こりうるものかもしれませんが、すべての宣言に対してテストする必要はありません。サーバ上で複数のことを行うコードがある場合を想定して、単一のテストだけで、何らかの理由でコードが適用されなかったことを示せば十分です。2つ目のリスクは、使用しているツールのバグに対する保護です。そのバグをツールの開発者が修正するか、チームがより信頼性の高いツールに切り替えるべきです。私は、チームが特定のバグを見つけ、それに対して自己保護をしたい場合に、このようなテストを使用するチームを見てきました。既知の問題をカバーするためにこれをテストすることは問題ありませんが、ツールにバグがあるかもしれないという理由だけで細かいテストをコード全体に適用するのは無駄です。最後のリスクは循環論理です。テストを削除すれば、それが対処しているリスクも取り除き、チームの作業量も減らすことができます。

宣言的テスト
テストを書くためには、Given、When、Thenフォーマットが有用です。宣言的なテストは「Given a thing, Then it has these characteristics」といった形式で、テストしているコードが変数の結果を生成しないということを示します。宣言的なテストは、宣言的なコードと同様に、多くのインフラストラクチャ用コードベースに適していますが、多くの動的コードのテスト手法やプラクティスは適切ではないことに注意してください。

宣言的コードをテストする場合には、いくつかの状況で有用です。一つは、宣言的なコードが異なる結果を作り出す場合です。もう一つは、複数の宣言を組み合わせた場合です。宣言的なコードの変数をテストする場合、前の例は単純で、値がハードコードされているため、コードの適用結果が明確です。変数を追加してみましょう。以下のようなシンプルな変数を追加した場合、このコードには既に適用されているツールによって管理されているリスクはほとんどありません。

subnet:
  name: ${MY_APP}-${MY_ENVIRONMENT}
  address_range: ${SUBNET_IP_RANGE}

誰かが変数を無効な値に設定した場合、ツールはエラーで失敗するはずです。このコードは、より多くの結果を生み出す場合にリスクが増します。次に、例にいくつかの条件付きコードを追加してみましょう。

subnet:
  name: ${MY_APP}-${MY_ENVIRONMENT}
  address_range: get_networking_subrange(
    get_vpc(${MY_ENVIRONMENT}),
    data_centers.howmany,
    data_centers.howmany++
  )

このコードにはいくつかのロジックがあり、テストする価値があるかもしれません。get_networking_subrangeとget_vpcという2つの関数を呼び出しており、どちらかが失敗したり、予期しない方法で他の関数と干渉したりする可能性があります。このコードを適用することの結果は、入力とコンテキストに基づいて変化するため、テストを書く価値があります。

Imagine that instead of calling these functions, you wrote the code to select a subset of the address range as a part of this declaration for your subnet. This is an example of mixing declarative and imperative code (as discussed in "Separate Declarative and Imperative Code"). The tests for the subnet code would need to include various edge cases of the imperative code — for example, what happens if the parent range is smaller than the range needed? If your declarative code is complex enough that it needs complex testing, it is a sign that you should pull some of the logic out of your declarations and into a library written in a procedural language. You can then write clearly separate tests for that function, and simplify the test for the subnet declaration.

宣言的コードの組み合わせをテストするもう一つの有用なシチュエーションは、複数のインフラストラクチャの宣言を組み合わせてより複雑な構造を作成する場合です。たとえば、アドレスブロック、ロードバランサー、ルーティングルール、およびゲートウェイを定義するコードがあるかもしれません。それぞれのコードはおそらく単純すぎてテストが不要ですが、これらの組み合わせによって望む結果が生み出されるかどうかをテストすることは価値があります。コードが宣言したものがツールによって作成されたかどうかをテストすることは通常それほど価値がありませんが、望む結果が実現されたかどうかをテストすることはより重要です。