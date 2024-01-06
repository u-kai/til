数十年にわたり、データ管理は常に重要なアーキテクチャ上の課題でした。私のキャリア初期には、単一の企業全体のデータモデルが非常に注目され、しばしば単一の企業全体のデータベースに格納されていました。
しかし、共有データストアにアクセスする多数のアプリケーションが存在することが、アドホックなカップリングの災害であることがすぐにわかりました。
それに加えて、より深刻な問題が存在しました。顧客などの中核的な企業概念は、異なるビジネスユニットで異なるデータモデルを必要とすることがありました。企業の買収は状況をさらに複雑にしました。
そのため、より賢明な企業は、データを分散させて、データの格納、モデル、管理を異なるビジネスユニットに委任しました。
このようにして、各領域でデータを最も理解している人々がそのデータを管理する責任を持つようになりました。彼らは他のドメインとよく定義された API を通じて協力します。これらの API には動作が含まれるため、データの共有方法やより重要な点であるデータ管理の進化に対して柔軟性を持つことができます。これは日常の業務においてますます重要な方法となってきましたが、データ分析は依然としてより集中した活動でした。
データウェアハウスは、企業の扱いやすい重要な情報のリポジトリを提供することを目指していました。
しかし、そのような中央集権的なグループは、データやその消費者のニーズを十分に理解していなかったため、その仕事と競合する顧客との折り合いがつかなかったのです。
データレイクは、生データへのアクセスを普及させることで解決を助けましたが、理解や起源の貧弱なデータスワンプとなることが簡単でした。
データメッシュは、運用データの世界において学んだ教訓を分析データの世界に適用しようとしています。
ビジネスユニットのドメインは、運用データと同様に、分析データを API を通じて公開する責任を持つようになります。
彼らはそのデータを第一級の製品として扱い、データの意味や起源を伝え、消費者と協力します。この目標を達成するためには、企業がこれらのデータ製品の構築と公開のためのプラットフォームを提供し、それを統合的なガバナンス構造で維持する必要があります
。すべてに渡って技術的な優れたものであることの重要性を認識しており、ビジネスのニーズが変わるにつれて、プラットフォームや製品を迅速に進化させることができます。
したがって、データメッシュは、分析データの世界において、既存のデータ管理の原則を当然のこととして適用したものです。
しかし、実際には、特に多くのベンダーの投資が中央集権的なモデルに焦点を当ててきたため、この方法を実現するためには多くの工夫が必要です（テスト、抽象化構築、リファクタリングなどの開発者が健全なソフトウェアに不可欠と知っている手法をサポートしていないことが一因です）。
Zhamak は、クライアントに先進的なアドバイスを行い、彼らの挫折と成功から学び、ツールを開発しやすくするためにベンダーを促してきました。
この本では、データメッシュが世界中で採用される初期だが重要な段階における彼女と同僚の知識をまとめています。私はこの本をレビューする過程で、これらの実践的な困難について多くのことを学びました。そして、データリソースを最大限に活用したいと思う組織の誰もが、私たちの理解する進むべき道をこの本で見出すでしょう。
