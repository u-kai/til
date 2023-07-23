If you work in a team that builds and runs IT infrastructure, then cloud and infrastructure automation technology should help you deliver more value in less time, and to do it more reliably. But in practice, it drives ever-increasing size, complexity, and diversity of things to manage. These technologies are especially relevant as organizations become digital. “Digital” is how people in business attire say that software systems are essential to what the organization does.1 The move to digital increases the pressure on you to do more and to do it faster. You need to add and support more services. More business activities. More employees. More customers, suppliers, and other stakeholders. Cloud and automation tools help by making it far easier to add and change infrastructure. But many teams struggle to find enough time to keep up with the infrastructure they already have. Making it easier to create even more stuff to manage is unhelpful. As one of my clients told me, “Using cloud knocked down the walls that kept our tire fire contained."2 Many people respond to the threat of unbounded chaos by tightening their change management processes. They hope that they can prevent chaos by limiting and controlling changes. So they wrap the cloud in chains. There are two problems with this. One is that it removes the benefits of using cloud technology; the other is that users want the benefits of cloud technology. So users bypass the people who are trying to limit the chaos. In the worst cases, people completely ignore risk management, deciding it’s not relevant in the brave new world of cloud. They embrace cowboy IT, which adds different problems.3 The premise of this book is that you can exploit cloud and automation technology to make changes easily, safely, quickly, and responsibly. These benefits don’t come out of the box with automation tools or cloud platforms. They depend on the way you use this technology.
In this chapter, I explain that modern, dynamic infrastructure requires a “Cloud Age” mindset. This mindset is fundamentally different from the traditional “Iron Age” approach we used with static pre-cloud systems. I define three core practices for implementing Infrastructure as Code: define everything as code, continuously test and deliver everything as you work, and build your system from small, loosely coupled pieces. Also in this chapter, I describe the reasoning behind the Cloud Age approach to infrastructure. This approach discards the false dichotomy of trading speed for quality. Instead, we use speed as a way to improve quality, and we use quality to enable delivery at speed. From the Iron Age to the Cloud Age Cloud Age technologies make it faster to provision and change infrastructure than traditional, Iron Age technologies (Table 1-1). Table 1-1. Technology changes in the Cloud Age Iron Age Cloud Age Physical hardware Virtualized resources Provisioning takes weeks Provisioning takes minutes Manual processes Automated processes However, these technologies don’t necessarily make it easier to manage and grow your systems. Moving a system with technical debt onto unbounded cloud infrastructure accelerates the chaos. Maybe you could use well-proven, traditional governance models to control the speed and chaos that newer technologies unleash. Thorough, up-front design, rigorous change review, and strictly segregated responsibilities will impose order! Unfortunately, these models optimize for the Iron Age, where changes are slow and expensive. They add extra work up front, hoping to reduce the time spent making changes later. This arguably makes sense when making changes later is slow and expensive. But cloud makes changes cheap and fast. You should exploit this speed to learn and improve your system continuously. Iron Age ways of working are a massive tax on learning and improvement. Rather than using slow-moving Iron Age processes with fast-moving Cloud Age technology, adopt a new mindset. Exploit faster-paced technology to reduce risk and improve quality. Doing this requires a fundamental change of approach and new ways of thinking about change and risk (Table 1-2).
Infrastructure as Code is a Cloud Age approach to managing systems that embraces continuous change for high reliability and quality. Infrastructure as Code Infrastructure as Code is an approach to infrastructure automation based on practices from software development. It emphasizes consistent, repeatable routines for provisioning and changing systems and their configuration. You make changes to code, then use automation to test and apply those changes to your systems. Throughout this book, I explain how to use Agile engineering practices such as Test Driven Development (TDD), Continuous Integration (CI), and Continuous Delivery (CD) to make changing infrastructure fast and safe. I also describe how modern software design can create resilient, well-maintained infrastructure. These practices and design approaches reinforce each other. Well-designed infrastructure is easier to test and deliver. Automated testing and delivery drive simpler and cleaner design. Benefits of Infrastructure as Code To summarize, organizations adopting Infrastructure as Code to manage dynamic infrastructure hope to achieve benefits, including: Using IT infrastructure as an enabler for rapid delivery of value
IT インフラストラクチャを構築・運用するチームで働く方は、クラウドとインフラ自動化技術を活用することで、より短い時間でより多くの価値を提供し、信頼性を高めることができるはずです。しかし、実際には、これらの技術は管理する対象のサイズ、複雑さ、多様性がますます増していく原因となります。これらの技術は、組織がデジタル化するにつれて特に関連性があります。ビジネス用語で言えば、「デジタル（digital）」というのは、ソフトウェアシステムが組織の中心的な役割を果たすことを意味します。デジタル化が進むことで、あなたはより多くのサービス、ビジネス活動、従業員、顧客、サプライヤー、その他の利害関係者を追加・サポートする必要があります。クラウドと自動化ツールは、インフラストラクチャの追加や変更をはるかに容易にすることで助けてくれます。しかし、多くのチームは既存のインフラストラクチャに追いつくための十分な時間を見つけることに苦労しています。ますます多くの管理対象を作成することを容易にすることは役に立ちません。クライアントの一人が私に語ったように、「クラウドを使うことで、私たちのタイヤ火事を抑えるための壁が取り壊された」ということです。多くの人々は、無制限の混沌の脅威に対処するために、変更管理プロセスを厳格化することで応えます。彼らは変更を制限し、管理することで混沌を防ぐことができると期待しています。彼らはクラウドを鎖で縛りたいのです。しかし、このアプローチには 2 つの問題があります。1 つは、クラウドテクノロジーの利点が失われることです。もう 1 つは、ユーザーがクラウドテクノロジーの利点を求めていることです。そのため、ユーザーは混沌を制限しようとしている人々をバイパスします。最悪の場合、人々はリスク管理を完全に無視し、クラウドの新たな世界では関係のないものと考えます。彼らはカウボーイ的な IT を受け入れますが、これは別の問題を引き起こします。本書の前提は、クラウドと自動化技術を利用することで、変更を容易に、安全に、迅速に、責任を持って行えるというものです。これらの利点は、自動化ツールやクラウドプラットフォームには最初から備わっているわけではありません。これらの利点は、技術の使用方法に依存します。
本章では、現代のダイナミックなインフラストラクチャには「クラウド時代」のマインドセットが必要であることを説明します。このマインドセットは、静的なクラウドプリクラウドシステムでは使用していた従来の「鉄時代」のアプローチとは根本的に異なります。私は、Infrastructure as Code を実装するための 3 つの核心プラクティスを定義します：すべてをコードとして定義し、作業中にすべてを継続的にテストして提供し、小さな相互に関連しない部分からシステムを構築します。また、本章では、インフラストラクチャに対するクラウド時代のアプローチの理論を説明します。このアプローチでは、速度と品質をトレードオフにするという誤った二分法を捨て、速度を品質向上の手段として利用し、品質を高めることで速度を実現します。鉄時代からクラウド時代への移行では、クラウド時代のテクノロジーによって、従来の鉄時代のテクノロジーよりもインフラストラクチャの提供と変更がより迅速に行われるようになります（表 1-1）。表 1-1.クラウド時代のテクノロジーの変化 TABLE 1-1. クラウド時代のテクノロジーの変化鉄時代クラウド時代物理的なハードウェア仮想化されたリソース提供に数週間かかる提供に数分かかる手動のプロセス自動化されたプロセスしかし、これらのテクノロジーはシステムの管理と成長を容易にするわけではありません。技術的な負債を抱えたシステムを無制限なクラウドインフラストラクチャに移行させることは、混沌を加速させます。新しい技術が解き放つ速度と混沌を制御するために、確立された伝統的なガバナンスモデルを使用することも考えられるかもしれません。徹底的な最初の設計、厳格な変更レビュー、厳密に分かれた責任などが秩序をもたらします！ただし、これらのモデルは、変更が遅くてコストがかかる鉄時代を最適化するためのものです。後で変更を行うことが遅くてコストがかかる場合、最初に追加の作業を行い、後で変更にかかる時間を減らすことを期待します。ただし、クラウドを利用すると変更が安価で迅速になります。あなたはこの速度を利用してシステムを継続的に学習し、改善する必要があります。鉄時代の働き方は学習と改善のための非常に大きな負担になります。素早いクラウド時代のテクノロジーを使って、リスクを減らし、品質を向上させるために、新しいマインドセットを取り入れてください。遅い鉄時代のプロセスを速いクラウド時代のテクノロジーと組み合わせて使用する代わりに、新しいアプローチを採用してください。高い信頼性と品質を実現するために、Infrastructure as Code を活用します。Infrastructure as Code は、ソフトウェア開発からのプラクティスを基にしたインフラストラクチャ自動化のアプローチです。これは、システムとその設定を供給および変更するための一貫した反復可能な手順を重視します。コードに変更を加え、その変更をシステムに対してテストし適用するために自動化を使用します。本書では、テスト駆動開発（TDD）、継続的インテグレーション（CI）、継続的デリバリー（CD）などのアジャイルなエンジニアリングプラクティスを使用して、インフラストラクチャの変更を迅速かつ安全に行う方法について説明します。また、現代のソフトウェア設計が、回復力のある、メンテナンス性の高いインフラストラクチャを作り出すことができる方法についても説明します。これらのプラクティスと設計アプローチは互いを補完し合います。よく設計されたインフラストラクチャは、テストと提供が容易になります。自動化されたテストと提供は、よりシンプルでクリーンな設計を促進します。インフラストラクチャのベネフィットを要約すると、動的なインフラストラクチャを管理するために Infrastructure as Code を採用する組織は、以下のような利点を得ることを期待しています：価値を迅速に提供するための IT インフラストラクチャの有効活用