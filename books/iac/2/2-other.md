Organizations have tension between allowing each team to choose technologies and solutions that are appropriate to their needs, versus keeping the amount of variation in the organization to a manageable level.
The automation fear spiral describes how many teams fall into configuration drift and technical debt. At an Open Space session on configuration automation at a DevOpsDays conference, I asked the group how many of them were using automation tools like Ansible, Chef, or Puppet. The majority of hands went up. I asked how many were running these tools unattended, on an automatic schedule. Most of the hands went down. Many people have the same problem I had in my early days of using automation tools. I used automation selectively — for example, to help build new servers, or to make a specific configuration change. I tweaked the configuration each time I ran it to suit the particular task I was doing. I was afraid to turn my back on my automation tools because I lacked confidence in what they would do. I lacked confidence in my automation because my servers were not consistent. My servers were not consistent because I wasn’t running automation frequently and consistently. This is the automation fear spiral, as shown in Figure 2-2. Infrastructure teams must break this spiral to use automation successfully. The most effective way to break the spiral is to face your fears. Start with one set of servers. Make sure you can apply, and then reapply, your infrastructure code to these servers. Then schedule an hourly process that continuously applies the code to those servers. Then pick another set of servers and repeat the process. Do this until every server is continuously updated. Good monitoring and automated testing builds the confidence to continuously synchronize your code. This exposes configuration drift as it happens, so you can fix it immediately.
組織は、各チームが自身のニーズに適したテクノロジーやソリューションを選択することと、組織内のバリエーションを管理可能なレベルに保つこととの間に緊張関係があります。
自動化の恐怖のスパイラルは、多くのチームが構成のドリフトや技術的負債に陥る様子を説明しています。DevOpsDaysカンファレンスでの構成自動化に関するOpen Spaceセッションで、Ansible、Chef、またはPuppetのような自動化ツールを使用している人はどれくらいいるかと参加者に尋ねました。多くの手が挙がりました。それから、これらのツールを無人で自動的に実行している人はどれくらいいるかと尋ねました。ほとんどの手が下がりました。多くの人々は、私が自動化ツールを初めて使用したときと同じ問題を抱えています。私は自動化を選択的に使用しました-例えば、新しいサーバーの構築や特定の構成変更のために。実行するたびに構成を微調整し、特定のタスクに合わせてしました。私は自動化ツールに背を向けるのを恐れていました。それは何をするか自信がなかったからです。私のサーバーは一貫していませんでした。サーバーが一貫していなかったのは、自動化を頻繁かつ一貫して実行していなかったからです。これが図2-2に示される自動化の恐怖のスパイラルです。インフラストラクチャチームは、このスパイラルを断ち切って自動化を成功させる必要があります。スパイラルを断ち切る最も効果的な方法は、自分の恐怖に立ち向かうことです。まず、1つのサーバーセットから始めます。そのサーバーセットに対してインフラストラクチャコードを適用し、再適用できることを確認します。その後、そのサーバーにコードを継続的に適用する毎時のプロセスをスケジュールします。それから別のサーバーセットを選び、同じプロセスを繰り返します。すべてのサーバーが継続的に更新されるまでこれを行います。良好なモニタリングと自動化されたテストは、コードを連続的に同期する自信を築きます。これにより、構成のドリフトを即座に修正することができます。