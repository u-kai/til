Designing for Modularity The goal of modularity is to make it easier and safer to make changes to a system. There are several ways modularity supports this goal. One is to remove duplication of implementations, to reduce the number of code changes you need to make to deliver a particular change. Another is to simplify implementation by providing components that can be assembled in different ways for different uses. A third way to make changes easier and safer is to design the system so that you can make changes to a smaller component without needing to change other parts of the system. Smaller pieces are easier, safer, and faster to change than larger pieces. Most design rules for modularity have a tension. Followed carelessly, they can actually make a system more brittle and harder to change. The four key metrics from “The Four Key Metrics” are a useful lens for considering the effectiveness of modularizing your system. Characteristics of Well-Designed Components Designing components is the art of deciding which elements of your system to group together, and which to separate. Doing this well involves understanding the relationships and dependencies between the elements. Two important design characteristics of a component are coupling and cohesion. The goal of a good design is to create low coupling and high cohesion. Coupling describes how often a change to one component requires a change to another component. Zero coupling isn’t a realistic goal for two parts of a system. Zero coupling probably means they aren’t a part of the same system at all. Instead, we aim for low, or loose coupling. A stack and a server image are coupled, because you may need to increase the memory allocated to the server instance in the stack when you upgrade software on the server. But you shouldn’t need to change code in the stack every time you update the server image. Low coupling makes it easier to change a component with little risk of breaking other parts of the system.
Cohesion describes the relationship between the elements within a component. As with coupling, the concept of cohesion relates to patterns of change. Changes to a resource defined in a stack with low cohesion are often not relevant to other resources in the stack. An infrastructure stack that defines separate networking structures for servers provisioned by two other stacks has low cohesion. Components with high cohesion are easier to change because they are smaller, simpler, and cleaner, with a lower blast radius (Blast Radius), than components that include a mash of loosely related things. Four Rules of Simple Design Kent Beck, the creator of XP and TDD, often cites four rules for making the design of a component simple. According to his rules, simple code should: Pass its tests (do what it is supposed to do) Reveal its intention (be clear and easy to understand) Have no duplication Include the fewest elements Rules for Designing Components Software architecture and design includes many principles and guidelines for designing components with low coupling and high cohesion. Avoid duplication The DRY (Don’t Repeat Yourself) principle says, “Every piece of knowledge must have a single, unambiguous, authoritative representation within a system.”1 Duplication forces people to make a change in
multiple places. For example, all of ShopSpinner’s stacks use a provisioner user account to apply configuration to server instances. Originally, the login details for the account are specified in every stack as well as the code that builds their base server image. When someone needs to change the login details for the user account, they need to find and change it in all of these locations in the codebase. So the team moves the login details in a central location, and each stack, plus the server image builder, refers to that location. Duplication Considered Useful The DRY principle discourages duplicating the implementation of a concept, which is not the same as duplicating literal lines of code. Having multiple components depend on shared code can create tight coupling, making it hard to change. I’ve seen teams insist on centralizing any code that looks similar; for example, having all virtual servers created using a single module. In practice, servers created for different purposes, such as application servers, web servers, and build servers, usually need to be defined differently. A module that needs to create all of these different types of servers can become overly complicated. When considering whether code is duplicated and should be centralized, consider whether the code truly represents the same concept. Does changing one instance of the code always mean the other instance should change as well? Also consider whether it’s a good idea to lock the two instances of code together into the same change cycle. Forcing every application server in the organization to upgrade at the same time may be unrealistic. Reuse increases coupling. So a good rule of thumb for reuse is to be DRY within a component, and wet across components. Rule of composition To make a composable system, make independent pieces. It should be easy to replace one side of a dependency relationship without disturbing the other.The ShopSpinner team starts with a single Linux application server image that they provision from different stacks. Later they add a Windows application server image. They design the code that provisions server images from any given stack so that they can easily switch between these two server images as needed for a given application. Single responsibility principle The single responsibility principle (SRP) says that any given component should have responsibility for one thing. The idea is to keep each component focused, so that its contents are cohesive.3 An infrastructure component, whether it’s a server, configuration library, stack component, or a stack, should be organized with a single purpose in mind. That purpose may be layered. Providing the infrastructure for an application is a single purpose that can be fulfilled by an infrastructure stack. You could break that purpose down into secure traffic routing for an application, implemented in a stack library; an application server, implemented by a server image; and a database instance, implemented by a stack module. Each component, at each level, has a single, easily understood purpose. Design components around domain concepts, not technical ones People are often tempted to build components around technical concepts. For example, it might seem like a good idea to create a component for defining a server, and reuse that in any stack that needs a server. In practice, any shared component couples all of the code it uses. A better approach is to build components around a domain concept. An application server is a domain concept that you may want to reuse for multiple applications. A build server is another domain concept, which you may want to reuse to give different teams their own instance. So these make better components than servers, which are probably used in different ways. Law of Demeter Also called the principle of least knowledge, the Law of Demeter says that a component should not have knowledge of how other components are implemented. This rules pushes for clear, simple interfaces between
between components. The ShopSpinner team initially violates this rule by having a stack that defines an application server cluster, and a shared networking stack that defines a load balancer and firewall rules for that cluster. The shared networking stack has too much detailed knowledge of the application server stack. Providers and Consumers In a dependency relationship between components a provider component creates or defines a resource that a consumer component uses. A shared networking stack may be a provider, creating networking address blocks such as subnets. An application infrastructure stack may be a consumer of the shared networking stack, provisioning servers and load balancers within the subnets managed by the provider. A key topic of this chapter is defining and implementing interfaces between infrastructure components.
モジュール性のデザイン モジュール性の目標は、システムの変更を容易かつ安全にすることです。モジュール性がこの目標を支持するための方法は複数あります。1つは実装の重複を削除することで、特定の変更を提供するために行う必要のあるコードの変更の数を減らすことです。もう1つは実装を単純化することで、異なる方法で異なる用途に組み立てることができるコンポーネントを提供することです。変更を容易かつ安全にする第3の方法は、システムを設計する際に他の部分を変更する必要なく、より小さなコンポーネントを変更できるようにすることです。小さな部分は大きな部分よりも変更が容易で安全で速いです。 モジュール性のためのほとんどの設計ルールは緊張を持っています。注意深くフォローされない場合、実際にはシステムをもっと壊れやすく、変更が難しくする可能性があります。 "The Four Key Metrics"からの4つの主要な指標は、システムのモジュール化の効果を考慮するための有用な手段です。 よく設計されたコンポーネントの特性 コンポーネントの設計は、システムのどの要素をグループ化し、どの要素を分離するかを決定する芸術です。これをうまく行うには、要素間の関係や依存関係を理解することが重要です。 コンポーネントの2つの重要な設計特性はカップリングと結束です。良い設計の目標は低いカップリングと高い結束を作り出すことです。 カップリングは、1つのコンポーネントの変更が他のコンポーネントの変更を必要とする頻度を説明します。システムの2つの部分の間でカップリングがゼロであることは現実的な目標ではありません。ゼロのカップリングはおそらくそれらがまったく同じシステムの一部ではないことを意味しています。代わりに、低いまたはゆるいカップリングを目指します。スタックとサーバーイメージはカップリングされています。スタックのサーバーインスタンスでソフトウェアをアップグレードする際に、スタック内でメモリを割り当てる必要があるかもしれません。ただし、サーバーイメージを更新するたびにスタックのコードを変更する必要はありません。低いカップリングは、システムの他の部分を壊すリスクが少ないままコンポーネントを変更することを容易にします。 結束は、コンポーネント内の要素間の関係を説明します。カップリングと同様に、結束の概念は変更のパターンに関係しています。結束が低いスタック内で定義されたリソースの変更は、スタックの他のリソースには通常関連性がありません。2つの別々のスタックによって提供される異なるサーバーのネットワーキング構造を定義するインフラストラクチャスタックには結束が低いです。高い結束を持つコンポーネントは、より小さく、シンプルで、クリーンであり、より低いブラスト半径（Blast Radius）を持つ、緩く関連付けられたもののまとまりを含むコンポーネントよりも変更が容易です。 シンプルな設計の4つのルール XPおよびTDDの作成者であるKent Beckは、コンポーネントの設計をシンプルに行うための4つのルールをよく引用します。彼のルールによると、シンプルなコードは次のような特徴を持つべきです。 テストに合格する（するべきことをする） 意図を明示する（明確で理解しやすい） 重複を含まない 最小の要素を含む コンポーネントの設計ルール 低いカップリングと高い結束を持つコンポーネントを設計するためのソフトウェアアーキテクチャとデザインには、多くの原則とガイドラインが含まれます。 重複を避ける DRY（Don't Repeat Yourself）の原則は、「システム内のすべての知識は、一意で明確に正確な表現を持たなければならない」と言っています。重複は人々に複数の場所で変更を行わせます。たとえば、ShopSpinnerのすべてのスタックは、サーバーインスタンスに設定を適用するためにプロビジョナーユーザーアカウントを使用しています。元々、アカウントのログイン詳細はすべてのスタックで指定され、ベースのサーバーイメージを構築するためのコードにも指定されていました。ユーザーアカウントのログインの詳細を変更する必要がある場合、コードベースのすべての場所でそれを見つけて変更する必要があります。したがって、チームはログインの詳細を中央の場所に移動し、各スタックとサーバーイメージビルダーはその場所を参照するようにしました。 重複は有用な考慮されるべきです DRYの原則は、コンセプトの実装の重複を防ぐことを推奨していますが、これは文字通りのコードの行の重複とは異なります。複数のコンポーネントが共有のコードに依存すると、強いカップリングが生じ、変更が困難になります。私はチームが類似したコードを必ず集中させることを主張するのを見てきました。たとえば、単一のモジュールを使用して作成されたすべての仮想サーバーを持つことです。実際には、アプリケーションサーバー、Webサーバー、ビルドサーバーなど、異なる目的のために作成されたサーバーは通常異なる方法で定義する必要があります。さまざまな種類のサーバーを作成する必要があるモジュールは過度に複雑になる可能性があります。コードが重複しているかどうかを検討し、それを集中させるべきかどうかを検討する際には、コードが本当に同じコンセプトを表しているかどうかを考慮してください。コードの1つのインスタンスを変更すると常に他のインスタンスも変更する必要があるのかを考慮してください。また、2つのコードのインスタンスを同じ変更サイクルに結びつけるのは良い考えなのかを考慮してください。組織内のすべてのアプリケーションサーバーを同時にアップグレードすることを強制することは現実的ではないかもしれません。 再利用はカップリングを増加させます。したがって、再利用のための良い指針は、コンポーネント内ではDRYであり、コンポーネント間ではWETであることです。 コンポーズルール 作曲可能なシステムを作るためには、独立した部分を作ります。依存関係の関係の片方を変更することなく、依存関係がある片方を交換できるようにするのが容易であるべきです。 ShopSpinnerチームは最初に個別のLinuxアプリケーションサーバーイメージを使用してプロビジョンするコードを作成します。後でWindowsアプリケーションサーバーイメージを追加します。どのアプリ