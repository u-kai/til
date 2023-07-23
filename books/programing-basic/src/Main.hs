module Main where


addToEach :: [[a]] -> a -> [[a]]
addToEach [] _ = []
addToEach (x:xs) y = (y : x) : addToEach xs y
prefix :: [a] -> [[a]]
prefix [] = []
prefix (first:rest) = [first] : addToEach (prefix rest) first
insert  :: (Ord a) => [a] -> a -> [a]
insert [] x = [x]
insert (x:xs) y = if y < x then y:x:xs else x:insert xs y
insSort :: (Ord a) => [a] -> [a]
insSort [] = []
insSort (x:xs) = insert (insSort xs) x

data Student = Student {name :: String, grade::String} deriving (Show, Eq, Ord)
count0 :: [Student] -> String -> Int
count0 [] _ = 0
count0 (x:xs) g = if grade x == g then 1 + count0 xs g else count0 xs g

countA :: [Student] -> Int
countA x = count0 x "A"

map' :: (a -> b) -> [a] -> [b]
map' _ [] = []
map' f (x:xs)  = f x : map' f xs

compose' :: (a -> a) -> (a -> a) -> (a -> a)
compose' f g x =  f (g x)
time2 :: Int -> Int
time2 x = x * 2
add3 :: Int -> Int
add3 x = x + 3
filter' :: [a] -> (a -> Bool) -> [a]
filter' [] _ = []
filter' (x:xs) f = if f x then x : filter' xs f else filter' xs f
fold' :: (b -> a -> b) -> b -> [a] -> b
fold' _ init [] = init
fold' f init (x:xs) = f (fold' f init xs) x
accStr :: String -> String -> String
accStr x y = x ++ y
accStrs :: [String] -> String
accStrs [] = ""
accStrs (x:xs) = x ++ accStrs xs

main :: IO ()
main = do
    print(addToEach [[1,2,3],[4,5,6],[7,8,9]] 0)
    print(prefix [1,2,3,4])
    print(insert [1,2,4,5,6] 3)
    print(insSort [5,3,8,1,7,4])
    print(count0 [Student "John" "A", Student "Bob" "B", Student "Alice" "A"] "A")
    print(countA [Student "John" "A", Student "Bob" "B", Student "Alice" "A"])
    print(countA [])
    print(compose' time2 add3 4)
    print(filter' [1,2,3,4,5,6,7,8,9,10] even)
    print(foldr accStr "" ["c","h","i","n","k","o"])