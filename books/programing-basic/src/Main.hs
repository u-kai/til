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

main :: IO ()
main = do
    print(addToEach [[1,2,3],[4,5,6],[7,8,9]] 0)
    print(prefix [1,2,3,4])
    print(insert [1,2,4,5,6] 3)
    print(insSort [5,3,8,1,7,4])