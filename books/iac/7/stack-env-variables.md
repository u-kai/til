Pattern: Stack Environment Variables The stack environment variables pattern involves setting parameter values as environment variables for the stack tool to use. This pattern is often combined with another pattern to set the environment variables. The environment variables are set beforehand, as shown in Example 7-4 (see “Implementation” for more on how). Example 7-4.
Motivation Most platforms and tools support environment variables, so it’s easy to do. Applicability If you’re already using environment variables in your system and have suitable mechanisms to manage them, you might find it convenient to use them for stack parameters. Consequences You need to use an additional pattern from this chapter to get the values to set. Doing this adds moving parts, making it hard to trace configuration values for any given stack instance, and more work to change the values.
Using environment variables directly in stack code, as in Example 7-5, arguably couples stack code too tightly to the runtime environment. Setting secrets in environment variables may expose them to other processes that run on the same system. Implementation Again, you need to set the environment variables to use, which means selecting another pattern from this chapter. For example, if you expect people to set environment variables in their local environment to apply stack code, you are using the manual stack parameters antipattern (see “Antipattern: Manual Stack Parameters”). You could set them in a script that runs the stack tool (the scripted parameters pattern, as discussed in “Pattern: Scripted Parameters”), or have the pipeline tool set them (see “Pattern: Pipeline Stack Parameters”). Another approach is to put the values into a script that people or instances import into their local environment. This is a variation of the stack configuration files pattern (see “Pattern: Stack Configuration Files”). The script to set the variables would be exactly like Example 7-4, and any command that runs the stack tool would import it into the environment: source ./environments/staging.env
stack up --source ./src Alternatively, you could build the environment values into a compute instance that runs the stack tool. For example, if you provision a separate CD agent node to run the stack tool to build and update stacks in each environment, the code to build the node could set the appropriate values as environment variables. Those environment variables would be available to any command that runs on the node, including the stack tool. But to do this, you need to pass the values to the code that builds your agent nodes. So you need to select another pattern from this chapter to do that. The other side of implementing this pattern is how the stack tool gets the environment values. Example 7-5 showed how stack code can directly read environment variables. But you could, instead, use a stack orchestration script (see “Using Scripts to Wrap Infrastructure Tools”) to read the environment variables and pass them to the stack tool on the command line. The code in the orchestration script would look like this.
This approach decouples your stack code from the environment it runs in. Related patterns Any of the other patterns in this chapter can be combined with this one to set environment values.
パターン：スタック環境変数
スタック環境変数パターンでは、パラメータの値をスタックツールが使用するための環境変数として設定します。このパターンは、環境変数を設定するために他のパターンと組み合わせることがよくあります。環境変数は事前に設定されます（詳細は「実装」を参照してください）。例7-4を参照してください。

動機
ほとんどのプラットフォームとツールは環境変数をサポートしているため、簡単に行うことができます。

適用性
すでにシステムで環境変数を使用しており、適切なメカニズムでそれらを管理する手段がある場合、スタックパラメータに使用するためにそれらを使用することが便利かもしれません。

結果
設定する値を取得するために、この章の他のパターンを使用する必要があります。これにより、構成値を特定のスタックインスタンスに追跡するのが難しくなり、値を変更するためにさらなる作業が必要になります。

例7-5のように環境変数を直接スタックコードで使用することは、スタックコードをランタイム環境に過度に結び付ける可能性があると言えます。環境変数でシークレットを設定すると、同じシステムで実行される他のプロセスにさらされる可能性があります。

実装
先に環境変数を設定する必要があります。これは、この章の他のパターンから別のパターンを選択することを意味します。たとえば、スタックコードを適用するためにローカル環境で人々が環境変数を設定することを期待する場合、手動スタックパラメータの反パターンを使用しています（「反パターン：手動スタックパラメータ」を参照）。スタックツールを実行するスクリプトにそれらを設定することもできます（「パターン：スクリプト付きパラメータ」を参照）、またはパイプラインツールがそれらを設定することもできます（「パターン：パイプラインスタックパラメータ」を参照）。別のアプローチとしては、人々やインスタンスがローカル環境にインポートするスクリプトに値を入れることができます。これはスタック構成ファイルのパターンのバリエーションです（「パターン：スタック構成ファイル」を参照）。変数を設定するスクリプトは、例7-4とまったく同じです。スタックツールを実行する任意のコマンドでそれを環境にインポートします：ソース./environments/staging.env
stack up --source./src

または、スタックツールを実行する計算インスタンスに環境値を組み込むこともできます。たとえば、各環境でスタックを構築および更新するために別のCDエージェントノードをプロビジョニングする場合、ノードを構築するコードは適切な値を環境変数として設定します。これらの環境変数は、スタックツールを含むノード上で実行される任意のコマンドで利用できます。ただし、これを行うためには、エージェントノードを構築するコードに値を渡す必要があります。したがって、これを行うためには、この章の他のパターンから別のパターンを選択する必要があります。このパターンの実装のもう一方の側面は、スタックツールが環境の値を取得する方法です。例7-5では、スタックコードが直接環境変数を読み取る方法を示しました。ただし、代わりにスタックオーケストレーションスクリプト（「インフラツールをラップするスクリプトの使用」を参照）を使用して、環境変数を読み取り、それらをスタックツールにコマンドラインで渡すことができます。オーケストレーションスクリプトのコードは次のようになります。

このアプローチにより、スタックコードと実行環境が分離されます。

関連パターン
この章の他のパターンのいずれかを組み合わせて、このパターンを使用して環境値を設定できます。