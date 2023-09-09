import requests


def main():
    flag = True
    filed = []

    for i in range(8):
        tmp = []
        for j in range(8):
            tmp.append(0)
        filed.append(tmp)

    while flag:
        print("Enter x:")
        input_x = input()
        print("Enter y:")
        input_y = input()

        if int(input_x) > 7 or int(input_y) > 7:
            print("Invalid input")
            print("Please enter again x and y")
            continue

        if filed[int(input_x)][int(input_y)] == 1:
            print("Already input")
            print("Please enter again x and y")
            continue

        filed[int(input_x)][int(input_y)] = 1

        for i in range(8):
            row_str = ""
            for j in range(8):
                if filed[i][j] == 1:
                    row_str += "|■"
                else:
                    row_str += "|□"
            row_str += "|"
            print(row_str)


if __name__ == "__main__":
    main()
