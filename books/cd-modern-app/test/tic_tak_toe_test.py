import unittest
from src.tic_tak_toe import tic_tak_toe


class TestTicTakToe(unittest.TestCase):
    def test_tik_tak_toeは複数のinputから勝敗を判定する(self):
        def test_勝敗は決まっていない(self):
            (result, player) = tic_tak_toe([[0, 0]])
            self.assertEqual(result, False)
            self.assertEqual(player, 0)

        def test_勝敗が決まっている(self):
            (result, player) = tic_tak_toe([[0, 0], [1, 0], [1, 1], [0, 1], [2, 2]])
            self.assertEqual(result, True)
            self.assertEqual(player, 1)

        test_勝敗は決まっていない(self)
        test_勝敗が決まっている(self)


if __name__ == "__main__":
    unittest.main()
