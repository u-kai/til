# IaC 本

## What is Infrastructure as Code?

---

# From the Iron Age to the Cloud Age 1

- Technology changes in the Cloud Age

| Iron Age                 | Cloud Age                  |
| ------------------------ | -------------------------- |
| Physical hardware        | Virtualized resources      |
| Provisioning takes weeks | Provisioning takes minutes |
| Manual processes         | Automated processes        |

---

# From the Iron Age to the Cloud Age 2

- Ways of working in the Cloud Age

| Iron Age                                                            | Cloud Age                                         |
| ------------------------------------------------------------------- | ------------------------------------------------- |
| Cost of change is high                                              | Cost of change is low                             |
| Changes represent failure (changes must be “managed,” “controlled”) | Changes represent learning and improvement        |
| Reduce opportunities to fail                                        | Maximize speed of improvement                     |
| Deliver in large batches, test at the end                           | Deliver small changes, test continuously          |
| Long release cycles                                                 | Short release cycles                              |
| Monolithic architectures (fewer, larger moving parts)               | Microservices architectures (more, smaller parts) |
| GUI-driven or physical configuration                                | Configuration as Code                             |

---

# インフラストラクチャのコード化の利点

- IT インフラストラクチャを価値の迅速な提供が可能になること
- インフラストラクチャの変更の労力とリスクを削減できること
- インフラストラクチャの利用者が必要なリソースを必要な時に入手できるようになること
- 開発、オペレーション、その他の関係者間で共通のツールを提供可能になること
- 信頼性のある、安全な、費用効果の高いシステムを作ることが可能になること
- ガバナンス、セキュリティ、コンプライアンスのコントロールを可視化可能になること
- 障害のトラブルシューティングと解決のスピードを向上できること

---

## まず構築してから自動化すべき？

- 前提として IaC を始めるのは険しい道のり
- この作業の価値は実際にそれを使ってサービスを構築，デプロイする前には明示的に示すことは難しい
- そのため，早急な構築を望むステークホルダーからは手作業でやることを求められてしまう

  - ステークホルダーから直接言われなくても，納期などのプレッシャーに負けてしまうことはありそう

---

### 構築されたものを後から自動化することは次の 3 つの理由から良いアイデアではない

1. ほとんどの作業が終わった後に自動化することは、多くの利点を犠牲にしている
1. 自動化により、構築したものに対して自動テストを容易に行うことができるが，手動で作成したものに対する自動テストは難しい
1. コードで記述していれば問題が見つかった場合に、素早く修正や再構築ができるようになるが，手動では難しい，問題の特定や再現も難しい

- 既存のシステムに自動化を導入するのは非常に困難

### 自動化はシステムの設計と実装の一部

---
