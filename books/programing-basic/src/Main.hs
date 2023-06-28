module Main where

add ::  Int ->  Int -> Int
add x y = x + y 
main :: IO ()
main = do
    let result =  add 1 2
    print (show result)