OO imposes discipline on function pointers.

“Huh?” you say. But that, in fact, is what OO comes down to. For all the hi-falutin’ rhetoric about OO and “real-world objects” and “programming closer to the way we think”, what OO really comes down to is that OO languages replace function pointers with convenient polymorphism. [2]

How do you implement polymorphism? You implement it with function pointers. OO languages simply do that implementation for you, and hide the function pointers from you. This is nice because function pointers are very difficult to manage well. Trying to write polymorphic code with function pointers (as in C) depends on complex and inconvenient conventions that everyone must follow in every case. This is usually unrealistic.

In Java, every function you call is polymorphic. There is no way you can call a function that is not polymorphic. And that means that every java function is called indirectly through a pointer to a function.[3]

If you wanted polymophism in C, you’d have to manage those pointers yourself; and that’s hard. If you wanted polymorphism in Lisp you’d have to manage those pointers yourself (pass them in as arguments to some higher level algorithm (which, by the way IS the Strategy pattern.)) But in an OO language, those pointers are managed for you. The language takes care to initialize them, and marshal them, and call all the functions through the

OOは、関数ポインタに規律を課す。

「え？」と思うかもしれませんが、それが実際にはOOの本質です。OOや「現実世界のオブジェクト」、「考え方に近いプログラミング」についての高尚なる言葉遣いにも関わらず、OOの真の要点は、OO言語が関数ポインタを便利なポリモーフィズムに置き換える、ということです。

どのようにポリモーフィズムを実装しますか？それは関数ポインタを使って実装します。OO言語は単純にその実装を代行し、関数ポインタを隠蔽します。それは関数ポインタの管理が非常に難しいからです。関数ポインタを使ってポリモーフィックなコードを書こうとすると（C言語のように）、全員が全ての場合に従わなければならない複雑で煩わしい規則に依存します。これは通常、非現実的です。

Javaでは、呼び出す全ての関数がポリモーフィックです。ポリモーフィックでない関数を呼び出す方法はありません。つまり、すべてのJava関数は関数ポインタを介して間接的に呼び出されます。

Cでポリモーフィズムを実現したい場合、自分でそのポインタを管理しなければならないでしょう。それは難しいです。 ポリモーフィズムをLispで望むなら、自分でそのポインタを管理する必要があります（それらを高レベルのアルゴリズムの引数として渡す必要があります（これがストラテジーパターンです。））。ただし、OO言語では、これらのポインタが自動で管理されます。言語はそれらを初期化し、並べ替え、すべての関数をその経由で呼び出します。