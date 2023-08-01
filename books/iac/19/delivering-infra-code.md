Delivering Infrastructure Code The pipeline metaphor describes how a change to infrastructure code progresses from the person making the change to production instances. The activities required for this delivery process influence how you organize your codebase. A pipeline for delivering versions of code has multiple types of activities, including build, promote, apply, and validate. Any given stage in the pipeline may involve multiple activities, as shown in Figure 19-1.
Building Prepares a version of the code for use and makes it available for other phases. Building is usually done once in a pipeline, each time the source code changes. Promoting Moves a version of the code between delivery stages, as described in “Progressive Testing”. For example, once a stack project version has passed its stack test stage, it might be promoted to show it’s ready for the system integration test stage. Applying Runs the relevant tooling to apply the code to the relevant infrastructure instance. The instance could be a delivery environment used for testing activities or a production instance. “Infrastructure Delivery Pipelines” has more detail on pipeline design. Here, we’re focusing on building and promoting code. Building an Infrastructure Project Building an infrastructure project prepares the code for use. Activities can include:
インフラストラクチャコードの提供
パイプラインのメタファーは、インフラストラクチャコードの変更が変更を行う人から本番インスタンスに進むまでの進行状況を表しています。この提供プロセスには、コードベースの組織方法に影響を与える必要な活動が含まれます。コードのバージョンを提供するためのパイプラインには、ビルド、プロモート、適用、および検証など、複数のタイプの活動が含まれます。図 19-1 に示されているように、パイプラインの各ステージでは複数の活動が行われる場合があります。
ビルドは、コードのバージョンを準備し、他のフェーズで使用できるようにします。ビルドは通常、パイプライン内で一度行われ、ソースコードが変更されるたびに実行されます。プロモートは、コードのバージョンを配信段階間で移動させることを意味し、”Progressive Testing”で説明されています。例えば、スタックのプロジェクトバージョンがスタックテスト段階を通過した場合、システム統合テスト段階に移行することができます。適用は、関連するツールを実行してコードを関連するインフラストラクチャインスタンスに適用します。インスタンスは、テスト活動に使用される配信環境または本番インスタンスを指すことができます。”Infrastructure Delivery Pipelines”では、パイプラインの設計に関する詳細情報が提供されています。ここでは、コードのビルドとプロモートに焦点を当てています。インフラストラクチャプロジェクトのビルドインフラストラクチャプロジェクトのビルドは、コードを使用する準備をします。活動には以下が含まれます：
ビルド
promoting
デリバリーステージ間でコードのバージョンを移動
例えば一つの Stack でのテストが終了すれば，それはシステム統合テストの段階に移動することができる
applying

- 考えまとめ
  - パイプラインに変更を本番に適用するためのたくさんのことを詰めることができる
  - 段階的にそのステージで行うべきことを行なっていく感じ
    - 逆に段階段階でするべきことはしっかり設計すべきかも
    - ステージごとに責務を設けていくことが大事そう
