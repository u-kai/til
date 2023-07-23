Governance and Server Images Traditional approaches to governance often include a manual review and approval process each time a new server image is built. As “Governance in a Pipeline-based Workflow” ex​plains, defining servers as code and using pipelines to deliver changes creates opportunities for stronger approaches. Rather than inspecting new server images, people responsible for governance can inspect the code that creates them. Better yet, they can collaborate with infrastructure and other teams involved to create automated checks that can run in pipelines to ensure compliance. Not only can your pipelines run these checks every time a new image is built, but they can also run them against new server instances to ensure that applying server configuration code hasn’t “unhardened” it. The ShopSpinner shopping application and search service frequently need to be scaled up and down to handle variable traffic loads at different times of the day. So it’s useful to have the images baked and ready. The team builds server instances for other applications and services from the more general application server image because the team doesn’t create instances of those as often. A separate pipeline builds each server image, triggering when the input image changes. The feedback loops can be very long. When the team changes the base Linux image, its pipeline completes and triggers the application server and container host pipelines. The application server image pipeline then triggers the pipelines to build new versions of the server images based on it. If each pipeline takes 10–15 minutes to run, this means it takes up to 45 minutes for the changes to ripple out to all of the images. An alternative is to use a shallower image building strategy. Sharing Code Across Server Images Even given an inheritance-based model where server roles inherit configuration from one another, you don’t necessarily need to build images in the layered model we just described. You can instead layer the server configuration code into server role definitions, and apply all of the code directly to build each image from scratch, as shown in Figure 13-10.
This strategy implements the same layered mode. But it does so by combining the relevant configuration code for each role and applying it all at once. So building the application server image doesn’t use the base server image as its source image, but instead applies all of the code used for the base server image to the application server image. Building images this way reduces the time it takes to build images at the bottom of the hierarchy. The pipeline that builds the image for the search service immediately runs when someone makes a change to the base Linux server configuration code.
ガバナンスとサーバーイメージ 伝統的なガバナンスの手法は、新しいサーバーイメージが作成されるたびに手動でレビューと承認プロセスが行われることがよくあります。『パイプラインベースのワークフローにおけるガバナンス』では、サーバーをコードとして定義し、変更を配信するためにパイプラインを使用することで、より強力なアプローチの機会が生まれます。新しいサーバーイメージを検査する代わりに、ガバナンスの責任者はそれらを作成するコードを検査することができます。さらに良いことに、ガバナンスに関与するインフラストラクチャや他のチームと協力して自動的なチェックを作成することができます。これらのチェックはパイプラインで実行され、コンプライアンスを確保することができます。パイプラインは新しいイメージが作成されるたびにこれらのチェックを実行するだけでなく、新しいサーバーインスタンスに対しても実行して、サーバー構成コードが「非保護」になっていないことを確認することもできます。ShopSpinnerのショッピングアプリケーションと検索サービスは、異なる時間帯の可変トラフィック負荷に対応するために頻繁にスケールアップやスケールダウンが必要です。そのため、イメージを事前に作成しておくことは便利です。チームは、より一般的なアプリケーションサーバーイメージから他のアプリケーションやサービスのサーバーインスタンスを作成するため、それらのインスタンスを頻繁に作成する必要がありません。各サーバーイメージごとに別個のパイプラインが構築され、入力イメージが変更された際にトリガーがかかります。フィードバックループは非常に長くなる場合があります。チームがベースとなるLinuxイメージを変更すると、そのパイプラインが完了し、アプリケーションサーバーやコンテナホストのパイプラインがトリガーされます。アプリケーションサーバーイメージのパイプラインは、それを基に新しいバージョンのサーバーイメージを構築するパイプラインをトリガーします。もし各パイプラインが10〜15分かかる場合、全てのイメージへ変更が波及するまで最大で45分かかります。別のアプローチとして、より浅いイメージ構築戦略を使用することもできます。サーバーロールが互いに設定を継承する継承ベースのモデルであっても、私たちが説明したレイヤードモデルでイメージを構築する必要はありません。代わりに、サーバーロールの定義にサーバー構成コードをレイヤー化し、すべてのコードを一度に適用して各イメージをゼロから構築することができます（図13-10参照）。
この戦略は同じレイヤードモデルを実装していますが、各ロールの関連する構成コードを組み合わせて一括で適用しています。そのため、アプリケーションサーバーイメージを構築する際には、ベースサーバーイメージをソースイメージとして使用せず、ベースサーバーイメージに使用されるすべてのコードをアプリケーションサーバーイメージに適用します。このようにイメージを構築することで、階層の一番下でのイメージ構築にかかる時間を短縮することができます。検索サービスのイメージを構築するパイプラインは、ベースLinuxサーバーの構成コードに変更が加えられた際にすぐに実行されます。