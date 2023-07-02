# 15 私のアプリケーションは API 呼び出しだらけです

- ライブラリなどのサードパーティコードはテストを必要ないと思いがち

  - 単にメソッドを呼び出すだけでコードは単純だと思い込みどんどんコードが育っていく

- そのうちに API と関係ないコードは存在するものの、テストできない寄せ集めのコードに埋め込まれた状態になる
  - 何かを変更する場合には常にアプリケーションを実行し、依然としてきちんと動くことを確認しなければいけない
- API 呼び出しだけのコードはいくつかの点で変更しにくくなる

  - API 呼び出ししかないためにたいていは構造を良くする方法を調べることが難しい.設計のヒントになりそうなものは何もない
  - API が自分たちの所有物ではない
    - もし自分たちのものならばインターフェースやクラスやメソッドの名前を理解しやすくなるように変更したり、コードの別の部分からも利用できるようにクラスにメソッドを追加したりすることができる

- 例としてメーリングリストサーバーのために書かれたもの

```java
import java.io.IOException;
import java.util.Properties;

import javax.mail.*;
import javax.mail.internet.*;

public class MailingListServer
{
    public static final String SUBJECT_MARKER = "[list]";
    public static final String LOOP_HEADER = "X-Loop";
    public static void main(String []args) {
        if (args.length != 8) {
            System.err.println("Usage: Java MailigList <popHost>" + "<smtpHost> <pop3user> <pop3password> " + "<smtpuser> <smtppassword> <listname> "+"<relayinterval>");
            return ;
        }
        HostInformation host = new HostInformation(args[0],args[1],args[2],args[3],args[4],args[5]);
        String listAddress = args[6];
        int interval = new Integer(args[7]).intValue();
        Roster roster = null;
        try {
            roster = new FileRoster("roster.txt");
        } catch (Exception e) {
            System.err.println("unable to open roster.txt");
            return ;
        }
        try {
            do {
                try {
                    Properties properties = System.getProperties();
                    Session session = Session.getDefaultInstance(properties,null);
                    Store store = session.getStore("pop3");
                    store.connect(host.pop3Host,-1,host.pop3User,host.pop3Password);
                    Folder defaultFolder = store.getDefaultFolder();
                    if (defaultFolder == null) {
                        System.err.println("Unable to open default folder");
                        return ;
                    }
                    Folder folder = defaultFolder.getFolder("INBOX");
                    if (folder == null) {
                        System.err.println("Unable to open default folder");
                        return ;
                    }
                    folder.open(Folder.READ_WRITE) ;
                    process(host,listAddress,roster,session,store,folder);
                } catch(Exception e) {
                    System.err.println(e);
                    System.err.println("retrying mail check");
                }
            }
        }catch(Exception e) {
            e.printStackTrace();
        }
    }
    private static void process(HostInformation host, String listAddress,Roster roster,Session session,Store store, Folder folder) {
        throws MessagingException {
            try {
                if (folder.getMessageCount() != 0) {
                    Message[] messages = folder.getMessages();
                    doMessage(host,listAddress,roster,session,folder,messages);
                }
            }catch(Exception e) {
                System.err.println("message handling error");
                e.printStackTrace(System.err);
            }finally{
                folder.close(true);
                store.close();
            }

        }
    }

...
}

```

- このコードをもっと良い構造にするには以下のステップを踏む
- 最初のステップはコードの主要な処理を識別すること

  - 上のコードの処理内容の簡単な説明
    - このコードは設定情報をコマンドラインから読み取り、電子メールアドレスのリストをファイルから読み取る
    - そしてメールを定期的に確認し、メールを見つけるとそのメールをファイル内の電子メールアドレスにそれぞれ転送する
  - 責務分けすると以下になる
    - 到着した各メッセージを受け取ってシステムに渡す責務
    - 単にメールのメッセージを送る責務
    - 到着した各メッセージと受取人名簿に基づいて、新しいメッセージを作る責務
    - ほとんどの時間をスリープしているが定期的に復帰してメールの到着を確認する責務
  - 上の中でもメールを利用するのは上 2 つ
  - そこはインターフェイス化して Fake にできるようにする

- ほとんどすべてのシステムには API 呼び出しから切り離せる主要なロジックがある
- API 呼び出しばかりのように見えるシステムは単に 1 つの大きなオブジェクトと取られて、20 章のこのクラスは大きすぎてもうこれ以上大きくしたくありません,で紹介されている責務を把握するための経験則を適用するとよい
- すぐに良い設計にはできないかもしれないが、責務を識別しておくことで、先に進んだ時に適切な判断がしやすくなる

- どうやって始めればよいかは基本的に二つのアプローチがある
  1. API をラップする
     - API が比較的小さいときに有効
     - サードパーティのライブラリのへの依存を完全に分離したいときに有効
     - テストがないことに加えて API を通じたテストが不可能ためテストを書くことができない時に有効
     - ラッパーから本物の API クラスへの委譲を行う薄いレイヤーを除いて、すべてのコードをテストで保護できる可能性がある
  1. 責務を基に抽出する
     - API が複雑であるときに有効
     - 安全なメソッド抽出をサポートするツールがあるか、手動での抽出を安全に行う自信があるときに有効
     - 上位レベルの名前を持つメソッドを抽出することから、API コードと一緒に自分たちの書いたロジックを抽出する結果になるかもしれない
     - これにより、コードを下位レベルの API 呼び出しではなく、上位レベルのインターフェイスに依存させることが可能になるが、抽出したコードはテストで保護できないかもしれない
- 多くのチームでは両方の手法を同時に使う
- その場合、薄いラッパーはテストのために、上位レベルのラッパーはアプリケーションの適切なインターフェイスを表現するために利用する
