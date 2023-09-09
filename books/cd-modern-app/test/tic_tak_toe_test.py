import unittest
from src.tic_tak_toe import tic_tak_toe, TicTakToePlayer


class TestTicTakToe(unittest.TestCase):
    def test_マス目の数を指定できる(self):
        def test_マス目が8の場合_勝敗がつかない(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")
            player = tic_tak_toe(
                player1,
                player2,
                [
                    [0, 0],
                    [1, 0],
                    [1, 1],
                    [0, 1],
                    [2, 2],
                    [0, 2],
                ],
                8,
            )
            self.assertEqual(player, None)

        def test_マス目が8の場合_勝敗がついている(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")
            player = tic_tak_toe(
                player1,
                player2,
                [
                    [0, 0],
                    [1, 0],
                    [1, 1],
                    [0, 1],
                    [2, 2],
                    [0, 2],
                    [3, 3],
                    [2, 1],
                    [4, 4],
                    [2, 0],
                    [5, 5],
                    [1, 2],
                    [6, 6],
                    [3, 4],
                    [7, 7],
                ],
                8,
            )
            self.assertEqual(player.name, player1.name)
            self.assertEqual(player.mark, player1.mark)

        test_マス目が8の場合_勝敗がつかない(self)
        test_マス目が8の場合_勝敗がついている(self)

    def test_tik_tak_toeは複数のinputから勝敗を判定する(self):
        def test_勝敗は決まっていない(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")

            player = tic_tak_toe(player1, player2, [[0, 0]])

            self.assertEqual(player, None)

        def test_勝敗が決まっている(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")

            player = tic_tak_toe(
                player1, player2, [[0, 0], [1, 0], [1, 1], [0, 1], [2, 2]]
            )
            self.assertEqual(player.name, "player1")

        def test_斜めに並んだ場合_勝敗が決まっている(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")

            player = tic_tak_toe(
                player1, player2, [[0, 2], [1, 0], [1, 1], [2, 1], [2, 0]]
            )

            self.assertEqual(player.name, "player1")

        def test_不正な値が入力された場合はその値を無視する(self):
            player1 = TicTakToePlayer("player1", "■")
            player2 = TicTakToePlayer("player2", "●")
            player = tic_tak_toe(
                player1,
                player2,
                [
                    [0, 0],
                    [1, 0],
                    # 不正値
                    [3, 3],
                    [1, 1],
                    [0, 1],
                    [2, 2],
                ],
            )
            self.assertEqual(player.name, "player1")

        test_勝敗は決まっていない(self)
        test_勝敗が決まっている(self)
        test_斜めに並んだ場合_勝敗が決まっている(self)
        test_不正な値が入力された場合はその値を無視する(self)


if __name__ == "__main__":
    unittest.main()
