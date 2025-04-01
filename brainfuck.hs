import System.IO
import Control.Monad
import Data.Char

data State = State {
    ptr :: Int,
    dat :: [Int]
} deriving (Eq, Show, Read)

replaceAt :: Int -> [a] -> a -> [a]
--replaceAt idx xs v = (map (\(i, e) -> if i == idx then v else e) (zip [0..] xs))
replaceAt idx xs v = (map (\(i, e) -> case (i == idx) of 
    True -> v
    False -> e) (zip [0..] xs))

evalState :: Char -> State -> State
evalState cIn (State ptrIn datIn) = case cIn of
    '>' -> State (ptrIn+1) datIn
    '<' -> State (ptrIn-1) datIn
    '+' -> State ptrIn (replaceAt ptrIn datIn (datIn!!ptrIn + 1))
    '-' -> State ptrIn (replaceAt ptrIn datIn (datIn!!ptrIn - 1))
    _ -> (State ptrIn datIn)

evalString :: String -> State -> Int -> [Int] -> IO ()
evalString tIn (st) ptrIn jmpList | ptrIn >= length(tIn) = putStr ""
                                  | otherwise = case tIn !! ptrIn of
                                        '.' -> do
                                            putChar $ chr $ (dat st)!!(ptr st)
                                            evalString tIn (st) (ptrIn+1) jmpList
                                        ',' -> do
                                            charIn <- getLine
                                            let charNum = ord $ case charIn of
                                                                  [] -> '\0'
                                                                  (x:xs) -> x
                                            evalString tIn (State (ptr st) (replaceAt (ptr st) (dat st) charNum)) (ptrIn+1) jmpList
                                        '[' -> evalString tIn (st) (ptrIn+1) (jmpList++[ptrIn])
                                        ']' -> if (dat st)!!(ptr st) /= 0 then
                                            evalString tIn (st) ((last jmpList)+1) jmpList
                                            else evalString tIn (st) (ptrIn+1) (init jmpList)
                                        _ -> evalString tIn (evalState (tIn !! ptrIn) st) (ptrIn+1) jmpList

main:: IO()
main = do
    handle <- openFile "helloWorld.bf" ReadMode
    contents <- hGetContents handle
    evalString contents (State 0 (take 10024 (repeat 0))) 0 []