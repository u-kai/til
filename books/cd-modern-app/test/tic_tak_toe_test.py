import unittest
from src.tic_tak_toe import tic_tak_toe, TicTakToePlayer, TicTakToeGame


class TestTicTakToe(unittest.TestCase):
    def test_turnは勝敗を返却する(self):
        masu_num = 3
        player1 = TicTakToePlayer("player1", "■")
        player2 = TicTakToePlayer("player2", "●")
        game = TicTakToeGame(player1, player2, masu_num)
        # put player1
        result = game.turn(0, 0)
        self.assertEqual(result, None)
        # put player2
        result = game.turn(1, 0)
        self.assertEqual(result, None)
        # put player1
        result = game.turn(1, 1)
        self.assertEqual(result, None)
        # put player2
        result = game.turn(0, 1)
        self.assertEqual(result, None)
        # put player1
        result = game.turn(2, 2)
        self.assertEqual(result.name, player1.name)
        self.assertEqual(result.mark, player1.mark)

    def test_マス目は自由に指定できる(self):
        masu_num = 8
        player1 = TicTakToePlayer("player1", "■")
        player2 = TicTakToePlayer("player2", "●")
        game = TicTakToeGame(player1, player2, masu_num)
        # put player1
        result = game.turn(0, 7)
        self.assertEqual(result, None)
        # put player2
        result = game.turn(7, 7)
        self.assertEqual(result, None)

    def test_マス目の数を指定できる(self):
        def test_マス目が8の場合_勝敗がつかない(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")
            game = TicTakToeGame(player1, player2, 8)
            # put player1
            result = game.turn(0, 0)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(1, 0)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(1, 1)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 1)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(2, 2)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 2)
            self.assertEqual(result, None)

        def test_マス目が8の場合_勝敗がついている(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")
            game = TicTakToeGame(player1, player2, 8)
            # put player1
            result = game.turn(0, 0)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(1, 0)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(1, 1)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 1)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(2, 2)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 2)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(3, 3)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 3)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(4, 4)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 4)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(5, 5)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 5)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(6, 6)
            self.assertEqual(result, None)
            # put player2
            result = game.turn(0, 6)
            self.assertEqual(result, None)
            # put player1
            result = game.turn(7, 7)
            self.assertEqual(result.name, player1.name)
            self.assertEqual(result.mark, player1.mark)

        test_マス目が8の場合_勝敗がつかない(self)
        test_マス目が8の場合_勝敗がついている(self)

    def test_不正な値が入力された場合は再度同じplayerが入力をやり直す(self):
        masu_num = 3
        player1 = TicTakToePlayer("player1", "■")
        player2 = TicTakToePlayer("player2", "●")
        game = TicTakToeGame(player1, player2, masu_num)
        # put player1
        result = game.turn(0, 0)
        self.assertEqual(result, None)
        # put player2
        result = game.turn(1, 0)
        self.assertEqual(result, None)
        # put player1
        result = game.turn(1, 1)
        self.assertEqual(result, None)
        # put player2
        result = game.turn(1, 2)
        self.assertEqual(result, None)
        # put player1
        result = game.turn(2, 3)
        self.assertEqual(result, None)
        # put player1
        result = game.turn(2, 2)
        self.assertEqual(result.name, player1.name)


if __name__ == "__main__":
    unittest.main()
