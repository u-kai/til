Using Scripts to Wrap Infrastructure Tools Most teams managing infrastructure code create custom scripts to orchestrate and run their infrastructure tools. Some use software build tools like Make, Rake, or Gradle. Others write scripts in Bash, Python, or PowerShell. In many cases, this supporting code becomes at least as complicated as the code that defines the infrastructure, leading teams to spend much of their time debugging and maintaining it. Teams may run these scripts at build time, delivery time, or apply time. Often, the scripts handle more than one of these project phases. The scripts handle a variety of tasks, which can include: Configuration Assemble configuration parameter values, potentially resolving a hierarchy of values. More on this shortly.
Dependencies Resolve and retrieve libraries, providers, and other code. Packaging Prepare code for delivery, whether packaging it into an artifact or creating or merging a branch. Promotion Move code from one stage to the next, whether by tagging or moving an artifact or creating or merging a branch. Orchestration Apply different stacks and other infrastructure elements in the correct order, based on their dependencies. Execution Run the relevant infrastructure tools, assembling command-line arguments and configuration files according to the instance the code is applied to. Testing Set up and run tests, including provisioning test fixtures and data, and gathering and publishing results.

インフラストラクチャのコードを管理するほとんどのチームは、インフラストラクチャツールのオーケストレーションや実行にカスタムスクリプトを作成します。いくつかのチームは、Make、Rake、または Gradle といったソフトウェアビルドツールを使用します。他のチームは、Bash、Python、または PowerShell でスクリプトを記述します。多くの場合、このサポートコードはインフラストラクチャの定義コードと同じくらい複雑になり、チームはデバッグやメンテナンスに多くの時間を費やすことになります。チームはこれらのスクリプトをビルド時間、デリバリー時間、または適用時間に実行することがあります。しばしば、これらのスクリプトは複数のプロジェクトフェーズを扱います。スクリプトはさまざまなタスクを処理します。これには以下のものが含まれます:

- 設定: 設定パラメータの値を組み立て、値の階層を解決することができます。
- 依存関係: ライブラリ、プロバイダ、および他のコードを解決および取得することができます。
- パッケージング: コードを配信用に準備し、それをアーティファクトにパッケージングしたり、ブランチを作成またはマージしたりします。
- プロモーション: タグ付けやアーティファクトの移動、またはブランチの作成やマージなどを使用して、コードを次のステージに移動します。
- オーケストレーション: 依存関係に基づいて、異なるスタックやその他のインフラストラクチャ要素を正しい順序で適用します。
- 実行: 関連するインフラストラクチャツールを実行し、コマンドライン引数と設定ファイルをコードが適用されるインスタンスに応じて組み立てます。
- テスト: テストフィクスチャやデータのプロビジョニング、結果の収集と公開など、テストのセットアップと実行を行います。

- 考えまとめ
  - どっかの記事で，Actions で実行される処理は Make ファイルにまとめて可読性や管理性を上げたみたいなことがあった
  - 確かにカスタムスクリプトを Pipeline の記述とは別に持っておくといいんかな？
