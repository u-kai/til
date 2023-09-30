Benefits of Polymorphism.
There really is only one benefit to Polymorphism; but it’s a big one. It is the inversion of source code and run time dependencies.

In most software systems when one function calls another, the runtime dependency and the source code dependency point in the same direction. The calling module depends on the called module. However, when polymorphism is injected between the two there is an inversion of the source code dependency. The calling module still depends on the called module at run time. However, the source code of the calling module does not depend upon the source code of the called module. Rather both modules depend upon a polymorphic interface.

This inversion allows the called module to act like a plugin. Indeed, this is how all plugins work.

Plugin architectures are very robust because stable high value business rules can be kept from depending upon volatile low value modules such as user interfaces and databases.

The net result is that in order to be robust a system must employ polymorphism across significant architectural boundaries.

Benefits of Immutability
The benefit of not using assignment statements should be obvious. You can’t have concurrent update problems if you never update anything.

Since functional programming languages do not have assignment statements, programs written in those languages don’t change the state of very many variables. Mutation is reserved for very specific sections of the system that can tolerate the high ceremony required. Those sections are inherently safe from multiple threads and multiple cores.

The bottom line is that functional programs are much safer in multiprocessing and multiprocessor environments.

The Deep Philosophies
Of course adherents to both Object Orientation and Functional Programming will protest my reductionist analysis. They will contend that there are deep philosophical, psychological, and mathematical reasons why their favorite style is better than the other.

My reaction to that is: Phooey!

Everybody thinks their way is the best. Everybody is wrong.

What about Design Principles, and Design Patterns?
The passage at the start of this article that irked me suggests that all the design principles and design patterns that we’ve identified over the last several decades apply only to OO; and that Functional Programming reduces them all down to: functions.

Wow! Talk about being reductionist!

This idea is bonkers in the extreme. The principles of software design still apply, regardless of your programming style. The fact that you’ve decided to use a language that doesn’t have an assignment operator does not mean that you can ignore the Single Responsibility Principle; or that the Open Closed Principle is somehow automatic. The fact that the Strategy pattern makes use of polymorphism does not mean that the pattern cannot be used in a good functional language[4].

The bottom, bottom line here is simply this. OO programming is good, when you know what it is. Functional programming is good when you know what it is. And functional OO programming is also good once you know what it is.

ポリモーフィズムの利点:
ポリモーフィズムの利点は基本的に一つですが、その影響は大きいです。それは、ソースコードと実行時の依存関係の反転です。

ほとんどのソフトウェアシステムでは、一つの関数が別の関数を呼び出すとき、ランタイム依存性とソースコード依存性は同じ方向を指しています。つまり呼び出すモジュールは呼び出されるモジュールに依存しています。しかし、ポリモーフィズムがこの間に注入されると、ソースコードの依存関係は反転します。ランタイムでは呼び出すモジュールが呼び出されるモジュールに依存していますが、呼び出すモジュールのソースコードは呼び出されるモジュールのソースコードに依存していません。どちらのモジュールもポリモーフィックなインターフェースに依存します。

この反転により、呼び出されるモジュールはプラグインのように動作し、これがすべてのプラグインがどのように動作するかです。 

プラグインアーキテクチャは非常に頑強で、安定した高価値のビジネスルールを、ユーザインターフェースやデータベースなど、変動しやすい低価値のモジュールから依存させないことが可能です。

結果として、頑強なシステムを実現するためには、主要なアーキテクチャ境界でポリモーフィズムを活用する必要があります。

不変性の利点:
代入文を使用しない利点は明らかで、何も更新しなければ同時更新の問題が生じません。

関数型プログラミング言語が代入文を持たないため、それらの言語で書かれたプログラムは変数の状態をあまり変更しません。変化は、高い手順を要求する特定のシステムのセクションで予約されています。これらのセクションは、複数のスレッドや複数のコアから自然に保護されています。

要するに、関数型プログラムはマルチプロセッシングやマルチプロセッサの環境でより安全です。

深淵な思想系:
もちろん、オブジェクト指向と関数型プログラミングの両方の支持者は、私の還元主義的な分析に反論するでしょう。彼らは哲学的、心理学的、数学的理由から、自分たちの好きなスタイルが他のものより優れていると主張します。

しかし私の反応は:それは馬鹿げている!

全員が自分たちの方法が最善だと考えています。全員が間違っています。

デザイン原則とデザインパターンについてはどうなのか？
先に挙げた文章で私が気になったのは、過去数十年にわたり我々が確認してきたすべてのデザイン原則やデザインパターンが、関数型プログラミングでは「関数」にまで還元されるとする考え方です。

これは極度に還元主義的です。

この考えは極めて奇妙です。プログラミングスタイルに関わらず、ソフトウェア設計の原理は依然として適用されます。代入演算子を持たない言語を使用することを決定したとしても、一責任原則を無視できるわけではないし、開放閉鎖原則が何らかの形で自動的に適用されるわけではありません。戦略パターンがポリモーフィズムを使用することで、そのパターンが良好な関数型言語で使用できないわけではありません。

ここでの最も重要な点は、オブジェクト指向プログラミングも関数型プログラミングも、どちらも理解した上で適切に使用すれば良いということです。そして、関数型オブジェクト指向プログラミングも同様です。