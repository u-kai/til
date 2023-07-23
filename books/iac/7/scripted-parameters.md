Pattern: Scripted Parameters Scripted parameters involves hardcoding the parameter values into a script that runs the stack tool. You can write a separate script for each environment or a single script that includes the values for all of your environments: if ${ENV} == "test"
stack up cluster_maximum=1 env="test"
elsif ${ENV} == "staging"
stack up cluster_maximum=3 env="staging"
elsif ${ENV} == "production"
stack up cluster_maximum=5 env="production"
end Motivation Scripts are a simple way to capture the values for each instance, avoiding the problems with the manual stack parameters antipattern (see “Antipattern: Manual Stack Parameters”). You can be confident that values are used consistently for each environment. By checking the script into version control, you ensure you are tracking any changes to the configuration values. Applicability A stack provisioning script is a useful way to set parameters when you have a fixed set of environments that don’t change very often. It doesn’t require the additional moving parts of some of the other patterns in this chapter.
Because it is wrong to hardcode secrets in scripts, this pattern is not suitable for secrets. That doesn’t mean you shouldn’t use this pattern, only that you’ll need to implement a separate pattern for dealing with secrets (see “Handling Secrets as Parameters” for suggestions). Consequences It’s common for the commands used to run the stack tool to become complicated over time. Provisioning scripts can grow into messy beasts. “Using Scripts to Wrap Infrastructure Tools” discusses how these scripts are used and outlines pitfalls and recommendations for keeping them maintainable. You should test provisioning scripts since they can be a source of issues with the systems they provision. Implementation There are two common implementations for this pattern. One is a single script that takes the environment as a command-line argument, with hardcoded parameter values for each environment. Example 7-6 is a simple example of this. Example 7-6. Example of a script that includes the parameters for multiple environments #!/bin/sh

case $1 in
test)
CLUSTER_MINIMUM=1
CLUSTER_MAXIMUM=1
;;
staging)
CLUSTER_MINIMUM=2
CLUSTER_MAXIMUM=3
;;
production)
CLUSTER_MINIMUM=2
CLUSTER_MAXIMUM=6
;;
\*)
echo "Unknown environment $1"
exit 1
;;
esac

stack
stack up \
 environment=$1 \
    cluster_minimum=${CLUSTER_MINIMUM} \
 cluster_maximum=${CLUSTER_MAXIMUM} Another implementation is a separate script for each stack instance, as shown in Example 7-7. Example 7-7. Example project structure with a script for each environment our-infra-stack/
├── bin/
│ ├── test.sh
│ ├── staging.sh
│ └── production.sh
├── src/
└── test/ Each of these scripts is identical but has different parameter values hardcoded in it. The scripts are smaller because they don’t need logic to select between different parameter values. However, they need more maintenance. If you need to change the command, you need to make it across all of the scripts. Having a script for each environment also tempts people to customize different environments, creating inconsistency. Commit your provisioning script or scripts to source control. Putting it in the same project as the stack it provisions ensures that it stays in sync with the stack code. For example, if you add a new parameter, you add it to the infrastructure source code and also to your provisioning script. You always know which version of the script to run for a given version of the stack code. “Using Scripts to Wrap Infrastructure Tools” discusses the use of scripts to run stack tools in much more detail. As mentioned earlier, you shouldn’t hardcode secrets into scripts, so you’ll need to use a different pattern for those. You can use the script to support that pattern. In Example 7-8, a command-line tool fetches the secret from a secrets manager, following the parameter registry pattern (see “Pattern: Stack Parameter Registry”). Example 7-8. Fetching a key from a secrets manager in a script ...

# (Set environment specific values as in other examples)

...

SSL_CERT_PASSPHRASE=$(some-tool get-secret id="/ssl_cert_passphrase/${ENV}")
stack up \
 environment=${ENV} \
    cluster_minimum=${CLUSTER_MINIMUM} \
 cluster_maximum=${CLUSTER_MAXIMUM} \
    ssl_cert_passphrase="${SSL_CERT_PASSPHRASE}" The some-tool command connects to the secrets manager and retrieves the secret for the relevant environment using the ID /ssl_cert_passphrase/${ENV}. This example assumes the session is authorized to use the secrets manager. An infrastructure developer may use the tool to start a session before running this script. Or the compute instance that runs the script may be authorized to retrieve secrets using secretless authorization (as I describe in “Secretless Authorization”). Related patterns Provisioning scripts run the command-line tool for you, so are a way to move beyond the manual stack parameters antipattern (see “Antipattern: Manual Stack Parameters”). The stack configuration files pattern extracts the parameter values from the script into separate files.
パターン: スクリプト化されたパラメータ
スクリプト化されたパラメータは、パラメータ値をスタックツールを実行するスクリプトにハードコーディングすることを意味します。各環境ごとに個別のスクリプトを作成するか、すべての環境の値が含まれた単一のスクリプトを作成することができます。

if ${ENV} == "test"
stack up cluster_maximum=1 env="test"
elsif ${ENV} == "staging"
stack up cluster_maximum=3 env="staging"
elsif ${ENV} == "production"
stack up cluster_maximum=5 env="production"
end

動機
スクリプトは、各インスタンスの値を取得するためのシンプルな方法であり、手動でスタックのパラメータを設定する際の問題を回避できます（「アンチパターン：手動スタックパラメータ」を参照）。各環境で一貫して値が使用されることを確認することができます。スクリプトをバージョン管理に組み込むことで、構成値の変更を追跡できるようにします。

適用範囲
スタックのプロビジョニングスクリプトは、あまり頻繁に変更されない固定の環境セットがある場合にパラメータを設定する便利な方法です。この章の他のパターンのような追加の要素は必要ありません。

ただし、スクリプトにシークレット情報をハードコードすることは間違っているため、このパターンはシークレットには適していません。これはこのパターンを使用しないべきではないことを意味しますが、シークレットを取り扱うための別のパターンを実装する必要があります（提案については「パラメータとしてのシークレットの処理」を参照）。

結果
スタックツールを実行するために使用されるコマンドが時間の経過とともに複雑になることは一般的です。プロビジョニングスクリプトは複雑になることがあります。これらのスクリプトが使用される方法と、保守性を保つための落とし穴や推奨事項については、「インフラストラクチャツールをラップするためのスクリプトの使用」で詳しく説明しています。プロビジョニングスクリプトはシステムの問題の原因になる可能性があるため、テストする必要があります。

実装
このパターンには2つの一般的な実装方法があります。1つは、環境をコマンドライン引数として受け取る単一のスクリプトで、各環境のハードコーディングされたパラメータ値があります。Example 7-6は、これのシンプルな例です。

Example 7-6. 複数の環境のパラメータを含むスクリプトの例

#!/bin/sh

case $1 in
test)
CLUSTER_MINIMUM=1
CLUSTER_MAXIMUM=1
;;
staging)
CLUSTER_MINIMUM=2
CLUSTER_MAXIMUM=3
;;
production)
CLUSTER_MINIMUM=2
CLUSTER_MAXIMUM=6
;;
\*)
echo "Unknown environment $1"
exit 1
;;
esac

stack
stack up \
 environment=$1 \
    cluster_minimum=${CLUSTER_MINIMUM} \
 cluster_maximum=${CLUSTER_MAXIMUM}

もう1つの実装方法は、各スタックインスタンスごとに個別のスクリプトを作成する方法です（Example 7-7を参照）。各スクリプトは同じですが、それぞれにハードコーディングされた異なるパラメータ値があります。これらのスクリプトは、さまざまなパラメータ値を選択するためのロジックが必要ないため、より小さいです。ただし、保守が必要です。コマンドを変更する必要がある場合は、すべてのスクリプトに変更を加える必要があります。各環境ごとにスクリプトを用意すると、異なる環境のカスタマイズが行われ、一貫性が失われる可能性があります。

プロビジョニングスクリプトまたはスクリプトをソースコントロールにコミットしてください。それをスタックを提供するのと同じプロジェクトに配置することで、スタックコードと同期することができます。たとえば、新しいパラメータを追加する場合は、インフラストラクチャのソースコードとプロビジョニングスクリプトの両方に追加する必要があります。特定のスタックコードのバージョンに対して実行するスクリプトのバージョンが常に分かります。

「インフラストラクチャツールをラップするためのスクリプトの使用」では、スクリプトを使用してスタックツールを実行する方法について詳しく説明しています。先ほども述べましたが、スクリプトにシークレット情報をハードコードすることは避けるべきですので、そのための異なるパターンを使用する必要があります。Example 7-8では、コマンドラインツールがシークレットマネージャーからキーを取得し、パラメータレジストリパターン（「パターン：スタックパラメータレジストリ」を参照）に従ってシークレットを取得しています。

Example 7-8. スクリプトでシークレットマネージャーからキーを取得する...

# (他の例と同様に環境固有の値を設定する)

...

SSL_CERT_PASSPHRASE=$(some-tool get-secret id="/ssl_cert_passphrase/${ENV}")
stack up \
 environment=${ENV} \
    cluster_minimum=${CLUSTER_MINIMUM} \
 cluster_maximum=${CLUSTER_MAXIMUM} \
    ssl_cert_passphrase="${SSL_CERT_PASSPHRASE}"

some-toolコマンドは、セッションがシークレットマネージャーを使用するための権限を持っていると想定しています。インフラストラクチャの開発者は、このスクリプトを実行する前にセッションを開始するためにこのツールを使用するかもしれません。または、スクリプトを実行するコンピューティングインスタンスがシークレットを取得するために秘密の認証を使用する権限を持っているかもしれません（「シークレットの認証なし」で説明する方法）。

関連するパターン
プロビジョニングスクリプトはコマンドラインツールを実行するため、手動のスタックパラメータのアンチパターンを回避する方法です（「アンチパターン：手動スタックパラメータ」を参照）。スタック構成ファイルパターンは、パラメータ値をスクリプトから別のファイルに抽出する方法です。