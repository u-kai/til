def main():
    inputs = []
    result_flag = False
    print("Start")
    while not result_flag:
        input_x = input("Please enter x")
        input_y = input("Please enter y")
        inputs.append([input_x, input_y])
        (result, player) = tic_tak_toe(inputs)
        if result:
            print("Player" + player + "win")
            result_flag = True


def switch_player(player1, player2):
    if player1:
        player1 = False
        player2 = True
    elif player2:
        player1 = True
        player2 = False


def field_str(filed):
    result = ""
    for i in range(3):
        row_str = ""
        for j in range(3):
            if filed[i][j] == 1:
                row_str += "|■"
            elif filed[i][j] == 2:
                row_str += "|●"
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
                    return (True, mark)
        return (False, 0)

    def align_row(self):
        for i in range(self.masu_num):
            mark = self.filed[i][0]
            if mark == 0:
                continue
            for j in range(self.masu_num):
                if self.filed[i][j] != mark:
                    break
                if j == self.masu_num - 1:
                    return (True, mark)
        return (False, 0)

    def align_diagonal(self):
        mark = self.filed[0][0]
        for i in range(self.masu_num):
            if mark == 0:
                break
            if self.filed[i][i] != mark:
                break
            if i == self.masu_num - 1:
                if mark == 1:
                    return (True, 1)
                else:
                    return (True, 2)
                break

        mark = self.filed[0][self.masu_num - 1]
        for i in range(self.masu_num):
            if mark == 0:
                return (False, 0)
            if self.filed[i][self.masu_num - 1 - i] != mark:
                return (False, 0)
            if i == self.masu_num - 1:
                if mark == 1:
                    return (True, 1)
                else:
                    return (True, 2)
        return (False, 0)

    def validate_inputs(self, input_x, input_y):
        if int(input_x) > self.masu_num - 1 or int(input_y) > self.masu_num - 1:
            return False
        if self.filed[int(input_x)][int(input_y)] != 0:
            return False
        return True


def tic_tak_toe(inputs, masu_num=3):
    player1 = True
    player2 = False

    masu = Masu(masu_num)

    for input in inputs:
        input_x = input[0]
        input_y = input[1]

        # 入力値のバリデーション
        if not masu.validate_inputs(input_x, input_y):
            print("Invalid input")
            print("Please enter again x and y")
            continue

        # フィールドに入力値を入れる
        if player1:
            masu.put(input_x, input_y, 1)
        elif player2:
            masu.put(input_x, input_y, 2)

        # プレイヤーの切り替え
        switch_player(player1, player2)

        print(field_str(masu.filed))

        # 勝敗判定
        # 行
        (result, player) = masu.align_row()
        if result:
            return (result, player)
        # 列
        (result, player) = masu.align_column()
        if result:
            return (result, player)

        (result, player) = masu.align_diagonal()
        if result:
            return (result, player)

    # ここまで来たら勝敗が決まっていないのでFalseを返す
    return (False, 0)
