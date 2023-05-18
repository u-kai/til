# Role

- Role と Cluster Role はいずれも Namespace スコープのリソースを対象とした許可設定を行うことができる
- 例えば Deployment リソースの削除および変更を許可すると言った設定
- ClusterRole の場合は他にも Node/Namespace/PersistentVolume といった Cluster スコープのリソースや，/version,/healthz といった k8s API の情報を取得する nonResourceURL に対する権限も設定できる
- Role および ClusterRole 作成時には主に apiGroups,resources,verbs の 3 つを指定する
- apiGroups と resources で指定されるリソースに対して，verbs の権限を許可する
  - 利用できる主な verb は create,delete,get,list,patch,update,watch,\*である

# RoleBinding

- roleRef で紐付ける Role を subjects に紐づける User や ServiceAccount を指定する
- 1RoleBinding につき 1Role だけだが subjects には複数の User や ServiceAccount を指定することも可能
- RoleBinding は特定の Namespace において Role または ClusterRole で定義された権限を付与し，ClusterRoleBinding は全ての Namespace において ClusterRole で定義された権限を付与する
- Namespace を新規で作成してしまうと，その Namespace にも同じ RoleBinding を追加する必要があるが，ClusterRoleBinding を利用することで，Namespace 横断で同じ権限を付与することができる
  - つまり ClusterRoleBinding は Namespace に紐づかない Cluster レベルのリソース
  - User に対して全ての Namespace で ClusterRole で定義された権限を付与する
- RoleBinding を利用する場合でも，クラスタで共通の権限として定義した Cluster Role を利用することで，同じような Role の乱立を防ぐことができる
