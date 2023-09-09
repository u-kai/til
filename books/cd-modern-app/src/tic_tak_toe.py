def main():
    inputs = []
    result_flag = False
    print("Start")
    player1 = TicTakToePlayer(
        "Player1",
        "■",
    )
    player2 = TicTakToePlayer(
        "Player2",
        "●",
    )
    while not result_flag:
        input_x = input("Please enter x")
        input_y = input("Please enter y")
        inputs.append([input_x, input_y])
        (result, player) = tic_tak_toe(inputs, player1, player2)
        if result:
            print(result.name + "win")
            result_flag = True


def field_str(masu):
    result = ""
    for i in range(masu.masu_num):
        row_str = ""
        for j in range(masu.masu_num):
            if masu.filed[i][j] != 0:
                row_str += "|" + str(masu.filed[i][j])
            else:
                row_str += "|□"
        row_str += "|"
        result += row_str + "\n"
    return result


class Masu:
    # __init__はコンストラクタ呼ばれるもので、インスタンス生成時に呼ばれる
    # そのため初期化処理を記述することが多い
    def __init__(self, masu_num):
        self.masu_num = masu_num

        self.filed = []
        for i in range(self.masu_num):
            tmp = []
            for j in range(self.masu_num):
                tmp.append(0)
            self.filed.append(tmp)

    def put(self, x, y, data):
        self.filed[x][y] = data

    def align_column(self):
        for i in range(self.masu_num):
            mark = self.filed[0][i]
            if mark == 0:
                continue
            for j in range(self.masu_num):
                if self.filed[j][i] != mark:
                    break
                if j == self.masu_num - 1:
                    return mark
        return None

    def align_row(self):
        for i in range(self.masu_num):
            mark = self.filed[i][0]
            if mark == 0:
                continue
            for j in range(self.masu_num):
                if self.filed[i][j] != mark:
                    break
                if j == self.masu_num - 1:
                    return mark
        return None

    def align_diagonal(self):
        mark = self.filed[0][0]
        for i in range(self.masu_num):
            if mark == 0:
                break
            if self.filed[i][i] != mark:
                break
            if i == self.masu_num - 1:
                return mark

        mark = self.filed[0][self.masu_num - 1]
        for i in range(self.masu_num):
            if mark == 0:
                return None
            if self.filed[i][self.masu_num - 1 - i] != mark:
                return None
            if i == self.masu_num - 1:
                return mark
        return None

    def validate_inputs(self, input_x, input_y):
        if int(input_x) > self.masu_num - 1 or int(input_y) > self.masu_num - 1:
            return False
        if self.filed[int(input_x)][int(input_y)] != 0:
            return False
        return True


class TicTakToePlayer:
    def __init__(
        self,
        name,
        mark,
    ):
        self.name = name
        self.mark = mark

    def is_player_mark(self, mark):
        return self.mark == mark


def tic_tak_toe(player1, player2, inputs, masu_num=3):

    masu = Masu(masu_num)
    player1_flag = True

    for input in inputs:
        input_x = input[0]
        input_y = input[1]

        # 入力値のバリデーション
        if not masu.validate_inputs(input_x, input_y):
            print("Invalid input")
            print("Please enter again x and y")
            continue

        # フィールドに入力値を入れる
        if player1_flag:
            masu.put(input_x, input_y, player1.mark)
        elif player2:
            masu.put(input_x, input_y, player2.mark)

        # プレイヤーの切り替え
        player1_flag = not player1_flag

        print(field_str(masu))

        # 勝敗判定
        # 行
        mark = masu.align_row()
        if mark is not None:
            if player1.is_player_mark(mark):
                return player1
            else:
                return player2
        # 列
        mark = masu.align_column()
        if mark is not None:
            if player1.is_player_mark(mark):
                return player1
            else:
                return player2

        mark = masu.align_diagonal()
        if mark is not None:
            if player1.is_player_mark(mark):
                return player1
            else:
                return player2

    # ここまで来たら勝敗が決まっていないのでFalseを返す
    return None
