def main():
    tik_tak_toe()


def tik_tak_toe():
    print("Start")
    flag = True
    player1 = True
    player2 = False
    filed = []

    for i in range(3):
        tmp = []
        for j in range(3):
            tmp.append(0)
        filed.append(tmp)

    while flag:
        print("Enter x:")
        input_x = input()
        print("Enter y:")
        input_y = input()

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

        for i in range(3):
            mark = filed[i][0]
            if mark == 0:
                continue
            for j in range(3):
                if filed[i][j] != mark:
                    break
                if filed[i][j] == mark and j == 2:
                    flag = False
                    if mark == 1:
                        print("Player1 win")
                    else:
                        print("Player2 win")
                    break

        for j in range(3):
            mark = filed[0][j]
            if mark == 0:
                continue
            for i in range(3):
                if filed[i][j] != mark:
                    break
                if filed[i][j] == mark and i == 2:
                    flag = False
                    if mark == 1:
                        print("Player1 win")
                    else:
                        print("Player2 win")
                    break

        for i in range(3):
            mark = filed[i][i]
            if mark == 0:
                continue
            if filed[i][i] != mark:
                break
            if filed[i][i] == mark and i == 2:
                flag = False
                if mark == 1:
                    print("Player1 win")
                else:
                    print("Player2 win")
                break


if __name__ == "__main__":
    main()
