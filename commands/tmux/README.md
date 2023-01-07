# tmux について

-   セッション

    -   1 つ以上のウィンドウを管理しているターミナル全体

-   ウィンドウ

    -   セッション内に開かれている 1 つ以上のペインを管理しているターミナル

-   ペイン

    -   ウィンドウないで分割されているウィンドウ

-   .tumx.conf に設定をかける

    ```
    # vimのキーバインドでペインを移動する
     bind h select-pane -L
     bind j select-pane -D
     bind k select-pane -U
     bind l select-pane -R
     # | でペインを縦分割する
     bind | split-window -h
     # - でペインを縦分割する
     bind - split-window -v
     bind -r H resize-pane -L 5
     bind -r J resize-pane -D 5
     bind -r K resize-pane -U 5
     bind -r L resize-pane -R 5
    ```

-   参考

    https://qiita.com/shin-ch13/items/9d207a70ccc8467f7bab
