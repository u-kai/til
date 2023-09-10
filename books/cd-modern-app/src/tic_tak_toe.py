def main():
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
    game = TicTakToeGame(player1, player2, 3)
    while not result_flag:
        input_x = int(input("Please enter x"))
        input_y = int(input("Please enter y"))
        result = game.turn(input_x, input_y)
        if result is not None:
            print(result.name + "win")
            result_flag = True


class TicTakToeGame:
    def __init__(self, player1, player2, masu_num):
        self.player_index = 0
        self.players = [player1, player2]
        self.masu = Masu(masu_num)

    def turn(self, input_x, input_y):
        if not self.masu.validate_inputs(input_x, input_y):
            self.invalid_input_alert()
            return None

        self.masu.put(input_x, input_y, self.players[self.player_index].mark)
        self.display_field()
        if self.is_finished():
            return self.players[self.player_index]

        self.switch_player()
        return None

    def is_finished(self):
        return (
            self.masu.align_column() is not None
            or self.masu.align_row() is not None
            or self.masu.align_diagonal() is not None
        )

    def invalid_input_alert(self):
        print("Invalid input")
        print("Please enter again x and y")

    def display_field(self):
        print(self.masu.field_str())

    def switch_player(self):
        self.player_index = (self.player_index + 1) % 2


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

    def field_str(self):
        result = ""
        for i in range(self.masu_num):
            row_str = ""
            for j in range(self.masu_num):
                if self.filed[i][j] != 0:
                    row_str += "|" + str(self.filed[i][j])
                else:
                    row_str += "|□"
            row_str += "|"
            result += row_str + "\n"
        return result

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

        # print(field_str(masu))

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
