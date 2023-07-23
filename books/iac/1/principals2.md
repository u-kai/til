Principle: Create Disposable Things Building a system that can cope with dynamic infrastructure is one level. The next level is building a system that is itself dynamic. You should be able to gracefully add, remove, start, stop, change, and move the parts of your system. Doing this creates operational flexibility, availability, and scalability. It also simplifies and de-risks changes. Making the pieces of your system malleable is the main idea of cloud native software. Cloud abstracts infrastructure resources (compute, networking, and storage) from physical hardware. Cloud native software completely decouples application functionality from the infrastructure it runs on.3 Cattle, Not Pets “Treat your servers like cattle, not pets,” is a popular expression about disposability.4 I miss giving fun names to each new server I create. But I don’t miss having to tweak and coddle every server in our estate by hand. If your systems are dynamic, then the tools you use to manage them need to cope with this. For example, your monitoring should not raise an alert every time you rebuild part of your system. However, it should raise a warning if something gets into a loop rebuilding itself.
People can take a while to get used to ephemeral infrastructure. One team I worked with set up automated infrastructure with VMware and Chef. The team deleted and replaced virtual machines as needed. A new developer on the team needed a web server to host files to share with teammates, so he manually installed an HTTP server on a development server and put the files there. A few days later, I rebuilt the VM, and his web server disappeared. After some confusion, the developer understood why this had happened. He added his web server to the Chef code, and persisted his files to the SAN. The team now had a reliable file-sharing service. Principle: Minimize Variation As a system grows, it becomes harder to understand, harder to change, and harder to fix. The work involved grows with the number of pieces, and also with the number of different types of pieces. So a useful way to keep a system manageable is to have fewer types of pieces — to keep variation low. It’s easier to manage one hundred identical servers than five completely different servers. The reproducibility principle (see “Principle: Make Everything Reproducible”) complements this idea. If you define a simple component and create many identical instances of it, then you can easily understand, change, and fix it. To make this work, you must apply any change you make to all instances of the component. Otherwise, you create configuration drift. Here are some types of variation you may have in your system: Multiple operating systems, application runtimes, databases, and other technologies. Each one of these needs people on your team to keep up skills and knowledge. Multiple versions of software such as an operating system or database. Even if you only use one operating system, each version may need different configurations and tooling. Different versions of a package. When some servers have a newer version of a package, utility, or library than others, you have risk. Commands may not run consistently across them, or older versions may have vulnerabilities or bugs. Organizations have tension between allowing each team to choose technologies and solutions that are appropriate to their needs, versus keeping the amount of variation in the organization to a manageable level.
原則：使い捨て可能なものを作成する
動的なインフラストラクチャに対応できるシステムを構築するのは一つのレベルです。次のレベルは、それ自体が動的なシステムを構築することです。システムの部品を優雅に追加、削除、開始、停止、変更、移動できるようにする必要があります。これにより、操作の柔軟性、可用性、拡張性が生まれます。また、変更が簡略化されリスクが軽減されます。システムの部品を可塑にすることが、クラウドネイティブソフトウェアの主要なアイデアです。クラウドは、物理的なハードウェアからインフラストラクチャリソース（計算、ネットワーキング、ストレージ）を抽象化します。クラウドネイティブソフトウェアは、アプリケーションの機能をインフラストラクチャから完全に切り離します。

牛ではなくペット「サーバーをペットではなく牛のように扱ってください」という言葉は、使い捨て可能性についての人気のある表現です。新しいサーバーを作成するたびに楽しい名前を付けることができないのは寂しいですが、手動でエステート内のすべてのサーバーを微調整しなければならないのは寂しくありません。システムが動的であれば、それを管理するためのツールもこれに対応する必要があります。例えば、システムの一部を再構築するたびに監視がアラートを上げるべきではありません。ただし、何かが自己再構築のループに入った場合は警告を上げるべきです。

人々は一時的なインフラストラクチャに慣れるまでに時間がかかることがあります。私が働いていたあるチームは、VMwareとChefを使用して自動化されたインフラストラクチャを設定しました。チームは必要に応じて仮想マシンを削除して置き換えました。チームの新しい開発者は、チームメートと共有するファイルをホストするためにWebサーバーが必要でしたので、開発用サーバーに手動でHTTPサーバーをインストールし、ファイルをそこに置きました。数日後、私はVMを再構築し、彼のWebサーバーが消えてしまいました。混乱の後、開発者はなぜこれが起こったのか理解しました。彼は自分のWebサーバーをChefのコードに追加し、ファイルをSANに保存しました。チームには信頼性の高いファイル共有サービスがありました。

原則：バリエーションの最小化
システムが成長すると、理解するのが難しく、変更するのが難しく、修正するのも難しくなります。関与する作業は、部品の数や異なる部品の数とともに増えていきます。したがって、システムを管理しやすくするための有用な方法は、部品の種類を少なくすることです。100個の同一のサーバーを管理する方が、5個の完全に異なるサーバーを管理するよりも簡単です。再現性の原則は、このアイデアを補完しています。シンプルなコンポーネントを定義し、それの同一のインスタンスを多数作成すると、それを容易に理解、変更、修正できます。これを実現するためには、行った変更をコンポーネントのすべてのインスタンスに適用する必要があります。そうしないと、設定のずれが生じます。システムには次のようなバリエーションの種類があります。複数のオペレーティングシステム、アプリケーションランタイム、データベース、その他のテクノロジー。これらのそれぞれに対して、チームのメンバーがスキルと知識を維持する必要があります。オペレーティングシステムやデータベースなどのソフトウェアの異なるバージョン。1つのオペレーティングシステムしか使用しない場合でも、それぞれのバージョンには異なる設定とツールが必要になる場合があります。パッケージの異なるバージョン。いくつかのサーバーに古いバージョンよりも新しいバージョンのパッケージ、ユーティリティ、またはライブラリがある場合、リスクがあります。コマンドが一貫して実行されず、古いバージョンには脆弱性やバグがある可能性があります。組織は、各チームが自分たちのニーズに合った技術とソリューションを選択できるようにするという目標と、組織内のバリエーションの量を管理可能なレベルに保つという目標との間で緊張関係があります。