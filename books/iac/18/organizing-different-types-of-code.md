Organizing Different Types of Code Different projects in an infrastructure codebase define different types of elements of your system, such as applications, infrastructure stacks, server configuration modules, and libraries. And these projects may include different types of code, including declarations, imperative code, configuration values, tests, and utility scripts. Having a strategy for organizing these things helps to keep your codebase maintainable. Project Support Files Generally speaking, any supporting code for a specific project should live with that project’s code. A typical project layout for a stack might look like Example 18-1. Example 18-1. Example folder layout for a project
├── build.sh
├── src/
├── test/
├── environments/
└── pipeline/ This project’s folder structure includes: src/ The infrastructure stack code, which is the heart of the project. test/ Test code. This folder can be divided into subfolders for tests that run at different phases, such as offline and online tests. Tests that use different tools, like static analysis, performance tests, and functional tests, probably have dedicated subfolders as well. environments/ Configuration. This folder includes a separate file with configuration values for each stack instance. pipeline/ Delivery configuration. The folder contains configuration files to create delivery stages in a delivery pipeline tool (see “Delivery Pipeline Software and Services”). build.sh/ Script to implement build activities. See “Using Scripts to Wrap Infrastructure Tools” for a discussion of scripts like this one. Of course, this is only an example. People organize their projects differently and include many other things than what’s shown here. The key takeaway is the recommendation that files specific to a project live with the project. This ensures that when someone checks out a version of the project, they know that the infrastructure code, tests, and delivery are all the same version, and so should work together. If the tests are stored in a separate
project it would be easy to mismatch them, running the wrong version of the tests for the code you’re testing. However, some tests, configuration, or other files might not be specific to a single project. How should you handle these?

インフラストラクチャコードベースのさまざまなプロジェクトは、アプリケーション、インフラストラクチャスタック、サーバーの設定モジュール、ライブラリなど、システムのさまざまな要素の種類を定義しています。また、これらのプロジェクトには、宣言、命令型コード、設定値、テスト、ユーティリティスクリプトなど、さまざまな種類のコードが含まれる場合があります。これらの要素を整理するための戦略を持つことは、コードベースの保守性を保つのに役立ちます。

プロジェクトのサポートファイル
一般的に、特定のプロジェクトのサポートコードは、そのプロジェクトのコードと一緒に保存されるべきです。スタックのための典型的なプロジェクトレイアウトは、Example 18-1 のようになるでしょう。

Example 18-1. プロジェクトのフォルダレイアウトの例
├── build.sh
├── src/
├── test/
├── environments/
└── pipeline/

このプロジェクトのフォルダ構造には以下が含まれています：
src/ インフラストラクチャのスタックコードです。これがプロジェクトの中心です。
test/ テストコード。このフォルダは、オフラインテストやオンラインテストなど、異なる段階で実行されるテストのためにサブフォルダに分割することができます。静的解析、パフォーマンステスト、機能テストなど、異なるツールを使用するテストは、おそらく専用のサブフォルダを持っています。
environments/ 設定。このフォルダには、各スタックインスタンスごとに設定値が含まれる別々のファイルがあります。
pipeline/ デリバリーの設定。このフォルダには、デリバリーパイプラインツールでデリバリーステージを作成するための設定ファイルが含まれています（詳しくは「デリバリーパイプラインソフトウェアとサービス」を参照）。
build.sh/ ビルドアクティビティを実装するスクリプトです。このようなスクリプトについての詳細は、「インフラストラクチャツールをラップするスクリプトの使用」を参照してください。

もちろん、これは単なる例です。人々は自分のプロジェクトを異なる方法で整理し、ここに表示されているもの以外の多くのものを含めることができます。重要なポイントは、プロジェクトに固有のファイルはプロジェクトと一緒に保存することが推奨されているということです。これにより、誰かがプロジェクトのバージョンをチェックアウトした場合、インフラストラクチャコード、テスト、デリバリーがすべて同じバージョンであることを確認し、互換性があるはずです。テストが別のプロジェクトに保存されている場合、テストするコードのバージョンとは異なるバージョンのテストを実行することが簡単になります。ただし、一部のテスト、設定、または他のファイルは単一のプロジェクトに固有ではない場合もあります。それらはどのように扱うべきでしょうか？

- これってアプリケーションコードはないのか？

  - このテストには機能テストってあるけどアプリケーションコードのテストなのか，インフラの機能テストなのか？

- 同じプロジェクトなら一つのリポジトリに保管して一緒にバージョン管理しましょうって感じね
