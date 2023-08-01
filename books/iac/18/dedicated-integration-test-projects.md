Dedicated Integration Test Projects An alternative approach is to create a separate project for integration tests, perhaps one for each integration
integration stage. This approach is common when a different team owns the integration tests, as predicted by Conway’s Law. Other teams do this when it’s not clear which project aligns with the integration tests. Versioning can be challenging when managing integration test suites separately from the code they test. People may confuse which version of the integration tests to run for a given version of the system code. To mitigate this, be sure to write and change tests along with the code, rather than separating those activities. And implement a way to correlate project versions; for example, using the fan-in approach described in the delivery-time integration pattern (see “Pattern: Delivery-Time Project Integration”).

専用の統合テストプロジェクト
別のアプローチとしては、統合テストのために独立したプロジェクトを作成する方法があります。統合テストを担当する別のチームが存在する場合、または統合テストがどのプロジェクトと関連しているか明確ではない場合に、このアプローチが一般的に使用されます。統合テストのコードとは別に管理することで、バージョン管理は難しくなるかもしれません。人々は、システムコードの特定のバージョンに対してどの統合テストのバージョンを実行すべきかを混同する可能性があります。これを緩和するために、コードと一緒にテストを書き換えるようにし、これらの活動を分離しないようにすることが重要です。また、プロジェクトのバージョンを関連付ける方法を実装する必要があります。たとえば、「パターン：デリバリータイムプロジェクト統合」に記載されているファンインアプローチを使用することが考えられます。

- これはいろいろ遅くなりそうだけど...
- 中央集権的にガンバナンスを聞かせるような統合テストならあり？
  - それでも相当遅くなりそう
