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


def make_tic_tak_toe_field():
    filed = []
    for i in range(3):
        tmp = []
        for j in range(3):
            tmp.append(0)
        filed.append(tmp)
    return filed


def is_win_at_row_of(filed):
    for i in range(3):
        mark = filed[i][0]
        if mark == 0:
            continue
        for j in range(3):
            if filed[i][j] != mark:
                break
            if filed[i][j] == mark and j == 2:
                if mark == 1:
                    return (True, 1)
                else:
                    return (True, 2)
                break
    return (False, 0)


def is_win_at_column_of(filed):
    for j in range(3):
        mark = filed[0][j]
        if mark == 0:
            continue
        for i in range(3):
            if filed[i][j] != mark:
                break
            if filed[i][j] == mark and i == 2:
                if mark == 1:
                    return (True, 1)
                else:
                    return (True, 2)
                break
    return (False, 0)


def is_win_at_diagonal_of(filed):
    mark = filed[0][0]
    for i in range(3):
        if mark == 0:
            break
        if filed[i][i] != mark:
            break
        if i == 2:
            if mark == 1:
                return (True, 1)
            else:
                return (True, 2)
            break

    mark = filed[0][2]
    for i in range(3):
        if mark == 0:
            return (False, 0)
        if filed[i][2 - i] != mark:
            return (False, 0)
        if i == 2:
            if mark == 1:
                return (True, 1)
            else:
                return (True, 2)
    return (False, 0)


def validate_inputs(filed, input_x, input_y):
    if int(input_x) > 2 or int(input_y) > 2:
        print("Invalid input")
        print("Please enter again x and y")
        return False
    if filed[int(input_x)][int(input_y)] == 1:
        print("Already input")
        print("Please enter again x and y")
        return False
    return True


def put(filed, input_x, input_y, player1, player2):
    if player1:
        filed[int(input_x)][int(input_y)] = 1
    elif player2:
        filed[int(input_x)][int(input_y)] = 2


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


def tic_tak_toe(inputs):
    flag = True
    player1 = True
    player2 = False

    filed = make_tic_tak_toe_field()

    for input in inputs:
        input_x = input[0]
        input_y = input[1]

        # 入力値のバリデーション
        if not validate_inputs(filed, input_x, input_y):
            continue

        # フィールドに入力値を入れる
        put(filed, input_x, input_y, player1, player2)

        # プレイヤーの切り替え
        switch_player(player1, player2)

        print(field_str(filed))

        # 勝敗判定
        # 行
        (result, player) = is_win_at_row_of(filed)
        if result:
            return (result, player)
        # 列
        (result, player) = is_win_at_column_of(filed)
        if result:
            return (result, player)

        (result, player) = is_win_at_diagonal_of(filed)
        if result:
            return (result, player)

    # ここまで来たら勝敗が決まっていないのでFalseを返す
    return (False, 0)
