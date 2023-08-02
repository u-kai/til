Who Writes Infrastructure Code? Here are a few different ways organizations answer the question of who writes and edits infrastructure code: Builders write code Some organizations try to keep traditional processes and team structures. So the team that builds (and perhaps supports) infrastructure uses Infrastructure as Code tools to optimize its work. Users request an environment, and the build team uses its tools and scripts to build it for them. See “Using Value Stream Mapping to Improve Workflows” for an example of how optimizing a build team’s process tends not to improve the end-to-end process, either for speed or quality. Users write code Many organizations enable application teams to define the infrastructure that their applications use. This aligns user needs to the solution. However, it either requires every team to include people with strong infrastructure expertise or tooling that simplifies defining infrastructure. The challenge with tooling is ensuring that it meets the needs of application teams, rather than constraining them. Toolmakers write code Specialist teams can create platforms, libraries, and tools that enable users to define the infrastructure they need. In these cases, the users tend to write more configuration than code. The difference between toolmakers and builders writing code is self-service. A builder team writes and uses code in response to requests from users to create or change an environment. A toolmaker team writes code that users can use to create or change their own environments. See “Building an Abstraction Layer” as an example of what toolmakers might build. Governance and testers write code
People who set policies and standards, and those who need to assure changes, can create tools that help other people to verify their own code. These people may become toolmakers or work closely with toolmakers. Using Value Stream Mapping to Improve Workflows Value stream mapping is a useful way to break down your lead time, so you can understand where the time goes.4 By measuring the time spent on various activities, including waiting, you can focus your improvements on areas that make the most difference. Too often, we optimize parts of our process that seem the most obviously inefficient but which have little impact on the total lead time. For example, I’ve seen teams implement automation to cut the time it takes to provision a server from eight hours to ten minutes. This is a massive 98% decrease in the time to provision a server. However, if users typically wait ten days to get their new server, the total decrease is a much less exciting 10% decrease. If server request tickets wait in a queue for an average of eight days, you should focus your efforts there instead. Value stream mapping makes the time to complete an action visible so that you can find the best opportunities for improvement. Continue to measure the end-to-end lead time, and other metrics such as failure rates, while you make improvements. Doing this helps avoid optimizations to one part of your process that make the full flow worse.

インフラストラクチャのコードを書くのは誰ですか？以下に、組織が誰がインフラストラクチャのコードを作成および編集するかという質問に対するいくつかの異なる方法を示します。

ビルダーがコードを作成します。一部の組織では、従来のプロセスとチーム構造を維持しようとします。したがって、インフラストラクチャを構築（およびサポート）するチームは、インフラストラクチャをコード化するツールを使用して作業を最適化します。ユーザーが環境をリクエストし、ビルドチームがツールとスクリプトを使用してそれを作成します。全体プロセスのスピードや品質において、ビルドチームのプロセスを最適化することは改善しない例を示すために、「バリューストリームマッピングを使用したワークフローの改善」を参照してください。

ユーザーがコードを書きます。多くの組織では、アプリケーションチームが使用するインフラストラクチャを定義できるようにしています。これにより、ユーザーのニーズがソリューションに合致します。ただし、それには強力なインフラストラクチャの専門知識を持つ人々またはインフラストラクチャを簡単に定義するツールが必要です。ツールの課題は、アプリケーションチームのニーズに合致するか、それを制約しないかを確実にすることです。

ツール作成者がコードを書きます。専門チームは、ユーザーが必要とするインフラストラクチャを定義できるようにするプラットフォーム、ライブラリ、およびツールを作成できます。これらの場合、ユーザーはコードよりも構成を書く傾向があります。ツール作成者がコードを書くこととビルダーチームがコードを書くことの違いは、セルフサービスです。ビルダーチームは、ユーザーのリクエストに応じて環境を作成または変更するためにコードを作成して使用します。ツール作成者チームは、ユーザーが自身の環境を作成または変更するために使用できるコードを作成します。これについては、ツール作成者が構築する可能性があるものの例として「抽象化レイヤーの構築」を参照してください。

ガバナンスおよびテスターがコードを書きます。方針と標準を設定する人々と変更を保証する必要がある人々は、他の人々が自分自身のコードを検証するのを助けるツールを作成できます。これらの人々はツール作成者になるか、ツール作成者と密接に連携することもあります。

バリューストリームマッピングを使用したワークフローの改善
バリューストリームマッピングは、リードタイムを細分化して、時間の経過を理解するための有用な手法です。さまざまな活動に費やされる時間、待機時間などを測定することにより、最も影響を与える領域に改善を焦点を当てることができます。私たちはしばしば、明らかに非効率なプロセスの一部を最適化しようとしますが、それが全体のリードタイムにほとんど影響を与えません。例えば、サーバーをプロビジョニングするのにかかる時間を 8 時間から 10 分に短縮するために自動化を導入するチームを見たことがあります。サーバーを提供するためにユーザーが通常 10 日待つのであれば、全体の減少率は 10％というあまり興奮しない数字です。もしサーバーリクエストのチケットが平均で 8 日間キューで待っているのであれば、そこに取り組むべきです。バリューストリームマッピングは、アクションを完了するための時間を可視化し、改善のための最適な機会を見つけるのに役立ちます。改善を行う間は、エンドツーエンドのリードタイムや障害率などのメトリクスを測定し続けることが重要です。これにより、プロセスの一部を最適化することで全体のフローが悪化することを防ぐことができます。

- ビルドチームがやるのは遅そう
  - 簡単なものであれば
- ユーザーが書くのが最近の流行りっぽそう

  - CDK とか Prumi とかはその一種

- ツール作成しや，ガバナンスチームが書くというのは大企業であれば目指されそう

  - うまく抽象化できれば，ユーザーが書くよりも簡単になるかもしれない

- バリューストリームマッピング活動も大切そう
  - 実際はどこが悪いんですか？と
  - これを図らないと，価値の薄いところを自動化してしまう
