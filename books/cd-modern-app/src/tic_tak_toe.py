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
                    # print("Player1 win")
                    # 値を返すように変更
                    return (True, 1)
                else:
                    # print("Player2 win")
                    # 値を返すように変更
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
                    # print("Player1 win")
                    # 値を返すように変更
                    return (True, 1)
                else:
                    # print("Player2 win")
                    # 値を返すように変更
                    return (True, 2)
                break
    return (False, 0)


def tic_tak_toe(inputs):
    flag = True
    player1 = True
    player2 = False

    filed = make_tic_tak_toe_field()

    for input in inputs:
        input_x = input[0]
        input_y = input[1]

        if int(input_x) > 2 or int(input_y) > 7:
            print("Invalid input")
            print("Please enter again x and y")
            continue

        if filed[int(input_x)][int(input_y)] == 1:
            print("Already input")
            print("Please enter again x and y")
            continue
        if player1:
            filed[int(input_x)][int(input_y)] = 1
            player1 = False
            player2 = True
        elif player2:
            filed[int(input_x)][int(input_y)] = 2
            player1 = True
            player2 = False

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
            print(row_str)

        # 勝敗判定
        # 行
        (result, player) = is_win_at_row_of(filed)
        if result:
            return (result, player)
        # 列
        (result, player) = is_win_at_column_of(filed)
        if result:
            return (result, player)

        for i in range(3):
            mark = filed[i][i]
            if mark == 0:
                continue
            if filed[i][i] != mark:
                break
            if filed[i][i] == mark and i == 2:
                flag = False
                if mark == 1:
                    # print("Player1 win")
                    # 値を返すように変更
                    return (True, 1)
                else:
                    # print("Player2 win")
                    # 値を返すように変更
                    return (True, 2)
                break

    # ここまで来たら勝敗が決まっていないのでFalseを返す
    return (False, 0)
