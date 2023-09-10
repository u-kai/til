def main():
    player1 = TicTakToePlayer(
        "Player1",
        "■",
    )
    player2 = TicTakToePlayer(
        "Player2",
        "●",
    )
    masu_num = 3
    game = TicTakToeGame(
        player1,
        player2,
        masu_num,
    )
    game.start(TerminalInputProvider())


# InputProviderを実装した具象クラス
class TerminalInputProvider:
    def provide_x(self):
        return int(input("Please enter x"))

    def provide_y(self):
        return int(input("Please enter y"))


# InputProviderを実装した具象クラス
# テスト用に入力値を事前に設定できるようにした
class FakeInputProvider:
    def __init__(self, all_inputs):
        self.all_inputs = all_inputs
        self.index = 0

    def provide_x(self):
        return self.all_inputs[self.index][0]

    def provide_y(self):
        result = self.all_inputs[self.index][1]
        self.index += 1
        return result


"""
interface InputProvider {
    provide_x():int
    provide_y():int
}

"""


class TicTakToeGame:
    def __init__(self, player1, player2, masu_num):
        self.player_index = 0
        self.players = [player1, player2]
        self.masu = Masu(masu_num)

    def start(self, provider):
        self.display_start()
        while not self.is_finished():
            x = provider.provide_x()
            y = provider.provide_y()
            result = self.turn(x, y)
            if result is not None:
                self.display_end(result)
                return result
        return None

    def display_start(self):
        print("Start")

    def display_end(self, player):
        print(player.name + "win")

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
