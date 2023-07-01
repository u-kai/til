module Main where

factorial :: Int -> Int
factorial 0 = 1
factorial n = n * factorial (n - 1)


newtype Circle = Circle {
    r :: Float
}

bmiTell :: Double ->  String
bmiTell bmi
    | bmi <= 18.5 = "You're underweight, you emo, you!"
    | bmi <= 25.0 = "You're supposedly normal. Pffft, I bet you're ugly!"
    | bmi <= 30.0 = "You're fat! Lose some weight, fatty!"
    | otherwise   = "You're a whale, congratulations!"

main :: IO ()
main = do
    print(factorial 0)
    print(factorial 1)
    print(factorial 2)
    print(factorial 3)
    print(factorial 4)